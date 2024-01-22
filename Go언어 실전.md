## 커맨드라인 플래그, 로깅, GODOC

`커맨드라인 플래그`란 커맨드라인 프로그램에서 옵션을 지정하는 방법

"ls -al" 처럼 커맨드 라인에서 명령어 뒤에 추가로 붙이는 옵션을 지정하는 일반적인 방법

```go
envPtr := flag.String("env", "prod", "환경변수")
cpuPtr := flag.Int("cpu", runtime.NumCPU(), "CPU수")
flag.Parse() // 플래그 값을 읽어 들이기 위해 반드시 필요
spew.Dump(envPtr)
spew.Dump(cpuPtr)
```

2가지 변수를 입력받을 수 있도록 하고 env를 생략하면 기본 값은 Prod, cpu를 생략하면 기본 값은 현재 시스템 cpu수

`로깅` 이란 문제가 생겼을 때 당시의 시스템 상태를 알기 위해 로그(log)를 남기는 활동을 말함

go언어에 로깅에 관련된 내장 패키지 있지만 rs/zerolog라는 로깅 패키지를 추천. (메모리 사용량도 적고 기능이 많음)

```go
func main() {
    if *envPtr != "prod" {
        log.Logger = log.Output(zerolog.ConsoleWriter{
            Out : os.StdOut,
            TimeFormat: time.RFC3339,
        })
    }
}
```

`GODOC` go언어로 작성한 주석으로부터 자동으로 문서화를 해주는 프로그램

https://pkg.go.dev/golang.org/x/tools/cmd/godoc

일단 패키지나 구조체,함수,주요변수 위에 아래 예제처럼 주석을 붙임. 보통 패키지에 대한 주석은 doc.go파일을 만들어

그 파일에 기록하는 관행이 있음. 패키지에 대한 설명일 때는 주석이 "Package 패키지명"으로 시작하며 구조체, 함수, 변수 등에 

대한 설명인경우 그 이름으로 시작

```go
/*
 Package string_util은 문자열형에 관한 유용한 유틸리티 함수들을 모아 놓은 패키지
*/
package string_util

// Stringer 인터페이스는 문자열형을 처리하는 인터페이스
type Stringer interface {
    String() string
}
```

---

## 파일과 디렉토리

빈 파일을 생성 후 파일에 내용 채워넣기

```go
func main() {
    const filename = "test.txt"
    fp, err := os.Crearte(filename)
    if err!= nul {
        log.Panic().Err(err).Msg("파일 생성 실패")
    }
    defer fp.Close()
    _, err = fmt.Fprintf(fp, "Hello World\n")
    if err != nil {
        log.Panic().Err(err).Msg("파일 기록 실패")    
    }
    contents, err := os.ReadFile(filename)
    if err != nil {
        log.Panic().Err(err).Msg("파일 내용 읽어오기 실패")
    }
    spew.Dump(contents)
}
```

os.Create는 여러 번 실행시켜도 파일 내용을 계속 덮어쓰기 때문에, 파일 내용을 append하는 법을 보자

```go
func main() {
    const filename = "test.txt"
    fp, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    // 이하 생략
}
```

파일을 다루려면 알아야하는게 있는데 파일 플래그를 보자

- O_RDONLY: read only
- O_WRONLY: writ eonly
- O_RDWR : read and write
- O_CREATE : 파일이 없다면 새로 생성
- O_EXCL : O_CREATE와 사용시 파일은 없어야 함
- O_TRUNC: 쓰기 가능한 모드로 파일이 열리면 파일의 내용을 전부 지워버림
- O_APPEND: 쓰기 할 때 파일의 내용을 추가 가능 하게 함. 즉 전부 지우지 않음

0644는 무슨의미냐면 0으로 시작했기에 8진수 즉 8진수 644란 의미. 각 자리마다 owner,group,others의 뜻을 가지고 있다
- owner : 파일의 소유자. 보통은 파일을 최초 생성한 사람
- group : 리눅스 시스템에서는 계정과 별도로 group으로 계정 그룹을 관리. 계정 그룹이 같은 사람들 의미
- other : 소유자도 아닌 같은 그룹도 아닌 사람들

각 자리마다 rwx를 2진수로 표시. r은 read w은 write x는 executable

101이면 read & executable이고 8진수로 표현하면 5. 111이면 read & write &executable이고 8진수로 표현시 7

일반적으로 디렉토리 만들 때는 0755, 일반 파일을 만들 때는 0644로 만드는 편


파일 이름을 변경하려면 os.Rename을 사용하면 됨

```go
func main() {
    oldName := "test.txt"
    newName := "testing.txt"
    if err := os.Rename(oldName, newName); err != nul {
        // 이하 생략
    }
}
```

`파일과 디렉토리`

- os.Create(name string) (*File,error) : name이름의 파일이 존재하지 않으면 생성
- os.Open(name string) (*file, error) : name의 이름을 읽기 전용 모드로 연다
- os.OpenFile(name string, flag int, perm FileMode) (*file, error) : 지정된 flag와 퍼미션으로 파일 열음
- os.Stat(name string) (*FileInfo, error) : 파일의 정보를 가져오거나, 파일이 존재하는지 체크하는데 사용
- os.Rename(oldpath, newpath string) error : oldpath에서 newpath로 이름을 변경
- os.Mkdir(name string, perm FileMode) error, os.Modkir(name string, perm FileMode) error : 둘다 디렉토리 
  
    만들지만 MkdirAll은 서브 티렉토리가 없으면 서브디렉토리도 만들음
- os.Remove(name string) error : 파일을 지웁니다
- os.Chmod(name string, mode FileMode) error : 특정 파일의 퍼미션을 변경


--- 

## 다양한 파일 읽고 쓰기


`json 읽고 쓰기`
```json
// test.json
{
  "employees": [
    {
      "name": "Kim",
      "email": "kim@email.com"
    },
    {
      "name": "Jun",
      "email": "jun@email.com"
    },
    {
      "name": "seong",
      "email" : "seong@email.com"
    }
  ]
}
```

위 json 데이터 구조를 표현하기 위해 2가지 구조체를 만들자

```Go
type Employees struct {
    Employees []Employee `json:"employees"`
}
type Employee struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {
    data1 , err := os.Readfile("test.json")
    if err != nil {
        panic(err)
    }
    emp := EMployees{}
    err = json.Unmarshal(data1, &emp)
    if err != nil {
        panic(err)
    }
    spew.Dump(emp)
    data2, err := json.MarshalIndent(emp, "","   ")
    if err != nil {
        panic(err)
    }
    err = os.WriteFile("test2.json",data2, 0644)
    if err != nil {
        panic(err)
    }

}
```

`Text 읽기/쓰기`

```go
func main() {
    // rfp, wfp 파일 여는 부분 생략
    rsc := bufio.NewScanner(rfp)
    rsc.Split(bufio.ScanLines)
    wsc := bufio.NewWriter(wfp)
    for rsc.Scan() {
        line := rsc.Text()
        fmt.Println(line)
        if _, err := wsc.WriteString(line + "\n"); err != nil {
            panic(err)
        }
        _ = wsc.Flush()
    }
}
```

`CSV 읽기/쓰기`

```csv
Class,Korean,English,Mathematics
Kim,90,80,60
Lee,100,90,90
Park,80,90,70
Jung,40,30,80
```

```go
func main() {
    // rfp, wfp 파일 여는 부분 생략
    reader := csv.NewReader(bufio.NewReader(rfp))
    writer := csv.NewWriter(bufio.NewWriter(wfp))

    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }

        err = writer.Write(row)
        if err != nil {
            panic(err)
        }
        writer.Flush()
    }
}
```