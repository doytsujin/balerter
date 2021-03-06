package modules

import (
	"github.com/balerter/balerter/internal/script/script"
	"github.com/stretchr/testify/mock"
	lua "github.com/yuin/gopher-lua"
)

type Module interface {
	Name() string
	GetLoader(script *script.Script) lua.LGFunction
	Stop() error
}

type ModuleMock struct {
	mock.Mock
}

func (m *ModuleMock) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *ModuleMock) GetLoader(script *script.Script) lua.LGFunction {
	args := m.Called(script)
	return args.Get(0).(lua.LGFunction)
}

func (m *ModuleMock) Stop() error {
	args := m.Called()
	return args.Error(0)
}
