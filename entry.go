package justonce

type Entry struct {
	ID        string `dynamo:"id"`
	Msg       []byte `dynamo:"msg"`
	CreatedAt int64  `dynamo:"created_at"`
	ExpiresAt int64  `dynamo:"expires_at"`
}
