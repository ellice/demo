package main

import (
	"fmt"
	"encoding/json"
)


func main(){

	//json eccode
	jstr := make(map[string]interface{})
	jstr["name"] = "恐怖份子"
	jstr["bolg_url"] = "http://www.fkbike.com/"
	str1, err := json.Marshal(jstr)
	if err != nil{
		panic(err)
	}
	
	fmt.Println(string(str1))
	
	
	
	// json decode	
	j2 := make(map[string]interface{})
    err = json.Unmarshal(str1, &j2)
    if err != nil {
        panic(err)
    }    
    fmt.Printf("%#v\n", j2)
	
}


