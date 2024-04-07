generate_auth:
	protoc -I protos protos/auth/auth.proto --go_out=./protos/gen/go/ --go_opt=paths=source_relative --go-grpc_out=./protos/gen/go/ --go-grpc_opt=paths=source_relative
start_auth_dev:
	cd ./services/auth/;go run ./cmd/main/main.go ../../config/auth/config_dev.yaml