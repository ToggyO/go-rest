package contracts

type IMapper interface {
	Map(source, dest interface{})
}

type IReturnableMapper interface {
	Map(source, dest interface{}) interface{}
}
