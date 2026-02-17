# go-walkthrough

I am an experienced programmer. I have worked in C,C++,Python. Through this repo I will try to learn go language.

[Hello World Program](/hello/)
In this i learnt that the entry point of the go codebase is "package main" and this package should have a `main` function 

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

---
Lesson 3


