package test

import (
	"github.com/balerter/balerter/internal/mock"
	"github.com/balerter/balerter/internal/modules"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
	"strings"
)

type modulesManager interface {
	Get() []modules.ModuleTest
}

// Test represents the Test core module
type Test struct {
	dsManager      modulesManager
	storageManager modulesManager
	mods           []modules.ModuleTest
	logger         *zap.Logger

	datasource map[string]modules.ModuleTest
	storage    map[string]modules.ModuleTest
}

// New creates new Test core module
func New(dsManager, storageManager modulesManager, mods []modules.ModuleTest, logger *zap.Logger) *Test {
	t := &Test{
		dsManager:      dsManager,
		storageManager: storageManager,
		mods:           mods,
		logger:         logger,

		datasource: make(map[string]modules.ModuleTest),
		storage:    make(map[string]modules.ModuleTest),
	}

	for _, module := range t.dsManager.Get() {
		t.datasource[module.Name()] = module
	}

	for _, module := range t.storageManager.Get() {
		t.storage[module.Name()] = module
	}

	return t
}

// Name returns the module name
func (t *Test) Name() string {
	return "test"
}

// Stop the module
func (t *Test) Stop() error {
	return nil
}

func (t *Test) getModule(moduleName string, j modules.Job) lua.LGFunction {
	return func(luaState *lua.LState) int {
		nameL := luaState.Get(1)
		if nameL.Type() == lua.LTNil {
			t.logger.Error("module should have 1 argument", zap.String("module", moduleName))
			return 0
		}

		name := strings.TrimSpace(nameL.String())
		if name == "" {
			t.logger.Error("module should have 1 not empty argument", zap.String("module", moduleName))
			return 0
		}

		var m modules.ModuleTest
		var ok bool

		switch moduleName {
		case "test.datasource":
			m, ok = t.datasource[name]
		case "test.storage":
			m, ok = t.storage[name]
		}

		if !ok {
			t.logger.Error("module not found", zap.String("module name", moduleName), zap.String("name", name))
			return 0
		}

		m.GetLoader(j)(luaState)

		return 1
	}
}

// GetLoader returns the lua loader
func (t *Test) GetLoader(j modules.Job) lua.LGFunction {
	return func(luaState *lua.LState) int {
		var exports = map[string]lua.LGFunction{
			"datasource": t.getModule("test.datasource", j),
			"storage":    t.getModule("test.storage", j),
		}

		for _, mod := range t.mods {
			exports[mod.Name()] = mod.GetLoader(j)
		}

		mod := luaState.SetFuncs(luaState.NewTable(), exports)

		mod.RawSetString("AnyValue", lua.LString(mock.AnyValue))

		luaState.Push(mod)
		return 1
	}
}

// Result is stub for ModuleTest interface
func (t *Test) Result() ([]modules.TestResult, error) {
	return nil, nil
}

// Clean is stub for ModuleTest interface
func (t *Test) Clean() {

}
