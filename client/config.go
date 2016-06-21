package main

import (
)

type Config struct {
	HeartbeatPeriod uint16
	HeartbeatTimeout uint16
	HeartbeatLimit int
	Addr string
	ConnId string
}

func InitConfig() *Config {
	return &Config{
		HeartbeatPeriod : 3,
		HeartbeatTimeout : 5,
		HeartbeatLimit : 5,
		Addr : "", //tcp server addr
		ConnId : "", //uniq device identifier (registered from admin page)
	}
}
