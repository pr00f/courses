package spells

import "fmt"

func Solution() {
	sp1()
	sp2()
	sp3()
	sp4()
	sp5()
	sp6()
	sp7()
	sp8()
	sp9()
	sp10()
	sp11()
	sp20()
}

func sp1() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if j <= i {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println()
	}
}

func sp2() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if j != i {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println()
	}
}

func sp3() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if 24-i == j {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp4() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if i < 6 || j < 24-i+6 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp5() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if j == i*2 || j == ((i*2)+1) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp6() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if i < 10 || j < 10 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp7() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if i > 15 && j > 15 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp8() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if i == 0 || j == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp9() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {

			if j >= i+11 || (i >= 11 && j+11 <= i) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp10() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {

			if j > i && j < (i+1)*2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp11() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if i == 1 || j == 1 || j == 23 || i == 23 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func sp20() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if (j-i)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}
