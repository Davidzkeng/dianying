package main

import (
	"encoding/json"
	"fmt"
)

func read(ch1 chan int,ch2 chan bool){
	for{
		v ,ok:= <- ch1
		if ok{
			fmt.Printf("read a int is %d\n",v)
		}else{
			fmt.Println("hello")
			ch2 <- true
		}
	}

}

func write(ch chan int){
	for i:=0;i<10;i++{
		ch <- i
	}
	close(ch)
}

type hello interface {
	Hello()
}

type Hello1 struct {
	Name string
	Age int
}


func (h Hello1) Hello(){
	h.Name = "34532"
	fmt.Println(h.Name)
}


func main() {
	a := "[1,2,3,4,5,6]"
	var b interface{}
	json.Unmarshal([]byte(a),&b)
	for _, value := range b.([]interface{}) {
		fmt.Println(typeof(value))
	}
	//fmt.Println(b)
	//hello := Hello1{
	//	Name:"test",
	//	Age:18,
	//}
	//
	//
	//test(&hello)
	//url := "http://localhost/api/test"
	//s := strings.TrimPrefix(url,"/")
	//fmt.Println(s)
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func test(h hello)  {
	h.Hello()
	fmt.Println(h)
}
