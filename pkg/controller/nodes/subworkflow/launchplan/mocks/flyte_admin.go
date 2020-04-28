// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

	launchplan "github.com/lyft/flytepropeller/pkg/controller/nodes/subworkflow/launchplan"

	mock "github.com/stretchr/testify/mock"
)

// FlyteAdmin is an autogenerated mock type for the FlyteAdmin type
type FlyteAdmin struct {
	mock.Mock
}

type FlyteAdmin_GetLaunchPlan struct {
	*mock.Call
}

func (_m FlyteAdmin_GetLaunchPlan) Return(_a0 *admin.LaunchPlan, _a1 error) *FlyteAdmin_GetLaunchPlan {
	return &FlyteAdmin_GetLaunchPlan{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *FlyteAdmin) OnGetLaunchPlan(ctx context.Context, launchPlanRef *core.Identifier) *FlyteAdmin_GetLaunchPlan {
	c := _m.On("GetLaunchPlan", ctx, launchPlanRef)
	return &FlyteAdmin_GetLaunchPlan{Call: c}
}

func (_m *FlyteAdmin) OnGetLaunchPlanMatch(matchers ...interface{}) *FlyteAdmin_GetLaunchPlan {
	c := _m.On("GetLaunchPlan", matchers...)
	return &FlyteAdmin_GetLaunchPlan{Call: c}
}

// GetLaunchPlan provides a mock function with given fields: ctx, launchPlanRef
func (_m *FlyteAdmin) GetLaunchPlan(ctx context.Context, launchPlanRef *core.Identifier) (*admin.LaunchPlan, error) {
	ret := _m.Called(ctx, launchPlanRef)

	var r0 *admin.LaunchPlan
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identifier) *admin.LaunchPlan); ok {
		r0 = rf(ctx, launchPlanRef)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.LaunchPlan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.Identifier) error); ok {
		r1 = rf(ctx, launchPlanRef)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type FlyteAdmin_GetStatus struct {
	*mock.Call
}

func (_m FlyteAdmin_GetStatus) Return(_a0 *admin.ExecutionClosure, _a1 error) *FlyteAdmin_GetStatus {
	return &FlyteAdmin_GetStatus{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *FlyteAdmin) OnGetStatus(ctx context.Context, executionID *core.WorkflowExecutionIdentifier) *FlyteAdmin_GetStatus {
	c := _m.On("GetStatus", ctx, executionID)
	return &FlyteAdmin_GetStatus{Call: c}
}

func (_m *FlyteAdmin) OnGetStatusMatch(matchers ...interface{}) *FlyteAdmin_GetStatus {
	c := _m.On("GetStatus", matchers...)
	return &FlyteAdmin_GetStatus{Call: c}
}

// GetStatus provides a mock function with given fields: ctx, executionID
func (_m *FlyteAdmin) GetStatus(ctx context.Context, executionID *core.WorkflowExecutionIdentifier) (*admin.ExecutionClosure, error) {
	ret := _m.Called(ctx, executionID)

	var r0 *admin.ExecutionClosure
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier) *admin.ExecutionClosure); ok {
		r0 = rf(ctx, executionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ExecutionClosure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowExecutionIdentifier) error); ok {
		r1 = rf(ctx, executionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type FlyteAdmin_Initialize struct {
	*mock.Call
}

func (_m FlyteAdmin_Initialize) Return(_a0 error) *FlyteAdmin_Initialize {
	return &FlyteAdmin_Initialize{Call: _m.Call.Return(_a0)}
}

func (_m *FlyteAdmin) OnInitialize(ctx context.Context) *FlyteAdmin_Initialize {
	c := _m.On("Initialize", ctx)
	return &FlyteAdmin_Initialize{Call: c}
}

func (_m *FlyteAdmin) OnInitializeMatch(matchers ...interface{}) *FlyteAdmin_Initialize {
	c := _m.On("Initialize", matchers...)
	return &FlyteAdmin_Initialize{Call: c}
}

// Initialize provides a mock function with given fields: ctx
func (_m *FlyteAdmin) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type FlyteAdmin_Kill struct {
	*mock.Call
}

func (_m FlyteAdmin_Kill) Return(_a0 error) *FlyteAdmin_Kill {
	return &FlyteAdmin_Kill{Call: _m.Call.Return(_a0)}
}

func (_m *FlyteAdmin) OnKill(ctx context.Context, executionID *core.WorkflowExecutionIdentifier, reason string) *FlyteAdmin_Kill {
	c := _m.On("Kill", ctx, executionID, reason)
	return &FlyteAdmin_Kill{Call: c}
}

func (_m *FlyteAdmin) OnKillMatch(matchers ...interface{}) *FlyteAdmin_Kill {
	c := _m.On("Kill", matchers...)
	return &FlyteAdmin_Kill{Call: c}
}

// Kill provides a mock function with given fields: ctx, executionID, reason
func (_m *FlyteAdmin) Kill(ctx context.Context, executionID *core.WorkflowExecutionIdentifier, reason string) error {
	ret := _m.Called(ctx, executionID, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowExecutionIdentifier, string) error); ok {
		r0 = rf(ctx, executionID, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type FlyteAdmin_Launch struct {
	*mock.Call
}

func (_m FlyteAdmin_Launch) Return(_a0 error) *FlyteAdmin_Launch {
	return &FlyteAdmin_Launch{Call: _m.Call.Return(_a0)}
}

func (_m *FlyteAdmin) OnLaunch(ctx context.Context, launchCtx launchplan.LaunchContext, executionID *core.WorkflowExecutionIdentifier, launchPlanRef *core.Identifier, inputs *core.LiteralMap) *FlyteAdmin_Launch {
	c := _m.On("Launch", ctx, launchCtx, executionID, launchPlanRef, inputs)
	return &FlyteAdmin_Launch{Call: c}
}

func (_m *FlyteAdmin) OnLaunchMatch(matchers ...interface{}) *FlyteAdmin_Launch {
	c := _m.On("Launch", matchers...)
	return &FlyteAdmin_Launch{Call: c}
}

// Launch provides a mock function with given fields: ctx, launchCtx, executionID, launchPlanRef, inputs
func (_m *FlyteAdmin) Launch(ctx context.Context, launchCtx launchplan.LaunchContext, executionID *core.WorkflowExecutionIdentifier, launchPlanRef *core.Identifier, inputs *core.LiteralMap) error {
	ret := _m.Called(ctx, launchCtx, executionID, launchPlanRef, inputs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, launchplan.LaunchContext, *core.WorkflowExecutionIdentifier, *core.Identifier, *core.LiteralMap) error); ok {
		r0 = rf(ctx, launchCtx, executionID, launchPlanRef, inputs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
