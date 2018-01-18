proto:
	rm -f main.pb.go  main.pb.gw.go main.validator.pb.go
	protoc \
		-I/usr/local/include \
		-I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		--go_out=plugins=grpc:. \
		--proto_path=${GOPATH}/src \
		--proto_path=${GOPATH}/src/github.com/google/protobuf/src \
		--proto_path=. \
		--govalidators_out=. \
		main.proto

serve:
	go run main.go  main.pb.go  main.pb.gw.go main.validator.pb.go

test:
	go run test.go