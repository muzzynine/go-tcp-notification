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
		Addr : "localhost:8001",
		ConnId : "68f0082b-38de-47ae-8464-bc6b2411a8ff",
	}
}
