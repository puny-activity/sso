package dbconnector

type Config interface {
	Database() string
	Host() string
	Port() int
	Name() string
	User() string
	Password() string
	MigrationsPath() string
}
