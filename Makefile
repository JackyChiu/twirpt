KUBE_YAML = k8s.yaml
DEPLOY = twirpt

proto:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. *.proto

up:
	minikube start
	kubectl config use-context minikube
	kubectl create -f $(KUBE_YAML)

down:
	minikube stop

build:
	docker build -t $(DEPLOY) .

run:
	kubectl delete deploy $(DEPLOY)
	kubectl apply -f $(KUBE_YAML)

all: build run
