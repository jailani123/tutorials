- Fiber vs Express
- run it for 1hr





INstall this - https://istio.io/latest/docs/tasks/traffic-management/request-routing/















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