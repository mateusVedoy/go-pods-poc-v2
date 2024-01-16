## Rodando no cluster
Opcionalmente há comandos Makefile para executar passos abaixo. 

#### #1 Criar cluster
```
kind create cluster
```

#### #2 Criar docker images dos serviços

##### Requisitos: estar em *./service-one*
```
docker build -t svc-1:v1 -f inCluster.Dockerfile .
```

##### Requisitos: estar em *./service-two*
```
docker build -t svc-1:v1 .
```

#### #3 Carregar imagens para cluster
```
kind load docker-image svc-1:v1
kind load docker-image svc-2:v1
```

#### #4 Aplicar arquivos de config K8S

##### Requisitos: estar em *./service-two/k8s*

```
kubetcl apply -f deployment.yaml
kubectl apply -f service.yaml
```

##### Requisitos: estar em *./service-one/k8s*

```
kubetcl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f configs.yaml
```

#### #5 Expor porta do serviço para máquina

##### Requisitos: estar em *./service-two/k8s*

```
kubectl port-forward svc/service-one 8080:8080
```

##### Requisitos: estar em *./service-one/k8s*

```
kubectl port-forward svc/service-one 8081:8081
```