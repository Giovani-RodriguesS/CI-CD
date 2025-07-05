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

Crie um `traefik-values.yaml` arquivo com o seguinte conteúdo:
```yaml
# Configure Network Ports and EntryPoints
# EntryPoints are the network listeners for incoming traffic.
ports:
  # Defines the HTTP entry point named 'web'
  web:
    port: 8087
    nodePort: 30000
    # Instructs this entry point to redirect all traffic to the 'websecure' entry point
    redirections:
      entryPoint:
        to: websecure
        scheme: https
        permanent: true

  # Defines the HTTPS entry point named 'websecure'
  websecure:
    port: 8443
    nodePort: 30001

# Enables the dashboard in Secure Mode
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
 
# Creates a BasiAuth Middleware and Secret for the Dashboard Security
extraObjects:
  - apiVersion: v1
    kind: Secret
    metadata:
      name: dashboard-auth-secret
    type: kubernetes.io/basic-auth
    stringData:
      username: admin
      password: "123456789"      # Replace with an Actual Password
  - apiVersion: traefik.io/v1alpha1
    kind: Middleware
    metadata:
      name: dashboard-auth
    spec:
      basicAuth:
        secret: dashboard-auth-secret

# We will route with Gateway API instead.
ingressClass:
  enabled: false

# Enable Gateway API Provider & Disables the KubernetesIngress provider
# Providers tell Traefik where to find routing configuration.
providers:
  kubernetesIngress:
     enabled: false
  kubernetesGateway:
     enabled: true

## Gateway Listeners
gateway:
  listeners:
    web:           # HTTP listener that matches entryPoint `web`
      port: 8087
      protocol: HTTP
      namespacePolicy: All

    websecure:         # HTTPS listener that matches entryPoint `websecure`
      port: 8443
      protocol: HTTPS  # TLS terminates inside Traefik
      namespacePolicy: All
      mode: Terminate
      certificateRefs:    
        - kind: Secret
          name: local-selfsigned-tls  # the Secret we created before the installation
          group: ""

# Enable Observability
logs:
  general:
    level: INFO
  # This enables access logs, outputting them to Traefik's standard output by default. The [Access Logs Documentation](https://doc.traefik.io/traefik/observability/access-logs/) covers formatting, filtering, and output options.
  access:
    enabled: true

# Enables Prometheus for Metrics
metrics:
  prometheus:
    enabled: true 
```

Agora instale a aplicação usando o Helm e os valores pré-definidos no arquivo acima:
```bash
helm install traefik traefik/traefik \
  --namespace traefik \
  --values traefik-values.yaml
```

### Configuração de Roteamento:
Configure o encaminhamento de tráfego:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: go-api-service
  namespace: dev # Namespace da aplicação
spec:
  parentRefs:
  - name: traefik-gateway
    namespace: traefik
  hostnames:
    - "go.docker.localhost"
  # Regras de roteamento
  rules:
    - matches:
      - path:
          type: PathPrefix
          value: /products
    - backendRefs:
      - name: go-api-service # Nome do serviço de backend
        port: 8080
```

Implantando HttpRoute
```bash
kubectl apply -f route-go.yaml
```
### Observações
- Nesse exemplo, implante o Gateway no **mesmo namespace** dos serviços que serão expostos
- Configure o Traefik para **Load Balancer**
- Use `ClusterIP` para os serviços de backend

Documentação
- [Traefik](https://doc.traefik.io/traefik/getting-started/kubernetes/)
- [HttpRoute](https://gateway-api.sigs.k8s.io/api-types/httproute/)
- [Gateway API](https://gateway-api.sigs.k8s.io/)