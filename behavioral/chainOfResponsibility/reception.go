package chainofresponsibility

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Registration Done")
		r.next.Execute(p)
	} else {
		fmt.Println("Reception Registering Patient")
		p.registrationDone = true
		r.next.Execute(p)
	}
}

func (r *Reception) SetNext(d Department) {
	r.next = d
}
