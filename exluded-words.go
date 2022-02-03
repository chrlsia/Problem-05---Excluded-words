package main

/* Assume we have 
sl1:=[]string{"This", "is", "a", "good", "message"}
sl2 := []string{"good", "is"}
We want to exlude every word from sl2 which is appeared in sl1
So the result would be: [This a message]
*/

import "fmt"

func main() {
	sl1 := []string{"This", "is", "a", "good", "message"}
	sl2 := []string{"good", "is"}

	//get element from sl2
	for _, ele2 := range sl2 {
		//get element from sl1
		for index1, ele1 := range sl1 {
			if ele2 == ele1 {
				//delete ele1 from sl1
				sl1 = append(sl1[:index1], sl1[index1+1:]...)
			}
		}
	}
	fmt.Printf("sl1 now is:%v\n", sl1)
}
