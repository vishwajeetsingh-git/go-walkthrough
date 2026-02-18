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



