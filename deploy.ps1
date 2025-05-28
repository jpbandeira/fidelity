# How to run

# May be required to change the execution policty
# Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass

# .\deploy-local.ps1

# Create pass env var powershell
# $env:MY_POSTGRES_PASSWORD = "suaSenhaSeguraAqui"

$projectRoot = Split-Path -Parent $MyInvocation.MyCommand.Definition
Write-Host "Projeto raiz: $projectRoot"

$dateTime = Get-Date -Format "yyyyMMdd-HHmm"
$imageBase = "fidelity"
$imageTag = "${imageBase}:${dateTime}"

Write-Host "=== Buildando imagem Docker $imageTag ==="
docker build -t $imageTag $projectRoot
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao buildar imagem Docker"
    exit 1
}

$deploymentFile = "$projectRoot\deploy\base\deployment.yaml"
Write-Host "=== Atualizando tag da imagem no arquivo $deploymentFile ==="
(Get-Content $deploymentFile) -replace 'image: fidelity:.*', "image: $imageTag" | Set-Content $deploymentFile

Write-Host "=== Criando namespace fidelity ==="
kubectl apply -f "$projectRoot\deploy\base\namespace.yaml"
if ($LASTEXITCODE -ne 0) {
    Write-Warning "Falha ao criar namespace fidelity, verifique se já existe e se está acessível"
}

if (-not $env:MY_POSTGRES_PASSWORD) {
    Write-Error "Variável de ambiente MY_POSTGRES_PASSWORD não está definida. Abortando."
    exit 1
}

Write-Host "=== Criando Secret postgres-secret no namespace fidelity ==="
kubectl delete secret postgres-secret -n fidelity --ignore-not-found
kubectl create secret generic postgres-secret -n fidelity --from-literal=POSTGRES_PASSWORD=$env:MY_POSTGRES_PASSWORD
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao criar a Secret postgres-secret"
    exit 1
}

$tempDir = Join-Path $env:TEMP "fidelity-k8s"
if (Test-Path $tempDir) { Remove-Item $tempDir -Recurse -Force }
New-Item -ItemType Directory -Path $tempDir | Out-Null

Write-Host "=== Gerando manifests Kubernetes via kustomize ==="
kubectl kustomize "$projectRoot\deploy\base" > "$tempDir\full-deployment.yaml"
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao gerar manifests com kustomize"
    exit 1
}

Write-Host "=== Aplicando manifest no namespace fidelity ==="
kubectl apply -n fidelity -f "$tempDir\full-deployment.yaml"
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao aplicar manifest no Kubernetes"
    exit 1
}

Write-Host "=== Limpando imagens Docker antigas (mantendo 2 últimas) ==="
$images = docker images --format "{{.Repository}}:{{.Tag}}" | Where-Object { $_ -like "${imageBase}:*" } | Sort-Object
if ($images.Count -gt 2) {
    $imagesToDelete = $images | Select-Object -First ($images.Count - 2)
    foreach ($img in $imagesToDelete) {
        Write-Host "Removendo imagem antiga: $img"
        docker rmi $img
    }
} else {
    Write-Host "Nenhuma imagem antiga para remover."
}

Write-Host "=== Pods atuais no namespace fidelity ==="
kubectl get pods -n fidelity

Write-Host "=== Services atuais no namespace fidelity ==="
kubectl get svc -n fidelity

Write-Host "=== Deploy completo no namespace fidelity! ==="