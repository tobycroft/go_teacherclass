package Redis

import (
	"context"
	"github.com/Unknwon/goconfig"
	"github.com/go-redis/redis/v9"
	"log"
	"main.go/config/app_conf"
	"time"
)

var goredis_ctx context.Context

var goredis *redis.Client

func init() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
	} else {
		value, err := cfg.GetSection("redis")
		if err != nil {
		} else {
			if value["address"] != "" && value["address"] != "" && value["database"] != "" {
				app_conf.Redicon_address = value["address"]
				app_conf.Redicon_port = value["port"]
				app_conf.Redicon_on = true
			}
		}
	}

	if !app_conf.Redicon_on {
		return
	}
	options := redis.Options{
		Addr:         app_conf.Redicon_address + ":" + app_conf.Redicon_port,
		Username:     app_conf.Redicon_username,
		Password:     app_conf.Redicon_password, // no password set
		Network:      app_conf.Redicon_proto,
		DB:           app_conf.Recion_db, // use default DB
		MinIdleConns: app_conf.Redicon_MinIdleConn,
		DialTimeout:  app_conf.Recion_timeout_dial,
		WriteTimeout: app_conf.Recion_timeout_write,
		ReadTimeout:  app_conf.Recion_timeout_read,
		PoolTimeout:  app_conf.Recion_timeout_pool,
	}
	if app_conf.Redicon_poolsize > 0 {
		options.PoolSize = app_conf.Redicon_poolsize
	}
	goredis = redis.NewClient(&options)
	goredis_ctx = goredis.Context()
	go func() {
		for {
			time.Sleep(3 * time.Second)
			ret, err := goredis.Ping(goredis_ctx).Result()
			if err != nil {
				log.Println("redis_out", ret, err)
				if app_conf.Recicon_panic_on_link_error {
					panic("redis_out")
				}
			}
		}
	}()
}
