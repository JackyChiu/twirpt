KUBE_YAML = k8s.yaml
SERVICE = twirpt
DEPLOY = twirpt

proto:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. *.proto

up:
	minikube start
	kubectl config use-context minikube

down:
	minikube stop

build:
	docker build -t $(DEPLOY) .

run:
	kubectl delete service $(SERVICE)
	kubectl delete deploy $(DEPLOY)
	kubectl apply -f $(KUBE_YAML)

run-client:
	go run cmd/twirpt_client/main.go

all: build run
