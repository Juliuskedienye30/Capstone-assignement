   Go Beginner Toolkit
   Title & Objective

Title: Getting Started with Go – A Beginner’s Toolkit

This project demonstrates how to install Go and build a simple interactive CLI application that accepts user input and performs basic operations.

  Technology Overview

Go (Golang) is a statically typed programming language developed by Google.
It is widely used for backend development, cloud computing, and command-line tools.

Real-world example: Tools like Docker and Kubernetes are built using Go.

   System Requirements
Linux OS
Go (v1.22 or higher)
Terminal
   Installation & Setup
sudo apt update
sudo apt install golang-go
go version
    Running the Project
go run main.go
   Minimal Working Example

This program:

Accepts user input (name)
Displays a welcome message
Performs a simple calculation

Code:

package main

import "fmt"

func main() {
    var name string
    var a, b int

    fmt.Println("=== Go Beginner Toolkit ===")

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Println("Welcome,", name, "🚀")

    fmt.Print("Enter first number: ")
    fmt.Scanln(&a)

    fmt.Print("Enter second number: ")
    fmt.Scanln(&b)

    fmt.Println("Sum is:", a+b)

    Full file structure
    go-toolkit/
├── go.mod
├── main.go
└── README.md
}
   AI Prompt Journal
Prompt 1:

How do I install Go on Linux?

Reflection:
The AI helped me install Go quickly without needing multiple tutorials.

Prompt 2:

How to create a simple Go program that takes user input?

Reflection:
I learned how to use fmt.Scanln to accept input from users.

Prompt 3:

Why am I getting "no such file or directory" in Go?

Reflection:
Helped me understand directory navigation and file structure.

     Common Errors & Fixes
Error:
stat main.go: no such file or directory

Cause: Running command in wrong folder

Fix:

cd go-toolkit
Error:
go: command not found

Fix:

sudo apt install golang-go
    Testing & Iteration

The program was tested using different inputs:

Name: Lucia → Correct output
Numbers: 5 and 3 → Output: 8
Numbers: 10 and 20 → Output: 30

This confirmed that:

Input handling works correctly
The logic is functioning as expected
   References
https://go.dev/doc/
https://go.dev/doc/tutorial/getting-started
   Step 3: Save it
CTRL + O → Enter
CTRL + X
    Step 4: Confirm
ls

You should now see:

go.mod
main.go
README.md
