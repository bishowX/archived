package main

import (
	"fmt"
	"os"

	"github.com/bishowX/archived/html_parser"
)

type ArchiveLink struct {
	Url string
	Id  string
}

func main() {
	// link := ArchiveLink{
	// 	Url: "https://go.dev/",
	// 	Id:  "go.dev",
	// }
	// fmt.Printf("URL: %s\n", link.Url)

	// resp, err := http.Get(link.Url)
	// if err != nil {
	// 	fmt.Println("got error: ", err)
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("got error: ", err)
	// }
	// // fmt.Println("Got response: ", string(body))
	// //
	// _, err = os.Stat(link.Id)
	// if err != nil {
	// 	fmt.Println("got error reading stuff ", err)
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		fmt.Println("creating dir")
	// 		err = os.Mkdir(link.Id, os.ModePerm)
	// 		if err != nil {
	// 			fmt.Println("err creating dir ", err)
	// 		}
	// 	}
	// }

	// html_parser.Parse(resp.Body)
	file, err := os.Open("html_parser/1.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	html_parser.Tokenize(file)

	// err = os.WriteFile(fmt.Sprintf("%s/index.html", link.Id), body, os.ModePerm)
	// if err != nil {
	// 	fmt.Println("couldn't write to file: ", err)
	// }
	// fmt.Println("Response saved to resp.html")
}
