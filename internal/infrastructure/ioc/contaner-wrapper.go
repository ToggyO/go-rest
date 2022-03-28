package ioc

import (
	"errors"
	"fmt"
	"go-rest/internal/infrastructure/config"
	"go-rest/internal/infrastructure/ioc/ioc_lib"
	"go.uber.org/dig"
	"reflect"
)

type containerWrapper struct {
	container *dig.Container
}

func NewIoc(configuration *config.Configuration) (ContainerWrapper, error) {
	container, err := ioc_lib.BuildDigIoc(configuration)
	if err != nil {
		return nil, err
	}
	return &containerWrapper{container}, nil
}

func (c *containerWrapper) GetService(function interface{}) error {
	err := c.container.Invoke(function)
	return err
}

func (c *containerWrapper) AddService(constructor interface{}) error {
	err := c.container.Provide(constructor)
	return err
}

func (c *containerWrapper) RunAfterBuild(functionList []func()) error {
	for _, f := range functionList {
		fType := reflect.TypeOf(f)
		if fType == nil {
			return errors.New("Can't invoke an untyped nil ")
		}

		if fType.Kind() != reflect.Func {
			// TODO: check replacement params
			return errors.New(fmt.Sprintf("can't invoke non-function %v (type %v)", f, fType))
		}

		f()
	}

	return nil
}
