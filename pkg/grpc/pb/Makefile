#Linux/wsl 專用
GOPATH := $(shell echo "$(GOPATH)")

all: stub gateway swagger


stub:
	protoc -I/usr/local/include -I. \
                  -I$(GOPATH)/src \
                  -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.10.0/third_party/googleapis \
                  --go_out=plugins=grpc:. \
                  template.proto
gateway:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
      -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.10.0/third_party/googleapis \
      --grpc-gateway_out=logtostderr=true:. \
      template.proto

swagger:
	protoc -I/usr/local/include -I. \
      -I$(GOPATH)/src \
      -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.10.0/third_party/googleapis \
      --swagger_out=logtostderr=true:. \
      template.proto