# Parar se ocorrer erro
$ErrorActionPreference = "Stop"

$namespace = "fidelity"
$imageName = "fidelity:local"

Write-Host "🚀 Iniciando undeploy total do namespace $namespace ..."

# 1. Deletar namespace fidelity
Write-Host "⏳ Deletando namespace $namespace..."
kubectl delete namespace $namespace --wait

# 2. Deletar PVCs restantes no namespace (se existir)
Write-Host "⏳ Deletando PVCs restantes no namespace $namespace (se houver)..."
try {
    kubectl delete pvc --all -n $namespace
} catch {
    Write-Host "Nenhum PVC encontrado para deletar."
}

# 3. Deletar imagem docker local fidelity:local
Write-Host "⏳ Deletando imagem Docker local $imageName..."
try {
    docker image rm $imageName -f
} catch {
    Write-Host "Imagem $imageName não encontrada localmente."
}

Write-Host "✅ Undeploy total concluído!"
