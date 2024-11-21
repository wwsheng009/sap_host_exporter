// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SUSE/sap_host_exporter/lib/sapcontrol (interfaces: WebService)

// Package mock_sapcontrol is a generated GoMock package.
package mock_sapcontrol

import (
	reflect "reflect"

	sapcontrol "github.com/SUSE/sap_host_exporter/lib/sapcontrol"
	gomock "github.com/golang/mock/gomock"
)

// MockWebService is a mock of WebService interface.
type MockWebService struct {
	ctrl     *gomock.Controller
	recorder *MockWebServiceMockRecorder
}

// MockWebServiceMockRecorder is the mock recorder for MockWebService.
type MockWebServiceMockRecorder struct {
	mock *MockWebService
}

// NewMockWebService creates a new mock instance.
func NewMockWebService(ctrl *gomock.Controller) *MockWebService {
	mock := &MockWebService{ctrl: ctrl}
	mock.recorder = &MockWebServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebService) EXPECT() *MockWebServiceMockRecorder {
	return m.recorder
}

// EnqGetStatistic mocks base method.
func (m *MockWebService) EnqGetStatistic() (*sapcontrol.EnqGetStatisticResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnqGetStatistic")
	ret0, _ := ret[0].(*sapcontrol.EnqGetStatisticResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqGetStatistic indicates an expected call of EnqGetStatistic.
func (mr *MockWebServiceMockRecorder) EnqGetStatistic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqGetStatistic", reflect.TypeOf((*MockWebService)(nil).EnqGetStatistic))
}

// GetCurrentInstance mocks base method.
func (m *MockWebService) GetCurrentInstance() (*sapcontrol.CurrentSapInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentInstance")
	ret0, _ := ret[0].(*sapcontrol.CurrentSapInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentInstance indicates an expected call of GetCurrentInstance.
func (mr *MockWebServiceMockRecorder) GetCurrentInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentInstance", reflect.TypeOf((*MockWebService)(nil).GetCurrentInstance))
}

// GetInstanceProperties mocks base method.
func (m *MockWebService) GetInstanceProperties() (*sapcontrol.GetInstancePropertiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceProperties")
	ret0, _ := ret[0].(*sapcontrol.GetInstancePropertiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceProperties indicates an expected call of GetInstanceProperties.
func (mr *MockWebServiceMockRecorder) GetInstanceProperties() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceProperties", reflect.TypeOf((*MockWebService)(nil).GetInstanceProperties))
}

// GetProcessList mocks base method.
func (m *MockWebService) GetProcessList() (*sapcontrol.GetProcessListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessList")
	ret0, _ := ret[0].(*sapcontrol.GetProcessListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessList indicates an expected call of GetProcessList.
func (mr *MockWebServiceMockRecorder) GetProcessList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessList", reflect.TypeOf((*MockWebService)(nil).GetProcessList))
}

// GetQueueStatistic mocks base method.
func (m *MockWebService) GetQueueStatistic() (*sapcontrol.GetQueueStatisticResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueueStatistic")
	ret0, _ := ret[0].(*sapcontrol.GetQueueStatisticResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueueStatistic indicates an expected call of GetQueueStatistic.
func (mr *MockWebServiceMockRecorder) GetQueueStatistic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueStatistic", reflect.TypeOf((*MockWebService)(nil).GetQueueStatistic))
}

// GetSystemInstanceList mocks base method.
func (m *MockWebService) GetSystemInstanceList() (*sapcontrol.GetSystemInstanceListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemInstanceList")
	ret0, _ := ret[0].(*sapcontrol.GetSystemInstanceListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemInstanceList indicates an expected call of GetSystemInstanceList.
func (mr *MockWebServiceMockRecorder) GetSystemInstanceList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemInstanceList", reflect.TypeOf((*MockWebService)(nil).GetSystemInstanceList))
}

// * Returns a list of the ABAP work processes (similar to SM50 transaction).
func (m *MockWebService) GetABAPWPTable() (*sapcontrol.GetABAPWPTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetABAPWPTable")
	ret0, _ := ret[0].(*sapcontrol.GetABAPWPTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
func (mr *MockWebServiceMockRecorder) GetABAPWPTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetABAPWPTable", reflect.TypeOf((*MockWebService)(nil).GetABAPWPTable))
}
