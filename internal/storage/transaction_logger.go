package storage

type EventByte byte

type Event struct {
	Sequence  uint64
	EventType EventByte
	Key       string
	Value     string
	Timestamp int64
}

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key string, value string)
	Err() <-chan error
	ReadEvents() (<-chan Event, <-chan error) //to read from file to replay the logs and recreate kv store
	Run()                                     //This will append to wal or any storage
}

const (
	EventPut    byte = 1
	EventDelete byte = 2
)
