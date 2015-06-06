package  mate

import (
	"hera"
)

func MatchAlgorithm(phone_number string)  []string {
	
	_ , err :=hera.Redis.DoCmd("sunionstore", "tmp", "like_" + phone_number, "unlike_" + phone_number)
	if err != nil {
		hera.Logger.Warn("sunionstore tmp like_"+ phone_number + " unlike_"+ phone_number)
	}

	ret , err :=hera.Strings(hera.Redis.DoCmd("sdiff", "tmp", "userid"))
	if err != nil {
		hera.Logger.Warn("sdiff tmp userid")
	}

	return ret
}
