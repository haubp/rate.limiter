package redis_worker

// RedisConfig: Redis cofig
type RedisConfig struct {
	ServerName string		`json:"serverName"`
	Port int				`json:"Port"`
	Domain string			`json:"Domain"`
	Algorithm string		`json:"Algorithm"`
	Unit string				`json:"Unit"`
	RequestPerUnit int 		`json:"RequestPerUnit"`
}