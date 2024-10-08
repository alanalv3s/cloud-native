# Kind

kind create cluster --config config.yaml

## Load Balancer

```sh
go install sigs.k8s.io/cloud-provider-kind@latest

kubectl label node kind-control-plane node.kubernetes.io/exclude-from-external-load-balancers-

bin/cloud-provider-kind

kubectl apply -f loadbalancer.yaml

LB_IP=$(kubectl get svc/lb-test -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')

for \_ in {1..10}; do curl ${LB_IP}; done
```

## NGINX Ingress

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
