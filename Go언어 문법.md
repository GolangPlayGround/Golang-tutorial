## 배열과 슬라이스, 맵

배열은 다음과 같이 사용

```go
var fruits [2]string

fruits[1] = "apple"

```

축약표현도 가능 대괄호안에 4를 채우지 않더라도
자동으로 ... 대신 4를 채워줌

```go
numbers := [...]int{4,5,6,7}
```

원소의 갯수는 len 함수로 알수있음


`슬라이스`는 배열과 같지만 초기화 시 원소의 개수가 정해져있지 않음
원소의 개수가 선언 시 지정되어 있다면 배열이고 슬라이스는 원소 개수가 정해져있지 않음
나중에 append 등을 사용하여 동적으로 추가
[...]string은 배열이고 []string은 슬라이스

```go
alphabet := []string{"a","b"}
orders := make([int],0)
orders = append(orders, 100, 101)
```
슬라이스는 다음 명령을 통해 쉽게 일부분을 취할 수 있다. 이를 슬라이싱이라함

`슬라이스이름[시작인덱스:종료인덱스]` 시작인덱스는 그값을 포함하고 종료인덱스는 그값보다 작은값

```go
	alphabet := []string{"a", "b", "c", "d", "e"}
	a1 := alphabet[0:4] // a,b,c,d
	a2 := alphabet[0:1] // a
	a3 := alphabet[:]   // a,b,c,d,e
	a4 := alphabet[1:3] // b,c
```    

go언어에서는 배열과 슬라이스의 복사는 다르게 동작한다.

`배열을 복사하면 다른 배열이 생성되어 값들이 복사되지만 슬라이스는 그 주소를 가리키게 된다`

```go
array1 := [...]int{1,2,3}
array2 := array1
array2[0] = 5
fmt.Println(&array1[0]) // 0x140000aa018
fmt.Println(&array2[0]) // 0x140000aa030
fmt.Println(array1[0]) // 1


slice1 := []int{1,2,3}
slice2 := slice1
slice2[0] = 5
fmt.Println(&slice1[0]) // 0x140000aa048
fmt.Println(&slice2[0]) // 0x140000aa048
fmt.Println(slice1[0]) // 5
```


맵은 다음처럼 선언하여 사용

```go
fruits := make(map[string]int)
fruits["apple"] = 1000
fruits["banana"] = 3000

fruits := map[string]int{
    "apple" : 1000,
    "banana" : 3000,
}
```

맵에서 key에 해당하는 value를 가져올땐 보통 이렇게함.

```go
fruits := make(map[string]int)

orange, ok := fruits["orange"]
if !ok {
    fmt.Println("no orange")
} else {
    fmt.Println(orange)
}
```

맵의 루프는 다음과 같이 사용

```go
fruits := make(map[string]int)
fruits["apple"] = 1000
fruits["banana"] = 3000
for key, val := range fruits {
    fmt.Printf("%s : %d\n" , key, val)
}
```

---

## 함수

인자값들을 넘겨줄때는 값을 넘겨줄 수도 있고 (Pass by Value), 주소를 넘겨줄수도 있다 (Pass by Reference)

값을 넘겨줄때는 그값이 복사되며, 주소를 넘겨줄 때는 값이 복사되지 않고 원래의 값이 바뀔 수도 있다

---

## 에러와 가변인자함수

Go에서는 여러 가지 에러가 생길 수 있는 함수들은 리턴값에 에러를 포함하고 있다.

```go
func Atoi(s string) (int, error)
```

그래서 에러가 널(nil)인지 체크를 하고, 널이 아니라면 그에 따른 에러 처리를 해줘야 한다

```go
num, err := strconv.Atoi("hello")
if err != nil {
    panic(err)
}
```

panic은 프로그램을 즉시 중단시키고 에러 메시지를 출력.

에러 함수는 아래 패키지를 사용하는 것을 권장. 이유는 에러의 wrapping과 추적이 되기 때문

`github.com/pkg/errors`

```go
func fn() error {
    e1 := errors.New("error")
    e2 := errors.Wrap(e1, "inner")
    return errors.Wrap(e2, "outer")
}

func main() {
    err := fn()
    fmt.Println("%+v", err)
}
```

`가변인자`란 고정되지 않은 가변 길이의 인자를 받는것

```go
func variadicExample(s ... string){
    // 생략
}

func main() {
    vavariadicExample("red")
    variadicExample("yellow", "green", "blue")
}
```

가변 인자의 정의는 ...을 사용하여 정의하고, 처리를 할 때는 슬라이스와 동일하게 처리


가변인자는 함수에서 1개만 사용할 수 있으며, `다른 인자와 섞어 사용할 때는 마지막에 사용해야함`

```go
func variadicExample(s string, i ...int) {
    fmt.Printf("%s\n%#v\n", s, i)
}
```

---

## 구조체

```go
type rectangle struct {
    length float64
    breadth float64
    color string
}

func main() {
    spew.Dump(rectangle{10.5, 12.5 , "red"})
}
```

`메소드 리시버`

```go
type rectangle struct {
    length float64
    breadth float64
    color string
}

func (r rectangle) area() float64 {
    return r.length * r.breadth
}

func (r recatngle) perimeter() float64 {
    return 2 * (r.length + r.breadth)
}
```

---

## 익명함수와 클로저 그리고 defer

`익명함수` 함수 이름 없이 선언된함수.

다른곳에서 재사용하지 않는다면 이론식으로 정의해서 사용해도 무방

```go
func main() {
    var area = func(l int, b int) int {
        return l * b
    }
    spew.Dump(area(20,30))

// 재사용이 필요 없다면 이렇게도 가능
    func(l int, b int){
        spew.Dump(l * b)
    }(20, 20)

}
```

익명함수는 보통 다른함수의 인자로 넘겨주거나 다른 함수의 리턴값으로 사용된다.

go 언어의 sort함수 정의는 아래와같다

```go
func Slice(x any, less func(i,j int) bool)

func main() {
    people := []struct {
        Name string
        Age int
    } {
        {"kim", 27},
        {"lee", 43},
        {"jung", 36},
    }
    sort.Slice(people , func(i, j int) bool {return people[i].Name < people[j].Name})
}
```

`클로저`란 익명함수의 특별한 경우로, 함수 외부의 변수에 액세스함

```go
func main(){
    for i:= 10.0; i< 100.0; i += 10.0 {
        cv := func() float64 {
            return i * 39.370
        }
        fmt.Printf("%.0f미터 => %6.1f인치", i, cv())
    }
}
```

### `defer`란 함수가 종료될 때까지 실행을 미루는 기능

```go
func main() {
    var fn = func(i int) {
        fmt.Println(i)
    }
    defer fn(3)
    defer fn(2)
    fn(1)
}
```

`defer`를 사용안할때

```go
func readWrite() bool {
    file.Open("file")
    if failure1 {
        file.Close()
        return false
    }
    if failure2 {
        file.Close()
        return false
    }
    // read and write
    file.Close()
    return true
}
```

`defer`를 사용할 때

```go
func readWrite() bool {
    file.Open("file")
    defer file.Close()
    if failure1 {
        return false
    }
    if failure2 {
        return false
    }
    // read and write
    return true
}

defer를 사용하였을 때의 장점은 잊지않고 close 함수류를 사용할 수 있다는점.
```

----

## 인터페이스


go언어에서는 extends나 implements가 없으며 만약 구조체를 employee 인터페이스로써 사용하려면 모든 메서드

(PrintName과 PrintAddress)를 리시버로 구현해야한다

```go
type Employee interface {
    PrintName() string
    PrintAddress() [3]string
}
```

하나의 사용자 정의 타입은 여러 개의 인터페이스를 구현할 수 있따

```go
type Polygons interface{
    Perimeter() float64
}
type Object interface {
    NumberOfSide() int
}

type Square float64

func (s Square) Perimeter() float64 {
    return float64(s * 4.0)
}

func (s Square) NumberOfSide() int {
    return 4
}

// go 언어에서는 Tpye Assertion을 하고 만족할 경우 사용할 수 있다

func main() {
    var s Polygons = Square(4.5)
    fmt.Println(s.Perimeter())
    var t = s.(object)
    fmt.Println(t.NumberOfSide())
}
```

`Pointer Receiver` 는 주소를 가리키므로 원래의 값을 변경할 수 있다.

`Value Reciver`는 원래의 값을 바꿀 수 없다. 또한 필드 이름이 소문자라면 private이므로 다른 패키지에서 접근할 수 없어 초기화 불가

```go
type Printer interface {
    Print()
}

type Emp struct {
    name, address string
}

func (e *Emp) Print() {
    fmt.Printf("Name: %s\n Address: %s\n", e.name, e.address)
}

func (e *Emp) Assign(n, a string) {
    e.name = n
    e.address = a
}


func main() {
    var e Emp
    e.Assign("John Doe", "very long address string")
    var p Printer
    p = &e
    p.Print()
}
```

### `빈 인터페이스(interface{})`는 모든 유형의 데이터를 담을 수 있다

```go
fun main() {
    var vary interface{}
    vary = 123
    spew.Dump(vary)
    vary = "good morning"
    spew.Dump(vary)
    vary = map[string]int{"Mark":10, "Jane": 20}
    spew.Dump(vary)
    vary = [3]string{"Korea", "Japan", "China"}
    spew.Dump(vary)
}
```