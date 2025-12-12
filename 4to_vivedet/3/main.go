package main

func main() {
	s := "test"
	println(s[0])

	//runes[0] = 'R' // что будет?

	s = "R" + s[1:]
	println(s)

}
