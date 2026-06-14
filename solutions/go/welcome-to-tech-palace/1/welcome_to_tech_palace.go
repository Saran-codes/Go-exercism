package techpalace


import(

    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    customer = strings.ToUpper(customer)
	final := "Welcome to the Tech Palace, " + customer
    return final
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	n := numStarsPerLine

    var tmp string
    
    for ;n>0;n-- {
        tmp+="*"
    }
    tmp += "\n" + welcomeMsg + "\n"

    n = numStarsPerLine
    for ;n>0;n-- {
        tmp+="*"
    }
    return tmp
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg=strings.ReplaceAll(oldMsg,"*","")
    oldMsg=strings.TrimSpace(oldMsg)
    return oldMsg
}
