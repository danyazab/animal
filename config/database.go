package config

type Database struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

func (cfg *Database) Cgf() *Database {
	return cfg
}
