// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/flyteorg/flyte/flyteplugins/go/tasks/plugins/presto/client"

	mock "github.com/stretchr/testify/mock"
)

// PrestoClient is an autogenerated mock type for the PrestoClient type
type PrestoClient struct {
	mock.Mock
}

type PrestoClient_ExecuteCommand struct {
	*mock.Call
}

func (_m PrestoClient_ExecuteCommand) Return(_a0 client.PrestoExecuteResponse, _a1 error) *PrestoClient_ExecuteCommand {
	return &PrestoClient_ExecuteCommand{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *PrestoClient) OnExecuteCommand(ctx context.Context, commandStr string, executeArgs client.PrestoExecuteArgs) *PrestoClient_ExecuteCommand {
	c_call := _m.On("ExecuteCommand", ctx, commandStr, executeArgs)
	return &PrestoClient_ExecuteCommand{Call: c_call}
}

func (_m *PrestoClient) OnExecuteCommandMatch(matchers ...interface{}) *PrestoClient_ExecuteCommand {
	c_call := _m.On("ExecuteCommand", matchers...)
	return &PrestoClient_ExecuteCommand{Call: c_call}
}

// ExecuteCommand provides a mock function with given fields: ctx, commandStr, executeArgs
func (_m *PrestoClient) ExecuteCommand(ctx context.Context, commandStr string, executeArgs client.PrestoExecuteArgs) (client.PrestoExecuteResponse, error) {
	ret := _m.Called(ctx, commandStr, executeArgs)

	var r0 client.PrestoExecuteResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, client.PrestoExecuteArgs) client.PrestoExecuteResponse); ok {
		r0 = rf(ctx, commandStr, executeArgs)
	} else {
		r0 = ret.Get(0).(client.PrestoExecuteResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, client.PrestoExecuteArgs) error); ok {
		r1 = rf(ctx, commandStr, executeArgs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type PrestoClient_GetCommandStatus struct {
	*mock.Call
}

func (_m PrestoClient_GetCommandStatus) Return(_a0 client.PrestoStatus, _a1 error) *PrestoClient_GetCommandStatus {
	return &PrestoClient_GetCommandStatus{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *PrestoClient) OnGetCommandStatus(ctx context.Context, commandID string) *PrestoClient_GetCommandStatus {
	c_call := _m.On("GetCommandStatus", ctx, commandID)
	return &PrestoClient_GetCommandStatus{Call: c_call}
}

func (_m *PrestoClient) OnGetCommandStatusMatch(matchers ...interface{}) *PrestoClient_GetCommandStatus {
	c_call := _m.On("GetCommandStatus", matchers...)
	return &PrestoClient_GetCommandStatus{Call: c_call}
}

// GetCommandStatus provides a mock function with given fields: ctx, commandID
func (_m *PrestoClient) GetCommandStatus(ctx context.Context, commandID string) (client.PrestoStatus, error) {
	ret := _m.Called(ctx, commandID)

	var r0 client.PrestoStatus
	if rf, ok := ret.Get(0).(func(context.Context, string) client.PrestoStatus); ok {
		r0 = rf(ctx, commandID)
	} else {
		r0 = ret.Get(0).(client.PrestoStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, commandID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type PrestoClient_KillCommand struct {
	*mock.Call
}

func (_m PrestoClient_KillCommand) Return(_a0 error) *PrestoClient_KillCommand {
	return &PrestoClient_KillCommand{Call: _m.Call.Return(_a0)}
}

func (_m *PrestoClient) OnKillCommand(ctx context.Context, commandID string) *PrestoClient_KillCommand {
	c_call := _m.On("KillCommand", ctx, commandID)
	return &PrestoClient_KillCommand{Call: c_call}
}

func (_m *PrestoClient) OnKillCommandMatch(matchers ...interface{}) *PrestoClient_KillCommand {
	c_call := _m.On("KillCommand", matchers...)
	return &PrestoClient_KillCommand{Call: c_call}
}

// KillCommand provides a mock function with given fields: ctx, commandID
func (_m *PrestoClient) KillCommand(ctx context.Context, commandID string) error {
	ret := _m.Called(ctx, commandID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, commandID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
