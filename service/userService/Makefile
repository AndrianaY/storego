build:
	protoc --proto_path=proto --proto_path=thirdParty -I. --go_out=plugins=grpc:proto --grpc-gateway_out=logtostderr=true:proto --swagger_out=logtostderr=true:proto proto/user.proto