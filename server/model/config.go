package model

type Config struct {
	Redis redis
}

type redis struct {
	DB       int    `json:"db" yaml:"db"`             // redis的哪个数据库
	Addr     string `json:"addr" yaml:"addr"`         // 服务器地址:端口
	Password string `json:"password" yaml:"password"` // 密码
}

type Website struct {
	LoginMode int    `json:"loginMode" yaml:"db"`
	Checking  bool   `json:"checking" yaml:"checking"`
	Cookie    string `json:"cookie" yaml:"cookie"`
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Hosts     string `json:"hosts" yaml:"hosts"`
	Notify    bool   `json:"notify" yaml:"notify"`
}
