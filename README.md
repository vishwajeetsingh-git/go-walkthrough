# go-walkthrough

I am an experienced programmer. I have worked in C,C++,Python. Through this repo I will try to learn go language.
I am following Go official tutorial for this. Every New thing I will learn. I will call it a Lesson. 

## Lesson 1

[Hello World Program](/hello/)
In this i learnt that the entry point of the go codebase is "package main" and this package should have a `main` function


## Lesson 2

[Creating a Module](/greetings/)
In this i learnt how to create a module.. Important thing to note the utility modules don't have `package main` on the top


We can import the module using `import "modulename"`
go compiler should be able to find the module 

We need to run 

```bash
go mod edit -replace modulename=/path/to/modulename
```

then run

```bash
go get modulename
```

## Lesson 3

In this i learnt that unlike C, Go functions can return multiple values in return statment. You can check the the 
`greetings` module `Hello` function now returning two values. string and error.

C also can return multiple values using struct or you can use pointers. But in case of Go, there is a straightforward approach. 

I also learnt there is a `nil` keyword in go. `nil` seems to be equivalent to C `NULL`. 

In this code we are returning `nil` as error when a valid string is provided to Hello function. 

## Lesson 4 

In go, there is a log package. We use this package for logging. 

`log.setPrefix` is used for setting prefix in log message.

in this library there is a function `log.Fatal` this accepts a string parameter which is printed to the console and then `os.exit(1)` is called which exits your application. 

Example Output:
```
greetings: empty name 
exit status 1
```

## Lesson 5 

I learnt how to define an array of strings. We already know that `:=` is used declaring and defining a variable. 

Given below is syntax to declare and initilize an array of strings. 

```go
greetings := []string{
    "Hello",
    "Welcome",
    "Hi"
} 

```

I learnt about a new module which is used for generating random numbers `math/rand`.

We can use fuction `func rand.Intn(n int) int` this accepts an non-zero positive integer and returns a random number <= n 

I just noticed one more thing, A single import statement can import multiple packages/modules.

For example:

```go
import (
    "fmt"
    "error"
    "math/rand"
)
```

## Lesson 6 

In this lesson I learnt that I was calling wrong this `array` in the previous lesson. Actually it is a `slice`. 

The key difference between an array and a slice is that an array has a fixed size which is part of its type, while a slice has dynamic size and can grow using append. You cannot resize an array, but you can append elements to a slice. 

Array Initilization:

```go 
a := [5]int{1,2,3,4,5}
```
Slice Initialization:

```go 
s := []int{1,2,3}
```

With var declaration:

Array:

```go 
var a [5]int
```

Slice:

```go
var s []int
```

I learnt about `map` in go. We can think of map equivalent to Python `dict`. map is used to store key-value pairs. 
The snippets below show different ways of creating and using maps. 

```go
messages := make(map[string]string) // this is correct syntax
// it initializes the map and returns a reference to it. The make function is used to create slices, maps, and channels in Go. When you use make to create a map, it allocates memory for the map and initializes it so that you can start adding key-value pairs to it immediately.
```

```go
messages := map[string]string{} // this is also correct syntax. It creates an empty map literal and assigns it to the variable messages. This syntax is more concise and achieves the same result as using make.

```

```go
var messages map[string]string // This declaration creates a nil map. It must be initialized before adding key-value pairs, otherwise writing to it causes a runtime panic.
```

if you don't want runtime panic while adding the value to last map. 
you need to initialize it using. 

```go 
messages = make(map[string]string)
```

## Lesson 7 

In this I learnt that you can write test for your go module in `modulename_test.go` file `_test.go` prefix tells go compiler to run the functions contained in this file to be executed when we run `go test`

## Lesson 8 

In go there is only one type of loop which is  `for` loop. The for loop syntax is quite similar to the `C for loop`

Example :

```go

for i := 0; i < 4; i++ {

}
```

Like C it also has three parts. 
initialization, condition and post statement 

Unlike C the paranthesis around the for loop components is not required. But the statement braces { } are required. 

The init and post statement in go for loop are optional. 

Example :

```go 

// for loop with no pre and post statements 
	j := 5
	for ; j > 0 ; {
		fmt.Printf("Countdown: %d\n", j)
		j--
	}

```

You can drop the semicolons from the previous snippet and it becomes equivalent to C's while loop

```go 
// for is go's while
	k := 0 
	for k < 5 {
		fmt.Printf("While loop iteration %d\n", k)
		k++
	}
```

You can drop the condition statement also. And the for loop becomes `forever` loop like 
while(1) or for(;;) in C

```go 
for {
		fmt.Println("This will run forever")
	}
```

## Lesson 9 

Like Go's for Go's if also doesn't need paranthesis surrounding the condition statement. But requires the braces. 

```go
if i % 2 == 0 {
    fmt.Println("even")
}
```

Like `for`, `if` statement can start with a short execution staement before condition statement.

```go 
if j := 6; j % 2 == 0 { 
    fmt.Println("even")
}
```

variable declared in short if statement are only accessible inside the if paranthesis. 

```go 

if i:=0; i < 5 {
    return 0
}

return i

```

The above snippet is invalid because i is being accessed out of if block. 

The `if` short statement variable can be accessed by else block if the `if` statement has an `else` block

```go 

if k := 7; k % 2 == 0 {
    fmt.Println("even : ", k)
} else {
    fmt.Println("odd : ", k)
}
```

## Lesson 10

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

## Lesson 11 

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

For example:

```go

package main

import "fmt"

func main() {

	defer fmt.Println("world")
	
	fmt.Println("hello")
}	

```

Output :

```bash
vsr@phoenix:~/go_tutorial/go-walkthrough/defer_in_go$ go run .
hello
world
```

### Stacking defers


Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

```go


package main

import "fmt"

func main() {

	defer fmt.Println("Counting done")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Counting...")
}	

```

Output:

```bash
vsr@phoenix:~/go_tutorial/go-walkthrough/defer_in_go$ go run .
Counting...
10
9
8
7
6
5
4
3
2
1
Counting done
```

## Lesson 12 

Pointers in Go 

Go has Pointers. A pointer holds memory address of another variable. 
The synax for declaration of Pointer in go is. 

```go
var p *TYPE 
```
where TYPE can be any data type, int , float , string , etc. 

Example:

```go 
var p *int
```

## Lesson 13 

Structs in Go.

A struct is a collection of fields. 

Given below is an example of Struct 

```go

type Student struct {
    name string 
    roll_no int
}

```

Structs Field are accessed usign dot. The following code snippet explains, 
How we can access/modify a struct field using dot.

```go 

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Employee struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(Vertex{1, 2})

	//struct fields are accessed using a dot.
	e := Employee{"Vishwajeet", 25}
	fmt.Println(e.Name)
	fmt.Println(e.Age)

	//change struct fields
	e.Age = 26
	fmt.Println(e.Age)
}


```

We can use pointers to struct. While accessing the member of struct using pointer we don't need `*`

```go

	// structs fields can be accessed through a struct pointer.
	// the struct pointer is automatically dereferenced.
	p := &e
	fmt.Println("Name using Pointer to Struct e :",p.Name)
	fmt.Println("Age using Pointer to Struct e :",p.Age)

```

## Lesson 14 

The type [n]T is an array of n values of type T.

The expression

```go
var a [10]int
```
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.