# twirpt

Twirp API On Kubernetes Test

## Server
```bash
make up     # starts minikube and sets the context
make build  # builds the docker image
make run    # applies the Service and Deployment onto the cluster
```

## Client
```bash
make run-client     # client makes request to server on k8s pod
```
