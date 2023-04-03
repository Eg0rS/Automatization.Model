package config

type Settings struct {
	Postgres   string `json:"postgres"`
	KafkaUrl   string `json:"kafka_url"`
	ReadTopic  string `json:"read_topic"`
	WriteTopic string `json:"write_topic"`
}
