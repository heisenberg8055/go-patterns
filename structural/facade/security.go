package facade

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode(securityCode int) error {
	if s.code != securityCode {
		return fmt.Errorf("wrong Security Code")
	}
	fmt.Printf("SecurityCode verified: %d", s.code)
	return nil
}
