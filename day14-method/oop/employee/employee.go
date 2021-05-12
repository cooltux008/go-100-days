package employee

import (
	"fmt"
)

type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

func New(firstName, lastName string, totalLeave, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeaves, leavesTaken}
	return e
}
func (e Employee) LeavesRemainning() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, e.TotalLeaves-e.LeavesTaken)
}
