package play_godoc_test

import (
	"fmt"
	play_godoc "play-godoc"
)

func Example() {
	play_godoc.InitTwoPeople()

	if play_godoc.Kazu.Sex.IsMan() {
		fmt.Println("カズは男です")
	} else {
		fmt.Println("カズは女です")
	}
	if play_godoc.Kazuko.Sex.IsWoman() {
		fmt.Println("カズコは女です")
	} else {
		fmt.Println("カズコは男です")
	}
	fmt.Printf("カズの挨拶: %s, カズコの挨拶: %s\n", play_godoc.Kazu.Greet(), play_godoc.Kazuko.Greet())
	fmt.Printf("カズとカズコの会話: %s", play_godoc.Kazu.Talk(play_godoc.Kazuko))
}

func ExampleGender_IsWoman() {
	w := play_godoc.WOMAN
	fmt.Println(w.IsWoman())
	fmt.Println(w.IsMan())
	// Output:
	// true
	// false
}

func ExampleGender_IsMan_true() {
	m := play_godoc.MAN
	fmt.Println(m.IsMan())
	// Output: true
}

func ExampleGender_IsMan_false() {
	w := play_godoc.WOMAN
	fmt.Println(w.IsMan())
	// Output: false
}

// もし、複数パターン用意したいのであれば、メソッド名や関数名の後ろに"_"をつけて、何かしらの言葉(キャメルケース)を入れる
// 上記の際に"_"の後に、何も言葉はないとExampleは表示されない
func ExampleJapanese_Talk_noOutput() {
	m := play_godoc.NewJapanes(20, "晴れですね", play_godoc.MAN)
	w := play_godoc.NewJapanes(20, "晴れですよ", play_godoc.WOMAN)
	fmt.Println(m.Talk(w))
}

func ExampleJapanese_Talk_output() {
	m := play_godoc.NewJapanes(20, "雨ですね", play_godoc.MAN)
	w := play_godoc.NewJapanes(20, "いいえ, 晴れですよ", play_godoc.WOMAN)
	fmt.Println(m.Talk(w))
	// Output:
	// Me: 雨ですね, Oponent: こんにちは, いいえ, 晴れですよ
}

func ExampleJapanese_Greet() {
	m := play_godoc.NewJapanes(20, "雨ですね", play_godoc.MAN)
	fmt.Println(m.Greet())
}
