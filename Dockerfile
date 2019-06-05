FROM golang:latest
WORKDIR /go/src/github.com/JackyChiu/twirpt/
COPY . .
RUN go install -v ./cmd/...
EXPOSE 8080
CMD ["twirpt"]
