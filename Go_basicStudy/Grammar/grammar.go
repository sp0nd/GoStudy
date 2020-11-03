package main

import ( // '(' 는 여러 패키지를 한번에 가져올수 있는 import 구문
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("hello,go")      // println 대소문자 구분 Println Print + ln 인데 ln은 줄바꿈 new line 을 뜻함
	strings.Title("go go go go") // 문자열을 받아서 각 단어의 첫문자만 대문자로 교체하는 함수 즉, 출력기능 없음
	math.Floor(20.75)            // 가장 가까운 정수로 내림 하여 값을 반환 ,출력 기능 없음
	fmt.Println(strings.Title("go go go go"), math.Floor(20.75))
	fmt.Println(4 * 6) // 사칙연산
	fmt.Println(4 < 6) // 비교연산자
	//타입
	fmt.Println(reflect.TypeOf(40))
	fmt.Println(reflect.TypeOf(3.141592))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("hello, go"))
	//변수 선언 타입 생략 가능
	var quantity = 4             //var a int =4
	var length, width = 1.2, 2.4 // var b,c float64 = 2.2,3.3
	var str = "Kim taejun"       //var d string = "asdf"
	var a int                    // 값을 할당 안했을 때 가비지 값을 반환하는 C/C++과 달리 0을 반환한다.
	Num := 100                   // 변수 선언과 동시에 값을 할당할 때 := 사용
	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(a)
	fmt.Println(Num)
	//타입변환
	var a1 int = 1
	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(float64(a1)))

}

//go 기본 형식
//package 절
// import 문
// 실제 코드
