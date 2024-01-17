# Golang-tutorial


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

