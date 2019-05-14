protoc --proto_path=proto --proto_path=thirdParty -I. --go_out=plugins=grpc:proto consignment.proto
protoc --proto_path=proto --proto_path=thirdParty -I. -I. -I. --grpc-gateway_out=logtostderr=true:proto consignment.proto
protoc --proto_path=proto --proto_path=thirdParty --swagger_out=logtostderr=true:proto consignment.proto