package printer

import "fmt"

func ListenMessage(prt string) {
	var top, bot string
	for i := 0; i < len(prt); i++ {
		top += "▀"
		bot += "▄"
	}
	fmt.Println("▛▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀" + top + "▀▜")
	fmt.Println("▌ Server listening on " + prt + " ▐")
	fmt.Println("▙▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄" + bot + "▄▟")
}
