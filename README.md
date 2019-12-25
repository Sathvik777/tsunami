# Tsunami
AiOps prometheus metric value estimator for time series data


### Requirements
```
* Go 
* Python
* Kubernetes
* Helm
* Prometheus
```

### Install Local Dependencies
 * [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
 * helm 
    * `curl -L https://git.io/get_helm.sh | bash`
    * helm init
 * Prometheus
    * `helm install prometheus-operator stable/prometheus-operator --namespace=monitoring`


### Examples
For example application implementation check `/sample-http`
