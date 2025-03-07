// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package modules

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
)

// ModuleMock is a mock implementation of Module.
//
// 	func TestSomethingThatUsesModule(t *testing.T) {
//
// 		// make and configure a mocked Module
// 		mockedModule := &ModuleMock{
// 			CoreApiHandlerFunc: func(method string, parts []string, params map[string]string, body []byte) (any, int, error) {
// 				panic("mock out the CoreApiHandler method")
// 			},
// 			GetLoaderFunc: func(j Job) lua.LGFunction {
// 				panic("mock out the GetLoader method")
// 			},
// 			NameFunc: func() string {
// 				panic("mock out the Name method")
// 			},
// 			StopFunc: func() error {
// 				panic("mock out the Stop method")
// 			},
// 		}
//
// 		// use mockedModule in code that requires Module
// 		// and then make assertions.
//
// 	}
type ModuleMock struct {
	// CoreApiHandlerFunc mocks the CoreApiHandler method.
	CoreApiHandlerFunc func(method string, parts []string, params map[string]string, body []byte) (any, int, error)

	// GetLoaderFunc mocks the GetLoader method.
	GetLoaderFunc func(j Job) lua.LGFunction

	// NameFunc mocks the Name method.
	NameFunc func() string

	// StopFunc mocks the Stop method.
	StopFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// CoreApiHandler holds details about calls to the CoreApiHandler method.
		CoreApiHandler []struct {
			// Method is the method argument value.
			Method string
			// Parts is the parts argument value.
			Parts []string
			// Params is the params argument value.
			Params map[string]string
			// Body is the body argument value.
			Body []byte
		}
		// GetLoader holds details about calls to the GetLoader method.
		GetLoader []struct {
			// J is the j argument value.
			J Job
		}
		// Name holds details about calls to the Name method.
		Name []struct {
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
	}
	lockCoreApiHandler sync.RWMutex
	lockGetLoader      sync.RWMutex
	lockName           sync.RWMutex
	lockStop           sync.RWMutex
}

// CoreApiHandler calls CoreApiHandlerFunc.
func (mock *ModuleMock) CoreApiHandler(method string, parts []string, params map[string]string, body []byte) (any, int, error) {
	if mock.CoreApiHandlerFunc == nil {
		panic("ModuleMock.CoreApiHandlerFunc: method is nil but Module.CoreApiHandler was just called")
	}
	callInfo := struct {
		Method string
		Parts  []string
		Params map[string]string
		Body   []byte
	}{
		Method: method,
		Parts:  parts,
		Params: params,
		Body:   body,
	}
	mock.lockCoreApiHandler.Lock()
	mock.calls.CoreApiHandler = append(mock.calls.CoreApiHandler, callInfo)
	mock.lockCoreApiHandler.Unlock()
	return mock.CoreApiHandlerFunc(method, parts, params, body)
}

// CoreApiHandlerCalls gets all the calls that were made to CoreApiHandler.
// Check the length with:
//     len(mockedModule.CoreApiHandlerCalls())
func (mock *ModuleMock) CoreApiHandlerCalls() []struct {
	Method string
	Parts  []string
	Params map[string]string
	Body   []byte
} {
	var calls []struct {
		Method string
		Parts  []string
		Params map[string]string
		Body   []byte
	}
	mock.lockCoreApiHandler.RLock()
	calls = mock.calls.CoreApiHandler
	mock.lockCoreApiHandler.RUnlock()
	return calls
}

// GetLoader calls GetLoaderFunc.
func (mock *ModuleMock) GetLoader(j Job) lua.LGFunction {
	if mock.GetLoaderFunc == nil {
		panic("ModuleMock.GetLoaderFunc: method is nil but Module.GetLoader was just called")
	}
	callInfo := struct {
		J Job
	}{
		J: j,
	}
	mock.lockGetLoader.Lock()
	mock.calls.GetLoader = append(mock.calls.GetLoader, callInfo)
	mock.lockGetLoader.Unlock()
	return mock.GetLoaderFunc(j)
}

// GetLoaderCalls gets all the calls that were made to GetLoader.
// Check the length with:
//     len(mockedModule.GetLoaderCalls())
func (mock *ModuleMock) GetLoaderCalls() []struct {
	J Job
} {
	var calls []struct {
		J Job
	}
	mock.lockGetLoader.RLock()
	calls = mock.calls.GetLoader
	mock.lockGetLoader.RUnlock()
	return calls
}

// Name calls NameFunc.
func (mock *ModuleMock) Name() string {
	if mock.NameFunc == nil {
		panic("ModuleMock.NameFunc: method is nil but Module.Name was just called")
	}
	callInfo := struct {
	}{}
	mock.lockName.Lock()
	mock.calls.Name = append(mock.calls.Name, callInfo)
	mock.lockName.Unlock()
	return mock.NameFunc()
}

// NameCalls gets all the calls that were made to Name.
// Check the length with:
//     len(mockedModule.NameCalls())
func (mock *ModuleMock) NameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockName.RLock()
	calls = mock.calls.Name
	mock.lockName.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *ModuleMock) Stop() error {
	if mock.StopFunc == nil {
		panic("ModuleMock.StopFunc: method is nil but Module.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//     len(mockedModule.StopCalls())
func (mock *ModuleMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}
