// dup1 выводит текст каждой строки, которая появляется в
// стандартном вводе более одного раза, а также количество ее появлений,
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Отображение (map) содержит набор пар "ключ-значение" и обеспечивает кон
	// стантное время выполнения операций хранения, извлечения или проверки наличия
	// элемента в множестве.
	// ключи представляют собой строки, а значения представлены типом int
	// Встроенная функция make создает новое пустое отображение
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Примечание: игнорируем потенциальные ошибки из input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
