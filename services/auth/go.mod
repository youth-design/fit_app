module github.com/youth-design/fit_app/services/auth

go 1.22

replace (
	github.com/youth-design/fit_app/lib/grpcServer v0.0.0 => ../../libs/grpcServer
	github.com/youth-design/fit_app/libs/logger v0.0.0 => ../../libs/logger
	github.com/youth-design/fit_app/protos v0.0.0 => ../../protos
)

require (
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/youth-design/fit_app/libs/logger v0.0.0
	github.com/youth-design/fit_app/protos v0.0.0
	google.golang.org/grpc v1.63.1
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	github.com/youth-design/fit_app/lib/grpcServer v0.0.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
