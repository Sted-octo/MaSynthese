terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project     = var.project_id
  region      = var.region
}

# Activate necessary APIs
resource "google_project_service" "containerregistry" {
  service            = "containerregistry.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "cloudbuild" {
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "run" {
  service            = "run.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "artifactregistry" {
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = false
  depends_on = [google_project_service.containerregistry, google_project_service.cloudbuild, google_project_service.run]
}

# Create Artifact Registry repository
resource "google_artifact_registry_repository" "default" {
  location      = var.region
  repository_id = var.artifact_repo_name
  format        = "DOCKER"
  description   = "Docker repository for octoptimist app"
  depends_on = [google_project_service.artifactregistry]
}

# Création du réseau VPC
resource "google_compute_network" "octoptimist_vpc" {
  name                    = "octoptimist-vpc"
  auto_create_subnetworks = false  # On gère les sous-réseaux explicitement
  project                 = var.project_id
}

# Création d'un sous-réseau
resource "google_compute_subnetwork" "octoptimist_subnet" {
  name                     = "octoptimist-subnet"
  ip_cidr_range            = "10.10.0.0/24"
  region                   = var.region
  network                  = google_compute_network.octoptimist_vpc.id
  private_ip_google_access = true
  project                  = var.project_id
  
  # Ajout de la configuration des journaux de flux VPC
  log_config {
    aggregation_interval = "INTERVAL_5_SEC"
    flow_sampling        = 0.5
    metadata             = "INCLUDE_ALL_METADATA"
  }
}

# Création d'une adresse IP pour le Cloud NAT
resource "google_compute_address" "octoptimist_nat_ip" {
  name         = "octoptimist-nat-ip"
  region       = var.region
  project      = var.project_id
}

# Création du Cloud NAT
resource "google_compute_router" "octoptimist_router" {
  name    = "octoptimist-router"
  region  = var.region
  network = google_compute_network.octoptimist_vpc.id
  project = var.project_id
}

resource "google_compute_router_nat" "octoptimist_nat" {
  name                               = "octoptimist-nat"
  router                             = google_compute_router.octoptimist_router.name
  region                             = var.region
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips                            = [google_compute_address.octoptimist_nat_ip.self_link]  # Changement de nat_ip_names à nat_ips et utilisation de self_link
  project                            = var.project_id

  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}

# Création du connecteur Serverless VPC Access
resource "google_vpc_access_connector" "octoptimist_connector" {
  name          = "octoptimist-connector"
  region        = var.region
  network       = google_compute_network.octoptimist_vpc.id
  ip_cidr_range = "10.8.0.0/28"  # Choisir une plage non chevauchante
  project       = var.project_id
}

# google_cloudbuild_trigger.default:
resource "google_cloudbuild_trigger" "default" {
    description        = "Trigger build d'octoptimist"
    location        = "global"
    name               = "octoptimist-trigger"
    project            = var.project_id
    filename = "cloudbuild.yaml"
    substitutions      = {
      _REGION = var.region
      _SERVICE_NAME = var.service_name
      _ARTIFACT_REPO_NAME = var.artifact_repo_name
      _TAG = var.tag
      _VPC_CONNECTOR = google_vpc_access_connector.octoptimist_connector.name
    }

    approval_config {
        approval_required = false
    }

    github {
        enterprise_config_resource_name = null
        name                            = "MaSynthese"
        owner                           = "Sted-octo"

        push {
            branch       = "^master$"
            invert_regex = false
            tag          = null
        }
    }
}