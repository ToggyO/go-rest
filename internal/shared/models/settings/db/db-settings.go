package settings

type IDbSettings interface {
	GetConnectionString() string
}

type DbSettings struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}
