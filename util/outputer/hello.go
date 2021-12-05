/*
Outpterパッケージは挨拶オブジェクトを提供するパッケージです。

How To Use

以下のように使います
	import ~/outpter

	user := "UserA"
	agree := outputer.Hello(user)
	fmt.Printf("content: %+v\n", agree)
-*/

package outputer

// Agree represent agree
type Agree struct {
	Word string
	Action Action
}

// Action represent agree action
type Action int

// enum for action
const (
	HandUp Action = iota
	HandShake
	Hug
)

// Hello is function for agree
func Hello(to string) Agree {
	a := Agree{Word: "Hello "+to, Action: Hug}
	return a
}