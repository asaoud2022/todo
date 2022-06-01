package repository

import (
	"github.com/asaoud2022/todo/config"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var (
	Storage *redis.Storage
	Store   *session.Store
)

func ConnectStore(c *config.Config) (*session.Store, error) {

	redisCfg := redis.Config{
		Host:     c.SessionStorageHost,
		Port:     c.SessionStoragePort,
		Username: c.SessionStorageUserName,
		Password: c.SessionStoragePassword,
		Database: c.SessionStorageDatabase,
		Reset:    c.SessionStorageReset,
	}

	//redisCfg.TLSConfig = c.SessionStorageTlsEnabled
	//if c.SessionStorageTlsEnabled {
	//	redisCfg.TLSConfig = c.GetClientTLSConfig()
	//}
	Storage = redis.New(redisCfg)

	//TODO:: Need to set all Session

	return session.New(session.Config{
		Storage: Storage,
	}), nil

	/*
		return session.Config{
			Expiration:     c.GetDuration("MW_FIBER_SESSION_EXPIRATION"),
			Storage: Storage,
			KeyLookup:      fmt.Sprintf("cookie:%s", c.SessionSesionCookieName),
			CookieDomain:   c.SessionCookieDomain,
			CookiePath:     c.SessionCookiePath,
			CookieSecure:   c.SessionCookieSecure,
			CookieHTTPOnly: c.SessionStorageCookieHttpOnly,
			CookieSameSite: c.SessionStorageCookieSameSite,
		}
	*/
}
