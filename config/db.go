package config

type DBConfig struct {
	Driver                      string
	Host                        string
	Port                        string
	UserName                    string
	Password                    string
	Database                    string
	TimeZone                    string
	Log                         bool
	DisableForeignKeyConstraint bool
}
