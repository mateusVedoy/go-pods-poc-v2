## Rodando localmente
Opcionalmente há comandos Makefile para executar passos abaixo. 

#### #1 Criar cluster
```
kind create cluster
```

#### #2 Criar kubeconfig
```
kind get kubeconfig > ./service-one/kubeconfig.yaml
```

#### #3 Dando permissões ao kubeconfig.yaml
``` 
chmod o+r kubeconfig.yaml
```

#### #4 Criar docker images dos serviços

##### Requisitos: estar em *./service-one*
```
docker build -t svc-1:v1 -f outCluster.Dockerfile .
```

##### Requisitos: estar em *./service-two*
```
docker build -t svc-1:v1 .
```

#### #5 Carregar imagens para cluster
```
kind load docker-image svc-1:v1
kind load docker-image svc-2:v1
```

#### #6 Aplicar arquivos de config K8S

##### Requisitos: estar em *./service-two/k8s*

```
kubetcl apply -f deployment.yaml
kubectl apply -f service.yaml
```

#### #7 Expor porta do serviço para máquina

##### Requisitos: estar em *./service-two/k8s*

```
kubectl port-forward svc/service-one 8080:8080
```
## Testando locamente
Para testar a poc é preciso executar comandos REST.

### Subindo service-one
```
docker run --name svc-one --network host svc-1:v1
```

### Endpoints

#### /health
```
http://localhost:8081/svc-one/health
```

#### /pods/amount/update/{amount} 
```
http://localhost:8081/svc-two/pods/amount/update/{amount}
```
