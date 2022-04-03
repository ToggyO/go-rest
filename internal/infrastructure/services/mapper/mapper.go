package mapper

import (
	"github.com/peteprogrammer/go-automapper"
	"go-rest/internal/application/contracts"
)

type Mapper struct{}

func NewMapper() contracts.IMapper {
	return &Mapper{}
}

func (m *Mapper) Map(source, dest interface{}) {
	automapper.Map(source, dest)
}
