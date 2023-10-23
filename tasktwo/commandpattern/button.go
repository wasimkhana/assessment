package commandpattern

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}
