package repositories

import (
	"encoding/json"
	"time"

	"github.com/pipusana/goapi/adapters"
)

type LoggerRepository interface {
	Log(route string) error
}

type Log struct {
	LogAdapter adapters.QueueAdapter
}

func NewLoggerRepository(logAdapter adapters.QueueAdapter) LoggerRepository {
	return &Log{
		LogAdapter: logAdapter,
	}
}

func (l *Log) Log(route string) error {
	t := time.Now()
	logItem := make(map[string]string)
	logItem["route"] = route
	logItem["count"] = "1"
	logItem["time"] = t.Format("2006-01-02 15:04:05")
	logMessage, err := json.Marshal(logItem)
	if err != nil {
		return err
	}
	return l.LogAdapter.PublishMessage(string(logMessage))
}
