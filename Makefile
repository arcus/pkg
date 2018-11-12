TEST ?= ./...

GO111MODULE ?= on

default: test

proto:
	cd service/internal/pb && \
		rm -f *.pb.go && \
		protoc --go_out=plugins=grpc:. *.proto
	cd transport/http/internal/test && \
		rm -f *.pb.go && \
		protoc --go_out=plugins=grpc:. *.proto

test:
	go test -v $(TEST)
