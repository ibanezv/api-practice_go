package settings

type Database struct {
	Username        string
	Password        string
	Host            string
	DatabaseName    string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	ConnMaxIdleTime int
	Timeout         int
	ReadTimeout     int
	WriteTimeout    int
	EnableMigrate   bool
}

func LoadConfigurationDB() Database {
	config := Database{}
	config.DatabaseName = "local"
	config.Host = "localhost:3306"
	config.Password = "Redric1973_"
	config.Username = "root"
	return config
}
