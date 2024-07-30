// TechPalace package is not a real Go package, just an exercise to practice Go.
package techpalace
import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars string = strings.Repeat("*", numStarsPerLine)
    return stars + "\n" + welcomeMsg + "\n" + stars
}

func CleanupMessage(oldMsg string) string {
    noStars := strings.ReplaceAll(oldMsg, "*", " ")
    return strings.TrimSpace(noStars)
}
