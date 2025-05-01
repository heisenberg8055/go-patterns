package main

import (
	abs "github.com/heisenberg8055/go-patterns/behavioral/template"
)

func main() {
	smsOtp := &abs.Sms{}

	o := abs.Otp{IO: smsOtp}

	o.GetAndSendOTP(5)

	emailOtp := &abs.Email{}

	o = abs.Otp{IO: emailOtp}

	o.GetAndSendOTP(3)

}
