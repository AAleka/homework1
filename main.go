package main

import( 
	"os"
	"fmt"
	//"log"
	"io/ioutil"
	"encoding/json"
	"sort"
)

func main(){
	fmt.Println("Enter two file paths")
    var input1, input2 string
	fmt.Scanln(&input1) 
	fmt.Scanln(&input2) 

	var result1 map[string]interface{}
	var result2 map[string]interface{}

	data1, err1 := ioutil.ReadFile(input1) 
	data2, err2 := ioutil.ReadFile(input2) 

	json.Unmarshal([]byte(data1), &result1)
	json.Unmarshal([]byte(data2), &result2)

    if err1 != nil || err2 != nil { 
		errorFile, _ := os.Create("C:\\Users\\mrall\\Documents\\GitHub\\homework1\\error.txt") 
		errorFile.WriteString("Error detected")
	} else {
		var samekeys []string
    	for k1 := range result1 {
			for k2 := range result2 {
				if k1 == k2 {
					if result1[k1] == result2[k2] {
						samekeys = append(samekeys, k1)
					}
				}
			}
		}
		sort.Strings(samekeys) 
		sameFile, _ := os.Create("C:\\Users\\mrall\\Documents\\GitHub\\homework1\\same.txt") 
		for _, i := range samekeys{
			sameFile.WriteString(i + "\n")
		}
		
		var flag = false
		var addkeys []string
		for k2 := range result2 {
			flag = false
			for k1 := range result1 {
				if k1 != k2 {
					flag = true
				} else {
					flag = false
					break
				}
			}
			if flag == true {
				addkeys = append(addkeys, k2)
			}
		}
		sort.Strings(addkeys) 
		addFile, _ := os.Create("C:\\Users\\mrall\\Documents\\GitHub\\homework1\\add.txt") 
		for _, i := range addkeys{
			addFile.WriteString(i + "\n")
		}

		var removekeys []string
		for k1 := range result1 {
			flag = false
			for k2 := range result2 {
				if k1 != k2 {
					flag = true
				} else {
					flag = false
					break
				}
			}
			if flag == true {
				removekeys = append(removekeys, k1)
			}
		}
		sort.Strings(removekeys) 
		removeFile, _ := os.Create("C:\\Users\\mrall\\Documents\\GitHub\\homework1\\remove.txt") 
		for _, i := range removekeys{
			removeFile.WriteString(i + "\n")
		}

		var changekeys []string
    	for k1 := range result1 {
			for k2 := range result2 {
				if k1 == k2 {
					if result1[k1] != result2[k2] {
						changekeys = append(changekeys, k1)
					}
				}
			}
		}
		sort.Strings(changekeys) 
		changeFile, _ := os.Create("C:\\Users\\mrall\\Documents\\GitHub\\homework1\\change.txt") 
		for _, i := range changekeys{
			changeFile.WriteString(i + "\n")
		}
	}
}