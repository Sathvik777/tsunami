# Tsunami
AiOps prometheus metric value estimator for time series data

## Concept

Extracting data (metrics) from Prometheus with timestamps. This data is time series, similar to stock market options. We are going to use a similar model to predict the direction of the metric based on time.

## Bolts and Pieces
* Data-extract
* Train
* Model management 
* Model service (host)


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
    * `helm install stable/prometheus-operator --name=prometheus-operator --namespace=monitoring -f values.yaml`


### Examples
For example application implementation check `/sample-http`
