// Code generated by MockGen. DO NOT EDIT.
// Source: container.go

// Package builder is a generated GoMock package.
package builder

import (
	context "context"
	io "io"
	reflect "reflect"

	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// MockContainerEngineClient is a mock of ContainerEngineClient interface.
type MockContainerEngineClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerEngineClientMockRecorder
}

// MockContainerEngineClientMockRecorder is the mock recorder for MockContainerEngineClient.
type MockContainerEngineClientMockRecorder struct {
	mock *MockContainerEngineClient
}

// NewMockContainerEngineClient creates a new mock instance.
func NewMockContainerEngineClient(ctrl *gomock.Controller) *MockContainerEngineClient {
	mock := &MockContainerEngineClient{ctrl: ctrl}
	mock.recorder = &MockContainerEngineClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerEngineClient) EXPECT() *MockContainerEngineClientMockRecorder {
	return m.recorder
}

// ContainerCreate mocks base method.
func (m *MockContainerEngineClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.ContainerCreateCreatedBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerCreate", ctx, config, hostConfig, networkingConfig, platform, containerName)
	ret0, _ := ret[0].(container.ContainerCreateCreatedBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerCreate indicates an expected call of ContainerCreate.
func (mr *MockContainerEngineClientMockRecorder) ContainerCreate(ctx, config, hostConfig, networkingConfig, platform, containerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerCreate", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerCreate), ctx, config, hostConfig, networkingConfig, platform, containerName)
}

// ContainerExecAttach mocks base method.
func (m *MockContainerEngineClient) ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecAttach", ctx, execID, config)
	ret0, _ := ret[0].(types.HijackedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerExecAttach indicates an expected call of ContainerExecAttach.
func (mr *MockContainerEngineClientMockRecorder) ContainerExecAttach(ctx, execID, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecAttach", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerExecAttach), ctx, execID, config)
}

// ContainerExecCreate mocks base method.
func (m *MockContainerEngineClient) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecCreate", ctx, container, config)
	ret0, _ := ret[0].(types.IDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerExecCreate indicates an expected call of ContainerExecCreate.
func (mr *MockContainerEngineClientMockRecorder) ContainerExecCreate(ctx, container, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecCreate", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerExecCreate), ctx, container, config)
}

// ContainerExecInspect mocks base method.
func (m *MockContainerEngineClient) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecInspect", ctx, execID)
	ret0, _ := ret[0].(types.ContainerExecInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerExecInspect indicates an expected call of ContainerExecInspect.
func (mr *MockContainerEngineClientMockRecorder) ContainerExecInspect(ctx, execID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecInspect", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerExecInspect), ctx, execID)
}

// ContainerExecStart mocks base method.
func (m *MockContainerEngineClient) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExecStart", ctx, execID, config)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerExecStart indicates an expected call of ContainerExecStart.
func (mr *MockContainerEngineClientMockRecorder) ContainerExecStart(ctx, execID, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExecStart", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerExecStart), ctx, execID, config)
}

// ContainerLogs mocks base method.
func (m *MockContainerEngineClient) ContainerLogs(ctx context.Context, containerName string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerLogs", ctx, containerName, options)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerLogs indicates an expected call of ContainerLogs.
func (mr *MockContainerEngineClientMockRecorder) ContainerLogs(ctx, containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerLogs", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerLogs), ctx, containerName, options)
}

// ContainerRemove mocks base method.
func (m *MockContainerEngineClient) ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerRemove", ctx, container, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerRemove indicates an expected call of ContainerRemove.
func (mr *MockContainerEngineClientMockRecorder) ContainerRemove(ctx, container, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerRemove", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerRemove), ctx, container, options)
}

// ContainerStart mocks base method.
func (m *MockContainerEngineClient) ContainerStart(ctx context.Context, containerName string, options types.ContainerStartOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerStart", ctx, containerName, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerStart indicates an expected call of ContainerStart.
func (mr *MockContainerEngineClientMockRecorder) ContainerStart(ctx, containerName, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerStart", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerStart), ctx, containerName, options)
}

// ContainerWait mocks base method.
func (m *MockContainerEngineClient) ContainerWait(ctx context.Context, containerName string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerWait", ctx, containerName, condition)
	ret0, _ := ret[0].(<-chan container.ContainerWaitOKBody)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// ContainerWait indicates an expected call of ContainerWait.
func (mr *MockContainerEngineClientMockRecorder) ContainerWait(ctx, containerName, condition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerWait", reflect.TypeOf((*MockContainerEngineClient)(nil).ContainerWait), ctx, containerName, condition)
}

// ImageBuild mocks base method.
func (m *MockContainerEngineClient) ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageBuild", ctx, context, options)
	ret0, _ := ret[0].(types.ImageBuildResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageBuild indicates an expected call of ImageBuild.
func (mr *MockContainerEngineClientMockRecorder) ImageBuild(ctx, context, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageBuild", reflect.TypeOf((*MockContainerEngineClient)(nil).ImageBuild), ctx, context, options)
}