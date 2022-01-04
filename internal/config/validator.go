package config

import "github.com/pkg/errors"

func validateConfig(app App) error {
	err := validateMasterMySqlConfig(app)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateMasterMySqlConfig(app App) error {

	return nil
}
