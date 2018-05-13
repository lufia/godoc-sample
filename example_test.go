package qiita_test

import (
	"fmt"

	"github.com/lufia/godoc-sample"
)

const defaultTitle = "untitled"

func Example_wholeFileExample() {
	a := qiita.NewArticle(defaultTitle)
	fmt.Println(a.Title)
	// Output: untitled
}
