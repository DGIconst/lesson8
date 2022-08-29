package configuration

import "github.com/kelseyhightower/envconfig"

type Config struct {
	port: 8080
db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
jaeger_url: http://jaeger:16686
sentry_url: http://sentry:9000
kafka_broker: kafka:9092
some_app_id: testid
some_app_key: testkey
}

func GetConfig() Config{
	return Config
}