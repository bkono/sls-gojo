package justonce

import (
	"time"
)

type Entry struct {
	ID        string `dynamo:"id"`
	Msg       string `dynamo:"msg"`
	TTL       int `dynamo:"ttl"`
	CreatedAt time.Time `dynamo:"created_at"`
}
