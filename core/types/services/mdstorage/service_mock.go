// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mdstorage

import (
	"context"
	"github.com/MaibornWolff/maDocK8s/core/types/protocol"
	"sync"
)

// Ensure, that MdstorageServerMock does implement MdstorageServer.
// If this is not the case, regenerate this file with moq.
var _ MdstorageServer = &MdstorageServerMock{}

// MdstorageServerMock is a mock implementation of MdstorageServer.
//
//     func TestSomethingThatUsesMdstorageServer(t *testing.T) {
//
//         // make and configure a mocked MdstorageServer
//         mockedMdstorageServer := &MdstorageServerMock{
//             DeleteMarkdownFunc: func(in1 context.Context, in2 *protocol.Markdown) (*RemoveResult, error) {
// 	               panic("mock out the DeleteMarkdown method")
//             },
//             StoreMarkdownFunc: func(in1 context.Context, in2 *protocol.Markdown) (*StoreResult, error) {
// 	               panic("mock out the StoreMarkdown method")
//             },
//             UpdateMarkdownFunc: func(in1 context.Context, in2 *protocol.Markdown) (*StoreResult, error) {
// 	               panic("mock out the UpdateMarkdown method")
//             },
//         }
//
//         // use mockedMdstorageServer in code that requires MdstorageServer
//         // and then make assertions.
//
//     }
type MdstorageServerMock struct {
	// DeleteMarkdownFunc mocks the DeleteMarkdown method.
	DeleteMarkdownFunc func(in1 context.Context, in2 *protocol.Markdown) (*RemoveResult, error)

	// StoreMarkdownFunc mocks the StoreMarkdown method.
	StoreMarkdownFunc func(in1 context.Context, in2 *protocol.Markdown) (*StoreResult, error)

	// UpdateMarkdownFunc mocks the UpdateMarkdown method.
	UpdateMarkdownFunc func(in1 context.Context, in2 *protocol.Markdown) (*StoreResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteMarkdown holds details about calls to the DeleteMarkdown method.
		DeleteMarkdown []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 *protocol.Markdown
		}
		// StoreMarkdown holds details about calls to the StoreMarkdown method.
		StoreMarkdown []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 *protocol.Markdown
		}
		// UpdateMarkdown holds details about calls to the UpdateMarkdown method.
		UpdateMarkdown []struct {
			// In1 is the in1 argument value.
			In1 context.Context
			// In2 is the in2 argument value.
			In2 *protocol.Markdown
		}
	}
	lockDeleteMarkdown sync.RWMutex
	lockStoreMarkdown  sync.RWMutex
	lockUpdateMarkdown sync.RWMutex
}

// DeleteMarkdown calls DeleteMarkdownFunc.
func (mock *MdstorageServerMock) DeleteMarkdown(in1 context.Context, in2 *protocol.Markdown) (*RemoveResult, error) {
	if mock.DeleteMarkdownFunc == nil {
		panic("MdstorageServerMock.DeleteMarkdownFunc: method is nil but MdstorageServer.DeleteMarkdown was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 *protocol.Markdown
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockDeleteMarkdown.Lock()
	mock.calls.DeleteMarkdown = append(mock.calls.DeleteMarkdown, callInfo)
	mock.lockDeleteMarkdown.Unlock()
	return mock.DeleteMarkdownFunc(in1, in2)
}

// DeleteMarkdownCalls gets all the calls that were made to DeleteMarkdown.
// Check the length with:
//     len(mockedMdstorageServer.DeleteMarkdownCalls())
func (mock *MdstorageServerMock) DeleteMarkdownCalls() []struct {
	In1 context.Context
	In2 *protocol.Markdown
} {
	var calls []struct {
		In1 context.Context
		In2 *protocol.Markdown
	}
	mock.lockDeleteMarkdown.RLock()
	calls = mock.calls.DeleteMarkdown
	mock.lockDeleteMarkdown.RUnlock()
	return calls
}

// StoreMarkdown calls StoreMarkdownFunc.
func (mock *MdstorageServerMock) StoreMarkdown(in1 context.Context, in2 *protocol.Markdown) (*StoreResult, error) {
	if mock.StoreMarkdownFunc == nil {
		panic("MdstorageServerMock.StoreMarkdownFunc: method is nil but MdstorageServer.StoreMarkdown was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 *protocol.Markdown
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockStoreMarkdown.Lock()
	mock.calls.StoreMarkdown = append(mock.calls.StoreMarkdown, callInfo)
	mock.lockStoreMarkdown.Unlock()
	return mock.StoreMarkdownFunc(in1, in2)
}

// StoreMarkdownCalls gets all the calls that were made to StoreMarkdown.
// Check the length with:
//     len(mockedMdstorageServer.StoreMarkdownCalls())
func (mock *MdstorageServerMock) StoreMarkdownCalls() []struct {
	In1 context.Context
	In2 *protocol.Markdown
} {
	var calls []struct {
		In1 context.Context
		In2 *protocol.Markdown
	}
	mock.lockStoreMarkdown.RLock()
	calls = mock.calls.StoreMarkdown
	mock.lockStoreMarkdown.RUnlock()
	return calls
}

// UpdateMarkdown calls UpdateMarkdownFunc.
func (mock *MdstorageServerMock) UpdateMarkdown(in1 context.Context, in2 *protocol.Markdown) (*StoreResult, error) {
	if mock.UpdateMarkdownFunc == nil {
		panic("MdstorageServerMock.UpdateMarkdownFunc: method is nil but MdstorageServer.UpdateMarkdown was just called")
	}
	callInfo := struct {
		In1 context.Context
		In2 *protocol.Markdown
	}{
		In1: in1,
		In2: in2,
	}
	mock.lockUpdateMarkdown.Lock()
	mock.calls.UpdateMarkdown = append(mock.calls.UpdateMarkdown, callInfo)
	mock.lockUpdateMarkdown.Unlock()
	return mock.UpdateMarkdownFunc(in1, in2)
}

// UpdateMarkdownCalls gets all the calls that were made to UpdateMarkdown.
// Check the length with:
//     len(mockedMdstorageServer.UpdateMarkdownCalls())
func (mock *MdstorageServerMock) UpdateMarkdownCalls() []struct {
	In1 context.Context
	In2 *protocol.Markdown
} {
	var calls []struct {
		In1 context.Context
		In2 *protocol.Markdown
	}
	mock.lockUpdateMarkdown.RLock()
	calls = mock.calls.UpdateMarkdown
	mock.lockUpdateMarkdown.RUnlock()
	return calls
}
