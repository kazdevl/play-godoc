package play_godoc

var (
	Kazu   *Japanese
	Kazuko *Japanese
)

// InitTwoPeople init Kazu and Kazuko
func InitTwoPeople() {
	Kazu = NewJapanes(20, "天気がいいですね!", MAN)
	Kazuko = NewJapanes(20, "最高の天気です！", WOMAN)
}
