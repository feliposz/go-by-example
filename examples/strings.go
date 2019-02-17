package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func stringFunctionsExample() {

	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	fmt.Println()

	fmt.Println("Len: ", len("hello"))
	fmt.Println("Char:", "hello"[1])

}

func formatingExample() {

	type point struct {
		x, y int
	}

	p := point{1, 2}

	// prints an instance point struct.
	fmt.Printf("%v\n", p)

	// include the struct’s field names
	fmt.Printf("%+v\n", p)

	// prints a Go syntax representation of the value
	fmt.Printf("%#v\n", p)

	// type of a value
	fmt.Printf("%T\n", p)

	// Formatting booleans
	fmt.Printf("%t\n", true)

	// Use %d for standard, base-10 formatting.
	fmt.Printf("%d\n", 123)

	// binary representation.
	fmt.Printf("%b\n", 14)

	// prints the character corresponding to the given integer.
	fmt.Printf("%c\n", 33)

	// %x provides hex encoding.
	fmt.Printf("%x\n", 456)

	// For basic decimal formatting use %f.
	fmt.Printf("%f\n", 78.9)

	// scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// basic string printing use %s.
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.
	fmt.Printf("%q\n", "\"string\"")

	// %x renders the string in base-16
	fmt.Printf("%x\n", "hex this")

	// representation of a pointer
	fmt.Printf("%p\n", &p)

	// number right-justified and padded with spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// also for floats with precision
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// to left-justify, use the - flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// formatting strings

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// left-justify use the - flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

func regexpExample() {
	r, err := regexp.Compile("a.*z")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.MatchString("az"), r.MatchString("za"))
}

func jsonExample() {

	type contact struct {
		AreaCode    int `json:"area_code"`
		PhoneNumber int `json:"phone_number"`
	}

	type person struct {
		Name        string `json:"name"`
		Age         int    `json:"age"`
		ContactInfo []contact
	}

	i1 := contact{AreaCode: 99, PhoneNumber: 2345678}
	i2 := contact{AreaCode: 88, PhoneNumber: 8765432}
	p := person{Name: "Bob", Age: 30, ContactInfo: []contact{i1, i2}}

	out, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	var decoded person
	json.Unmarshal(out, &decoded)
	fmt.Printf("%#v\n", decoded)
}

// StringExamples contains examples of manipulating strings
func StringExamples() {
	fmt.Println("\nStrings")
	fmt.Println("=======")

	stringFunctionsExample()
	formatingExample()
	regexpExample()
	jsonExample()
}
