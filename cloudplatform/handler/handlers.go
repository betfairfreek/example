package handler

import (
	"fmt"
	"net/http"
	"github.com/example/cloudplatform/client"
	"io/ioutil"
//	"bytes"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------->Connecting to Google")

	client:=client.HttpClient()
	req,_:=http.NewRequest("GET","http://google.bg",nil)
	resp,_:=client.Do(req)
	//	resp, _ := client.Get("http://google.bg")

	body,_:=ioutil.ReadAll(resp.Body)

	fmt.Println("Proto  ", resp.Proto)
	fmt.Println("Header ", resp.Header)
	fmt.Println("Body ",string(body) )
	fmt.Println("---------------------->End with status code", resp.StatusCode)
}

func Loggin(w http.ResponseWriter, r *http.Request)  {
	client:= client.HttpClient();
	json:=[]byte(`username:freekman,password:freek321man678`)

	req,_:=http.NewRequest("POST","https://identitysso.betfair.com/api/login",nil)
	req.PostForm.Add("username","freekman")
	req.PostForm.Add("password","freek321man678")
	req.Header.Add("Accept","application/json")
	req.Header.Add("X-Application","7LKDnmqcaXsraTC6")
		fmt.Println("Post data",string(json))

	resp,_:=client.Do(req)

	body,_:=ioutil.ReadAll(resp.Body)

	fmt.Println("Proto  ", resp.Proto)
	fmt.Println("Header ", resp.Header)
//	fmt.Println("Body ",string(body) )
	fmt.Println("---------------------->End with status code", resp.StatusCode)
}

