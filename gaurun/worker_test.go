package gaurun

import (
	"errors"
	"testing"

	"github.com/RobotsAndPencils/buford/push"
	"github.com/stretchr/testify/assert"
)

func TestIsExternalServerError(t *testing.T) {
	cases := []struct {
		Err      error
		Platform int
		Expected bool
	}{
		{push.ErrIdleTimeout, PlatformIos, true},
		{push.ErrShutdown, PlatformIos, true},
		{push.ErrInternalServerError, PlatformIos, true},
		{push.ErrServiceUnavailable, PlatformIos, true},
		{errors.New("no error"), PlatformIos, false},

		{errors.New("Unavailable"), PlatformAndroid, true},
		{errors.New("InternalServerError"), PlatformAndroid, true},
		{errors.New("Timeout"), PlatformAndroid, true},
		{errors.New("no error"), PlatformAndroid, false},

		{errors.New("no error"), 100 /* neither iOS nor Android */, false},
	}

	for _, c := range cases {
		actual := isExternalServerError(c.Err, c.Platform)
		assert.Equal(t, actual, c.Expected)
	}
}
