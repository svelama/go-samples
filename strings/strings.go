package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	city := "London"
	description := "London is the capital and largest city of England"

	fmt.Println("Comparing Strings")
	fmt.Println("Contains:", strings.Contains(city, "on"))
	fmt.Println("ContainsAny:", strings.ContainsAny(city, "cd"))
	fmt.Println("ContainsRune:", strings.ContainsRune(city, 'o'))
	fmt.Println("EqualFold:", strings.EqualFold(city, "LONDON"))
	fmt.Println("HasPrefix:", strings.HasPrefix(city, "L"))
	fmt.Println("HasSuffix:", strings.HasSuffix(city, "on"))

	fmt.Printf("\n\n")
	fmt.Println("Converting String Case")
	fmt.Println("ToLower Case: e.g. London >> ", strings.ToLower("London"))
	fmt.Println("ToUpper Case: e.g. London >> ", strings.ToUpper("London"))
	fmt.Println("Title Case: e.g. London is the capital and largest city of England >> ", strings.Title(description))
	fmt.Println("ToTitle Case: e.g. London is the capital and largest city of England >> ", strings.ToTitle(description))

	fmt.Printf("\n\n")
	fmt.Println("Converting Character Case")
	for _, char := range "Cat" {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
		fmt.Println(string(char), "Lower case:", unicode.IsLower(char))
	}

	fmt.Printf("\n\n")
	fmt.Println("Inspecting strings")
	fmt.Println("Count:", strings.Count(description, "a"))
	fmt.Println("Index:", strings.Index(description, "c"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "c"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abc"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abc"))

	fmt.Printf("\n\n")
	fmt.Println("Splitting strings")
	fmt.Println("Count:", strings.Count(description, "a"))
	splits := strings.Split(description, " ")
	for _, s := range splits {
		fmt.Println("Split >>" + s + "<<")
	}

	splitsAfter := strings.SplitAfter(description, " ")
	for _, sa := range splitsAfter {
		fmt.Println("SplitAfter >>" + sa + "<<")
	}

	splitsN := strings.SplitN(description, " ", 2)
	for _, sn := range splitsN {
		fmt.Println("Split N >>" + sn + "<<")
	}

	fmt.Println("Splitting on Whitespace Characters")
	splitsF := strings.Fields(description)
	for _, x := range splitsF {
		fmt.Println("Field >>" + x + "<<")
	}

	fmt.Printf("\n\n")
	fmt.Println("Trimming Strings")
	fmt.Println("Trimming Character: ", strings.Trim(description, "Li "))
	fmt.Println("Trimming Whitespace: ", strings.TrimSpace(description))
	fmt.Println("Trimming Substrings: ", strings.TrimPrefix(description, "London"))

	fmt.Printf("\n\n")
	fmt.Println("Altering Strings")
	fmt.Println("Replace: ", strings.Replace(description, "England", "United Kingdom", 1))
	fmt.Println("Replace All: ", strings.ReplaceAll(description, "nd", "ND"))
	
	fmt.Printf("\n\n")
	fmt.Println("Building and Generating Strings")
	elements := strings.Fields(description)
	fmt.Println("Joined: ", strings.Join(elements, "--"))

	//TODO: String Builder & Regex 

}
