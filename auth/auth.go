package auth

import (
	"encoding/json"
	"fmt"
	"os"
)

// {
//"accessToken":"5a59369066a235f2c5cb74e06df0886c3a748c26d970cf891f7a94f0c4dc0685",
//"refreshToken":"38d0a930576b13799ddd3e59813ec28966484cb84c8632b2438ebe9cdc7f8de8",
//"expiresAt":1700434525923
//}

type Callback func(info Info)

type Info struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int    `json:"expiresAt"`
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

func (a *Auth) Bearer() string {
	return "Bearer " + a.info.AccessToken
}

func (a *Auth) Refresh() string {
	return a.info.RefreshToken
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
