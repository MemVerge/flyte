// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/flyteorg/flyte/flyteadmin/pkg/workflowengine/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// WorkflowExecutor is an autogenerated mock type for the WorkflowExecutor type
type WorkflowExecutor struct {
	mock.Mock
}

type WorkflowExecutor_Abort struct {
	*mock.Call
}

func (_m WorkflowExecutor_Abort) Return(_a0 error) *WorkflowExecutor_Abort {
	return &WorkflowExecutor_Abort{Call: _m.Call.Return(_a0)}
}

func (_m *WorkflowExecutor) OnAbort(ctx context.Context, data interfaces.AbortData) *WorkflowExecutor_Abort {
	c_call := _m.On("Abort", ctx, data)
	return &WorkflowExecutor_Abort{Call: c_call}
}

func (_m *WorkflowExecutor) OnAbortMatch(matchers ...interface{}) *WorkflowExecutor_Abort {
	c_call := _m.On("Abort", matchers...)
	return &WorkflowExecutor_Abort{Call: c_call}
}

// Abort provides a mock function with given fields: ctx, data
func (_m *WorkflowExecutor) Abort(ctx context.Context, data interfaces.AbortData) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.AbortData) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type WorkflowExecutor_Execute struct {
	*mock.Call
}

func (_m WorkflowExecutor_Execute) Return(_a0 interfaces.ExecutionResponse, _a1 error) *WorkflowExecutor_Execute {
	return &WorkflowExecutor_Execute{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *WorkflowExecutor) OnExecute(ctx context.Context, data interfaces.ExecutionData) *WorkflowExecutor_Execute {
	c_call := _m.On("Execute", ctx, data)
	return &WorkflowExecutor_Execute{Call: c_call}
}

func (_m *WorkflowExecutor) OnExecuteMatch(matchers ...interface{}) *WorkflowExecutor_Execute {
	c_call := _m.On("Execute", matchers...)
	return &WorkflowExecutor_Execute{Call: c_call}
}

// Execute provides a mock function with given fields: ctx, data
func (_m *WorkflowExecutor) Execute(ctx context.Context, data interfaces.ExecutionData) (interfaces.ExecutionResponse, error) {
	ret := _m.Called(ctx, data)

	var r0 interfaces.ExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.ExecutionData) interfaces.ExecutionResponse); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(interfaces.ExecutionResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.ExecutionData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type WorkflowExecutor_ID struct {
	*mock.Call
}

func (_m WorkflowExecutor_ID) Return(_a0 string) *WorkflowExecutor_ID {
	return &WorkflowExecutor_ID{Call: _m.Call.Return(_a0)}
}

func (_m *WorkflowExecutor) OnID() *WorkflowExecutor_ID {
	c_call := _m.On("ID")
	return &WorkflowExecutor_ID{Call: c_call}
}

func (_m *WorkflowExecutor) OnIDMatch(matchers ...interface{}) *WorkflowExecutor_ID {
	c_call := _m.On("ID", matchers...)
	return &WorkflowExecutor_ID{Call: c_call}
}

// ID provides a mock function with given fields:
func (_m *WorkflowExecutor) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
