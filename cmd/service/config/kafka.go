package config

// Kafka конфиги кафки
type Kafka struct {
	Brokers []string
}

var kafka Kafka

// GetKafkaConfig чтение конфигов кафки
func GetKafkaConfig() Kafka {
	return kafka
}

func init() {
	kafka.Brokers = []string{"localhost:9095"}
}
