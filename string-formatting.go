package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

var print = fmt.Printf

func main() {
    p := point{1, 2}
    print("%v\n", p) //Prints an instance
    print("%+v\n", p) //Prints the field names
    print("%#v\n", p) //Prints the Go syntax representation
    print("%T\n", p) //Prints the type of a value
    print("%t\n", true) //Prints a boolean
    print("%d\n", 123) //Prints integer in decimal form
    print("%b\n", 14) //Prints integer in binary
    print("%c\n", 33) //Prints the character representation
    print("%x\n", 456) //Prints the integer in hex
    print("%f\n", 78.9) //Prints the float in decimal form
    print("%e\n", 123400000.0) //Scientific notation
    print("%E\n", 123400000.0) //Slightly different scientific notation
    print("%s\n", "\"string\"") //Basic string printing
    print("%q\n", "\"string\"") //Double-quoting
    print("%x\n", "hex this") //Hex conversion - two chars/byte of input
    print("%p\n", &p) //Pointer representation
    print("|%6d|%6d\n", 12, 345) //Specifies the width of the integer.
                               //Right-justifies and pads with spaces
    print("|%6.2f|%6.2f\n", 1.2, 3.45) //Specifies width and decimal precision
    print("|%-6s|%-6s\n", "foo", "bar")
    s := fmt.Sprintf("a %s", "string") //Returns formatted string
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "an %s\n", "error") //Print to io.Writers
}
