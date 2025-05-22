# Configuração Kubernetes - Kustomize

## Base

### Objetivo: 
Armazenar configurações **base** para ser usada no cluster K8S, como a configuração base da API que será implantada em diferentes ambientes.

### Detalhes:
- **database**: armazena a configuração do **service.yaml**, **deployment.yaml** e **configmap.yaml** para o banco de dados que a API acessará, além de conter o **kustomizition.yaml**, que referencia a base para o **overlays**.

- **app**: configuração do **service.yaml**, **deployment.yaml** e **configmap.yaml** para a API, além de conter o **kustomizition.yaml**, que referencia a base para o **overlays**.

## Overlays

### Objetivo: 
Armazenar configurações **customizadas** que será aplicada ao cluster K8S. Aqui definiremos os **ambientes de implantação**, como **dev** e **stg**.

### Detalhes:
- **database**: armazena a configuração para definir **namespace** e outras configurações personalizadas, como variáveis de ambiente.

- **dev**: contém a configuração para definir **namespace** e outras definições personalizadas, como variáveis de ambiente.

