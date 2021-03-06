package main

import (
	"fmt"
	"time"
)

/*
Эта программа выводит «from 1» каждые 2 секунды и «from 2» каждые 3 секунды.
Оператор select выбирает первый готовый канал, и получает сообщение из него,
или же передает сообщение через него. Когда готовы несколько каналов,
получение сообщения происходит из случайно выбранного готового канала.
Если же ни один из каналов не готов, оператор блокирует ход программы до тех пор,
пока какой-либо из каналов будет готов к отправке или получению.
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
