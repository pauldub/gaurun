package gaurun

import (
	"net"
	"net/http"
	"time"

	twilio "github.com/kevinburke/twilio-go"
	"github.com/mercari/gaurun/gcm"
)

func keepAliveInterval(keepAliveTimeout int) int {
	const minInterval = 30
	const maxInterval = 90
	if keepAliveTimeout <= minInterval {
		return keepAliveTimeout
	}
	result := keepAliveTimeout / 3
	if result < minInterval {
		return minInterval
	}
	if result > maxInterval {
		return maxInterval
	}
	return result
}

// InitGCMClient initializes GCMClient which is globally declared.
func InitGCMClient() error {
	// By default, use FCM endpoint. If UseFCM is explicitly disabled via configuration,
	// use GCM endpoint.
	url := gcm.FCMSendEndpoint
	if !ConfGaurun.Android.UseFCM {
		url = gcm.GCMSendEndpoint
	}

	var err error
	GCMClient, err = gcm.NewClient(url, ConfGaurun.Android.ApiKey)
	if err != nil {
		return err
	}

	transport := &http.Transport{
		MaxIdleConnsPerHost: ConfGaurun.Android.KeepAliveConns,
		Dial: (&net.Dialer{
			Timeout:   time.Duration(ConfGaurun.Android.Timeout) * time.Second,
			KeepAlive: time.Duration(keepAliveInterval(ConfGaurun.Android.KeepAliveTimeout)) * time.Second,
		}).Dial,
		IdleConnTimeout: time.Duration(ConfGaurun.Android.KeepAliveTimeout) * time.Second,
	}

	GCMClient.Http = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(ConfGaurun.Android.Timeout) * time.Second,
	}

	return nil
}

func InitTwilioClient() error {
	transport := &http.Transport{
		MaxIdleConnsPerHost: ConfGaurun.Twilio.KeepAliveConns,
		Dial: (&net.Dialer{
			Timeout:   time.Duration(ConfGaurun.Twilio.Timeout) * time.Second,
			KeepAlive: time.Duration(keepAliveInterval(ConfGaurun.Twilio.KeepAliveTimeout)) * time.Second,
		}).Dial,
		IdleConnTimeout: time.Duration(ConfGaurun.Twilio.KeepAliveTimeout) * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(ConfGaurun.Twilio.Timeout) * time.Second,
	}

	TwilioClient = twilio.NewClient(ConfGaurun.Twilio.AccountSID, ConfGaurun.Twilio.AuthToken, httpClient)

	return nil
}

func InitAPNSClient() error {
	var err error
	APNSClient, err = NewApnsClientHttp2(
		ConfGaurun.Ios.PemCertPath,
		ConfGaurun.Ios.PemKeyPath,
		ConfGaurun.Ios.PemKeyPassphrase,
	)
	if err != nil {
		return err
	}
	return nil
}
