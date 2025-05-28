# ===============================================
# Script de deploy do PostgreSQL no namespace 'database'
# ===============================================

# Verificar variável de ambiente para senha do banco
if (-not $env:MY_POSTGRES_PASSWORD) {
    Write-Error "Variável de ambiente MY_POSTGRES_PASSWORD não está definida. Abortando."
    exit 1
}

# Criar o namespace 'database' se não existir
Write-Host "Verificando namespace 'database'..."
if (-not (kubectl get namespace database -o json 2>$null)) {
    Write-Host "Criando namespace 'database'..."
    kubectl create namespace database
} else {
    Write-Host "Namespace 'database' já existe."
}

# Criar Secret postgres-secret no namespace database
kubectl delete secret postgres-secret -n database --ignore-not-found
kubectl create secret generic postgres-secret `
    -n database `
    --from-literal=password=$env:MY_POSTGRES_PASSWORD
if ($LASTEXITCODE -ne 0) {
    Write-Error "Erro ao criar a Secret postgres-secret"
    exit 1
}

# Aplicar PVC
Write-Host "Aplicando postgres-pvc.yaml..."
kubectl apply -f .\postgres-pvc.yaml -n database

# Aplicar Service
Write-Host "Aplicando postgres-service.yaml..."
kubectl apply -f .\postgres-service.yaml -n database

# Aplicar Deployment
Write-Host "Aplicando postgres-deployment.yaml..."
kubectl apply -f .\
