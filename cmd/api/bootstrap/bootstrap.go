package bootstrap

import "github.com/EdwBaeza/inhouse/internal/platform/server"

const (
	host = "0.0.0.0" // For binding to all interfaces.
	port = 3000
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
