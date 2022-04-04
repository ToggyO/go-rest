package settings

import "fmt"

type PgDbSettings struct {
	DbSettings

	template string
}

func (s *PgDbSettings) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		s.DbHost, s.DbPort, s.DbUser, s.DbPassword, s.DbName)
}
