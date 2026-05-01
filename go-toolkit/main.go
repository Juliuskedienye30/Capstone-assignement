package main // Defines the main package. Every executable Go program must have a main package.

import "fmt" // Imports the fmt package, which is used for input and output (printing to the console).

func main() { // The main function is the entry point of the program. Execution starts here.

    var name string // Declares a variable called 'name' of type string to store user input.

    fmt.Println("=== Go Beginner Toolkit ===") 
    // Prints a line of text to the console with a newline at the end.

    fmt.Print("Enter your name: ") 
    // Prints text to the console WITHOUT a newline, so the user can type on the same line.

    fmt.Scanln(&name) 
    // Reads user input from the keyboard and stores it in the 'name' variable.
    // The '&' symbol means we are passing the memory address of 'name'.

    fmt.Println("Welcome,", name, "🚀") 
    // Prints a welcome message including the user's name.

    fmt.Println("This is your first Go CLI app!") 
    // Prints a final message to the user.

} 