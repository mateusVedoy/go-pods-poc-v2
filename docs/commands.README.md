## Comandos úteis

### Recuperando kubeconfig do cluster
```
<nome-cluster> get kubeconfig
```

### Criando cluster via Kind
```
kind create cluster --name <cluster-name>
```

### Deletando cluster via Kind
```
kind delete cluster --name <cluster-name>
```

### Gerando imagem local para pod
```
docker build -t <name>:<tag> /path/to/Dockerfile
```

### Importando docker container image para cluster
```
kind load docker-image <name>:<tag> --name <cluster-name>
```

### Se precisar limpar docker images e containers
```
docker rmi -f $(docker images -aq)
docker container prune
```

### Criando pod dentro do cluster
```
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

### Verificando pod
```
kubectl get po
```

### Verificando service
```
kubectl get svc
```

### deletando pod
```
kubectl delete pod <pod-name>

```

### Assistir pods em tempo de exec
```
watch 'kubectl get po'

```

### Acessando cluster (docker container)
```
docker exec -it <container-name> sh
```

### Acessando pod dentro do cluster
Primeiro executa passo acima

```
kubectl exec -it <pod-name> -- sh
```

### Bind de portas pod
```
kubectl port-forward svc/<service-name> <port-local>:<port-exposed-by-docker-image>
```

### Deletando imagens do cluster


### Verificando docker-images no cluster
```
docker exec -it <cluster-name>-control-plane crictl images
```
- As docker-images estarão no diretório *docker.io/library/{image}*

### Verificando namespaces do cluster
```
kubectl get namespaces
```

### Copiando kubeconfig para projeto
```
<cluster-name> get kubeconfig > kubeconfig.yaml
```

### Deletando regra de deployment dentro do cluster
```
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
```

### Vendo log dos pods
```
kubectl logs <pod-name>
```

### Rodando cluster.Dockerfile
```
docker run -t <image-name> -f cluster.Dockerfile .
```

### Dando permissões para kubeconfig.yaml
```
chmod o+r kubeconfig.yaml
```