package local

import (
	"context"
	uow "lazy/common"
	"sync"
)

type localUow struct {
	mutex sync.Mutex
}

func NewUnitOfWork() uow.UnitOfWork {
	return &localUow{}
}

func (uow *localUow) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	uow.mutex.Lock()
	defer uow.mutex.Unlock()

	err := fn(ctx)
	if err != nil {
		return err
	}

	return nil
}
