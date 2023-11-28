package auth

import (
	"encoding/json"
	"os"
)

type Option func(a *Auth) error

func WithFile(file string) Option {
	return func(a *Auth) error {
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(data, &(a.info)); err != nil {
			return err
		}

		a.file = file

		return nil
	}
}

func WithInfo(info Info) Option {
	return func(a *Auth) error {
		a.info = info

		return nil
	}
}

func WithInfoUpdateCallback(callback Callback) Option {
	return func(a *Auth) error {
		a.callback = callback

		return nil
	}
}
