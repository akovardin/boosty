package auth

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
}

func (s *AuthTestSuite) SetupTest() {
}

func (s *AuthTestSuite) TestNew() {
	tests := map[string]struct {
		options []Option
		wantErr bool
	}{
		"success without options": {
			options: []Option{},
			wantErr: false,
		},
		"success with info": {
			options: []Option{
				WithInfo(Info{
					AccessToken:  "test_access_token",
					RefreshToken: "test_refresh_token",
					ExpiresAt:    1234567890,
					DeviceId:     "test_device_id",
				}),
			},
			wantErr: false,
		},
	}

	for name, test := range tests {
		s.T().Run(name, func(t *testing.T) {
			_, err := New(test.options...)

			if test.wantErr {
				s.Assert().Error(err)
			} else {
				s.Assert().NoError(err)
			}
		})
	}
}

func (s *AuthTestSuite) TestInfo() {
	expectedInfo := Info{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresAt:    1234567890,
		DeviceId:     "test_device_id",
	}

	auth, err := New(WithInfo(expectedInfo))
	s.Assert().NoError(err)

	info := auth.Info()
	s.Assert().Equal(expectedInfo, info)
}

func (s *AuthTestSuite) TestUpdate() {
	initialInfo := Info{
		AccessToken:  "initial_access_token",
		RefreshToken: "initial_refresh_token",
		ExpiresAt:    1234567890,
		DeviceId:     "initial_device_id",
	}

	updatedInfo := Info{
		AccessToken:  "updated_access_token",
		RefreshToken: "updated_refresh_token",
		ExpiresAt:    9876543210,
		DeviceId:     "updated_device_id",
	}

	callbackCalled := false
	callbackInfo := Info{}

	auth, err := New(
		WithInfo(initialInfo),
		WithInfoUpdateCallback(func(info Info) {
			callbackCalled = true
			callbackInfo = info
		}),
	)
	s.Assert().NoError(err)

	auth.Update(updatedInfo)

	info := auth.Info()
	s.Assert().Equal(updatedInfo, info)
	s.Assert().True(callbackCalled)
	s.Assert().Equal(updatedInfo, callbackInfo)
}

func (s *AuthTestSuite) TestBearerHeader() {
	expectedToken := "test_access_token"
	expectedHeader := "Bearer " + expectedToken

	auth, err := New(WithInfo(Info{AccessToken: expectedToken}))
	s.Assert().NoError(err)

	header := auth.BearerHeader()
	s.Assert().Equal(expectedHeader, header)
}

func (s *AuthTestSuite) TestRefreshToken() {
	expectedToken := "test_refresh_token"

	auth, err := New(WithInfo(Info{RefreshToken: expectedToken}))
	s.Assert().NoError(err)

	token := auth.RefreshToken()
	s.Assert().Equal(expectedToken, token)
}

func (s *AuthTestSuite) TestDeviceId() {
	expectedDeviceId := "test_device_id"

	auth, err := New(WithInfo(Info{DeviceId: expectedDeviceId}))
	s.Assert().NoError(err)

	deviceId := auth.DeviceId()
	s.Assert().Equal(expectedDeviceId, deviceId)
}

func (s *AuthTestSuite) TestSave() {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "auth_test")
	s.Assert().NoError(err)
	defer os.Remove(tmpFile.Name())

	expectedInfo := Info{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresAt:    1234567890,
		DeviceId:     "test_device_id",
	}

	// First create an auth instance with info
	auth, err := New(WithInfo(expectedInfo))
	s.Assert().NoError(err)

	// Set the file path directly for testing
	auth.file = tmpFile.Name()

	err = auth.Save()
	s.Assert().NoError(err)

	// Read the file and check its contents
	data, err := os.ReadFile(tmpFile.Name())
	s.Assert().NoError(err)

	var savedInfo Info
	err = json.Unmarshal(data, &savedInfo)
	s.Assert().NoError(err)

	s.Assert().Equal(expectedInfo, savedInfo)
}

func (s *AuthTestSuite) TestSaveWithoutFile() {
	auth, err := New()
	s.Assert().NoError(err)

	// Should not return an error when no file is specified
	err = auth.Save()
	s.Assert().NoError(err)
}

func (s *AuthTestSuite) TestWithFile() {
	// Create a temporary file with auth data
	tmpFile, err := os.CreateTemp("", "auth_test")
	s.Assert().NoError(err)
	defer os.Remove(tmpFile.Name())

	expectedInfo := Info{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresAt:    1234567890,
		DeviceId:     "test_device_id",
	}

	data, err := json.Marshal(expectedInfo)
	s.Assert().NoError(err)

	err = os.WriteFile(tmpFile.Name(), data, 0644)
	s.Assert().NoError(err)

	auth, err := New(WithFile(tmpFile.Name()))
	s.Assert().NoError(err)

	info := auth.Info()
	s.Assert().Equal(expectedInfo, info)
}

func (s *AuthTestSuite) TestWithFileError() {
	// Test with non-existent file
	_, err := New(WithFile("/non/existent/file"))
	s.Assert().Error(err)
}

func (s *AuthTestSuite) TestWithInfo() {
	expectedInfo := Info{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresAt:    1234567890,
		DeviceId:     "test_device_id",
	}

	auth, err := New(WithInfo(expectedInfo))
	s.Assert().NoError(err)

	info := auth.Info()
	s.Assert().Equal(expectedInfo, info)
}

func (s *AuthTestSuite) TestWithInfoUpdateCallback() {
	callbackCalled := false
	callbackInfo := Info{}

	expectedInfo := Info{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
		ExpiresAt:    1234567890,
		DeviceId:     "test_device_id",
	}

	auth, err := New(
		WithInfo(expectedInfo),
		WithInfoUpdateCallback(func(info Info) {
			callbackCalled = true
			callbackInfo = info
		}),
	)
	s.Assert().NoError(err)

	// Trigger the callback by updating the info
	auth.Update(expectedInfo)

	s.Assert().True(callbackCalled)
	s.Assert().Equal(expectedInfo, callbackInfo)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
