package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
)


// struct
type BaseSt struct{
	A int
}

//内嵌
type DerserveSt struct{
	BaseSt
	B string
}

// func
func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// handler
func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "lsklsklsk")
}

func main() {
	// fmt
	fmt.Printf("hello world\n")
	fmt.Printf("1111\n")

	// goruntime
	// net/http
	go func(){
		http.HandleFunc("/", IndexHandler)
		http.ListenAndServe("192.168.115.128:9090", nil)
	}()
 	var (
 		name string
 		age int
 		married bool
 
 	)

	value := &DerserveSt{}
	value.A = 1
	value.B = "lsklskkkkkkk"
	fmt.Print("A:%d, B:%s\n", value.A, value.B)

	// channel
	chInt := make(chan int, 100)
	go func(){
		chInt <- 1
		fmt.Println("write channel 1")
	}()
	b := <- chInt
	fmt.Printf("channel:%d\n", b)
	// scan
	fmt.Scan(&name,&age,&married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t\t",name,age,married)
}
