package config

// Kafka конфиг Kafka
type Kafka struct {
	Brokers []string
	Topics  []string
	Group   string
}

var configKafka Kafka

// GetConfigKafka чтение конфига Kafka
func GetConfigKafka() Kafka {
	return configKafka
}

func init() {
	configKafka.Brokers = []string{"localhost:9095", "localhost:9096", "localhost:9097"}
	configKafka.Topics = []string{"income_add_request", "income_delete_request",
		"income_update_request", "income_get_request", "income_list_request"}
	configKafka.Group = "task_request_consumer"
}
