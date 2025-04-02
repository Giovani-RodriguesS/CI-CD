# Traefik 

### Permissões
Para usar a API do Kubernetes, o Traefik precisa de algumas permissões. Esse mecanismo de permissão é baseado em funções definidas pelo administrador do cluster.

A função é então vinculada a uma conta usada por um aplicativo, nesse caso, o Traefik Proxy.

O primeiro passo é criar a função. O ClusterRole enumera os recursos e ações disponíveis para a função

### Vinculação de permissões 
O próximo passo é criar uma conta de serviço dedicada para o Traefik. E então, vincule a função na conta para aplicar as permissões e regras na última.


### Implantação do Traefik
O `Ingress Controller` é um software que roda da mesma forma que qualquer outro aplicativo em um cluster.

Para iniciar o Traefik no cluster do Kubernetes, um Deployment deve existir para descrever como configurar e escalar contêineres horizontalmente para suportar cargas de trabalho maiores.

A implantação contém um atributo importante para personalizar o Traefik: **args**.

Esses argumentos são a **configuração estática do Traefik**.
A partir daqui, é possível habilitar o painel, configurar pontos de entrada, selecionar provedores de configuração dinâmica e muito mais 

### Observações
- Implante o Ingress no **mesmo namespace** dos serviços que serão expostos
- Configure o Traefik para **Load Balancer**

### Links
- [Docs](https://doc.traefik.io/traefik/)