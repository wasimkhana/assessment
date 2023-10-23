package commandpattern

type PromoteRoleCommand struct {
	role     Role
	newTitle string
}

func (c *PromoteRoleCommand) Execute() {
	c.role.Promote(c.newTitle)
}
