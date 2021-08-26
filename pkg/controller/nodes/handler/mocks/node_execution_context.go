// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	events "github.com/flyteorg/flytepropeller/pkg/controller/events"
	executors "github.com/flyteorg/flytepropeller/pkg/controller/executors"

	handler "github.com/flyteorg/flytepropeller/pkg/controller/nodes/handler"

	io "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/io"

	ioutils "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/ioutils"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flytestdlib/storage"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// NodeExecutionContext is an autogenerated mock type for the NodeExecutionContext type
type NodeExecutionContext struct {
	mock.Mock
}

type NodeExecutionContext_ContextualNodeLookup struct {
	*mock.Call
}

func (_m NodeExecutionContext_ContextualNodeLookup) Return(_a0 executors.NodeLookup) *NodeExecutionContext_ContextualNodeLookup {
	return &NodeExecutionContext_ContextualNodeLookup{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnContextualNodeLookup() *NodeExecutionContext_ContextualNodeLookup {
	c := _m.On("ContextualNodeLookup")
	return &NodeExecutionContext_ContextualNodeLookup{Call: c}
}

func (_m *NodeExecutionContext) OnContextualNodeLookupMatch(matchers ...interface{}) *NodeExecutionContext_ContextualNodeLookup {
	c := _m.On("ContextualNodeLookup", matchers...)
	return &NodeExecutionContext_ContextualNodeLookup{Call: c}
}

// ContextualNodeLookup provides a mock function with given fields:
func (_m *NodeExecutionContext) ContextualNodeLookup() executors.NodeLookup {
	ret := _m.Called()

	var r0 executors.NodeLookup
	if rf, ok := ret.Get(0).(func() executors.NodeLookup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executors.NodeLookup)
		}
	}

	return r0
}

type NodeExecutionContext_CurrentAttempt struct {
	*mock.Call
}

func (_m NodeExecutionContext_CurrentAttempt) Return(_a0 uint32) *NodeExecutionContext_CurrentAttempt {
	return &NodeExecutionContext_CurrentAttempt{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnCurrentAttempt() *NodeExecutionContext_CurrentAttempt {
	c := _m.On("CurrentAttempt")
	return &NodeExecutionContext_CurrentAttempt{Call: c}
}

func (_m *NodeExecutionContext) OnCurrentAttemptMatch(matchers ...interface{}) *NodeExecutionContext_CurrentAttempt {
	c := _m.On("CurrentAttempt", matchers...)
	return &NodeExecutionContext_CurrentAttempt{Call: c}
}

// CurrentAttempt provides a mock function with given fields:
func (_m *NodeExecutionContext) CurrentAttempt() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type NodeExecutionContext_DataStore struct {
	*mock.Call
}

func (_m NodeExecutionContext_DataStore) Return(_a0 *storage.DataStore) *NodeExecutionContext_DataStore {
	return &NodeExecutionContext_DataStore{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnDataStore() *NodeExecutionContext_DataStore {
	c := _m.On("DataStore")
	return &NodeExecutionContext_DataStore{Call: c}
}

func (_m *NodeExecutionContext) OnDataStoreMatch(matchers ...interface{}) *NodeExecutionContext_DataStore {
	c := _m.On("DataStore", matchers...)
	return &NodeExecutionContext_DataStore{Call: c}
}

// DataStore provides a mock function with given fields:
func (_m *NodeExecutionContext) DataStore() *storage.DataStore {
	ret := _m.Called()

	var r0 *storage.DataStore
	if rf, ok := ret.Get(0).(func() *storage.DataStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.DataStore)
		}
	}

	return r0
}

type NodeExecutionContext_EnqueueOwnerFunc struct {
	*mock.Call
}

func (_m NodeExecutionContext_EnqueueOwnerFunc) Return(_a0 func() error) *NodeExecutionContext_EnqueueOwnerFunc {
	return &NodeExecutionContext_EnqueueOwnerFunc{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnEnqueueOwnerFunc() *NodeExecutionContext_EnqueueOwnerFunc {
	c := _m.On("EnqueueOwnerFunc")
	return &NodeExecutionContext_EnqueueOwnerFunc{Call: c}
}

func (_m *NodeExecutionContext) OnEnqueueOwnerFuncMatch(matchers ...interface{}) *NodeExecutionContext_EnqueueOwnerFunc {
	c := _m.On("EnqueueOwnerFunc", matchers...)
	return &NodeExecutionContext_EnqueueOwnerFunc{Call: c}
}

// EnqueueOwnerFunc provides a mock function with given fields:
func (_m *NodeExecutionContext) EnqueueOwnerFunc() func() error {
	ret := _m.Called()

	var r0 func() error
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

type NodeExecutionContext_EventsRecorder struct {
	*mock.Call
}

func (_m NodeExecutionContext_EventsRecorder) Return(_a0 events.TaskEventRecorder) *NodeExecutionContext_EventsRecorder {
	return &NodeExecutionContext_EventsRecorder{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnEventsRecorder() *NodeExecutionContext_EventsRecorder {
	c := _m.On("EventsRecorder")
	return &NodeExecutionContext_EventsRecorder{Call: c}
}

func (_m *NodeExecutionContext) OnEventsRecorderMatch(matchers ...interface{}) *NodeExecutionContext_EventsRecorder {
	c := _m.On("EventsRecorder", matchers...)
	return &NodeExecutionContext_EventsRecorder{Call: c}
}

// EventsRecorder provides a mock function with given fields:
func (_m *NodeExecutionContext) EventsRecorder() events.TaskEventRecorder {
	ret := _m.Called()

	var r0 events.TaskEventRecorder
	if rf, ok := ret.Get(0).(func() events.TaskEventRecorder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(events.TaskEventRecorder)
		}
	}

	return r0
}

type NodeExecutionContext_ExecutionContext struct {
	*mock.Call
}

func (_m NodeExecutionContext_ExecutionContext) Return(_a0 executors.ExecutionContext) *NodeExecutionContext_ExecutionContext {
	return &NodeExecutionContext_ExecutionContext{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnExecutionContext() *NodeExecutionContext_ExecutionContext {
	c := _m.On("ExecutionContext")
	return &NodeExecutionContext_ExecutionContext{Call: c}
}

func (_m *NodeExecutionContext) OnExecutionContextMatch(matchers ...interface{}) *NodeExecutionContext_ExecutionContext {
	c := _m.On("ExecutionContext", matchers...)
	return &NodeExecutionContext_ExecutionContext{Call: c}
}

// ExecutionContext provides a mock function with given fields:
func (_m *NodeExecutionContext) ExecutionContext() executors.ExecutionContext {
	ret := _m.Called()

	var r0 executors.ExecutionContext
	if rf, ok := ret.Get(0).(func() executors.ExecutionContext); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executors.ExecutionContext)
		}
	}

	return r0
}

type NodeExecutionContext_InputReader struct {
	*mock.Call
}

func (_m NodeExecutionContext_InputReader) Return(_a0 io.InputReader) *NodeExecutionContext_InputReader {
	return &NodeExecutionContext_InputReader{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnInputReader() *NodeExecutionContext_InputReader {
	c := _m.On("InputReader")
	return &NodeExecutionContext_InputReader{Call: c}
}

func (_m *NodeExecutionContext) OnInputReaderMatch(matchers ...interface{}) *NodeExecutionContext_InputReader {
	c := _m.On("InputReader", matchers...)
	return &NodeExecutionContext_InputReader{Call: c}
}

// InputReader provides a mock function with given fields:
func (_m *NodeExecutionContext) InputReader() io.InputReader {
	ret := _m.Called()

	var r0 io.InputReader
	if rf, ok := ret.Get(0).(func() io.InputReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.InputReader)
		}
	}

	return r0
}

type NodeExecutionContext_MaxDatasetSizeBytes struct {
	*mock.Call
}

func (_m NodeExecutionContext_MaxDatasetSizeBytes) Return(_a0 int64) *NodeExecutionContext_MaxDatasetSizeBytes {
	return &NodeExecutionContext_MaxDatasetSizeBytes{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnMaxDatasetSizeBytes() *NodeExecutionContext_MaxDatasetSizeBytes {
	c := _m.On("MaxDatasetSizeBytes")
	return &NodeExecutionContext_MaxDatasetSizeBytes{Call: c}
}

func (_m *NodeExecutionContext) OnMaxDatasetSizeBytesMatch(matchers ...interface{}) *NodeExecutionContext_MaxDatasetSizeBytes {
	c := _m.On("MaxDatasetSizeBytes", matchers...)
	return &NodeExecutionContext_MaxDatasetSizeBytes{Call: c}
}

// MaxDatasetSizeBytes provides a mock function with given fields:
func (_m *NodeExecutionContext) MaxDatasetSizeBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type NodeExecutionContext_Node struct {
	*mock.Call
}

func (_m NodeExecutionContext_Node) Return(_a0 v1alpha1.ExecutableNode) *NodeExecutionContext_Node {
	return &NodeExecutionContext_Node{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNode() *NodeExecutionContext_Node {
	c := _m.On("Node")
	return &NodeExecutionContext_Node{Call: c}
}

func (_m *NodeExecutionContext) OnNodeMatch(matchers ...interface{}) *NodeExecutionContext_Node {
	c := _m.On("Node", matchers...)
	return &NodeExecutionContext_Node{Call: c}
}

// Node provides a mock function with given fields:
func (_m *NodeExecutionContext) Node() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

type NodeExecutionContext_NodeExecutionMetadata struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeExecutionMetadata) Return(_a0 handler.NodeExecutionMetadata) *NodeExecutionContext_NodeExecutionMetadata {
	return &NodeExecutionContext_NodeExecutionMetadata{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeExecutionMetadata() *NodeExecutionContext_NodeExecutionMetadata {
	c := _m.On("NodeExecutionMetadata")
	return &NodeExecutionContext_NodeExecutionMetadata{Call: c}
}

func (_m *NodeExecutionContext) OnNodeExecutionMetadataMatch(matchers ...interface{}) *NodeExecutionContext_NodeExecutionMetadata {
	c := _m.On("NodeExecutionMetadata", matchers...)
	return &NodeExecutionContext_NodeExecutionMetadata{Call: c}
}

// NodeExecutionMetadata provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeExecutionMetadata() handler.NodeExecutionMetadata {
	ret := _m.Called()

	var r0 handler.NodeExecutionMetadata
	if rf, ok := ret.Get(0).(func() handler.NodeExecutionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handler.NodeExecutionMetadata)
		}
	}

	return r0
}

type NodeExecutionContext_NodeID struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeID) Return(_a0 string) *NodeExecutionContext_NodeID {
	return &NodeExecutionContext_NodeID{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeID() *NodeExecutionContext_NodeID {
	c := _m.On("NodeID")
	return &NodeExecutionContext_NodeID{Call: c}
}

func (_m *NodeExecutionContext) OnNodeIDMatch(matchers ...interface{}) *NodeExecutionContext_NodeID {
	c := _m.On("NodeID", matchers...)
	return &NodeExecutionContext_NodeID{Call: c}
}

// NodeID provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type NodeExecutionContext_NodeStateReader struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeStateReader) Return(_a0 handler.NodeStateReader) *NodeExecutionContext_NodeStateReader {
	return &NodeExecutionContext_NodeStateReader{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeStateReader() *NodeExecutionContext_NodeStateReader {
	c := _m.On("NodeStateReader")
	return &NodeExecutionContext_NodeStateReader{Call: c}
}

func (_m *NodeExecutionContext) OnNodeStateReaderMatch(matchers ...interface{}) *NodeExecutionContext_NodeStateReader {
	c := _m.On("NodeStateReader", matchers...)
	return &NodeExecutionContext_NodeStateReader{Call: c}
}

// NodeStateReader provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeStateReader() handler.NodeStateReader {
	ret := _m.Called()

	var r0 handler.NodeStateReader
	if rf, ok := ret.Get(0).(func() handler.NodeStateReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handler.NodeStateReader)
		}
	}

	return r0
}

type NodeExecutionContext_NodeStateWriter struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeStateWriter) Return(_a0 handler.NodeStateWriter) *NodeExecutionContext_NodeStateWriter {
	return &NodeExecutionContext_NodeStateWriter{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeStateWriter() *NodeExecutionContext_NodeStateWriter {
	c := _m.On("NodeStateWriter")
	return &NodeExecutionContext_NodeStateWriter{Call: c}
}

func (_m *NodeExecutionContext) OnNodeStateWriterMatch(matchers ...interface{}) *NodeExecutionContext_NodeStateWriter {
	c := _m.On("NodeStateWriter", matchers...)
	return &NodeExecutionContext_NodeStateWriter{Call: c}
}

// NodeStateWriter provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeStateWriter() handler.NodeStateWriter {
	ret := _m.Called()

	var r0 handler.NodeStateWriter
	if rf, ok := ret.Get(0).(func() handler.NodeStateWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handler.NodeStateWriter)
		}
	}

	return r0
}

type NodeExecutionContext_NodeStatus struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeStatus) Return(_a0 v1alpha1.ExecutableNodeStatus) *NodeExecutionContext_NodeStatus {
	return &NodeExecutionContext_NodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeStatus() *NodeExecutionContext_NodeStatus {
	c := _m.On("NodeStatus")
	return &NodeExecutionContext_NodeStatus{Call: c}
}

func (_m *NodeExecutionContext) OnNodeStatusMatch(matchers ...interface{}) *NodeExecutionContext_NodeStatus {
	c := _m.On("NodeStatus", matchers...)
	return &NodeExecutionContext_NodeStatus{Call: c}
}

// NodeStatus provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeStatus() v1alpha1.ExecutableNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

type NodeExecutionContext_OutputShardSelector struct {
	*mock.Call
}

func (_m NodeExecutionContext_OutputShardSelector) Return(_a0 ioutils.ShardSelector) *NodeExecutionContext_OutputShardSelector {
	return &NodeExecutionContext_OutputShardSelector{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnOutputShardSelector() *NodeExecutionContext_OutputShardSelector {
	c := _m.On("OutputShardSelector")
	return &NodeExecutionContext_OutputShardSelector{Call: c}
}

func (_m *NodeExecutionContext) OnOutputShardSelectorMatch(matchers ...interface{}) *NodeExecutionContext_OutputShardSelector {
	c := _m.On("OutputShardSelector", matchers...)
	return &NodeExecutionContext_OutputShardSelector{Call: c}
}

// OutputShardSelector provides a mock function with given fields:
func (_m *NodeExecutionContext) OutputShardSelector() ioutils.ShardSelector {
	ret := _m.Called()

	var r0 ioutils.ShardSelector
	if rf, ok := ret.Get(0).(func() ioutils.ShardSelector); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ioutils.ShardSelector)
		}
	}

	return r0
}

type NodeExecutionContext_RawOutputPrefix struct {
	*mock.Call
}

func (_m NodeExecutionContext_RawOutputPrefix) Return(_a0 storage.DataReference) *NodeExecutionContext_RawOutputPrefix {
	return &NodeExecutionContext_RawOutputPrefix{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnRawOutputPrefix() *NodeExecutionContext_RawOutputPrefix {
	c := _m.On("RawOutputPrefix")
	return &NodeExecutionContext_RawOutputPrefix{Call: c}
}

func (_m *NodeExecutionContext) OnRawOutputPrefixMatch(matchers ...interface{}) *NodeExecutionContext_RawOutputPrefix {
	c := _m.On("RawOutputPrefix", matchers...)
	return &NodeExecutionContext_RawOutputPrefix{Call: c}
}

// RawOutputPrefix provides a mock function with given fields:
func (_m *NodeExecutionContext) RawOutputPrefix() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type NodeExecutionContext_TaskReader struct {
	*mock.Call
}

func (_m NodeExecutionContext_TaskReader) Return(_a0 handler.TaskReader) *NodeExecutionContext_TaskReader {
	return &NodeExecutionContext_TaskReader{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnTaskReader() *NodeExecutionContext_TaskReader {
	c := _m.On("TaskReader")
	return &NodeExecutionContext_TaskReader{Call: c}
}

func (_m *NodeExecutionContext) OnTaskReaderMatch(matchers ...interface{}) *NodeExecutionContext_TaskReader {
	c := _m.On("TaskReader", matchers...)
	return &NodeExecutionContext_TaskReader{Call: c}
}

// TaskReader provides a mock function with given fields:
func (_m *NodeExecutionContext) TaskReader() handler.TaskReader {
	ret := _m.Called()

	var r0 handler.TaskReader
	if rf, ok := ret.Get(0).(func() handler.TaskReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handler.TaskReader)
		}
	}

	return r0
}
