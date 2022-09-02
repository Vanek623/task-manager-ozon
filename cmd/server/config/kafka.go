package config

type Kafka struct {
	Brokers []string
}

var kafka Kafka

func GetKafkaConfig() Kafka {
	return kafka
}

func init() {
	kafka.Brokers = []string{"localhost:9095"}
}
