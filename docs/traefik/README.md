## Traefik - Ingress Controller

Este guia descreve como instalar e configurar o Traefik como gateway api para Ingress em um cluster Kubernetes provisionado pelo Docker Desktop. O objetivo é rotear tráfego externo (via `localhost`) para serviços internos do cluster.

### Pré-requisitos
- **Docker Desktop** com Kubernetes habilitado:
  - Abra o Docker Desktop > **Settings > Kubernetes** > Marque **Enable Kubernetes** > **Apply & Restart**.
- **Helm** instalado:
  - Baixe em [helm.sh](https://helm.sh/docs/intro/install/) ou use um gerenciador de pacotes (ex.: `choco install kubernetes-helm` no Windows).

Seguindo o tutorial da Docuemntação oficial do Traefik:
### Instalação
Geração de certificado TLS para `websecure`:
```bash
# Gerar um certificado autoassinado válido para *.docker.localhost
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout tls.key -out tls.crt \
  -subj "/CN=*.docker.localhost"

# Crie o segredo TLS no namespace traefik
kubectl create secret tls local-selfsigned-tls \
  --cert=tls.crt --key=tls.key \
  --namespace traefik
```

Crie um `values.yaml` arquivo com o seguinte conteúdo:
```yaml
# Configurar portas de rede e pontos de entrada
# Os EntryPoints são os ouvintes de rede para o tráfego de entrada.
ports:
  # Define o ponto de entrada HTTP denominado 'web'
  web:
    port: 80
    nodePort: 30000
    # Instrui este ponto de entrada a redirecionar todo o tráfego para o ponto de entrada 'websecure'
    redirections:
      entryPoint:
        to: websecure
        scheme: https
        permanent: true

  # Define o ponto de entrada HTTPS denominado 'websecure'
  websecure:
    port: 443
    nodePort: 30001

# Habilita o painel no Modo Seguro
api:
  dashboard: true
  insecure: false

ingressRoute:
  dashboard:
    enabled: true
    matchRule: Host(`dashboard.docker.localhost`)
    entryPoints:
      - websecure
    middlewares:
      - name: dashboard-auth

# Cria um Middleware BasiAuth e um Segredo para a Segurança do Painel
extraObjects:
  - apiVersion: v1
    kind: Secret
    metadata:
      name: dashboard-auth-secret
    type: kubernetes.io/basic-auth
    stringData:
      username: admin
      password: "P@ssw0rd"
  - apiVersion: traefik.io/v1alpha1
    kind: Middleware
    metadata:
      name: dashboard-auth
    spec:
      basicAuth:
        secret: dashboard-auth-secret

# Em vez disso, faremos o roteamento com a API Gateway.
ingressClass:
  enabled: false

# Habilita o Provedor de API do Gateway e desabilita o provedor KubernetesIngress
# Os provedores informam ao Traefik onde encontrar a configuração de roteamento.
providers:
  kubernetesIngress:
     enabled: false
  kubernetesGateway:
     enabled: true

## Gateway Listeners
gateway:
  listeners:
    web: # Ouvinte HTTP que corresponde ao entryPoint `web`
      port: 80
      protocol: HTTP
      namespacePolicy: All

    websecure:         # Ouvinte HTTPS que corresponde ao entryPoint `websecure`
      port: 443
      protocol: HTTPS  # TLS termina dentro do Traefik
      namespacePolicy: All
      mode: Terminate
      certificateRefs:    
        - kind: Secret
          name: local-selfsigned-tls  # O segredo que criamos antes da instalação
          group: ""

# Habilita Observabilidade
logs:
  general:
    level: INFO
  # Isso habilita os logs de acesso, enviando-os para a saída padrão do Traefik por padrão. A [Documentação de Logs de Acesso](https://doc.traefik.io/traefik/observability/access-logs/) aborda formatação, filtragem e opções de saída.
  access:
    enabled: true

# Habilita métricas para Prometheus
metrics:
  prometheus:
    enabled: true
```

Agora instale a aplicação usando o Helm e os valores pré-definidos no arquivo acima:
```bash
helm install traefik traefik/traefik \
  --namespace traefik \
  --values values.yaml
```

### Configuração de Roteamento:
Configure o encaminhamento de tráfego:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: whoami
  namespace: traefik
spec:
  parentRefs:
    - name: traefik-gateway # Nome do Gateway que o Traefik cria quando você habilita o provedor da API do Gateway
  hostnames:
    - "whoami.docker.localhost"
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /
      backendRefs:
        - name: whoami
          port: 80
```

Implantando HttpRoute
```bash
kubectl apply -f httproute.yaml
```
### Observações
- Nesse exemplo, implante o Gateway no **mesmo namespace** dos serviços que serão expostos
- Configure o Traefik para **Load Balancer**
- Use `ClusterIP` para os serviços de backend

Documentação
- [Traefik](https://doc.traefik.io/traefik/getting-started/kubernetes/)
- [HttpRoute](https://gateway-api.sigs.k8s.io/api-types/httproute/)
- [Gateway API](https://gateway-api.sigs.k8s.io/)