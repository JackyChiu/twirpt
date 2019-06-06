# twirpt

A little test repo to play around with Twirp and Kubernetes.

## Server
The server runs as 3 replicas on k8s.
```bash
make up     # starts minikube and creates k8s resources
make run    # applies the Service and Deployment onto the cluster
```

## Client
The client runs as your local machine.
```bash
make run-client     # client makes request to server on k8s pod
```
