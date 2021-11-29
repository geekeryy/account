// Package redis @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/11/27 10:05 下午
package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
}

func NewClient(confStr string) *Client {
	opt := Options{}
	if err := json.Unmarshal([]byte(confStr), &opt); err != nil {
		panic(err.Error())
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err.Error())
	}
	return &Client{
		rdb,
	}
}

type Options struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
