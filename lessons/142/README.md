- Fiber vs Express
- run it for 1hr





INstall this - https://istio.io/latest/docs/tasks/traffic-management/request-routing/

Traffic Shifting - https://istio.io/latest/docs/tasks/traffic-management/traffic-shifting/

As a security best practice, it is recommended to deploy the gateway in a different namespace from the control plane.


In this task you migrated traffic from an old to new version of the reviews service using Istio’s weighted routing feature. Note that this is very different than doing version migration using the deployment features of container orchestration platforms, which use instance scaling to manage the traffic.

With Istio, you can allow the two versions of the reviews service to scale up and down independently, without affecting the traffic distribution between them.

By default, Istio tracks the server workloads migrated to Istio proxies, and configures client proxies to send mutual TLS traffic to those workloads automatically, and to send plain text traffic to workloads without sidecars.
https://istio.io/latest/docs/tasks/security/authentication/authn-policy/#auto-mutual-tls















## Istiod
k apply -f /Users/antonputra/devel/tutorials/lessons/142/istio-crds.yaml
k create ns istio-system
k apply -f /Users/antonputra/devel/tutorials/lessons/142/istiod

export INGRESS_HOST=$(kubectl get gtw bookinfo-gateway -o jsonpath='{.status.addresses[*].value}')
export INGRESS_PORT=$(kubectl get gtw bookinfo-gateway -o jsonpath='{.spec.listeners[?(@.name=="http")].port}')
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT

 the Istio APIs (Gateway, VirtualService and DestinationRule will be depreciated
 Just like Kubernetes intends to support the Ingress API for many years after the Gateway API goes stable, the Istio APIs (Gateway, VirtualService and DestinationRule) will remain supported for the foreseeable future.


 Other parts of Istio functionality, including policy and telemetry, will continue to be configured using Istio-specific APIs while we work with SIG Network on standardization of these use cases.

  We also encourage early adopters to start experimenting with the Gateway API for mesh (service-to-service) use, and we will move that support to Beta when SIG Network has standardized the required semantics.


  Linkerd, Consul and Open Service Mesh
  https://gateway-api.sigs.k8s.io/contributing/gamma/


  Gateway API to configure internal mesh - https://istio.io/latest/docs/tasks/traffic-management/request-routing/












kubectl -n monitoring port-forward svc/prometheus-operated 9090

kubectl -n staging port-forward svc/service-a 8080
kubectl -n istio-system port-forward svc/kiali 20001

kubectl get httproute service-b -o yaml

hey -n 100000000 -c 1 -q 10 http://localhost:8080/api/devices









## Deployed
terraform apply