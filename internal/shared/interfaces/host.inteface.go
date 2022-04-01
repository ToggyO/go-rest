package interfaces

import "context"

type IHost interface {
	Start() error
	Stop(ctx context.Context) error
}
