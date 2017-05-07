package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-xz]+)ch")
    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("peach punch")) //Returns the matching strings
    fmt.Println(r.FindStringIndex("peach punch")) //Returns start, end indices

    fmt.Println(r.FindStringSubmatch("peach punch")) //Returns info for whole-pattern (i.e. p([a-z]+)ch) matches and submatches (i.e. [a-z]+)
    fmt.Println(r.FindStringSubmatchIndex("peach punch")) //Infer from contecxt

    fmt.Println(r.FindAllString("peach punch pinch", -1)) //The All variants find a specified number of matches
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2))
    fmt.Println(r.Match([]byte("peach"))) //Dont quite get this; specify input is list of bytes?

    r = regexp.MustCompile("p([a-z]+)ch") //Returns only one value
    fmt.Println(r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //Replace all matches

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)

    fmt.Println(string(out))

}
