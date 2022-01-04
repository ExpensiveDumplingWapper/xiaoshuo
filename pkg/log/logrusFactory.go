package log

import (
	"xiaoshuo/internal/config"
	"xiaoshuo/pkg/log/logrus"

	"github.com/pkg/errors"
)

type LogrusFactory struct{}

func (mf *LogrusFactory) Build(lc *config.LogConfig) error {
	err := logrus.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
