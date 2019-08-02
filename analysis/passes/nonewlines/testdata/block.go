package testdata

import (
	"fmt"
)

func BlockFirstEmptyLine() { // want "block statements should not start or end with empty lines"

	// do something here
	fmt.Println("something")
	fmt.Println("something")
}

func BlockMultiLineFirstEmptyLine(
	a string,
) { // want "block statements should not start or end with empty lines"

	// do something here
	fmt.Println("something")
	fmt.Println("something")
}

func BlockLastEmptyLine() { // want "block statements should not start or end with empty lines"
	// do something here
	fmt.Println("something")
	fmt.Println("something")

}

func BlockBothFirstAndLastEmptyLine() { // want "block statements should not start or end with empty lines"

	// do something here
	fmt.Println("something")
	fmt.Println("something")

}

func BlockBothMultiLineFirstAndLastEmptyLine(
	a string,
) { // want "block statements should not start or end with empty lines"

	// do something here
	fmt.Println("something")
	fmt.Println("something")

}

func BlockEmpty() {

}

func BlockOneLine() { fmt.Println("something") }

func BlockInline() {
	{ // want "block statements should not start or end with empty lines"

		// do something here
		fmt.Println("something")
		fmt.Println("something")
	}

	{ // want "block statements should not start or end with empty lines"
		// do something here
		fmt.Println("something")
		fmt.Println("something")

	}

	{ // want "block statements should not start or end with empty lines"

		// do something here
		fmt.Println("something")
		fmt.Println("something")

	}

	{
		// do something here
		fmt.Println("something")
		fmt.Println("something")
	}
}
