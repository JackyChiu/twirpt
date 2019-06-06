# twirpt

A little test repo to play around with Twirp and Kubernetes.

## Server
The server runs as 3 replicas on k8s.

### Start

```bash
make up     # starts minikube and creates k8s resources
make run    # applies the Service and Deployment onto the cluster
```

### Expected Output
```bash
kubectl get service
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP          96m
twirpt       NodePort    10.109.50.116   <none>        8080:30000/TCP   8m10s

kubectl get deploy
NAME     READY   UP-TO-DATE   AVAILABLE   AGE
twirpt   3/3     3            3           8m27s

kubectl get pods
NAME                     READY   STATUS    RESTARTS   AGE
twirpt-c445b6595-7j8g9   1/1     Running   1          8m43s
twirpt-c445b6595-b49fc   1/1     Running   1          8m43s
twirpt-c445b6595-pwbbl   1/1     Running   1          8m43s
```

### Stop
```
make down   # stops and deletes the minikube VM
```

## Client
The client runs as your local machine.
```bash
make run-client     # client makes request to server on k8s pod
```
