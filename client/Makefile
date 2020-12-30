export GO111MODULE=on
export GOBIN=$HOME/go/bin:$PATH

PROTO_DIR=car_service_client
PROTOC=protoc -I $(PROTO_DIR)

.PHONY: go_proto
go_proto:
	$(PROTOC) $(PROTO_DIR)/car_service.proto --go_out=plugins=grpc:$(PROTO_DIR) --gofullmethods_out=$(PROTO_DIR)