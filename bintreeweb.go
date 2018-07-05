package main

import (
	"encoding/json"
	"fmt"
	"time"
	"net/http"
)


type TestObject struct {       
	Value   int    `json:"id"`
	LeftNode  *TestObject `json:"LeftReference"`
	RightNode *TestObject `json:"RightReference"`
}

//rootnd := &TestObject{}
//var rootnd *TestObject

func marshalthecontent(testobj *TestObject){
	b, err := json.MarshalIndent(testobj, "", "\t")
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b[:]))
}


func createtree(arr []int) *TestObject{
	rootnd := &TestObject{}
	rootnd.Value = arr[0]
	rootnd.LeftNode  = nil
	rootnd.RightNode = nil
	for _, num := range arr[1:]{
	
		 rootnd =  inserttotree(rootnd,num)	
		 //go func(){
		//	rootnd =  inserttotree(rootnd,num)	
		//}()
	}
	return rootnd
}


func inserttotree(root *TestObject, num int) *TestObject{
	time.Sleep(100 * time.Millisecond)
	fmt.Println("inserttotree control")
	if num > root.Value {
		if root.RightNode == nil {
			rnode := TestObject{Value:num,LeftNode :nil,RightNode:nil}
			root.RightNode = &rnode
			return root
		} else {
			root.RightNode = inserttotree(root.RightNode,num)
			return root
		}
	} else {
		if root.LeftNode  == nil {
			lnode := TestObject{Value:num,LeftNode :nil,RightNode:nil}
			root.LeftNode  = &lnode
			return root
		} else {
			root.LeftNode  = inserttotree(root.LeftNode ,num)
			return root
		}
	}
}

func main() {
	nodearray := []int{2,7,5,3,9,4,6,1}
	//nodearray1 := []int{2,7,5,1,10,15,6,1}
	//nodearray2 := []int{12,7,5,11,10,15,16,1}

	
	//line#77 - 83 & #93 [ListenAndServe] to be used to web binary tree. Accessbile at: http://localhost:1706/processtree
	//Also, import net/http package.
	http.HandleFunc("/processtree", func(w http.ResponseWriter, r *http.Request) {
			arrivedobj := &TestObject{}
			arrivedobj = createtree(nodearray)
			json.NewEncoder(w).Encode(arrivedobj)
			marshalthecontent(arrivedobj)
			
		})
	t := time.Now()

	http.ListenAndServe(":1706", nil)
		time.Sleep(1000 * time.Millisecond)

	taken := time.Since(t)
		
		fmt.Println("value of time taken: ", taken)
	
}
