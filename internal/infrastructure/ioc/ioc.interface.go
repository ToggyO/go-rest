package ioc

type ContainerWrapper interface {
	GetService(function interface{}) error
	AddService(constructor interface{}) error
	RunAfterBuild(functionList []func()) error
}
