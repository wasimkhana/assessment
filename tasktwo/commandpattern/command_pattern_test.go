package commandpattern

import (
	"testing"
)

// I am declaring global variable to demonstrate command pattern
// with execution of invoker (Button) based on concrete commands
// and that include concrete receivers.
var empl = &Employee{ID: 1, name: "Elon", title: "Board Member", company: "Twitter"}

func TestPromoteEmployeeCommand(t *testing.T) {

	promoted_role := "CEO"
	promoteCommand := &PromoteRoleCommand{
		role:     empl,
		newTitle: promoted_role,
	}

	promoteCommand.Execute()
	promote_button := &Button{
		command: promoteCommand,
	}

	promote_button.Press()
	if empl.title != promoted_role {
		t.Errorf("Expected title to be 'Senior Assistant', but got '%s'", empl.title)
	}

}

func TestDemoteEmployeeCommand(t *testing.T) {

	demoted_role := "Board Member"
	demoteCommand := &DemoteRoleCommand{
		role:     empl,
		newTitle: demoted_role,
	}

	promote_button := &Button{
		command: demoteCommand,
	}

	promote_button.Press()
	if empl.title != demoted_role {
		t.Errorf("Expected title to be 'Assistant', but got '%s'", empl.title)
	}
}
