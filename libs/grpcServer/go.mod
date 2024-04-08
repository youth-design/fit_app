module github.com/youth-design/fit_app/lib/grpcServer

go 1.22

require (
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/youth-design/fit_app/libs/logger v0.0.0
	google.golang.org/grpc v1.63.1
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/youth-design/fit_app/libs/logger v0.0.0 => ../logger
