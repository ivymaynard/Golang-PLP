# Golang Tutorial - Personal Language Project

## History of Golang (Go)!

Go is a language designed by Robert Griesemer, Rob Pike, and Ken Thompson. It was developed by Google and publically annpunced in 2009. It is a general-purpouse programming language built to imporve productivity and asses some of the criticisims of other languages used at Google, while keeping the desired features. It mostly stemmed from issues the designres had with the language C++

### Uses of Go

Go is primarily used for web backends, deleopment automation tools, distributed systems, and building scalable cloud-native infrastructure. It has fast startup times which make it great for building scalable microservices, scalable APIS, web servers, and more. Go is one of the languages used at google for site realiability engineering and large scale data processing and is part of the software that runs Google Cloud. Uber also uses Go to handle tasks like matching riders with drivers. Netflix uses Go to handle heavy data processing demands.

### Resources
[Official Go tutorial](https://go.dev/tour/welcome/1)
[Go Development Page](https://go.dev/)
[Geeks for Geeks Intro to Go](https://www.geeksforgeeks.org/go-language/go-programming-language-introduction/)

## Getting Started with Go!
<details><summary></summary>

Before you start coding, you need to install Go. You can do so using the [Download and Install](https://go.dev/doc/install) instructions here.

### Hello World! Tutorial

**Step 1:**
cd to your home directory using a command prompt

On Windows:
```
cd %HOMEPATH%
```

On Mac or Linux:
```
cd
```

**Step 2:**
Create a hello directory for your Go cource code
```
mkdir hello
cd hello
```
You should now see in your terminal that you are working in the hello directory!

**Step 3:**
Enable dependency tracking

When importing packages contained in other modules to be used in your code, you must manage those dependencies through your code's module. This is defined by a go.mod file which tracks the modules that provide those packages.

Running this code will enabel dependency tracking by crearing a go.mod file. The name of the module is its path ie. github.com/mymodule. In this case, we will be using example/hello as our module path, but if you intend on publishing your module for others to use,the path must lead to a real location.
```
go mod init example/hello
go: creating new go.mod: module example/hello // this is the output you ger from running the first line of code in the terminal
```

**Step 4:**
Create a file hello.go in which to write your code. This must be done in a text editor such as VSCode, GoLand, or Vim. It is VERY important to pay attention to the location of your file. Make sure that hello.go is located in your recently created hello folder which also contais your go.mod file.

**Step 5:**
Paste in the Heelo, World! code to your hello.go file and save it.
```
package main // declares a main package

import "fmt"

func main() { // implements a main function. executes by default when you run a function
    fmt.Println("Hello, World!")
}
```
Notes:
- A package is a way to group functions, and it's made up of all the files in the same directory
- The ftm package contains funcctions to format text and notably includes a print function. This is a standard package that is automatically installed when you install Go.

**Step 6:**
Run your code!

Go back to the terminal and run the code below. You will see Hello, World! as the output.
```
go run .  // necessary command to get your code to run
```

This help feature will give you a list of all other commands that may be helpful to you when executing other Go scripts.
```
go help
```

</details>
  






