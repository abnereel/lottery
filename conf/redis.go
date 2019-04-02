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
		Host:      "localhost",
		Port:      6379,
		User:      "",
		Pwd:       "password",
		IsRunning: true,
	},
}


var RdsCache RdsConfig = RdsCacheList[0]