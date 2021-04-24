package main

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
// Rust 50, Typescript 20, Kotlin 50, dart 30
