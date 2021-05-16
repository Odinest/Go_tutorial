package main

import "fmt"

const MAXMAX = "AOI";


//Go 언어에서는 함수에 '{' 기호를 내려서 쓰지 못하게 하고 있다...나쁜놈들
func main() {
    fmt.Println("Hello, 세계");

	//var ccc 의 경우 오류가 난다. 값지정 없이 변수만 생성하고 싶다면 필수 적으로 데이터 타입을 작성해야 한다. 
	//-> var ccc int; 같은 형태가 될 것이다.
	var aaa = 10;
	var aa int = 20;

	//string 
	var str string;
	str = "123";

	//`` 의 경우 엔터를 인식하여 저장하게 된다.
	var chr string;
	chr = `
	123
	`;

	//String Length
	var Length_A int = len(str);	

	fmt.Println(aaa + aa);
	fmt.Println(str + chr);
	fmt.Println(Length_A);
	fmt.Println(MAXMAX);
}