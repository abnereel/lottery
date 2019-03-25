package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	isRunning bool
}

var DbMasterList = []DbConfig {
	{
		Host:      "47.96.7.111",
		Port:      3306,
		User:      "lottery",
		Pwd:       "lottery@123QWE",
		Database:  "lottery",
		isRunning: true,
	},
}

var DbMaster DbConfig = DbMasterList[0]