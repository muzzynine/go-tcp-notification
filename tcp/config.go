package main

type Config struct {
	RcvBufferSize int
	WrtBufferSize int
	RPCAddr string
	TCPAddr string
	HeartbeatTimeout uint16
	HeartbeatLimit int
	ConnId string
	DBAddr string
	DBUser string
	DBPasswd string
	DBName string
}

func InitConfig() *Config {
	return &Config{
		RcvBufferSize : 1024,
		WrtBufferSize : 1024,
		TCPAddr : "192.168.0.8:8010",
		RPCAddr : "localhost:8002",
		HeartbeatLimit : 5,
		DBAddr : "signboard.cqm2majqgqx4.ap-northeast-1.rds.amazonaws.com:3306",
		DBUser : "muzzynine",
		DBPasswd : "su1c1delog1c",
		DBName : "signboard",
	}
}