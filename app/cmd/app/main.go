package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/kurochkinivan/Note_taker/internal/config"
	"github.com/kurochkinivan/Note_taker/pkg/postgresql"
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
	cfg := config.GetConfig()

	cfgPSQL := cfg.PostgreSQL
	pgcfg := postgresql.NewPgConfig(cfgPSQL.Username, cfgPSQL.Password, cfgPSQL.Host, cfgPSQL.Port, cfgPSQL.Database)

	client, err := postgresql.NewClient(context.TODO(), 5, pgcfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(client)

	logrus.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
