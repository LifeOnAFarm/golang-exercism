/*
Name:			LifeOnAFarm
Date:			16/01/17
Program:		A working clock
 */
package clock

// Test Version
const testVersion = 4


type Clock int16

// This function converts the time to minutes
func New(hour, minute int)  Clock  {
	allMins := hour*60 + minute
	

	return Clock(allMins)
}

func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}
