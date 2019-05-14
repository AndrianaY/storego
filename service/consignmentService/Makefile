build:
	protoc --proto_path=proto --proto_path=$(GOPATH)\pkg\mod\google.golang.org -I. --go_out=plugins=grpc:proto --grpc-gateway_out=logtostderr=true:proto --swagger_out=logtostderr=true:proto consignment.proto
#	protoc -I . -I proto --go_out=plugins=grpc:%GOPATH%/src/ crm/v1/crm.proto