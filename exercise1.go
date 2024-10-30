"""
변수 선언
"""
var message1 string // type 선언
var message2 = "Hi, Welcome!" // 값 대입 시, type 생략
message3 := "Hi, Welcome!" // shortcut 선언 

"""
[환경 설정 및 패키지 테스트]
0) mkdir -p {basic,greetings}
1) cd greetings && go mod init github.com/devtae/go-exercise/greetings
2) cd ../basic && go mod init github.com/devtae/go-exercise/basic
3) cd .. && go work init && go work use ./greetings && go work use ./basic
4) cd basic && go mod edit -replace github.com/devtae/go-exercise/greetings=../greetings
5) go mod tidy
6) go run main.go
"""

"""
[폴더구조]
basic
├── go.mod
└── main.go
greetings
├── go.mod
└── greetings.go
go.work
"""