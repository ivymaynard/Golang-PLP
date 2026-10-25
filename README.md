# Golang Tutorial - Personal Language Project

## History of Golang (Go)!

Go is a language designed by Robert Griesemer, Rob Pike, and Ken Thompson. It was developed by Google and publically annpunced in 2009. It is a general-purpouse programming language built to imporve productivity and asses some of the criticisims of other languages used at Google, while keeping the desired features. It mostly stemmed from issues the designres had with the language C++

### Uses of Go

Go is primarily used for web backends, deleopment automation tools, distributed systems, and building scalable cloud-native infrastructure. It has fast startup times which make it great for building scalable microservices, scalable APIS, web servers, and more. Go is one of the languages used at google for site realiability engineering and large scale data processing and is part of the software that runs Google Cloud. Uber also uses Go to handle tasks like matching riders with drivers. Netflix uses Go to handle heavy data processing demands.

## Getting Started with Go!

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

## Data Types

Go is a statically types programming language. This means that the type of variables is known at compile time. Each one is assigned a sata type, which determines it's size, memory, operations that can be performed on it, and the values it can hold.

The four categories of data types are:
1. Basic Type
2. Aggregate Type
3. Reference Type
4. Interface Type

### Basic Type:
The basic type, also refered to as primative type, includes numbers, strings, and booleans. These categories contan various data types listed below.

#### Numeric Types
In go, numeric types store various types of numbers. This includes standard integers, floats, and complex numbers.
- Integers
- Signed Integers: int, int8, int16, int32, int64
- Unsigned Integers: uint, uint8, uint16, uint32, uint64
- Floating-Point Numbers
- float32: 32-bit floating-point number
- float64: 64-bit floating-point number (double precision)
- Complex Numbers
- complex64: Complex number with float32 real and imaginary parts
- complex128: Complex number with float64 real and imaginary parts

```
var x int = 24
var y float32 = 4.44
var z complex64 = 6 + 7i
```

Numeric types follow standard rules for **artithmetic operators**. For example, addition, subtraction, multiplicaton, division, and remainder division can be performed on integers, whereas all but remainder division can be performed on floats.


#### Booleans
The boolean type can only hold true or false and is represented by "bool." You cannot perform arithmetic operations on booleans.

```
var isTrue bool = true
```

#### Strings
Strings are immutable sequences of characters.
```
var name string = "Jane"
```

The only arithmetic operator that can be used on strings is +, which concatenates separate, whole, strings into one.
```
string1 = "Hello! "
string2 = "World"
fmt.Println(string1 + string2) // Output: Hello! World

```

### Aggregate Type
These are also called derived or composite data types. These aggreate types are used to build more complex data structures.

#### Arrays
These are sequences of a particular data type. The size of the array is defined in the programming stage and cannot be changed. Arithmetic operations cannot be directly performed on array data types, but they can be performed on the values within the array, either through indexing or looping.
```
// array of 4 integers
var nums[4]int = [5]int{1, 2, 3,4}

// array of 3 strings
var names[3]string = [3]string{"Connor", "Mack", "Will"}

// accessing an element
ftm.Println(names[1]) // Output: Mack
```

#### Slices
Similar to arrays, slices are sequences of a particular data type, but they are more flexible as they are not of a fixed size. Arithmetic operations also cannot be directly performed on slices, only thier values.
```
var floats []float32 = []float32{4.10, 2.71, 2.97}
floats = append(floats, 8.70) // Adding a value to the array
```
#### Structs
A struct (abbreviation for structure) is used to create custom data types by grouping together variables of different data types. They are used to represent real-world entities with sets of properties. Arithmetic operations cannot be performed on the structs as a whole, but they can be used on the values of the properties stored within the struct.

This data type is similar to classes in object-oriented languages like Python. You can read more about Structs on [Geeks for Geeks](https://www.geeksforgeeks.org/go-language/structures-in-golang/).

```
type Person struct {
    name string
    age int
    weight float
    dogOwner bool
```

Defining and adding to a structure:
```
var p // initializing a blank struct

var p = Person{"Adriana", 21, 150, true} // initializing a variable of a struct
```

### Reference Type

#### Maps 
Also known as a dictionary, maps are unordered collections of key-value pairs. Each key is unique, and can correspond to one or many values (or arrays of values). These values can be retrived quickly using their associated keys. Arithmetic operations cannot be used on maps as a whole. They can be used on the values within the map, either in isolation, or through iteration over all the key-value pairs.

```
var salary map[string]int = make(map[string]int) // instantiates the map
salary["Sid"] = 8
salary["Leo"] = 14

fmt.Println(salary["Leo"]) //Output: 14
```

#### Pointers
These are variables used to store the memory address of another variable. The memory address of these pointers is always found in hexidecimal form.

## Reserved Words
There are 25 reserved words in Go. These words have predefined meaning that cannot be used as variable names or function names. There are four cetegories of key words.

**Declarations:** const, func, import, package, type, var. These allow you to import packages, organize code, and define variables, constants, types, and functions.

**Composite Types:** chan, interface, map, struct. These are used to declare data structures, method sets, or communication channels.

**Control Flow:** break, case, continue, default, else, fallthrough, for, goto, if, range, return, switch. These words establish and control loops and conditional statments.

**Concurrency:** defer, go, select. These handle asynchronus execution, cleanups, and milti-channel monitoring.

## Naming Requierments and Convention
An important feature of Go is that access control is built directly into the case of an identifier. An exported, or public, identifier begins with an uppercase while an unexported, or private, identifier starts with a lowercase.

Go also strictly enforces the use of either camelCase or PascalCase for identifiers. Intitials or acronyms must be consistent in their original case however. For example, UserID and HTTPClient are correct while UserId and HttpClient are incorrect. Don't use ALL_CAPS, ALLCAPS, or snake_case.

Conventionally, the smaller the scope the shorter the name. This means that local variables should have very short names, such as i and x. The same rules apply to functions. If you have a function that reads, instead of calling it reader, you should name it r. Also, get functions in Go should not use the Get prefix. For example user.GetName() is incorrect. Instead, use user.Name().

More nuanced information about naming conventions and requierments can be found in the [Go Documentation](https://go.dev/doc/effective_go).

## Binding
Identifiers are bound to their respective entities during compile time, link time, or run time.

**1. Static Binding**
 Most identifier binding in C++ is static, meaning it is bound before the program runs. Static binding happens in compile time or link time. Local variables and function calls are exampels of binding happening in compile time. Global variables and functions are bound during link time.

**2. Dynamic Binding**
This occurs during run time when the exact entity that an identifier refers to cannot be known until the program is running.


## Limitations of Go
Go is widley known to have fast compilation, exceptional performance, and simplistic design, as it was designed to improve upon some of the issues of C++; however, it's minimalism causes some limitations.

Becasue Go is statically and strongly typed, you cannot add variables of different types, including adding ints to floats.
```
var num int = 5
var dec float64 = 2.71
ftm.Println(num + dec) // results in an error
```

The way to work around this is by explicitly converting the integer to a float and then adding them.
```
var num int = 5
var dec float64 = 2.71

result := float64(num) + dec
ftm.Println(result) // output: 7.71
```

Thinking about the limitation above, what would happen if you tried to run the code below? Can you add strings to integers? What about other data types?
```
var word string = "Hi"
var num int = 5

ftm.Println(word + num)
```

## Data Types: Sample Code

To see an example of how all these data types interact, download this code [INSERT HERE] and try it for yourself! Just load it into a compiler and run it through your terminal as stated in the tutorial, and get started using Go.

## Conditionals, Logical Operators, and Short-Circut Evaluation

Conditional statements iin Go control the flow of program execution baed on boolean conditions. These conditions evaluate to true or false and are used in if/else/else if statements as well as switch statements, all of which will be covered in this section.

As you will see below, go requires curly braces {} around each code block attatched to a conditional statement, but parentheses () around conditions are not required.

### One-Condition if/else Statement

Conditional statements in Go use boolean expressions to dictate the flow of program control. They check the specified conditions and run code chunks depending on the evaluated value of the boolean statement. In general, the format of an if/else statement in Go looks like:

```
if condition {
  // code to be executed if the condition is true
} else {
  // code to be executed if the condition is false
}
```

Here is an example of a simple if/else statement being used to determine if a student passed or failed an exam. Notice how the else statement is used as a catch-all. Anything grade that does not evaluate to true, meaning anything not lower than a 51, will evaluate as false and the else statement will run.

```
grade := 95
if grade < 51 {
  fmt.Println("Try again next time...")
} else {
  fmt.Println("Congrats! You passed.")
}
```

You can also use an if statement without an else statement. If the conditions evaluates to true, the block will run. If it evaluates to false, the flow of programming control will continue on.

```
grade := 27
if grade > 0 {
  fmt.Println("You got at least one question right")
}
```

In some languages, there is the "dangling else" problem which occurs when the else statement can logically attatch to multiple if statements due to syntaxical errors. In Go, this is not an issue. Because of the strict formatting of code blocks being contained in curly braces {}, there is no ambiguity that can lead to the dangling else.

### if/else/else if Statements

If you have more than one condition you wish to check, you can use an if/else/else if statement. The condition in the if statement will be evaluated, and if it is not true, it will be checked against the else if statement. If that is not true, it will move on until it either is true or it reaches to else statement.

```
grade := 95
if grade > 90 {
  fmt.Println("You got an A!")
} else if grade > 80 {
  fmt.Println("That's a B!")
} else if grade > 70 {
  fmt.Println("C for you.")
} else if grade > 60 {
  fmt.Println("You earned a D.")
} else {  \\ true if grade is lower than 60
  fmt.Println("That's and F...")
}
```


### Multi-Condition if/else Statement
does go short circut? include example 

Instead of writing an one condition if/else/else if statement in go, you can combine conditions in one statement. There are a few ways to do this

#### The AND Operator: &&

A statement will evaluate to true if both conditions connected by the && are true. In this case, the code block will run if condition1 and condition2 equal true. If either condition1 or condition2 are false, the entire statement will be false. If both  condition1 and condition2 are false, the result is false.
```
if (condition1) && (condition2) {
......
}
```

Here is an example. If the student scored high enough on the exam AND it was a bio exam, a message is printed for the student. If either or both are false, no message will be printed.

```
grade := 95
subject := "Biology"
if grade > 90 && subject == "Biology" {
    fmt.Println("You will make a great scientist!")
}
```

This is where the idea of short-circut evaluation comes into play. Short-circut evaluation is when the righthand side of a condition is skipped because the lefthand side's evaluation determined the outcome of the whole statement. 

For the && operator in Go, we know that both conditions must be true on order for the whole statement to be true. If the first condition is false, there is no need for the other to be evaluated, and it will be skipped. This is short-circuting. If the first condition is true, however, the second must be evaluated to determine the outcome.

For example, this code would short-circut. Because the grade is lower than 90, the first condition fails and the second is not checked.
```
grade := 45
subject := "Biology"
if grade > 90 && subject == "Biology" {
    fmt.Println("You will make a great scientist!")
}
```

#### The OR Operator: ||

The OR operator combines conditions and evaluates to true if either the left or the right condition are true. If both are true, the statement is true. If condition1 is false and condition2 is true, the statement is true. If condition1 is false and condition2 is true, the statement is true. If both are false, the statement is false.

```
if (condition1) ||  (condition2) {
......
}
```

In this example, a message will print for the student if they scored higher than a 90 in either math or literature.
```
mathGrade := 45
litGrade := 97
if mathGrade > 90 || litGrade > 90 {
    fmt.Println("You have found at least one strong suit!")
}
```

Short-circut evaluation also applies to the OR operator. Since only one condition needs to be true for the statement to be true, if the right condition is true, the left will not be checked. This is shor-circuting. If the right is false, the left must be checked to see if it is true.

In this example, the literature grade won't be checked because the math grade, whcih is condition1, is high enough to evaluate as tru
```
mathGrade := 99
litGrade := 12
if mathGrade > 90 || litGrade > 90 {
    fmt.Println("You have found at least one strong suit!")
}
```


### Switch-Case Statements

Go supports something called a switch statement, which is a shorter way to write a sequence of if/else statements. It works by selecting one of many code blocks to be executed based on the value or type of an expression. t is similar to switch statements in C++, java, and PHP, except it does not require break statements since it only evaluates the matched case.

The syntax of a switch statement in Go looks like:

```
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default: // optional
   // code block
}
```

The expression in the switch statement is evaluated once. Then the value or type of the switch expression is compared with each case. If there is a match, the associated block of code will run. The default case is optional and it specifies what your program will do if there is no case match. It works similarly to an else statement in a series of if/else/else if statements.

This is an example of a value based switch statement. A message is printed out based on the mood of the user. Because the mood is 4, case 4 will be executed, and the associated message will run. There is no default statement, meaning that any mood value not between 1 and 5 will result in nothing being executed.
```
mood := 4

switch day {
case 1:
    fmt.Println("Sad day..")
case 2:
    fmt.Println("Not feeling great")
case 3:
    fmt.Println("Very average")
case 4:
    fmt.Println("Its a good day!")
case 5:
    fmt.Println("BEST DAY EVER!!")
}
```

Short-circut evaluation applies to switch statements as well. Because Go executes these statements in a top to bottom order, it will stop immediatley after a case is matched. In the last example, the statement short-circuted after case 4 was matched. case 5 never needed to be evaluated because a match had already been found.


## Resources
[Official Go tutorial](https://go.dev/tour/welcome/1)

[Go Development Page](https://go.dev/)

[Intro to Go - Geeks for Geeks](https://www.geeksforgeeks.org/go-language/go-programming-language-introduction/)

[Data Types in Go - Geeks for Geeks](geeksforgeeks.org/go-language/data-types-in-go/)

[Golang Data Types - Medium](https://jyos-sw.medium.com/golang-data-types-429ba314f10a)

[Keywords and Identifiers in Go - Go101](https://go101.org/article/keywords-and-identifiers.html)

[If Else Statments - Geeks for Geeks](https://www.w3schools.com/go/go_else_statement.php)

[Pointers in Golang - Geeks for Geeks](https://www.geeksforgeeks.org/go-language/pointers-in-golang/#google_vignette)

[Go if else Statement](https://www.w3schools.com/go/go_else_statement.php)

[If/Else - Go by Example](https://gobyexample.com/if-else)

[Combining Conditional Statements in Golang](https://www.geeksforgeeks.org/go-language/combining-conditional-statements-in-golang/)

[Go Switch Statement - W3 Schools](https://www.w3schools.com/go/go_switch.php)
