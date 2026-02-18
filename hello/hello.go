package main
import "fmt"
import "greetings"
import "log"

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Greet("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}