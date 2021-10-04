package enum

import (
	"os"
	"sync"
)

type env struct {
	PORT,
	AccountName,
	AppKey,
	AppToken string
}

var Env *env
var once sync.Once

func GetEnv() {
	once.Do(func() {
		Env = &env{
			PORT:        os.Getenv("PORT"),
			AccountName: os.Getenv("AccountName"),
			AppKey:      os.Getenv("AppKey"),
			AppToken:    os.Getenv("AppToken"),
		}
	})
}
