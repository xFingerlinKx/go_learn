// выводит в одну строку аргументы, переданные в командной строке
package main

import (
	"fmt"
	"os"
)

func main() {
	//s := ""
	//var s string
	//var s = ,,n
	//var s string = ""
	s, sep := "", ""
	// в каждой итерации цикла range производит пару значений: индекс и значение элемента с этим индексом.
	// В данном примере мы не нуждаемся в индексе, но синтаксис цикла по диапазону требует,
	// чтобы, имея дело с элементом, мы работали и с индексом. Одно из решений заключается в том,
	// чтобы присваивать значение индекса временной переменной с очевидным именем наподобие temp
	// и игнорировать его. Однако Go не допускает наличия неиспользуемых локальных переменных,
	// так что этот способ приведет к ошибке компиляции. РЕШЕНИЕ - использовать _ (пустой идентификатор).
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}