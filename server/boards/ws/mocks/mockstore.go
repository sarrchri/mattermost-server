// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-server/v6/server/boards/ws (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v6/server/boards/model"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetBlock mocks base method.
func (m *MockStore) GetBlock(arg0 string) (*model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(*model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockStoreMockRecorder) GetBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockStore)(nil).GetBlock), arg0)
}

// GetMembersForBoard mocks base method.
func (m *MockStore) GetMembersForBoard(arg0 string) ([]*model.BoardMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembersForBoard", arg0)
	ret0, _ := ret[0].([]*model.BoardMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembersForBoard indicates an expected call of GetMembersForBoard.
func (mr *MockStoreMockRecorder) GetMembersForBoard(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembersForBoard", reflect.TypeOf((*MockStore)(nil).GetMembersForBoard), arg0)
}
