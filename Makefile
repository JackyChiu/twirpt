proto:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. *.proto

up:
	minikube start
	kubectl config use-context minikube

down:
	minikube stop

run:
	skaffold dev

manual-build:
	docker build -t twirpt .

manual-run:
	kubectl delete deploy twirpt
	kubectl apply -f k8s.yaml

all: manual-build manual-run
