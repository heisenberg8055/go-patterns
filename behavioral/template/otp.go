package template

type IOtp interface {
	getRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	IO IOtp
}

func (o *Otp) GetAndSendOTP(otpLength int) error {
	otp := o.IO.getRandomOTP(otpLength)
	o.IO.saveOTPCache(otp)
	message := o.IO.getMessage(otp)
	err := o.IO.sendNotification(message)
	if err != nil {
		panic(err)
	}
	return nil
}
