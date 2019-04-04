package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
	var str string
	for i := 0; i < len(sliceToConcat); i++ {
		str += string(sliceToConcat[i])
		if i != len(sliceToConcat)-1 {
			str += "-"
		}
	}
	return str
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i := 0; i < len(sliceToEncrypt); i++ {
		sliceToEncrypt[i] += byte(ceaserCount)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
