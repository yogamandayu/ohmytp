package config

type DBConfig struct {
	Driver                      string
	Host                        string
	Port                        string
	User                        string
	Name                        string
	Password                    string
	TimeZone                    string
	Log                         bool
	DisableForeignKeyConstraint bool
}
