package airportrobot
import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type GermanGreeter struct {
	    
}

func (g GermanGreeter) LanguageName() string{
	return "German"
}

func (g GermanGreeter) Greet(name string) string{
	return fmt.Sprintf("Hallo %s!", name)
}

type Italian struct {
	    
}

func (g Italian) LanguageName() string{
	return "Italian"
}

func (g Italian) Greet(name string) string{
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {
	    
}

func (g Portuguese) LanguageName() string{
	return "Portuguese"
}

func (g Portuguese) Greet(name string) string{
	return fmt.Sprintf("Olá %s!", name)
}

func SayHello(name string, g Greeter)string {
    return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}