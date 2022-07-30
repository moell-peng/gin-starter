package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func InitLogger() {
	path := "./logs/error.log"

	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),                                  // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),        // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second), // 日志切割时间间隔
	)

	logrus.SetOutput(io.MultiWriter(os.Stdout, writer))
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

}