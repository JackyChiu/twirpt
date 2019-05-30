FROM golang:latest
WORKDIR /go/src/github.com/JackyChiu/twirpt/cmd/twirpt
COPY . .
RUN go get
RUN go install -v .
EXPOSE 8080
CMD ["twirpt"]
