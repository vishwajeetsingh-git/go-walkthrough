package main
import "fmt"
import "greetings"
import "log"

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	persons := []string{"Vishwajeet", "Vaishnavi", "Vansh"} 
	messages, err := greetings.Hellos(persons)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}