package main

import "testing"

func TestNewKSCEmployee(t *testing.T) {
	expLastName := "Kerman"
	var Jobs = [...]string{
		"Pilot",
		"Scientist",
		"Engineer",
	}
	expMinAge := 22
	expMaxAge := 39

	/* Get a random employee */
	employee := newKSCEmployee()

	/* Test if the has the best last name on Kerbin */
	if employee.LastName != expLastName {
		t.Errorf("LastName of Kerbal (%s) was not %s", employee.LastName, expLastName)
	}

	/* Test if the employee has a valid job that is available at KSC */
	validJob := false
	for _, job := range Jobs {
		if job == employee.Position {
			validJob = true
		}
	}
	if validJob != true {
		t.Errorf("Position of Kerbal (%s) is not a valid job at KSC", employee.Position)
	}

	/* Test if the employee has a valid age */
	if employee.Age < expMinAge || employee.Age > expMaxAge {
		t.Errorf("Age of Kerbal (%d) is not between %d and %d", employee.Age, expMinAge, expMaxAge)
	}
}
