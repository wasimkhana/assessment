package commandpattern

type Role interface {
	Promote(newTitle string)
	Demote(newTitle string)
}
