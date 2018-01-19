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

build:
	go build -o agent agent.go common.go  main.pb.go  main.pb.gw.go main.validator.pb.go
	go build -o agent_tester agent_tester.go
	go build -o worker worker.go common.go main.pb.go

serve_agent:
	make build
	./agent -c agent.json

test_agent:
	make build
	./agent_tester -c agent.json

serve_mosca:
	node mosca.js -c server1.json &
	node mosca.js -c server2.json &

stop_mosca:
	ps aux  |  grep -i '[n]ode mosca.js'  |  awk '{print $$2}'  |  xargs kill -9

serve_worker:
	make build
	./worker -c worker.json
