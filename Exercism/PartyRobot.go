// Party Robot is not a real Go package.
package partyrobot

import "fmt"

func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %[1]s!\nYou have been assigned to table %03d. Your table is %[4]s, exactly %.1f meters from here.\nYou will be sitting next to %[3]s.", name, table, neighbor, direction, distance)
}
