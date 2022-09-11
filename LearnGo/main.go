package main

func main() {
	// First Program Golang
	// fmt.Println("Hello, I am programmer!")

	// Data Type Golang
	// var isLogin bool = false
	// fmt.Println(isLogin)
	// var prime int = 7
	// fmt.Println(prime)
	// var decimal float64 = 45.6
	// fmt.Println(decimal)
	// var hello string = "Hello World"
	// fmt.Println(hello)
	// const pi = 3.14
	// fmt.Println(pi)

	// Operator dan Operand Golang
	// x := 1 + 2
	// fmt.Println(x)
	// Expression
	// a := 5
	// b := 6
	// c := a + b
	// fmt.Println(c)
	// al := 10
	// ti := 15
	// L := (al * ti) / 2
	// fmt.Println(L)
	// String Operation
	// helloworld := "hello" + " " + "world!!!"
	// fmt.Println(helloworld)

	// Branching (IF, ELSE IF, ELSE)
	// hour := 20
	// if hour < 12 {
	// 	fmt.Println("Selamat Pagi!!!")
	// } else if hour < 18 {
	// 	fmt.Println("Selamat Sore!!!")
	// } else {
	// 	fmt.Println("Selamat Malam!!!")
	// }
	// Switch
	// var today int = 3
	// switch today {
	// case 1:
	// 	fmt.Println("Today is Monday")
	// case 2:
	// 	fmt.Println("Today is Tuesday")
	// case 3:
	// 	fmt.Println("Today is Wednesday")
	// default:
	// 	fmt.Println("Invalid date")
	// }

	// Looping
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// Looping with continue and break
	// for i := 0; i < 5; i++ {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	if i > 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// Looping Over String
	// sentence := "Hello"
	// for i := 0; i < len(sentence); i++ {
	// 	fmt.Printf(string(sentence[i]) + "-")
	// }
	// fmt.Println("     ----->>>")
	// for pos, char := range sentence {
	// 	fmt.Printf("Character %c start at byte position %d\n", char, pos)
	// }

	// var number []int = make([]int, 4)
	// for i := 0; i <= 5; i++ {
	// 	number = append(number, i)
	// }
	// var capacity int = cap(number)
	// fmt.Println(capacity)

	// Array
	// var primes [5]int
	// var countries [5]string
	// fmt.Println(reflect.ValueOf(primes).Kind())
	// fmt.Println(reflect.ValueOf(countries).Kind())
	// x := [5]int{1, 3, 5, 7, 9}
	// var y [5]int = [5]int{0, 2, 4}
	// fmt.Println(x)
	// fmt.Println(y)
	// primes := [5]int{2, 3, 5}
	// Looping cara-1
	// for index := 0; index < len(primes); index++ {
	// 	fmt.Println(primes[index])
	// }
	// Looping cara-2
	// for i, e := range primes {
	// 	fmt.Println(i, "->", e)
	// }
	// Looping cara-3
	// i := 0
	// for range primes {
	// 	fmt.Println(primes[i])
	// 	i++
	// }

	// Slice
	// slice A = {1,2,3,4,5} len = 5 cap = 5
	// ditambah 1 dibelakang slice A = {1,2,3,4,5,6} len = 6 cap = 10(kelipatan 2)
	// var even_numbers []int
	// var odd_numbers []int{1,3,5,7,9}
	// numbers := []int{1,2,3,4,5}
	// // penggunaan keyword function make
	// var primes = make([]int,5,10)
	// append & copy
	// var colors = []string{"red", "green", "yellow"}
	// colors = append(colors, "purple")
	// copied_colors := make([]string, 10)
	// copy(copied_colors, colors)
	// fmt.Println(copied_colors)

}
