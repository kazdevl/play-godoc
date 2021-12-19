/*
GoDocの練習のために書いております。
長い文章を書くときは、doc.goに分離して書くことがおすすめです。
ここはoverviewとして表示されます。

How To Use

インデントを下げてテキストを記述すると、コードブロックが生成され、そこにテキストが記述されます。
	sample := "sample"
	fmt.Println(sample)
上記のインデントが下がっている部分がコードブロックです。

Header

Headerにするには、前後に空白行を挟みます。
また、Headerは日本語ではなく、英語でないとHeaderになりません。
ここに関してはgodocのソースコードを読む必要がありそうです。
Headerの後にHeaderを記述したい場合は、間に文章が必要です。

About paragraph

コード上のコメントでは
"この部分"は
段落を分けていますが、ドキュメントには反映されていません

段落を分けるためには、空白行を残す必要性があります。なので、段落を分けたいときは、このように、
空白行を残すようようにしましょう。

How to Link

このように、単にurlを記述すると遷移できる文字になります。

ref: http://localhost:6060/pkg/net

Reference

example: https://go.dev/blog/examples

書き方: https://go.dev/blog/godoc∂
*/
package play_godoc
