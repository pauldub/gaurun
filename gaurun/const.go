package gaurun

const (
	Version = "0.10.0"
)

const (
	PlatformIos = iota + 1
	PlatformAndroid
	PlatformTwilio
)

const (
	StatusAcceptedPush  = "accepted-push"
	StatusSucceededPush = "succeeded-push"
	StatusFailedPush    = "failed-push"
	StatusDisabledPush  = "disabled-push"
)
