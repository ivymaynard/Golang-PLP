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


## Resources
[Official Go tutorial](https://go.dev/tour/welcome/1)

[Go Development Page](https://go.dev/)

[Intro to Go - Geeks for Geeks](https://www.geeksforgeeks.org/go-language/go-programming-language-introduction/)

[Data Types in Go - Geeks for Geeks](geeksforgeeks.org/go-language/data-types-in-go/)

[Golang Data Types - Medium](https://jyos-sw.medium.com/golang-data-types-429ba314f10a)

[Keywords and Identifiers in Go - Go101](https://go101.org/article/keywords-and-identifiers.html)
