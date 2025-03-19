variable "project_id" {
  type        = string
  description = "ID du projet GCP"
}

variable "region" {
  type        = string
  description = "Région GCP"
  default     = "europe-west1"
}

variable "service_name" {
  type        = string
  description = "Nom du service Cloud Run"
  default     = "octoptimist-service"
}

variable "artifact_repo_name" {
  type        = string
  description = "Nom du repository Artifact Registry"
  default     = "octoptimist-repo"
}

variable "tag" {
  type        = string
  description = "Tag de l'image Docker"
  default     = "latest"
}

variable "github_owner" {
  type        = string
  description = "Propriétaire du dépôt GitHub (e.g., your-org)"
}

variable "github_repo_name" {
  type        = string
  description = "Nom du dépôt GitHub (e.g., octoptimist)"
}