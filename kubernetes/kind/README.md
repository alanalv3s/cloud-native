# Kind

kind create cluster --config config.yaml

## Load Balancer

```sh
go install sigs.k8s.io/cloud-provider-kind@latest

kubectl label node kind-control-plane node.kubernetes.io/exclude-from-external-load-balancers-

sudo cloud-provider-kind

kubectl apply -f loadbalancer.yaml

LB_IP=$(kubectl get svc/lb-test -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')

for \_ in {1..10}; do curl ${LB_IP}; done
```

## NGINX Ingress

kubectl apply -f ingress-nginx.yaml

# Metrics Server

helm upgrade --install metrics-server metrics-server/metrics-server -f metrics-server-values.yaml

# Cloud Native PG

helm repo add cnpg https://cloudnative-pg.github.io/charts
helm upgrade --install cnpg --namespace cnpg-system --create-namespace cnpg/cloudnative-pg
