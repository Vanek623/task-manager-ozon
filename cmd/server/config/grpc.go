package config

type ServiceGRPC struct {
	Host, ConnectionType string
}

type StorageGRPC struct {
	Host string
}

var serviceGRPC ServiceGRPC
var storageGRPC StorageGRPC

func GetServiceGRPCConfig() ServiceGRPC {
	return serviceGRPC
}

func GetStorageGRPCConfig() StorageGRPC {
	return storageGRPC
}

func init() {
	serviceGRPC.Host = "localhost:8081"
	serviceGRPC.ConnectionType = "tcp"

	storageGRPC.Host = "localhost:8082"
}
