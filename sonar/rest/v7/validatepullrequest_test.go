package v7

import (
	"errors"
	"testing"
)

func Test_ValidatePullRequest_CheckError(t *testing.T) {
	restApi := NewRestApi(&mockConnection{
		request: func(route string) (<-chan []byte, <-chan error) {
			chErr := make(chan error, 1)
			chErr <- errors.New("error-test")
			return nil, chErr
		},
	})

	err := restApi.ValidatePullRequest("project", "pull-request")
	if err == nil || err.Error() != "error-test" {
		t.Errorf("ValidatePullRequest() not returned expected error")
	}
}
