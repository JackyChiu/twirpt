proto:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. *.proto

up:
	minikube start
	kubectl config use-context minikube

down:
	minikube stop

run:
	skaffold dev

