package entities

import "time"

type Config struct {
	Port        string `mapstructure:"port"`
	MongoURI    string `mapstructure:"mongo_uri"`
	RabbitMQURI string `mapstructure:"rabbitmq_uri"`
	QueueName   string `mapstructure:"queue_name"`
	DATABASE    string `mapstructure:"database"`
	COLLECTION  string `mapstructure:"collection"`
}

type Nisit struct {
	ID        string `json:"id"`
	FristName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type NisitUpdate struct {
	FristName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type Log struct {
	Route string    `json:"route"`
	Time  time.Time `josn:"time"`
}
