/*
Outpterパッケージは挨拶オブジェクトを提供するパッケージです。

How To Use

	import ~/outpter

	user := "UserA"
	agree := outputer.Hello(user)
	fmt.Printf("content: %+v\n", agree)


*/

package outputer

type Agree struct {
	Word string
	Action Action
}

type Action int
const (
	HandUp Action = iota
	HandShake
	Hug
)

func Hello(to string) Agree {
	a := Agree{Word: "Hello "+to, Action: Hug}
	return a
}