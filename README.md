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
