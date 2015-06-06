package  mate

import (
	"fmt"
	"hera"
)

func MatchAlgorithm(phone_number string)  []string {
	
	_ , err :=hera.Redis.DoCmd("sunion", "tmp", "like_" + phone_number, "unlike_" + phone_number)
	if err != nil {
		fmt.Println("[warn] sunion, value: "+ phone_number )
	}

	ret , err :=hera.Redis.DoCmd("sdiff", "tmp", "userid")
	if err != nil {
		fmt.Println("[warn] sunion, value: "+ phone_number )
	}
	value,ok := ret.([]string)

	fmt.Printf("%v,%v",value,ok)

	return value
}
