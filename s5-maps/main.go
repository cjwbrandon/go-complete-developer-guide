package main

import "fmt"

func main() {
	// Literal syntax -> map[{keyType}]{valueType}{{values}}
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#jasldfas",
		"white": "#jgklsadfs",
	}

	// var syntax
	// var colors map[string]string

	// Using make
	// colors := make(map[string]string)

	// Manipulating Maps -> adding or changing
	colors["white"] = "#ffdsfasdfas"
	colors["black"] = "#askdflas"
	// colors.white // Cannot work since keys in maps are typed colors.10 does not make sense

	fmt.Println(colors)

	// delete({map}, {key})
	delete(colors, "black") // Remove values

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
