package main

type Config struct {
	RpcAddr string
	Port string
	DBAddr string
	DBUser string
	DBPasswd string
	DBName string
}

func InitConfig() *Config {
	return &Config{
		RpcAddr : ""//rpcaddr,
		Port : ""//webservport,
		DBAddr : ""//dbaddr,
		DBUser : ""//dbusername,
		DBPasswd : ""//dbpassword,
		DBName : ""//dbname,
	}
}
