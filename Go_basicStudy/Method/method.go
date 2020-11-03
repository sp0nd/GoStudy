//메서드 호출하기
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"log"
)

func main() {
	var now time.Time = time.Now() // time.Now()는 현재 날짜 및 시간을 나태느는 Time 값을 반환해서 now에 저장
	var year int = now.Year()      //now가 참조하고 있는 값에서 Year()를 호출해서 현재 시간중 year 을 int 타입으로 반환
	fmt.Println(year)

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "0") // Replacer을 설정하는 메서드(NewReplacer)를 통해 모든 #을 0으로 치환하도록 설정해 strings.Replace를 반환한다.
	fixed := replacer.Replace(broken)         // strings.Replace의 Replace 메서드를 호출해 치환할 문자열broken을 아규먼트로 넣고 치환된 문자열을 fixed에 반환한다.
	fmt.Printf(fixed)

	//package.func1()
	// func1함수가 package에 속한다는 뜻
	//value.func2()
	//func2() 메서드는 개별 값에 속한다.  ### 즉 .(dot)오른쪽에 있는 무언가가 왼쪽의 무언가에 속해 있음을 나타낸다

	////////////pass_fail
	fmt.Println("Enter a grade : ")     // 성적을 입력하라는 문구 출력
	reader := bufio.NewReader(os.Stdin) // 키보드로 부터 텍스트를 읽어오기 위한 '버퍼 리더'를 만든다.
	// 1-1 input := reader.ReadString('\n')    	// '\n'눌리기 직전까지 입력된 모든 문자를 반환한다.
	// 1-1 assignment mismatch: 1 variable but reader.ReadString returns 2 values reader.ReadString 의 반환갑이 두개인데 하나만 반환해서 에러가 뜬다
	//input, err := reader.ReadString('\n')  	// err 변수를 할당하지만 err을 사용하지 않아서 에러가 뜬다. go에서 선언한 변수는 반드시 사용해야한다.
	input, _ := reader.ReadString('\n') // 빈식별자를 이용해(_) 필요한 값을 버린다. 빈식별자에 값을 할당하면 값은 버려진다.
	fmt.Println(input) //사용자가 입력한 값을 출력한다.

	
}
