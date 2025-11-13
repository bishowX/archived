package html_parser

import (
	"fmt"
	"io"
)

type Parsed struct {
	LinkTags []string
}

func Parse(r io.Reader) {
	content, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("Error reading content:", err)
	}

	// smth := "<link"
	fmt.Println(string(content))

}
