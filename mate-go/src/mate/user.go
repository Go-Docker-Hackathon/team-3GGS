package example

import (
	"fmt"
	"hera"
)

type UserREST struct {
}

//curl 'localhost:8083/Hello/Get?fd=123'
func (this *UserREST) Login(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]

	return c.Success("phone_number : " + phone_number)
}

//curl 'localhost:8083/Hello/Set?fd=123'
func (this *UserREST) MateList(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	return c.Success("phone_number : " + phone_number)
}

func (this *UserREST) MatedList(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	return c.Success("phone_number : " + phone_number)
}

//curl 'localhost:8083/Hello/Set?fd=123'
func (this *UserREST) Like(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	return c.Success("phone_number : " + phone_number)

}

func (this *UserREST) Unlike(c *hera.Context) error {
	params := c.Params
	phone_number :=  params["phone_number"]
	fmt.Println("HelloREST::Set")
	return c.Success("phone_number : " + phone_number)
}

func init() {
	hera.NewRouter().AddRouter(&UserREST{})
}


//login

//getmatelist

//nosure

//sure



