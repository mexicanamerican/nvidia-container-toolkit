// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package info

import (
	"sync"
)

// Ensure, that PropertyExtractorMock does implement PropertyExtractor.
// If this is not the case, regenerate this file with moq.
var _ PropertyExtractor = &PropertyExtractorMock{}

// PropertyExtractorMock is a mock implementation of PropertyExtractor.
//
//	func TestSomethingThatUsesPropertyExtractor(t *testing.T) {
//
//		// make and configure a mocked PropertyExtractor
//		mockedPropertyExtractor := &PropertyExtractorMock{
//			HasDXCoreFunc: func() (bool, string) {
//				panic("mock out the HasDXCore method")
//			},
//			HasNvmlFunc: func() (bool, string) {
//				panic("mock out the HasNvml method")
//			},
//			HasTegraFilesFunc: func() (bool, string) {
//				panic("mock out the HasTegraFiles method")
//			},
//			IsTegraSystemFunc: func() (bool, string) {
//				panic("mock out the IsTegraSystem method")
//			},
//			UsesOnlyNVGPUModuleFunc: func() (bool, string) {
//				panic("mock out the UsesOnlyNVGPUModule method")
//			},
//		}
//
//		// use mockedPropertyExtractor in code that requires PropertyExtractor
//		// and then make assertions.
//
//	}
type PropertyExtractorMock struct {
	// HasDXCoreFunc mocks the HasDXCore method.
	HasDXCoreFunc func() (bool, string)

	// HasNvmlFunc mocks the HasNvml method.
	HasNvmlFunc func() (bool, string)

	// HasTegraFilesFunc mocks the HasTegraFiles method.
	HasTegraFilesFunc func() (bool, string)

	// IsTegraSystemFunc mocks the IsTegraSystem method.
	IsTegraSystemFunc func() (bool, string)

	// UsesOnlyNVGPUModuleFunc mocks the UsesOnlyNVGPUModule method.
	UsesOnlyNVGPUModuleFunc func() (bool, string)

	// calls tracks calls to the methods.
	calls struct {
		// HasDXCore holds details about calls to the HasDXCore method.
		HasDXCore []struct {
		}
		// HasNvml holds details about calls to the HasNvml method.
		HasNvml []struct {
		}
		// HasTegraFiles holds details about calls to the HasTegraFiles method.
		HasTegraFiles []struct {
		}
		// IsTegraSystem holds details about calls to the IsTegraSystem method.
		IsTegraSystem []struct {
		}
		// UsesOnlyNVGPUModule holds details about calls to the UsesOnlyNVGPUModule method.
		UsesOnlyNVGPUModule []struct {
		}
	}
	lockHasDXCore           sync.RWMutex
	lockHasNvml             sync.RWMutex
	lockHasTegraFiles       sync.RWMutex
	lockIsTegraSystem       sync.RWMutex
	lockUsesOnlyNVGPUModule sync.RWMutex
}

// HasDXCore calls HasDXCoreFunc.
func (mock *PropertyExtractorMock) HasDXCore() (bool, string) {
	if mock.HasDXCoreFunc == nil {
		panic("PropertyExtractorMock.HasDXCoreFunc: method is nil but PropertyExtractor.HasDXCore was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHasDXCore.Lock()
	mock.calls.HasDXCore = append(mock.calls.HasDXCore, callInfo)
	mock.lockHasDXCore.Unlock()
	return mock.HasDXCoreFunc()
}

// HasDXCoreCalls gets all the calls that were made to HasDXCore.
// Check the length with:
//
//	len(mockedPropertyExtractor.HasDXCoreCalls())
func (mock *PropertyExtractorMock) HasDXCoreCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHasDXCore.RLock()
	calls = mock.calls.HasDXCore
	mock.lockHasDXCore.RUnlock()
	return calls
}

// HasNvml calls HasNvmlFunc.
func (mock *PropertyExtractorMock) HasNvml() (bool, string) {
	if mock.HasNvmlFunc == nil {
		panic("PropertyExtractorMock.HasNvmlFunc: method is nil but PropertyExtractor.HasNvml was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHasNvml.Lock()
	mock.calls.HasNvml = append(mock.calls.HasNvml, callInfo)
	mock.lockHasNvml.Unlock()
	return mock.HasNvmlFunc()
}

// HasNvmlCalls gets all the calls that were made to HasNvml.
// Check the length with:
//
//	len(mockedPropertyExtractor.HasNvmlCalls())
func (mock *PropertyExtractorMock) HasNvmlCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHasNvml.RLock()
	calls = mock.calls.HasNvml
	mock.lockHasNvml.RUnlock()
	return calls
}

// HasTegraFiles calls HasTegraFilesFunc.
func (mock *PropertyExtractorMock) HasTegraFiles() (bool, string) {
	if mock.HasTegraFilesFunc == nil {
		panic("PropertyExtractorMock.HasTegraFilesFunc: method is nil but PropertyExtractor.HasTegraFiles was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHasTegraFiles.Lock()
	mock.calls.HasTegraFiles = append(mock.calls.HasTegraFiles, callInfo)
	mock.lockHasTegraFiles.Unlock()
	return mock.HasTegraFilesFunc()
}

// HasTegraFilesCalls gets all the calls that were made to HasTegraFiles.
// Check the length with:
//
//	len(mockedPropertyExtractor.HasTegraFilesCalls())
func (mock *PropertyExtractorMock) HasTegraFilesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHasTegraFiles.RLock()
	calls = mock.calls.HasTegraFiles
	mock.lockHasTegraFiles.RUnlock()
	return calls
}

// IsTegraSystem calls IsTegraSystemFunc.
func (mock *PropertyExtractorMock) IsTegraSystem() (bool, string) {
	if mock.IsTegraSystemFunc == nil {
		panic("PropertyExtractorMock.IsTegraSystemFunc: method is nil but PropertyExtractor.IsTegraSystem was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsTegraSystem.Lock()
	mock.calls.IsTegraSystem = append(mock.calls.IsTegraSystem, callInfo)
	mock.lockIsTegraSystem.Unlock()
	return mock.IsTegraSystemFunc()
}

// IsTegraSystemCalls gets all the calls that were made to IsTegraSystem.
// Check the length with:
//
//	len(mockedPropertyExtractor.IsTegraSystemCalls())
func (mock *PropertyExtractorMock) IsTegraSystemCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsTegraSystem.RLock()
	calls = mock.calls.IsTegraSystem
	mock.lockIsTegraSystem.RUnlock()
	return calls
}

// UsesOnlyNVGPUModule calls UsesOnlyNVGPUModuleFunc.
func (mock *PropertyExtractorMock) UsesOnlyNVGPUModule() (bool, string) {
	if mock.UsesOnlyNVGPUModuleFunc == nil {
		panic("PropertyExtractorMock.UsesOnlyNVGPUModuleFunc: method is nil but PropertyExtractor.UsesOnlyNVGPUModule was just called")
	}
	callInfo := struct {
	}{}
	mock.lockUsesOnlyNVGPUModule.Lock()
	mock.calls.UsesOnlyNVGPUModule = append(mock.calls.UsesOnlyNVGPUModule, callInfo)
	mock.lockUsesOnlyNVGPUModule.Unlock()
	return mock.UsesOnlyNVGPUModuleFunc()
}

// UsesOnlyNVGPUModuleCalls gets all the calls that were made to UsesOnlyNVGPUModule.
// Check the length with:
//
//	len(mockedPropertyExtractor.UsesOnlyNVGPUModuleCalls())
func (mock *PropertyExtractorMock) UsesOnlyNVGPUModuleCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockUsesOnlyNVGPUModule.RLock()
	calls = mock.calls.UsesOnlyNVGPUModule
	mock.lockUsesOnlyNVGPUModule.RUnlock()
	return calls
}
