package template

import "fmt"

type Sms struct {
	Otp
}

func (s *Sms) getRandomOTP(len int) string {
	randonOTP := "1234"
	fmt.Printf("OTP Generated SMS: %s\n", randonOTP)
	return randonOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: Saved %s\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return fmt.Sprintf("SMS OTP %s", otp)
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: send notification %s", message)
	return nil
}
