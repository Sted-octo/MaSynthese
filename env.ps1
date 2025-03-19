# Définition des variables d'environnement pour Terraform et gcloud

$env:TF_VAR_project_id           = "bsjxygz-gcp-octo-lille"
$env:TF_VAR_region               = "europe-west9"
$env:TF_VAR_service_name         = "octoptimist-service"
$env:TF_VAR_artifact_repo_name   = "octoptimist-repo"
$env:TF_VAR_tag                  = "latest"
$env:TF_VAR_github_owner         = "Sted-octo"
$env:TF_VAR_github_repo_name     = "MaSynthese"

# Affichage des variables pour vérification (facultatif)
Write-Host "Variables d'environnement définies :"
Write-Host "project_id          : $($env:TF_VAR_project_id)"
Write-Host "region              : $($env:TF_VAR_region)"
Write-Host "service_name        : $($env:TF_VAR_service_name)"
Write-Host "artifact_repo_name  : $($env:TF_VAR_artifact_repo_name)"
Write-Host "tag                 : $($env:TF_VAR_tag)"
Write-Host "github_owner        : $($env:TF_VAR_github_owner)"
Write-Host "github_repo_name    : $($env:TF_VAR_github_repo_name)"

# Vous pouvez maintenant exécuter vos commandes terraform et gcloud
# Exemple : terraform apply
#          gcloud builds triggers create ... (comme montré précédemment)
