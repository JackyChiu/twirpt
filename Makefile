KUBE_YAML = k8s.yaml
SERVICE = twirpt
DEPLOY = twirpt

GOPATH := $(GOPATH)

deps:
	go get -u ./...
	go mod vendor

proto:
	protoc --proto_path=$(GOPATH)/src:. --twirp_out=. --go_out=. rpc/**/*.proto

up:
	minikube start
	kubectl config use-context minikube
	kubectl create -f $(KUBE_YAML)

down:
	minikube stop

build:
	go build ./... # validate the Go code before attempting to build the image
	docker build -t $(DEPLOY) .

run:
	kubectl delete service $(SERVICE)
	kubectl delete deploy $(DEPLOY)
	kubectl apply -f $(KUBE_YAML)

client:
	go run cmd/twirpt_client/main.go

all: deps build run
