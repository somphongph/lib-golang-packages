package xlogger

import (
	"encoding/json"
	"os"
	"time"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func AppInfo(appLog AppLog) {
	saveLog(appLog, AppLogLevelInfo)
}

func AppWarn(appLog AppLog) {
	saveLog(appLog, AppLogLevelWarn)
}

func AppError(appLog AppLog) {
	saveLog(appLog, AppLogLevelError)
}

func saveLog(appLog AppLog, level AppLogLevel) {
	var logRecord logRecord
	copier.Copy(&logRecord, &appLog)

	logRecord.Timestamp = time.Now().Format(time.RFC3339)
	logRecord.Level = level
	logRecord.TraceId = appLog.SpanId
	logRecord.ServiceName = os.Getenv("TGRENV_APP__NAME")
	logRecord.Environment = os.Getenv("TGRENV_APP__ENV")

	jsonLog, err := json.Marshal(logRecord)
	if err != nil {
		SysErrorf("%v", err)
	}

	zap.L().Info(string(jsonLog))
}
