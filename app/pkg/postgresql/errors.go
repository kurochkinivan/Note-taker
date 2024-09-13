package postgresql

import (
	"github.com/pkg/errors"
)

func ErrDoQuery(err error) error {
	return errors.Wrap(err, "failed to do query")
}

func ErrCreateQuery(err error) error {
	return errors.Wrap(err, "failed to create query")
}

func ErrScan(err error) error {
	return errors.Wrap(err, "failed to scan")
}

func ErrExec(err error) error {
	return errors.Wrap(err, "failed to execute query")
}
