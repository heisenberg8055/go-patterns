package factory

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "mp5" {
		return newMp5(), nil
	}
	return nil, fmt.Errorf("wrong gun")
}
