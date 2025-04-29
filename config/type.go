package config

type config struct {
	Etcd  etcd
	MySQL mySQL
	OSS   oss
	Redis redis
}
type etcd struct {
	Addr string
}
type mySQL struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}
type redis struct {
	Addr     string
	Username string
	Password string
}
type oss struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Domain    string
}
type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}
