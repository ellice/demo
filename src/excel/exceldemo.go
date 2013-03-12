package main

import (
	"os"
	"encoding/csv"
)


func main(){
	f,err := os.Create("ellice.csv");
	if err!=nil{panic(err)}
	
	defer f.Close();
	
	f.WriteString("\xEF\xbb\xBF") //写入UTF_8 BOM
	
	w := csv.NewWriter(f) //创建一个文件
	w.Write([]string{"编号","姓名","年龄"})
	w.Write([]string{"01","恐怖份子","28"})
	w.Write([]string{"02","寻找动力","28"})
	w.Write([]string{"03","小白","12"})
	w.Write([]string{"04","zzj","11"})
	w.Flush()

}

