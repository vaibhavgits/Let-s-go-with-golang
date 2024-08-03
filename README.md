# golang


![golang](https://github.com/user-attachments/assets/adf44975-31ab-472e-b1aa-cf11366388b2)

## Introduction

Golang, also known as Go, is an open-source programming language created by Google.

- Case Sensitive - Go is a case-sensitive language. This means that the language distinguishes between uppercase and lowercase letters in identifiers such as variable names, function names, and type names. In Go, foo, Foo, and FOO would be considered three different identifiers.
- Garbage Collection: Go includes a garbage collector that helps manage memory automatically, reducing the complexity of manual memory management.
- Cross-Platform - Go is designed to be cross-platform, allowing developers to write code on one operating system and compile it to run on various other operating systems with minimal changes.
- Compiled Lang - Go is a compiled language. This means that Go source code is translated into machine code by a compiler before it is executed.
- Statically typed - Go is a statically typed language. This means that the type of a variable is known at compile time, and type checking is performed during compilation. In contrast, dynamically typed languages perform type checking at runtime.

## go mod init

The go.mod file in Go serves a similar purpose to the requirements.txt file in Python, but there are some differences in how they work and what they contain.

1. Dependency Versioning: Go modules help you specify and use specific versions of dependencies, avoiding issues with 2. incompatible or breaking changes.
2. Consistency: Ensures that anyone building your project uses the same versions of dependencies.
3. Isolation: Allows you to work with multiple projects with different dependency versions without conflicts.
4. This file is more structured and includes not only the list of dependencies but also metadata about the module, such as the module path and the version of Go being used. It looks like this:
   
    ![go mod 1](/01hello/image.png)

5. It maintains a go.sum file alongside go.mod to ensure exact versions and integrity of the dependencies.
6. You can automatically update and clean up dependencies using commands like go get and go mod tidy.
7. Example - go mod init example.com/greetings
   

## go.sum

- The go.sum file in Go is used alongside the go.mod file to ensure the integrity and reproducibility of your project's dependencies.
- When you first add a dependency (e.g., by running go get), Go records the version of the module in the go.mod file and the checksum of that version in the go.sum file.
- If someone else clones your project and runs go build, go test, or go mod tidy, Go will check the go.sum file to ensure the integrity of the dependencies being downloaded.
- Example of go.sum file - 
  
  ![go.sum](/01hello/image-1.png)
- Each line contains:
    - The module path and version.
    - A hash type (h1) and the checksum.
- You don’t need to manually edit the go.sum file.It’s automatically generated and updated by Go commands like go mod tidy, go get, go build, and go test.
  

## GOPATH

- GOPATH is an environment variable that specifies the location of your workspace, where your Go code, dependencies, and binaries are stored.
- Example Structure
  
  ![GOPATH Structure](/01hello/image-2.png)
  
    src:
    - Contains the source code of Go projects.
    - Each project resides in its own directory under src.
    
    pkg:
    - Contains compiled package objects, organized by operating system and architecture.
    - This helps in speeding up the build process by reusing compiled packages.

    bin:
    - Contains compiled binaries (executables).
    - When you run go install, the binaries are placed here.
  


## Transition to Go Modules

- Project Location - With Go modules, your projects no longer need to be located within the GOPATH/src directory. You can create and work on Go projects anywhere on your filesystem.
- Go modules manage dependencies on a per-project basis using go.mod and go.sum files.This eliminates the need for a centralized GOPATH/pkg directory for dependency management.
- Each project can have its own set of dependencies, versions, and configurations, independent of other projects.
- Example Structure
  
  ![Go Module Structure](/01hello/image-3.png)

- GOPATH still serves as a default module cache location and supports backward compatibility, but its role in Go development has been greatly minimized. 


## Lexer

Imagine you're reading a book. As you read, you naturally break down the text into words and sentences. You recognize where one word ends and another begins, and you understand the structure of the sentences. A lexer (short for lexical analyzer) does something similar for code.

- a lexer breaks down a program's source code into smaller pieces called tokens. Tokens are the building blocks of the code, such as keywords (like if, for), identifiers (variable names), operators (+, -), and symbols ({,}).
- The lexer reads the raw source code (which is just a long string of characters) and identifies meaningful units (tokens). This process is called lexical analysis.
- After the lexer breaks the code into tokens, it passes these tokens to the next stage called the parser. The parser then uses these tokens to understand the structure of the code and to check if it follows the language's grammar rules.
- Example
  
  ![Lexer Example](/01hello/image-4.png)

  When a lexer processes this code, it breaks it down into tokens like:

  - package
  - main
  - import
  - "fmt"
  - func
  - main
  - (
  - )
  - {
  - fmt
  - .
  - Println
  - (
  - "Hello, World!"
  - )
  - }
  

## Does semi-colon necessary in go?

In Go, semicolons (;) are technically necessary to terminate statements, but the Go compiler inserts them automatically at the end of each line during the compilation process. This means you typically don't need to write semicolons yourself except in a few specific cases.

If a line ends with a token that could legally end a statement (such as an identifier, a literal, or a keyword like break or return), the compiler inserts a semicolon at the end of the line.

## When to Use Semicolons Explicitly

- Multiple Statements on One Line: If you want to place multiple statements on a single line, you need to separate them with semicolons.

![Multi line statment](/01hello/image-5.png)
  
- For Loop Initialization and Post Statements: In for loops, semicolons are required to separate the initialization, condition, and post statements.
  
![For Loop](/01hello/image-6.png)

## Variable

- The default value of any vairable is 0. There is no garbage value. For string, it's ""(empty string).


### Numeric Types

- uint8       the set of all unsigned  8-bit integers (0 to 255)
- uint16      the set of all unsigned 16-bit integers (0 to 65535)
- uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
- uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
  
<br/>



- int8        the set of all signed  8-bit integers (-128 to 127)
- int16       the set of all signed 16-bit integers (-32768 to 32767)
- int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
- int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

<br/>

- float32     5 points after the decimal eg. 123.12345, the set of all IEEE-754 32-bit floating-point numbers
- float64     10 points after the decimal eg. 123.1234567890, the set of all IEEE-754 64-bit floating-point numbers

<br/>

- complex64   the set of all complex numbers with float32 real and imaginary parts
- complex128  the set of all complex numbers with float64 real and imaginary parts

<br/>

- byte        alias for uint8
- rune        alias for int32


### There is also a set of predeclared integer types with implementation-specific sizes:

- uint     either 32 or 64 bits
- int      same size as uint


### Walrus Operaor

Walrus operator(:=) can only be used inside a method and not for a global variable. When using wa;rus operator, you need not to add "var" keyword. 


### const

If your variable has first letter as capital, it is treated as public varible. eg. const LoginToken string = "jbsafkjbjsfs"

## comma OK

- There’s no Try Catch block in Go.
- It’s used in 4 different scenarios:
  - Test for error on a function return
  - Testing if a key-value item exists in a Go map
  - Testing if an interface variable is of a certain type
  - Testing if a channel is closed
  

## Input

- You can search for any package on pkg.go.dev


## Packages

- bufio - This package is used for taking input.
- os - This package is used to deal with programming related to reading and writing files.


## Conversion

### strconv

- This package is used for conversion of to and from string. 


### common error in goland

- You can't use input like this:- 

    numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

    Suppose, you entered 4. This will output and error saying:- 
    strconv.ParseFloat:  parsing "4\n":  invalid syntax

    You can see that "\n" is added automatically after the rating 4 is entered. 

- To avoid this error, strings package is used strings.TrimSpace() method is used.
    
    Example:

    numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

    TrimSpace returns a slice of the string s, with all leading and trailing white space removed



## Build

- You can build executables for Linux, Windows and Mac using golang. The command is - **go build**
- Go searches "GOOS" environment variable, suppose GOOS="darwin", then it will build executables for Mac. 

    Note - "darwin" is used for Mac.

- If you want to build executables for linux, you can run command - **GOOS="linux" go build**


## Time


### Formatting Time


- In Go, the layout string for formatting and parsing dates and times must follow a specific pattern based on a reference date, **Mon Jan _2 15:04:05 MST 2006**

- In Go, time formatting is done using a specific layout string. The string "Mon Jan _2 15:04:05 MST 2006" is a layout that represents a specific date and time format in Go, with no fractional seconds.



- If you want to modify the present time in human-readble or simple form, you must have to write the following:-s

    fmt.Println(presentTime.Format("01-02-2006 15:04:05 MST Mon"))


- What if we change the the string "Mon Jan _2 15:04:05 MST 2006"  to random date or time say "Wed Jan _5 17:19:05 IST 2023"

    This code will produce an error because the layout string does not follow the reference date pattern.



### Date

- In date, month is written in the form of time.\<name of month\>. e.g. time.November

## Memory Management

- As programs run they write objects to memory. At some point these objects should be removed when they’re not needed anymore. This process is called memory management.
- Memory allocation and Deallocation happens automatically


### new()

- Allocates Memory: When you use new(), you reserve space in memory for a variable.
- Does Not Initialize: The reserved space is filled with zero values (default values) for the type, but you haven’t set it to any specific value.

- Example
  - Imagine you are reserving a parking spot:
    - Reserve the spot: p := new(int) (You have a parking spot reserved for an integer).
    - Spot is empty: The spot is reserved, but no car is parked in it yet (it's just an empty spot).
    - Zeroed Storage: The spot is ready to hold a car (it’s clean and empty).
  

### make()
  
- Allocates Memory: When you use make(), you reserve space in memory for a specific type of data structure (slices, maps, or channels).
- Initializes it: The reserved space is also prepared with an initial setup (like setting up a parking lot with specific spots for cars).

- Example of make()
    - Imagine you are setting up a parking lot:
      - Reserve and set up spots: s := make([]int, 10) (You have a parking lot with 10 spots for integers).
      - Each spot is ready: Each spot is reserved and ready to park cars in it (it’s clean and set up to hold cars).



### What does this mean -  "You can allocate memory but not initialize it"?

- Think of allocating memory as reserving a space in your computer’s memory (like reserving a parking spot). When you allocate memory, you’re saying, "I need this much space for my data."

- Initializing memory means putting actual values into the reserved space (like parking a car in the reserved spot). When you initialize memory, you’re saying, "Here’s the data that will fill this space."
new() Function in Go


## GOGC Variable

The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. 


## Pointers

- A pointer is a variable that stores the memory address of another variable. 
- When you want to pass on actual value to avoid copying of a varible, use pointer.  

### Why Use Pointers?
Pointers are useful because they allow you to:

- Directly modify a variable: By pointing to a variable's address, you can change its value directly.
- Pass large structures efficiently: Passing a pointer is more efficient than passing large data structures because you only pass the address rather than copying the entire structure.

### Key Points
- Example
  - var a int = 42
  - var p *int = &a
- Pointer Declaration: var p *int means p is a pointer to an int OR p will be storing an integer value.
- Getting Address: &a gets the memory address of a.
- **& means referencing.**
- Dereferencing: *p accesses the value at the address p
- Default value of a pointer is **<nil>**


### So pointer refers to an address of another variable or it refers to the value of a variable accessing by its address?

- A pointer itself refers to the address of another variable, but you can use that address to access or modify the value stored at that address.
- When you create a pointer, it holds the memory address of another variable.
- To access or modify the value at the address the pointer is pointing to, you dereference the pointer using the * operator.
- Summary
  - Pointer: Refers to the address of a variable.
  - Dereferencing: Accesses or modifies the value at that address.


## Array

### len()

hen you declare an array with a specific size, the length of the array is fixed and does not depend on the number of values you initially provide. The len() function will always return the size of the array as declared.


## Slices

- A slice is a dynamically-sized. Unlike arrays, slices can grow and shrink, making them more versatile.

### Range

- Range are not inclusive in golang.
- Example fruitList[1:3], This will append the fruitList elements from 1st position(index 1) to 2nd position( index 2).

### Creating Slices
You can create slices in several ways:

- From an Array:

  - var arr = [5]int{1, 2, 3, 4, 5}
  - var slice []int = arr[1:4] // Creates a slice from elements 1 to 3 (index 1 to 3, not including 4)
  - fmt.Println(slice) // Output: [2 3 4]


- Using make:

  - slice := make([]int, 3) // Creates a slice of length 3, initialized with zero values
  - fmt.Println(slice) // Output: [0 0 0]
  - If you want to add more items, you can use append() method. **If you use slice[4]=431, it will raise an error!**


- Slice Literals:

  - slice := []int{1, 2, 3} // Creates a slice with initial values
  - fmt.Println(slice) // Output: [1 2 3]


### Length and Capacity
- Length (len): Number of elements in the slice.
- Capacity (cap): Number of elements in the underlying array, starting from the first element in the slice.



### Append Method
- This method is used for adding elements to the slices and can also be used in removing an element



