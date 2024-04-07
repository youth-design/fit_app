module github.com/youth-design/fit_app/services/auth

go 1.22

replace (
	github.com/youth-design/fit_app/libs/logger v0.0.0 => ../../libs/logger
	github.com/youth-design/fit_app/protos v.0.0.0 => ../../protos
)

require (
	github.com/youth-design/fit_app/libs/logger v0.0.0
	github.com/youth-design/fit_app/proto v0.0.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)
