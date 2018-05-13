package qiita_test

import (
	"fmt"
	"log"

	"github.com/lufia/godoc-sample"
)

func Example() {
	a := qiita.NewArticle("テスト")
	fmt.Println(a.Title)
	// Output: テスト
}

func Example_other() {
	a := qiita.NewArticle("テスト")
	a.Body = "サンプル"
	fmt.Println(a.Body)
	// Output: サンプル
}

func ExampleNewArticle() {
	a := qiita.NewArticle("テスト")
	fmt.Println(a.Status)
	// Output: 0
}

func ExampleNewArticle_otherStatus() {
	a := qiita.NewArticle("テスト")
	a.Status = qiita.StatusPublish
	fmt.Println(a.Status)
	// Output: 1
}

func ExampleArticle_Save() {
	a := qiita.NewArticle("テスト")
	a.Save()
}

func ExampleArticle_Save_errorHandling() {
	a := qiita.NewArticle("エラー")
	if err := a.Save(); err != nil {
		log.Fatalln(err)
	}
}
