package  mate

import (
	"hera"
)

type UserREST struct {
}

//curl 'localhost:8083/User/Login?phone_number=123'
func (this *UserREST) Login(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	
	hera.Logger.Info("have sucess vistited User::Login() interface, value: "+ phone_number)
	
	//todo  add user info
	_ , err :=hera.Redis.DoCmd("sadd", "userid", phone_number)
	if err != nil {
		hera.Logger.Warn("sadd  userid error, value: "+ phone_number)
	}

	data := make(map[string]interface{})
	data["phone_number"] = phone_number	
	data["auth_code"] = IdentifyCode()

	return c.Success(data)
}

//curl 'localhost:8083/User/MateList?phone_number=123'
func (this *UserREST) MateList(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]

	hera.Logger.Info("have sucess vistited User::MateList() interface, value: "+ phone_number)
	ret := MatchAlgorithm(phone_number)

	return c.Success(ret)
}

//curl 'localhost:8083/User/MatedList?phone_number=123'
func (this *UserREST) MatedList(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	
	hera.Logger.Info("have sucess vistited User::MatedList() interface, value: "+ phone_number)
	ret , err :=hera.Strings(hera.Redis.DoCmd("smembers", "like_" + phone_number))
	if err != nil {
		hera.Logger.Warn("smembers like_"+ phone_number)
	}

	return c.Success(ret)
}

//curl 'localhost:8083/User/Like?phone_number=123&like_id=12312'
func (this *UserREST) Like(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	like_id :=  params["like_id"]
	
	hera.Logger.Info("have sucess vistited User::Like() interface, value: "+ phone_number + " " + like_id)

	_ , err :=hera.Redis.DoCmd("sadd", "like_" + phone_number, like_id)
	if err != nil {
		hera.Logger.Warn("sadd like_"+ phone_number + " " + like_id)
	}
	return c.Success()
}

//curl 'localhost:8083/User/Unlike?phone_number=123&like_id=12312'
func (this *UserREST) Unlike(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	like_id :=  params["like_id"]
	
	hera.Logger.Info("have sucess vistited User::Unlike() interface, value: "+ phone_number + " " + like_id)
	_ , err :=hera.Redis.DoCmd("sadd", "unlike_" + phone_number, like_id)
	if err != nil {
		hera.Logger.Warn("sadd unlike_"+ phone_number + " " + like_id)
	}
	return c.Success()
}

func init() {
	hera.NewRouter().AddRouter(&UserREST{})
}

