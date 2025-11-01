package logger

import (
	"go.uber.org/zap"
)

// NewLogger cria e retorna um logger configurado com zap
func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}
