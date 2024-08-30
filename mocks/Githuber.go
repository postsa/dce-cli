// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Githuber is an autogenerated mock type for the Githuber type
type Githuber struct {
	mock.Mock
}

// DownloadGithubReleaseAsset provides a mock function with given fields: assetName, dceVersion
func (_m *Githuber) DownloadGithubReleaseAsset(assetName string, dceVersion string) error {
	ret := _m.Called(assetName, dceVersion)

	if len(ret) == 0 {
		panic("no return value specified for DownloadGithubReleaseAsset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(assetName, dceVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewGithuber creates a new instance of Githuber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGithuber(t interface {
	mock.TestingT
	Cleanup(func())
}) *Githuber {
	mock := &Githuber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
