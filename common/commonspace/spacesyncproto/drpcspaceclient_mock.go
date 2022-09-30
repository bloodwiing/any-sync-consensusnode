// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto (interfaces: DRPCSpaceClient)

// Package spacesyncproto is a generated GoMock package.
package spacesyncproto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	drpc "storj.io/drpc"
)

// MockDRPCSpaceClient is a mock of DRPCSpaceClient interface.
type MockDRPCSpaceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDRPCSpaceClientMockRecorder
}

// MockDRPCSpaceClientMockRecorder is the mock recorder for MockDRPCSpaceClient.
type MockDRPCSpaceClientMockRecorder struct {
	mock *MockDRPCSpaceClient
}

// NewMockDRPCSpaceClient creates a new mock instance.
func NewMockDRPCSpaceClient(ctrl *gomock.Controller) *MockDRPCSpaceClient {
	mock := &MockDRPCSpaceClient{ctrl: ctrl}
	mock.recorder = &MockDRPCSpaceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDRPCSpaceClient) EXPECT() *MockDRPCSpaceClientMockRecorder {
	return m.recorder
}

// DRPCConn mocks base method.
func (m *MockDRPCSpaceClient) DRPCConn() drpc.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DRPCConn")
	ret0, _ := ret[0].(drpc.Conn)
	return ret0
}

// DRPCConn indicates an expected call of DRPCConn.
func (mr *MockDRPCSpaceClientMockRecorder) DRPCConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DRPCConn", reflect.TypeOf((*MockDRPCSpaceClient)(nil).DRPCConn))
}

// HeadSync mocks base method.
func (m *MockDRPCSpaceClient) HeadSync(arg0 context.Context, arg1 *HeadSyncRequest) (*HeadSyncResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadSync", arg0, arg1)
	ret0, _ := ret[0].(*HeadSyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadSync indicates an expected call of HeadSync.
func (mr *MockDRPCSpaceClientMockRecorder) HeadSync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadSync", reflect.TypeOf((*MockDRPCSpaceClient)(nil).HeadSync), arg0, arg1)
}

// PushSpace mocks base method.
func (m *MockDRPCSpaceClient) PushSpace(arg0 context.Context, arg1 *PushSpaceRequest) (*PushSpaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushSpace", arg0, arg1)
	ret0, _ := ret[0].(*PushSpaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PushSpace indicates an expected call of PushSpace.
func (mr *MockDRPCSpaceClientMockRecorder) PushSpace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushSpace", reflect.TypeOf((*MockDRPCSpaceClient)(nil).PushSpace), arg0, arg1)
}

// Stream mocks base method.
func (m *MockDRPCSpaceClient) Stream(arg0 context.Context) (DRPCSpace_StreamClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(DRPCSpace_StreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream.
func (mr *MockDRPCSpaceClientMockRecorder) Stream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockDRPCSpaceClient)(nil).Stream), arg0)
}
