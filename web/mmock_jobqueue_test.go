// Generated by running the following command and fixed the type
// signatures to resolve external packages:
//   mockgen -source ../jobqueue/jobqueue.go -destination mmock_jobqueue_test.go -package web
// i.e., the following command with
// https://github.com/golang/mock/pull/38 merged should generate the
// same thing:
//   mockgen -source ../jobqueue/jobqueue.go -destination mmock_jobqueue_test.go -package web -source_package jobqueue
package web

import (
	reflect "reflect"

	"github.com/fireworq/fireworq/jobqueue"

	gomock "github.com/golang/mock/gomock"
)

// MockJobQueue is a mock of JobQueue interface
type MockJobQueue struct {
	ctrl     *gomock.Controller
	recorder *MockJobQueueMockRecorder
}

// MockJobQueueMockRecorder is the mock recorder for MockJobQueue
type MockJobQueueMockRecorder struct {
	mock *MockJobQueue
}

// NewMockJobQueue creates a new mock instance
func NewMockJobQueue(ctrl *gomock.Controller) *MockJobQueue {
	mock := &MockJobQueue{ctrl: ctrl}
	mock.recorder = &MockJobQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobQueue) EXPECT() *MockJobQueueMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockJobQueue) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockJobQueueMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockJobQueue)(nil).Start))
}

// Stop mocks base method
func (m *MockJobQueue) Stop() <-chan struct{} {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockJobQueueMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJobQueue)(nil).Stop))
}

// Push mocks base method
func (m *MockJobQueue) Push(job jobqueue.IncomingJob) (uint64, error) {
	ret := m.ctrl.Call(m, "Push", job)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Push indicates an expected call of Push
func (mr *MockJobQueueMockRecorder) Push(job interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockJobQueue)(nil).Push), job)
}

// Pop mocks base method
func (m *MockJobQueue) Pop(limit uint) ([]jobqueue.Job, error) {
	ret := m.ctrl.Call(m, "Pop", limit)
	ret0, _ := ret[0].([]jobqueue.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pop indicates an expected call of Pop
func (mr *MockJobQueueMockRecorder) Pop(limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pop", reflect.TypeOf((*MockJobQueue)(nil).Pop), limit)
}

// Complete mocks base method
func (m *MockJobQueue) Complete(job jobqueue.Job, res *jobqueue.Result) {
	m.ctrl.Call(m, "Complete", job, res)
}

// Complete indicates an expected call of Complete
func (mr *MockJobQueueMockRecorder) Complete(job, res interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Complete", reflect.TypeOf((*MockJobQueue)(nil).Complete), job, res)
}

// Name mocks base method
func (m *MockJobQueue) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockJobQueueMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockJobQueue)(nil).Name))
}

// IsActive mocks base method
func (m *MockJobQueue) IsActive() bool {
	ret := m.ctrl.Call(m, "IsActive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActive indicates an expected call of IsActive
func (mr *MockJobQueueMockRecorder) IsActive() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockJobQueue)(nil).IsActive))
}

// Node mocks base method
func (m *MockJobQueue) Node() (*jobqueue.Node, error) {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(*jobqueue.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Node indicates an expected call of Node
func (mr *MockJobQueueMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockJobQueue)(nil).Node))
}

// Stats mocks base method
func (m *MockJobQueue) Stats() *jobqueue.Stats {
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(*jobqueue.Stats)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockJobQueueMockRecorder) Stats() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockJobQueue)(nil).Stats))
}

// Inspector mocks base method
func (m *MockJobQueue) Inspector() (jobqueue.Inspector, bool) {
	ret := m.ctrl.Call(m, "Inspector")
	ret0, _ := ret[0].(jobqueue.Inspector)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Inspector indicates an expected call of Inspector
func (mr *MockJobQueueMockRecorder) Inspector() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspector", reflect.TypeOf((*MockJobQueue)(nil).Inspector))
}

// FailureLog mocks base method
func (m *MockJobQueue) FailureLog() (jobqueue.FailureLog, bool) {
	ret := m.ctrl.Call(m, "FailureLog")
	ret0, _ := ret[0].(jobqueue.FailureLog)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FailureLog indicates an expected call of FailureLog
func (mr *MockJobQueueMockRecorder) FailureLog() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureLog", reflect.TypeOf((*MockJobQueue)(nil).FailureLog))
}

// MockImpl is a mock of Impl interface
type MockImpl struct {
	ctrl     *gomock.Controller
	recorder *MockImplMockRecorder
}

// MockImplMockRecorder is the mock recorder for MockImpl
type MockImplMockRecorder struct {
	mock *MockImpl
}

// NewMockImpl creates a new mock instance
func NewMockImpl(ctrl *gomock.Controller) *MockImpl {
	mock := &MockImpl{ctrl: ctrl}
	mock.recorder = &MockImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImpl) EXPECT() *MockImplMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockImpl) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockImplMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockImpl)(nil).Start))
}

// Stop mocks base method
func (m *MockImpl) Stop() <-chan struct{} {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockImplMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockImpl)(nil).Stop))
}

// Push mocks base method
func (m *MockImpl) Push(job jobqueue.IncomingJob) (jobqueue.Job, error) {
	ret := m.ctrl.Call(m, "Push", job)
	ret0, _ := ret[0].(jobqueue.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Push indicates an expected call of Push
func (mr *MockImplMockRecorder) Push(job interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockImpl)(nil).Push), job)
}

// Pop mocks base method
func (m *MockImpl) Pop(limit uint) ([]jobqueue.Job, error) {
	ret := m.ctrl.Call(m, "Pop", limit)
	ret0, _ := ret[0].([]jobqueue.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pop indicates an expected call of Pop
func (mr *MockImplMockRecorder) Pop(limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pop", reflect.TypeOf((*MockImpl)(nil).Pop), limit)
}

// Delete mocks base method
func (m *MockImpl) Delete(job jobqueue.Job) {
	m.ctrl.Call(m, "Delete", job)
}

// Delete indicates an expected call of Delete
func (mr *MockImplMockRecorder) Delete(job interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImpl)(nil).Delete), job)
}

// Update mocks base method
func (m *MockImpl) Update(job jobqueue.Job, next jobqueue.NextInfo) {
	m.ctrl.Call(m, "Update", job, next)
}

// Update indicates an expected call of Update
func (mr *MockImplMockRecorder) Update(job, next interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockImpl)(nil).Update), job, next)
}

// IsActive mocks base method
func (m *MockImpl) IsActive() bool {
	ret := m.ctrl.Call(m, "IsActive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsActive indicates an expected call of IsActive
func (mr *MockImplMockRecorder) IsActive() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockImpl)(nil).IsActive))
}
