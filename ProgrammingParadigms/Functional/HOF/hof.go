package main

import "fmt"

func showName(f func(string), name string) {
	f(name)
}

func printInConsole(name string) {
	fmt.Println(name)
}

func getClicker() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	showName(printInConsole, "Venky")
	click1 := getClicker()
	click2 := getClicker()

	fmt.Println(click1())
	fmt.Println(click2())

	fmt.Println(click1())
	fmt.Println(click2())

}
