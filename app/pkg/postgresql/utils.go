package postgresql

import "time"

type PgConfig struct {
	userName string
	password string
	host     string
	port     string
	database string
}

func NewPgConfig(userName, password, host, port, database string) *PgConfig {
	return &PgConfig{
		userName: userName,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

func doWithAttempts(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return err
}
