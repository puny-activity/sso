package config

type Database struct {
	database       string
	host           string
	port           int
	name           string
	user           string
	password       string
	migrationsPath string
}

func (c Database) Database() string {
	return c.database
}

func (c Database) Host() string {
	return c.host
}

func (c Database) Port() int {
	return c.port
}

func (c Database) Name() string {
	return c.name
}

func (c Database) User() string {
	return c.user
}

func (c Database) Password() string {
	return c.password
}

func (c Database) MigrationsPath() string {
	return c.migrationsPath
}
