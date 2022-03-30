package di

type ServiceDescriptorOptions interface {
}

type ServiceDescriptor struct {
	Service interface{}
	Options interface{}
}
