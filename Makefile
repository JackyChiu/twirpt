KUBE_YAML = k8s.yaml
SERVICE = twirpt
DEPLOY = twirpt

GOPATH := $(GOPATH)

deps:
	$(eval export GO111MODULE=on)
	go get -u ./...
	go mod vendor

proto:
	protoc --proto_path=$(GOPATH)/src:. --twirp_out=. --go_out=. rpc/**/*.proto

up: deps build
	minikube start
	kubectl config use-context minikube
	kubectl apply -f $(KUBE_YAML)

down:
	minikube stop

build:
	$(eval $(minikube docker-env))
	go build ./... # validate the Go code before attempting to build the image
	docker build -t $(DEPLOY) .

run: build
	kubectl delete service $(SERVICE)
	kubectl delete deploy $(DEPLOY)
	kubectl apply -f $(KUBE_YAML)

client:
	go run cmd/twirpt_client/main.go
