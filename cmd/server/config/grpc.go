package config

// ServiceGRPC конфиги GRPC сервера сервиса
type ServiceGRPC struct {
	Host, ConnectionType string
}

// StorageGRPC конфиги GRPC клиента хранилища
type StorageGRPC struct {
	Host string
}

var serviceGRPC ServiceGRPC
var storageGRPC StorageGRPC

// GetServiceGRPCConfig чтение конфигов GRPC сервера сервиса
func GetServiceGRPCConfig() ServiceGRPC {
	return serviceGRPC
}

// GetStorageGRPCConfig чтение конфигов GRPC клиента хранилища
func GetStorageGRPCConfig() StorageGRPC {
	return storageGRPC
}

func init() {
	serviceGRPC.Host = "localhost:8081"
	serviceGRPC.ConnectionType = "tcp"

	storageGRPC.Host = "localhost:8082"
}
