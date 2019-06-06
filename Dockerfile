FROM golang:1.12
ENV GO111MODULE=on
WORKDIR /go/src/github.com/JackyChiu/twirpt/
COPY . .
RUN go install -v ./cmd/...
EXPOSE 8080
CMD ["twirpt"]
