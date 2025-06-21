package logger

import "go.uber.org/zap"

type ILoggerService interface {
	Init()
	Log(message string)
}

type LoggerService struct {
	logger *zap.SugaredLogger
}

func (ls *LoggerService) Init() {
	logger, _ := zap.NewDevelopment()
	ls.logger = logger.Sugar()
}

func (ls *LoggerService) Log(message string) {
	ls.logger.Info(message)
}

func NewLoggerService() *LoggerService {
	ls := &LoggerService{}
	ls.Init()
	return ls
}
