package config

// GRPC конфиг GRPC
type GRPC struct {
	Host, ConnectionType string
}

var configGRPC GRPC

// GetConfigGRPC чтение конфига GRPC
func GetConfigGRPC() GRPC {
	return configGRPC
}

func init() {
	configGRPC.Host = "localhost:8082"
	configGRPC.ConnectionType = "tcp"
}
