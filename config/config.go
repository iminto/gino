package config

type WebConfig struct {
	HostAddr     string
	DbType       string
	DSN          string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
	SQLDebug     bool
}

var Config = WebConfig{
	HostAddr:     ":8090",
	DbType:       "mysql",
	DSN:          "www:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
	MaxIdleConns: 5,
	MaxLifetime:  7200,
	MaxOpenConns: 5,
	SQLDebug:     true,
}
