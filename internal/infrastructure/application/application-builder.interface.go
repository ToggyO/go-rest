package application

import "go-rest/internal/shared/interfaces"

type IApplicationBuilder interface {
	SetServiceProvider(serviceProvider interfaces.IServiceProvider) IApplicationBuilder
	SetWebHost(host interfaces.IHost) IApplicationBuilder
	Build() (IApplication, error)
}
