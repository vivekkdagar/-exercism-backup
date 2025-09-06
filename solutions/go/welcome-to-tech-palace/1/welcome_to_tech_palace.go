package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var border = ""
    border += strings.Repeat("*", numStarsPerLine)
    border += "\n"
    border += welcomeMsg
    border += "\n"
    border += strings.Repeat("*", numStarsPerLine)
    return border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// First, remove the asterisks.
	noStars := strings.ReplaceAll(oldMsg, "*", "")
	
	// Then, trim any remaining leading or trailing whitespace.
	cleanedMsg := strings.TrimSpace(noStars)

	return cleanedMsg
}
