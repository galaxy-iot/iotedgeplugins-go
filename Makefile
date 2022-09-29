PROTOBUFDIR=$(CURDIR)/protobuf
TARGETS := $(shell ls $(PROTOBUFDIR))

protobuf: $(TARGETS)

%.proto:
	protoc --go_out=. --go-grpc_out=. --proto_path=$(PROTOBUFDIR) $@

.PHONY: protobuf