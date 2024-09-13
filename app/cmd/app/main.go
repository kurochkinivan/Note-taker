package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/kurochkinivan/Note-taker/internal/app"
	"github.com/kurochkinivan/Note-taker/internal/config"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	logrus.SetOutput(os.Stdout)
}

func main() {
	logrus.Info("getting config")
	cfg := config.GetConfig()

	app, err := app.NewApp(cfg)
	if err != nil {
		logrus.Fatalf("failed to start app, err: %v", err)
	}

	err = app.Run(context.Background())
	if err != nil {
		logrus.Fatalf("server trouble, err: %v", err)
	}
}
