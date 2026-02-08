package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan string)

	go TakeInterview(ch)

	for text := range ch {
		fmt.Println("Человек сообщил следующее: ", text)
	}

}

func TakeInterview(ch chan string) {
	timeSlice := []int{300, 400, 500, 600, 700}
	timeIndex := rand.Intn(len(timeSlice))
	randLetters := "asdfqwrexvbaserqwetasdf2345kljfaslha"
	N := rand.Intn(30)
	for i := range N {
		if timeIndex+i < len(timeSlice)-1 {
			timeIndex += i
		}
		time.Sleep(time.Duration(timeSlice[timeIndex]) * time.Millisecond)

		wordSlice := make([]byte, 5)
		for i := range 5 {
			letterIndex := rand.Intn(len(randLetters) - 1)
			wordSlice[i] = randLetters[letterIndex]
		}
		ch <- string(wordSlice)
	}
	close(ch)
}
