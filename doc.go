// +build darwin

/*
Godoc はGoのパッケージドキュメント情報をQiitaで紹介するためのパッケージです。そのままサンプルとして使えます。
リポジトリは https://github.com/lufia/godoc-sample です。

How to read a document

ドキュメントを読むためにはgo docコマンドまたはgodocコマンドが使えます。

How to write a document

パッケージドキュメントはpackage句の直前に書く必要がありますが、
Build constraintsはpackageよりも前に書かなければなりません。
そのため、記述する順番としては、Build constraintsが先になります。

空行を入れると、別の段落として区切ることができます。

	インデントすると、ソースコードのような
	整形されたテキストも書けます。

Heading

アルファベットの大文字で始まり、句読点を含まない1行だけの段落があれば、
それはヘッダとして装飾されます。ただし、ヘッダを2つ以上続けることはできません。
次の行はヘッダになりますが、その次は同じルールにも関わらず普通の段落です。

This is a header

This is not a header

*/
package qiita

// PostStatus は記事の投稿状態を表現します。
type PostStatus int

// 記事の投稿状態。
const (
	StatusDraft   PostStatus = iota // 下書き
	StatusPublish                   // 公開済み
)

// DefaultTags はデフォルトで記事に設定されるタグの配列です。
var DefaultTags []string

// Article は1つの記事を表します。
type Article struct {
	// 記事のタイトル
	Title string

	// 記事本文
	Body string

	// 状態
	Status PostStatus // Draft or Publish
}

// NewArticle はタイトルをtitleに設定した新しい記事を作成します。
func NewArticle(title string) *Article {
	return &Article{Title: title}
}

// Save は、記事aの状態をデータベースに保存します。
func (a *Article) Save() error {
	// BUG(lufia): 保存機能は未実装です。
	// TODO(lufia): 実装する。
	return nil
}
