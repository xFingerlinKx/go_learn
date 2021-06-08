package main

/*
Структуры и интерфейсы

Несмотря на то, что вполне можно писать программы на Go используя только встроенные типы,
в какой-то момент это станет очень утомительным занятием. Вот пример — программа,
которая взаимодействует с фигурами:

package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}
func circleArea(x, y, r float64) float64 {
    return math.Pi * r*r
}
func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    var cx, cy, cr float64 = 0, 0, 5

    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
    fmt.Println(circleArea(cx, cy, cr))
}

Отслеживание всех переменных мешает нам понять, что делает программа, и наверняка приведет к ошибкам.

Структуры

С помощью структур эту программу можно сделать гораздо лучше.
Структура — это тип, содержащий именованные поля. Например, мы можем представить круг таким образом:

type Circle struct {
    x float64
    y float64
    r float64
}

Ключевое слово type вводит новый тип. За ним следует имя нового типа (Circle)
и ключевое слово struct, которое говорит, что мы определяем структуру и
список полей внутри фигурных скобок. Каждое поле имеет имя и тип.
Как и с функциями, мы можем объединять поля одного типа:

type Circle struct {
    x, y, r float64
}

Инициализация

Мы можем создать экземпляр нового типа Circle несколькими способами:

var c Circle

Подобно другим типами данных, будет создана локальная переменная типа Circle,
чьи поля по умолчанию будут равны нулю (0 для int, 0.0 для float, "" для string, nil для указателей, …).
Также, для создания экземпляра можно использовать функцию new.

c := new(Circle)

Это выделит память для всех полей, присвоит каждому из них нулевое значение и вернет указатель (*Circle).
Часто, при создании структуры мы хотим присвоить полям структуры какие-нибудь значения.
Существует два способа сделать это. Первый способ:
c := Circle{x: 0, y: 0, r: 5}

Второй способ — мы можем опустить имена полей, если мы знаем порядок в котором они определены:
c := Circle{0, 0, 5}

Поля

Получить доступ к полям можно с помощью оператора . (точка):

fmt.Println(c.x, c.y, c.r)
c.x = 10
c.y = 5

Давайте изменим функцию circleArea так, чтобы она использовала структуру Circle:

func circleArea(c Circle) float64 {
    return math.Pi * c.r*c.r
}

В функции main у нас будет:

c := Circle{0, 0, 5}
fmt.Println(circleArea(c))

Очень важно помнить о том, что аргументы в Go всегда копируются.
Если мы попытаемся изменить любое поле в функции circleArea,
оригинальная переменная не изменится. Именно поэтому мы будем писать функции так:
func circleArea(c *Circle) float64 {
    return math.Pi * c.r*c.r
}

И изменим main:
c := Circle{0, 0, 5}
fmt.Println(circleArea(&c))

Методы

Несмотря на то, что программа стала лучше, мы все еще можем значительно её улучшить, используя метод — функцию особого типа:

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

Между ключевым словом func и именем функции мы добавили «получателя». Получатель похож на параметр — у него есть имя и тип, но объявление функции таким способом позволяет нам вызывать функцию с помощью оператора .:

fmt.Println(c.area())

Это гораздо проще прочесть, нам не нужно использовать оператор & (Go автоматически предоставляет доступ к указателю на Circle для этого метода), и поскольку эта функция может быть использована только для Circle мы можем назвать её просто area.

Давайте сделаем то же самое с прямоугольником:

type Rectangle struct {
    x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

В main будет написано:

r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())

Встраиваемые типы

Обычно, поля структур представляют отношения принадлежности (включения).
Например, у Circle (круга) есть radius (радиус). Предположим, у нас есть структура Person (личность):

type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

И если мы хотим создать новую структуру Android, то можем сделать так:

type Android struct {
    Person Person
    Model string
}

Это будет работать, но мы можем захотеть создать другое отношение.
Сейчас у андроида «есть» личность, можем ли мы описать отношение андроид «является» личностью?
Go поддерживает подобные отношения с помощью встраиваемых типов, также называемых анонимными полями.
Выглядят они так:

type Android struct {
    Person
    Model string
}

Мы использовали тип (Person) и не написали его имя.
Объявленная таким способом структура доступна через имя типа:

a := new(Android)
a.Person.Talk()

Но мы также можем вызвать любой метод Person прямо из Android:

a := new(Android)
a.Talk()

Это отношение работает достаточно интуитивно: личности могут говорить,
андроид это личность, значит андроид может говорить.

Интерфейсы

Вы могли заметить, что названия методов для вычисления площади круга и прямоугольника совпадают.
Это было сделано не случайно. И в реальной жизни и в программировании отношения могут быть очень похожими.
В Go есть способ сделать эти случайные сходства явными с помощью типа называемого интерфейсом.
Пример интерфейса для фигуры (Shape):

type Shape interface {
    area() float64
}

Как и структуры, интерфейсы создаются с помощью ключевого слова type,
за которым следует имя интерфейса и ключевое слово interface.
Однако, вместо того, чтобы определять поля, мы определяем «множество методов».
Множество методов - это список методов, которые будут использоваться для «реализации» интерфейса.

В нашем случае у Rectangle и Circle есть метод area, который возвращает float64,
получается они оба реализуют интерфейс Shape. Само по себе это не очень полезно,
но мы можем использовать интерфейсы как аргументы в функциях:

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

Мы будем вызывать эту функцию так:

fmt.Println(totalArea(&c, &r))

Интерфейсы также могут быть использованы в качестве полей:

type MultiShape struct {
    shapes []Shape
}

Мы можем даже хранить в MultiShape данные Shape, определив в ней метод area:

func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

Теперь MultiShape может содержать Circle, Rectangle и даже другие MultiShape.
*/