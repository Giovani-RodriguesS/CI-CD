# CI/CD - Repositório de Estudos

## Visão Geral

Este projeto tem como objetivo aprofundar meus estudos CI/CD de maneira prática implementando uma pipeline **CI/CD** utilizando de ferramentas como GitHub Actions, um cluster Kubernetes local e o ArgoCD para **GitOps**.

## Descrição
Desenvolvi uma simples aplicação Go que será implantada no cluster K8s local por meio do CI/CD e usará um banco SQL apenas para armazenar os dados. Primeiro, ao fazer push no repositório*, a pipeline inicia fazendo testes no código e fazendo o *build* da imagem. Em seguida configura credenciais para enviar ao Docker Hub (meu Registry nesse exemplo), e envia a imagem. 

No mesmo repositório, criei uma pasta para armazenar os arquivos de configuração do Kubernetes, usando Kustomize para personalizar a configuração de diferentes ambientes, como **dev** e **stg**. Quando o ArgoCd estiver criado e configurado no cluster, ele escutará as alterações nos manifestos **application** e implantará a nova definição do cluster. Assim, implementei a técnica **GitOps** para gerenciamento do Kubernetes.

Além disso, busquei implementar um **Proxy Reverso** como ponto de entrada e para lidar com balanceamento de carga entre instâncias.

---

**Obs**: Não é uma boa prática armazenar o código da aplicação no mesmo repositório de manifestos Kubernetes, mas para fins de estudo, decidi colocar ambos juntos.

---

## Estrutura do Projeto

- **app/**: Código da API Go.
  - **api/**: Contém `Dockerfile`, `go.mod`, `go.sum`, código principal (`cmd/main.go`), e pacotes internos (`controller/`, `initializers/`, `migrate/`, `model/`, `repository/`, `router/`, `usecase/`, `utils/`).
  - **readme.md**: Documentação da aplicação.
- **application.yaml**: Configuração do ArgoCD para implantação.
- **config/**: Configurações de ambiente.
  - **database/**: Configurações do banco de dados (`application.yaml`).
  - **dev/**: Configurações do ambiente de desenvolvimento (`application.yaml`).
- **ingress/**: Configurações do Ingress.
  - **ingresses.yaml**: Definições do Ingress no K8s.
  - **traefik/**: Configurações do Traefik (`account.yaml`, `kustomization.yaml`, `role-binding.yaml`, `role.yaml`, `traefik-services.yaml`, `traefik.yaml`).
- **k8s/**: Manifestos Kubernetes.
  - **base/**: Configurações base para aplicação (`app/`) e banco de dados (`database/`), com `configmap.yaml`, `deployment.yaml`, `service.yaml`, `kustomization.yaml`, `netPolicy.yaml`, `secret.yaml`.
  - **overlays/**: Sobreposições para ambientes (`database/`, `dev/`, `stg/`), com `kustomization.yaml`, `namespace.yaml`, e patches.
- **tests/**: Configurações de testes.
  - **sonarqube/**: Integração com SonarQube (`docker-compose.yaml`, `sonar-scanner.sh`).

## Pré-requisitos

- Kubernetes (Kind)
- ArgoCD
- Go (versão 1.23)
- Banco de dados relacional (PostgreSQL)
- SonarQube
- Traefik
- Docker DeskTop

## Instalação

1. Clone o repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   ```
2. Configure o cluster Kubernetes local.
- Instale o ArgoCD:
   ```bash
   kubectl create namespace argocd
   kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml -n argocd
   ```
- Obtenha o valor da senha
  ```bash
  kubectl get secret argocd-initial-admin-secret -n argocd -o yaml 
  ```
- Copie o valor retornado na saída do comando: 
  ```yaml
   data:
    password: <SENHA_BASE64>
  ```

- Decodifique:
  ```bash
  echo <SENHA_BASE64> | base64 --decode
  ```
- Com a senha, rode o comando que encaminha requisição para a porta do serviço ArgoCD:
  ```bash
  kubectl port-forward svc/argocd-server 8000:80 -n argocd
  # acesse localhost:8000
  ```

- Cole o valor com o seguite login:
  ```textplain
  username: admin
  password: <SENHA_DECODIFICADA>
  ```

3. Agora, aplique a aplicação para que os ambientes de **database** e **dev** estejam sincronizados com o repositório remoto
- Execute o seguinte:
  ```bash
  kubectl apply -f application.yaml
  ```

4. Configure o tarefik em [docs/traefik/](docs\traefik\README.md)

## Licença
MIT