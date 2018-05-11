// +build darwin

/*
Godoc はGoのパッケージドキュメント情報をQiitaで紹介するためのパッケージです。そのままサンプルとして使えます。

Reading

ドキュメントを読むためにはgo docコマンドまたはgodocコマンドが使えます。

Writing

書くためには、

	整形されたもの
	ソースコードとか

テスト

*/
package qiita

// Doc はドキュメントを表現します。
type Doc struct {
	Title string
}

const (
	RoleAuthor = iota
	RoleContributor
)
