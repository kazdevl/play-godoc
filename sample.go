/*
以下のように、コメントアウトの最後の部分と、packageの間に空白があると、
コメントはgodocに反映されません。
*/

package play_godoc

import "fmt"

type Human interface {
	Talk(Human) string // Talk print out each words
	Greet() string     // Greet print out greet word
}

// Gender is a type for judging sex
type Gender int

// IsWoman check gender is woman
func (g *Gender) IsWoman() bool {
	return *g == WOMAN
}

// IsMan check gender is man
func (g *Gender) IsMan() bool {
	return *g == MAN
}

type Japanese struct {
	Age      int
	position int
	Word     string
	Sex      Gender
}

//  Talk print out each words
func (j *Japanese) Talk(oponent Human) string {
	return fmt.Sprintf("Me: %s, Oponent: %s", j.Word, oponent.Greet())
}

// Greet print out greet word
func (j *Japanese) Greet() string {
	return fmt.Sprintf("こんにちは, %s", j.Word)
}

const (
	WOMAN Gender = iota
	MAN
)

// NewJapanese Create Instance
func NewJapanes(age int, word string, gender Gender) *Japanese {
	return &Japanese{
		Age:  age,
		Word: word,
		Sex:  gender,
	}
}
