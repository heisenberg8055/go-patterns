package main

import (
	abs "github.com/heisenberg8055/go-patterns/behavioral/chainOfResponsibility"
)

func main() {
	cashier := &abs.Cashier{}

	medical := &abs.Medical{}
	medical.SetNext(cashier)

	doctor := &abs.Doctor{}
	doctor.SetNext(medical)

	reception := &abs.Reception{}
	reception.SetNext(doctor)

	patient := &abs.Patient{Name: "yes"}

	reception.Execute(patient)
}
