package main

import "fmt"

//создаёт новый обьъект string величиной 100 символов,
// с сылкой на  целый объект justString в 1024 символа, который остаётся в буфере памяти,
//на протяжении всего срока жизни переменной justString.
// Потенциальная утечка памяти, при многократном вызове функции

//var justString string

//	func someFunc() {
//		v := createHugeString(1 << 10)
//		justString = v[:100]
//
// }
// /////////////////////////////////////////////
func createHugeString(size int) string {
	runes := make([]rune, size)
	for i := 0; i < size; i++ {
		runes[i] = 'а'
	}
	return string(runes)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	if len(runes) > 100 {
		return string(runes[:100])
	}
	return v
}

func main() {
	justString := someFunc()
	fmt.Println("Длина результата:", len([]rune(justString)))
}
