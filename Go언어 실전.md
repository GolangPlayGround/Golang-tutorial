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