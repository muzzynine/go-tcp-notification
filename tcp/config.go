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
		TCPAddr : ""//tcpaddr,
		RPCAddr : ""//rpcaddr,
		HeartbeatLimit : 5,
		DBAddr : "",//dbaddr 
		DBUser : "",//dbusername
		DBPasswd : "",//dbpassword
		DBName : ""//dbname,
	}
}
