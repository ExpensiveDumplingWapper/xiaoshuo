package log

import (
	"xiaoshuo/internal/config"
	"xiaoshuo/pkg/log/zap"

	"github.com/pkg/errors"
)

type ZapFactory struct{}

func (mf *ZapFactory) Build(lc *config.LogConfig) error {
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
