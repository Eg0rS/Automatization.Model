package config

type Settings struct {
	Port       int    `json:"port"`
	Postgres   string `json:"postgres"`
	WriteTopic string `json:"kafka_topic"`
	KafkaUrl   string `json:"kafka_url"`
}
