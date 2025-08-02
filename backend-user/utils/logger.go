package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

// InitLogger inicializa o logger da aplicação
func InitLogger() {
	Logger = logrus.New()

	// Configurar formato JSON para produção
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	// Configurar nível de log baseado na variável de ambiente
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

	// Configurar output (arquivo ou stdout)
	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			Logger.SetOutput(file)
		}
	}

	// Adicionar campos padrão
	Logger.AddHook(&DefaultFieldsHook{})
}

// DefaultFieldsHook adiciona campos padrão aos logs
type DefaultFieldsHook struct{}

func (h *DefaultFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldsHook) Fire(entry *logrus.Entry) error {
	entry.Data["service"] = "backend-user"
	entry.Data["version"] = "1.0.0"
	return nil
}

// LogRequest loga informações da requisição HTTP
func LogRequest(method, path, ip string, statusCode int, duration float64) {
	Logger.WithFields(logrus.Fields{
		"method":   method,
		"path":     path,
		"ip":       ip,
		"status":   statusCode,
		"duration": duration,
		"type":     "http_request",
	}).Info("HTTP Request")
}

// LogError loga erros da aplicação
func LogError(err error, context map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"error":   err.Error(),
		"type":    "application_error",
		"context": context,
	}).Error("Application Error")
}

// LogInfo loga informações gerais
func LogInfo(message string, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"type":   "info",
		"fields": fields,
	}).Info(message)
}

// LogDebug loga informações de debug
func LogDebug(message string, fields map[string]interface{}) {
	Logger.WithFields(logrus.Fields{
		"type":   "debug",
		"fields": fields,
	}).Debug(message)
}
