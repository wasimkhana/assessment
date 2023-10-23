package commandpattern

import "fmt"

type Employee struct {
	ID      int
	name    string
	title   string
	company string
}

func (r *Employee) Promote(newTitle string) {
	r.title = newTitle
	fmt.Printf("Promoted %s to %s\n", r.name, r.title)
}

func (r *Employee) Demote(newTitle string) {
	r.title = newTitle
	fmt.Printf("Demoted %s to %s\n", r.name, r.title)
}
