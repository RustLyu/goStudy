package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"container/list"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
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

type JsonSt struct{
	Name string `json:"name"`
	Pwd  int	`json:"pwd"`
	Login bool `json:"login"`
}
/*************************************************************/
type FuncSt struct{
	name string
	Pwd int
}

// struct Public Function
func (self *FuncSt) UpdateName(name string) bool {
	if self.name == name {
		return false
	}
	self.name = name
	return true
}

func (self FuncSt) GetMyName() string{
	return self.name
}
/*************************************************************/
// interface
type DateWrite interface{
	TWrite(data interface{}) error
}

type FileSt struct{
}

func (self *FileSt) TWrite(data interface{}) error{
	fmt.Println("writedata", data)
	return nil
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

var (
	host = "127.0.0.1"
	user = "root"
	pwd = "root"
	db = "test"
)

func main() {

	db,err := sql.Open("mysql",user + ":" + pwd +"@tcp(" + host + ")/" + db)
	if err != nil{
		fmt.Println(err)
		fmt.Println("open mysql error")
		return;
	}else{
		rows,terr := db.Query("select * from data")
		if terr == nil{
			for rows.Next(){
				var id int
				err := rows.Scan(&id)
				if err == nil{
					fmt.Println(id)
				}
			}
		}else{
			fmt.Println(ret)
		}
	}
	// interface
	tFile := &FileSt{}
	var writer DateWrite
	writer = tFile
	writer.TWrite("lsk dataWrite test")
	// panic
	// panic("DDDDDDDDDDDDDDDDDDDDDDD")
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

	// array
	var num [3]int
	// for
	for i,v := range num{
		fmt.Printf("数组下表：%d,值:%d\n",i,v)
	}

	// slice
	lsklsk := num[2:3]

	for i,v := range lsklsk{
		fmt.Printf("数组下表：%d,值:%d\n",i,v)
	}

	// map
	testMap := make(map[int]string)

	testMap[1] = "121"
	testMap[2] = "122"
	testMap[3] = "123"
	testMap[4] = "124"
	testMap[5] = "125"

	for i,v := range testMap{
		fmt.Printf("index: %d, value: %s\n", i, v)
	}

	delete(testMap, 5)
	for i,v := range testMap{
		fmt.Printf("index: %d, value: %s\n", i, v)
	}
	// list

	tList := list.New()
	elem := tList.PushBack("124")
	tList.PushFront(456)
	tList.InsertAfter("789",elem)
	tList.InsertBefore(909,elem)

	// json Marshal
	jsonTest := &JsonSt{
		Name: "lsk",
		Pwd: 233,
		Login: false,
	}

	se, err:= json.Marshal(jsonTest)
	if err != nil{
		fmt.Println("serilize error")
	}
	fmt.Println(string(se))

	// json Unmarshal
	jsonObj := &JsonSt{}
	unMErr := json.Unmarshal(se, &jsonObj)

	if unMErr != nil{
		fmt.Println("Unmarshl error")
	}
	fmt.Println(jsonObj)

	// scan
	fmt.Scan(&name,&age,&married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t\t",name,age,married)
}
