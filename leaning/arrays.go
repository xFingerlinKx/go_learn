package main

import "fmt"

// программа получает среднее значение всех чисел массива,
// состоящего из 5 чисел
func main() {
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0

	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	//var total float64 = 0
	//for i := 0; i < 5; i ++ {
	//	total += x[i]
	//}

	//fmt.Println(total / 5)
}

// var total float64 = 0
// for _, value := range x {
//    total += value
// }
// fmt.Println(total / float64(len(x)))

// Одиночный символ подчеркивания _ используется,
// чтобы сказать компилятору, что переменная нам не нужна
// (в данном случае нам не нужна переменная итератора).
// value будет тем же самым что и x[i]. Мы использовали ключевое слово range
// перед переменной, по которой мы хотим пройтись циклом.
//
// А еще в Go есть короткая запись для создания массивов:
//
// x := [5]float64{ 98, 93, 77, 82, 83 }
//
// Указывать тип не обязательно — Go сам может его выяснить по содержимому массива.
