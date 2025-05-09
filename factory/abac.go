package factory

import (
	"github.com/ddan1l/tega-backend/abac"
)

type ABACFactory interface {
	CreateABAC() abac.Engine
}

func (f *DefaultFactory) CreateABAC() abac.Engine {
	return abac.NewEngine(f.db)
}
