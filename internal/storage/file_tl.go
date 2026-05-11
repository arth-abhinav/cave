package storage

import (
	"os"
)

type FileTransactionLogger struct {
	events       chan Event
	errors       chan error
	lastSequence uint64
	file         *os.File
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{Key: key, EventType: EventByte(EventDelete)}
}

func (l *FileTransactionLogger) WritePut(key string, value string) {
	l.events <- Event{Key: key, EventType: EventByte(EventPut)}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}
