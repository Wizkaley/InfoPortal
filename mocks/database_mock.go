// Code generated by MockGen. DO NOT EDIT.
// Source: mongoDAL.go

// Package mongodal is a generated GoMock package.
package mongodal

import (
	mongodal "RESTApp/utils/mongodal"
	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
	reflect "reflect"
)

// MockMgoSessionDAL is a mock of MgoSessionDAL interface
type MockMgoSessionDAL struct {
	ctrl     *gomock.Controller
	recorder *MockMgoSessionDALMockRecorder
}

// MockMgoSessionDALMockRecorder is the mock recorder for MockMgoSessionDAL
type MockMgoSessionDALMockRecorder struct {
	mock *MockMgoSessionDAL
}

// NewMockMgoSessionDAL creates a new mock instance
func NewMockMgoSessionDAL(ctrl *gomock.Controller) *MockMgoSessionDAL {
	mock := &MockMgoSessionDAL{ctrl: ctrl}
	mock.recorder = &MockMgoSessionDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMgoSessionDAL) EXPECT() *MockMgoSessionDALMockRecorder {
	return m.recorder
}

// DB mocks base method
func (m *MockMgoSessionDAL) DB(db string) mongodal.MgoDBDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB", db)
	ret0, _ := ret[0].(mongodal.MgoDBDAL)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockMgoSessionDALMockRecorder) DB(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockMgoSessionDAL)(nil).DB), db)
}

// MockMgoDBDAL is a mock of MgoDBDAL interface
type MockMgoDBDAL struct {
	ctrl     *gomock.Controller
	recorder *MockMgoDBDALMockRecorder
}

// MockMgoDBDALMockRecorder is the mock recorder for MockMgoDBDAL
type MockMgoDBDALMockRecorder struct {
	mock *MockMgoDBDAL
}

// NewMockMgoDBDAL creates a new mock instance
func NewMockMgoDBDAL(ctrl *gomock.Controller) *MockMgoDBDAL {
	mock := &MockMgoDBDAL{ctrl: ctrl}
	mock.recorder = &MockMgoDBDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMgoDBDAL) EXPECT() *MockMgoDBDALMockRecorder {
	return m.recorder
}

// C mocks base method
func (m *MockMgoDBDAL) C(collection string) mongodal.MgoCollectionDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "C", collection)
	ret0, _ := ret[0].(mongodal.MgoCollectionDAL)
	return ret0
}

// C indicates an expected call of C
func (mr *MockMgoDBDALMockRecorder) C(collection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockMgoDBDAL)(nil).C), collection)
}

// MockMgoCollectionDAL is a mock of MgoCollectionDAL interface
type MockMgoCollectionDAL struct {
	ctrl     *gomock.Controller
	recorder *MockMgoCollectionDALMockRecorder
}

// MockMgoCollectionDALMockRecorder is the mock recorder for MockMgoCollectionDAL
type MockMgoCollectionDALMockRecorder struct {
	mock *MockMgoCollectionDAL
}

// NewMockMgoCollectionDAL creates a new mock instance
func NewMockMgoCollectionDAL(ctrl *gomock.Controller) *MockMgoCollectionDAL {
	mock := &MockMgoCollectionDAL{ctrl: ctrl}
	mock.recorder = &MockMgoCollectionDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMgoCollectionDAL) EXPECT() *MockMgoCollectionDALMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockMgoCollectionDAL) Insert(docs ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range docs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockMgoCollectionDALMockRecorder) Insert(docs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMgoCollectionDAL)(nil).Insert), docs...)
}

// Find mocks base method
func (m *MockMgoCollectionDAL) Find(query interface{}) mongodal.MgoQueryDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", query)
	ret0, _ := ret[0].(mongodal.MgoQueryDAL)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockMgoCollectionDALMockRecorder) Find(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMgoCollectionDAL)(nil).Find), query)
}

// Upsert mocks base method
func (m *MockMgoCollectionDAL) Upsert(selector, update interface{}) (*mgo_v2.ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", selector, update)
	ret0, _ := ret[0].(*mgo_v2.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockMgoCollectionDALMockRecorder) Upsert(selector, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockMgoCollectionDAL)(nil).Upsert), selector, update)
}

// Update mocks base method
func (m *MockMgoCollectionDAL) Update(selectot, update interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", selectot, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockMgoCollectionDALMockRecorder) Update(selectot, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMgoCollectionDAL)(nil).Update), selectot, update)
}

// UpdateAll mocks base method
func (m *MockMgoCollectionDAL) UpdateAll(selector, update interface{}) (*mgo_v2.ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAll", selector, update)
	ret0, _ := ret[0].(*mgo_v2.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockMgoCollectionDALMockRecorder) UpdateAll(selector, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockMgoCollectionDAL)(nil).UpdateAll), selector, update)
}

// Remove mocks base method
func (m *MockMgoCollectionDAL) Remove(selector interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", selector)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockMgoCollectionDALMockRecorder) Remove(selector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockMgoCollectionDAL)(nil).Remove), selector)
}

// RemoveAll mocks base method
func (m *MockMgoCollectionDAL) RemoveAll(selctor interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", selctor)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockMgoCollectionDALMockRecorder) RemoveAll(selctor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockMgoCollectionDAL)(nil).RemoveAll), selctor)
}

// Pipe mocks base method
func (m *MockMgoCollectionDAL) Pipe(pipeline interface{}) mongodal.MgoPipeDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipe", pipeline)
	ret0, _ := ret[0].(mongodal.MgoPipeDAL)
	return ret0
}

// Pipe indicates an expected call of Pipe
func (mr *MockMgoCollectionDALMockRecorder) Pipe(pipeline interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockMgoCollectionDAL)(nil).Pipe), pipeline)
}

// MockMgoQueryDAL is a mock of MgoQueryDAL interface
type MockMgoQueryDAL struct {
	ctrl     *gomock.Controller
	recorder *MockMgoQueryDALMockRecorder
}

// MockMgoQueryDALMockRecorder is the mock recorder for MockMgoQueryDAL
type MockMgoQueryDALMockRecorder struct {
	mock *MockMgoQueryDAL
}

// NewMockMgoQueryDAL creates a new mock instance
func NewMockMgoQueryDAL(ctrl *gomock.Controller) *MockMgoQueryDAL {
	mock := &MockMgoQueryDAL{ctrl: ctrl}
	mock.recorder = &MockMgoQueryDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMgoQueryDAL) EXPECT() *MockMgoQueryDALMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockMgoQueryDAL) All(result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockMgoQueryDALMockRecorder) All(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockMgoQueryDAL)(nil).All), result)
}

// One mocks base method
func (m *MockMgoQueryDAL) One(result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockMgoQueryDALMockRecorder) One(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockMgoQueryDAL)(nil).One), result)
}

// Count mocks base method
func (m *MockMgoQueryDAL) Count() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockMgoQueryDALMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockMgoQueryDAL)(nil).Count))
}

// Skip mocks base method
func (m *MockMgoQueryDAL) Skip(n int) mongodal.MgoQueryDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Skip", n)
	ret0, _ := ret[0].(mongodal.MgoQueryDAL)
	return ret0
}

// Skip indicates an expected call of Skip
func (mr *MockMgoQueryDALMockRecorder) Skip(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockMgoQueryDAL)(nil).Skip), n)
}

// Limit mocks base method
func (m *MockMgoQueryDAL) Limit(n int) mongodal.MgoQueryDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", n)
	ret0, _ := ret[0].(mongodal.MgoQueryDAL)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockMgoQueryDALMockRecorder) Limit(n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockMgoQueryDAL)(nil).Limit), n)
}

// Sort mocks base method
func (m *MockMgoQueryDAL) Sort(fields ...string) mongodal.MgoQueryDAL {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sort", varargs...)
	ret0, _ := ret[0].(mongodal.MgoQueryDAL)
	return ret0
}

// Sort indicates an expected call of Sort
func (mr *MockMgoQueryDALMockRecorder) Sort(fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockMgoQueryDAL)(nil).Sort), fields...)
}

// Apply mocks base method
func (m *MockMgoQueryDAL) Apply(change mgo_v2.Change, result interface{}) (*mgo_v2.ChangeInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", change, result)
	ret0, _ := ret[0].(*mgo_v2.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply
func (mr *MockMgoQueryDALMockRecorder) Apply(change, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockMgoQueryDAL)(nil).Apply), change, result)
}

// MockMgoPipeDAL is a mock of MgoPipeDAL interface
type MockMgoPipeDAL struct {
	ctrl     *gomock.Controller
	recorder *MockMgoPipeDALMockRecorder
}

// MockMgoPipeDALMockRecorder is the mock recorder for MockMgoPipeDAL
type MockMgoPipeDALMockRecorder struct {
	mock *MockMgoPipeDAL
}

// NewMockMgoPipeDAL creates a new mock instance
func NewMockMgoPipeDAL(ctrl *gomock.Controller) *MockMgoPipeDAL {
	mock := &MockMgoPipeDAL{ctrl: ctrl}
	mock.recorder = &MockMgoPipeDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMgoPipeDAL) EXPECT() *MockMgoPipeDALMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockMgoPipeDAL) All(result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockMgoPipeDALMockRecorder) All(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockMgoPipeDAL)(nil).All), result)
}

// Iter mocks base method
func (m *MockMgoPipeDAL) Iter() mongodal.MgoIterDAL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(mongodal.MgoIterDAL)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockMgoPipeDALMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockMgoPipeDAL)(nil).Iter))
}

// One mocks base method
func (m *MockMgoPipeDAL) One(result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockMgoPipeDALMockRecorder) One(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockMgoPipeDAL)(nil).One), result)
}

// MockMgoIterDAL is a mock of MgoIterDAL interface
type MockMgoIterDAL struct {
	ctrl     *gomock.Controller
	recorder *MockMgoIterDALMockRecorder
}

// MockMgoIterDALMockRecorder is the mock recorder for MockMgoIterDAL
type MockMgoIterDALMockRecorder struct {
	mock *MockMgoIterDAL
}

// NewMockMgoIterDAL creates a new mock instance
func NewMockMgoIterDAL(ctrl *gomock.Controller) *MockMgoIterDAL {
	mock := &MockMgoIterDAL{ctrl: ctrl}
	mock.recorder = &MockMgoIterDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMgoIterDAL) EXPECT() *MockMgoIterDALMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockMgoIterDAL) All(result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockMgoIterDALMockRecorder) All(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockMgoIterDAL)(nil).All), result)
}

// Close mocks base method
func (m *MockMgoIterDAL) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMgoIterDALMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMgoIterDAL)(nil).Close))
}

// Err mocks base method
func (m *MockMgoIterDAL) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockMgoIterDALMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockMgoIterDAL)(nil).Err))
}

// Next mocks base method
func (m *MockMgoIterDAL) Next(result interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", result)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockMgoIterDALMockRecorder) Next(result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockMgoIterDAL)(nil).Next), result)
}

// Timeout mocks base method
func (m *MockMgoIterDAL) Timeout() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Timeout indicates an expected call of Timeout
func (mr *MockMgoIterDALMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockMgoIterDAL)(nil).Timeout))
}