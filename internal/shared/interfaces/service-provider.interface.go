package interfaces

type IServiceProvider interface {
	GetService(function interface{}) error
	AddService(constructor interface{}) error
	RunAfterBuild(functionList []func()) error
}
