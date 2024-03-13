package auth

import (
	"encoding/json"
	"fmt"
	"os"
)

type Callback func(info Info)

type Info struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
	DeviceId     string `json:"deviceId"`
}

type Auth struct {
	file     string
	info     Info
	callback Callback
}

func New(options ...Option) (*Auth, error) {
	auth := &Auth{
		file: "",
		info: Info{},
	}

	for _, o := range options {
		if err := o(auth); err != nil {
			return nil, err
		}
	}

	return auth, nil
}

func (a *Auth) Info() Info {
	return a.info
}

func (a *Auth) Update(info Info) {
	a.info = info
	if a.callback != nil {
		a.callback(info)
	}
}

func (a *Auth) BearerHeader() string {
	return "Bearer " + a.info.AccessToken
}

func (a *Auth) RefreshToken() string {
	return a.info.RefreshToken
}

func (a *Auth) DeviceId() string {
	return a.info.DeviceId
}

func (a *Auth) Save() error {
	if a.file == "" {
		return nil
	}

	data, err := json.Marshal(a.info)
	if err != nil {
		return fmt.Errorf("error on marshal auth data: %w", err)
	}
	return os.WriteFile(a.file, data, 0644)
}
