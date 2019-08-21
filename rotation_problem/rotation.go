package main

import (
	"strings"	
)

// func main(){
// 	NoOfSteps("aaxaabc", "aabcaax")
	
// }


func NoOfSteps(string1 string, string2 string) int {
	count := 0
	if checkLength(string1, string2) == true{
		if string1 != string2{
			for i := 0; i < len(string1); i++ {
				if string1[i] != string2[i] {					
					count++
					ch := string1[:1]
					string1 = strings.Replace(string1, ch, "", 1)
					string1 = string1 + ch
				}
			}
		}
	}
	return count
}

func checkLength(string1 string, string2 string) bool {
	if len(string1) == len(string2) {
		return true
	}
	return false
}