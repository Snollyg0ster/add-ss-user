package config

type ConfigType struct {
	Server       string            `json:"server"`
	LocalAddress string            `json:"local_address"`
	LocalPort    int32             `json:"local_port"`
	PortPassword map[string]string `json:"port_password"`
	Timeout      int64             `json:"timeout"`
	Method       string            `json:"method"`
	Fast_open    bool              `json:"fast_open"`
	Users        map[string]string `json:"users"`
	MyServerIp   string            `json:"my_server_ip"`
}

var Config ConfigType
