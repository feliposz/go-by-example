package examples

import (
	"fmt"
	"os"
	"sort"
)

func sortingExample() {
	str := []string{"microsoft", "google", "apple", "amazon", "facebook"}
	sort.Strings(str)
	fmt.Println("Sorted strings:", str)

	ints := []int{4, 5, 2, 4, 3, 2, 32, 4, 21, 57, 7}
	sort.Ints(ints)
	fmt.Println("Sorted ints", ints)
	fmt.Println("Sorted? ", sort.IntsAreSorted(ints))
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i int, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func customSortExample() {
	str := []string{"microsoft", "google", "apple", "amazon", "facebook"}
	sort.Sort(byLength(str))
	fmt.Println("Sorted by length: ", str)
}

func panicExample() {
	defer recoverExample()
	panic("Help, something is wrong!!!")
}

func recoverExample() {
	if r := recover(); r != nil {
		fmt.Println("recovered from:", r)
	}
}

func deferExample() {
	const filename = "tmp_file"

	// Create a file
	func() {
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Fprintln(f, "test")
	}()

	// Read its contents
	func() {
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		var content string
		fmt.Fscanln(f, &content)
		fmt.Println("File content:", content)
	}()

	// Clean up
	func() {
		err := os.Remove(filename)
		if err != nil {
			panic(err)
		}
	}()
}

// DataExamples contains examples of manipulating data
func DataExamples() {
	fmt.Println("\nData")
	fmt.Println("====")

	sortingExample()
	customSortExample()
	panicExample()
	deferExample()
}
