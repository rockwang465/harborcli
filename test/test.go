package main


import (
	"fmt"
	"github.com/rockwang465/harborcli"
)

const (
	harborUrl  = "https://harbor.xxx.com.cn:5443"
	harborUser = "admin"
	harborPwd  = "xxx@harbor"
)

func main(){
	harborURL := harborUrl
	auth := harborcli.LoginForm{
		Username: harborUser,
		Password: harborPwd,
	}
	client, err := harborcli.NewHarborClient(harborURL, auth)
	if err != nil {
		fmt.Println("Error: new harbor client err: ", err)
	}
	if err = client.Login(); err != nil {
		fmt.Println("Error: login failure err: ", err)
	} else {
		fmt.Println("Info: login successful")
	}
}
