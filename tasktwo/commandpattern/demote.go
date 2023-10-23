package commandpattern

type DemoteRoleCommand struct {
	role     Role
	newTitle string
}

func (c *DemoteRoleCommand) Execute() {
	c.role.Demote(c.newTitle)
}
