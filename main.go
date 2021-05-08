/* package main

//// Go 언어는 사용하지 않는 모듈이 있으면 오류를 내뿜는다!
import (
	"encoding/json"
	"fmt" //formating - standard input output
)

func main() { //main 먼저 실행
	//어떤 모듈에 export되는 function은 무조건 대문자로 시작
	fmt.Println("hi") //세미콜론  X

	test := map[string]int{"a": 10, "b": 11, "c": 12} //map이라는 파이썬에서 딕셔너리처럼 기능하는 변수를 만듬 string아 key의 type이고 intrk value의 type
	value, err := json.Marshal(test)                  //join으로 변환하는 함수인 Marshal에 넣어줌

	//모든 error를 프로그래머가 직접 제어해야함
	if err != nil {
		fmt.Println(value)
	}

	fmt.Println(value)
}

//2009년 구글에서 개발
// Go 언어에 대한 설교
// 2009년 구글에서 개발된 언어
// 개발진: C++ 개발자, ~~~~~~~~
// Stackoverflow 2020 - Most Wanted 3위
// " - Most Loved 5위
// " - 평균 연봉 세계 3위($74k = $74000 = 82,658,000)
// 2009, 2016 올해의 언어
// 언어 순위 11위
// Rust 50, Typescript 20, Kotlin 50, dart 30 */

//20210508
package main

import "fmt"

var age int64 = 18

//age := 18 (X)

/* func main() {
	//var  변수명 변수타입 =
	//var name string = "jaeeun"

	//이름 := 값
	name := "jaeeun"

	fmt.Println(name)
} */

func multiply(a int, b int) (total int) {
	//total = a * b
	//return
	return a * b
}

func addAll(inp ...int) (total int) {
	total = 0
	for _, val := range inp {
		total += val
	}
	fmt.Println(inp)
	//inp type list

	/* for{
		//while문
	} */

	/* for i:=; i<10;i++{

	} */

	return
}

func main() {
	fmt.Println(multiply(2, 3))
	fmt.Println(addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	age := 20

	if manAge := age - 1; manAge >= 19 {
		fmt.Println("You can buy an alchol")
	} else {
		fmt.Println("You cannot buy an alchol")
	}

	switch manAge := age - 1; {
	case manAge <= 13:
		fmt.Println("어린이 요금")
	case manAge < 65:
		fmt.Println("일반 요금")
	case manAge >= 65:
		fmt.Println("노약자 요금")
	}

	a := 10
	b := &a
	a = 11

	fmt.Println(a, *b) //11 11

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for v := range arr {
		fmt.Println(v)
	}
	fmt.Println(arr)
	arr = append(arr, 11)
	fmt.Println(arr)
}
