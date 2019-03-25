package conf

type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool
}

var RdsCacheList = []RdsConfig {
	{
		Host:      "47.96.7.111",
		Port:      6379,
		User:      "",
		Pwd:       "redis123qwe",
		IsRunning: true,
	},
}


var RdsCache RdsConfig = RdsCacheList[0]