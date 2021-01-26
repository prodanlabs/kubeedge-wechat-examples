COMPONENTS?=kubeedge-wechat-app
BINARYNAME?=app
TAG?=latest

.PHONY : gen
gen:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative  --go-grpc_opt=paths=source_relative  stream/stream.proto

.PHONY : clean
clean:
	rm -f stream/*.go

.PHONY : build
build:
	CGO_ENABLED=1 GOOS=linux go build -o ${BINARYNAME} main.go
# 交叉编译
.PHONY : build-client-arm
build-client-arm:
	GOARCH=arm \
    GOOS="linux" \
    GOARM=6 \
    CGO_ENABLED=1 \
    CC=arm-linux-gnueabi-gcc \
    go build -o pi-player-app edgecore-client-pi/main.go
.PHONY : image
image:
	docker build . -t prodan/${COMPONENTS}:${TAG}

.PHONY : run
run:
	CGO_ENABLED=1 GOOS=linux go run main.go