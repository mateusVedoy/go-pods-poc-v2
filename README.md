# Go Pods Poc
Este projeto é uma poc em Go que busca gerenciar pods de uma aplicação a partir de outra aplicação.

![poc](./docs/img/poc.png)

## Dependências
* Go versão 20
* [Kind](https://kind.sigs.k8s.io/)
* [Kubectl](https://kubernetes.io/pt-br/docs/tasks/tools/#kubectl)

## Formas de rodar
* Localmente
* Cluster

### Rodando localmente
[Doc Aqui](./docs/Local.README.md)

### Rodando no cluster
[Doc Aqui](./docs/cluster.README.md)

## Respostas

#### Sucesso

#### /pods/amount/update/{amount} 
```
Pods updated: {amount}
```

#### /health
```
Hey, I'm alive and running at local machine
```

```
Hey, I'm alive and running at Cluster
```

#### Erro
Dependerá de onde estourar erro, mas via de regra virá uma string contendo uma descrição do erro acompanhado de sua causa.
Ex.:
```
Error building config from flags. Reason: stat ./kubeconfig.yaml: no such file or directory
```

### Acompanhando os pods em tempo real

Requisitos: Rodar comando via terminal
```
watch 'kubectl get po'
```
Haverá algo similar à imagem abaixo:
![watch-pods](./docs/img/watch-pods.png)

### Lista de comandos úteis
[Doc Aqui](./docs/commands.README.md)