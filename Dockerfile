FROM golang:latest
ENV GO111MODULE="" \
    CGO_ENABLED="1" \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8080
WORKDIR /go/src/app
VOLUME ["/go/src/app"]
COPY . .
RUN go mod vendor
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build -o build/ttu-service"  --command=./build/ttu-service

