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
		Host:      "localhost",
		Port:      3306,
		User:      "lottery",
		Pwd:       "password",
		Database:  "lottery",
		isRunning: true,
	},
}

var DbMaster DbConfig = DbMasterList[0]