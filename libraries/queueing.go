package libraries

import (
	"github.com/adjust/rmq"
	"github.com/jacky-htg/api-news/config"
)

var rmqConnection = rmq.OpenConnection("producer", config.GetString("database.redis.protocol"), config.GetString("database.redis.address"), config.GetInt("database.redis.db"))

func OpenQueue(key string) rmq.Queue {
	return rmqConnection.OpenQueue(key)
}
