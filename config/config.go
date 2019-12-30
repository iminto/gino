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
	DSN:          "web:web123@tcp(10.180.249.73:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
	MaxIdleConns: 5,
	MaxLifetime:  7200,
	MaxOpenConns: 5,
	SQLDebug:     true,
}
