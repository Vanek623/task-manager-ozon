package server

// GRPC and HTTP config
const (
	address = "localhost"

	portGRPC    = ":8081"
	addressGRPC = address + portGRPC

	portHTTP    = ":8080"
	addressHTTP = address + portHTTP

	connectionType = "tcp"
)
