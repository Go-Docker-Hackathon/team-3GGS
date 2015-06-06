package  mate

import (
	"fmt"
	"hera"
)

func MatchAlgorithm(phone_number string)  []string {
	
	_ , err :=hera.Redis.DoCmd("sunionstore", "tmp", "like_" + phone_number, "unlike_" + phone_number)
	if err != nil {
		fmt.Println("[warn] sunion, value: "+ phone_number )
	}

	val,_ :=hera.Redis.DoCmd("smembers", "tmp")

	ret , err :=hera.Strings(hera.Redis.DoCmd("sdiff", "tmp", "userid"))
	if err != nil {
		fmt.Println("[warn] sunion, value: "+ phone_number )
	}

	return ret
}
