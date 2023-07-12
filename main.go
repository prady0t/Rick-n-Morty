package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type result struct{
 Id            int                        `json:"id"`
 Name          string                     `json:"name"`
 Status        string                     `json:"status"`
 Species       string                     `json:"species"`
 Type          string                     `json:"type"`
 Gender        string                     `json:"gender"`
 Origin        map[string]string          `json:"origin"`
 Location      map[string]string          `json:"location"`
 Image         string                     `json:"image"`
}

// type info struct{
// 	Info       map[string]json.RawMessage              `json:"info"`
// }


func random(x int) int{
	return rand.Intn(x)
}

func main(){
	
	resp := GETresponse(linkGenerator()) 

	defer resp.Body.Close()

	jsonData := ReadData(resp)

	ParsedData := ParsedData(jsonData)

	fmt.Println("Name-> ",ParsedData.Name)
	fmt.Println("Status-> ",ParsedData.Status)
	fmt.Println("Species-> ",ParsedData.Species)
	fmt.Println("Type-> ",ParsedData.Type)
	fmt.Println("Gender-> ",ParsedData.Gender)
	fmt.Println("Origin-> ",ParsedData.Origin["name"])
	fmt.Println("Location-> ",ParsedData.Location["name"])
	
}

// func getInfo() *info{
// 	resp := GETresponse("https://rickandmortyapi.com/api/character")
// 	jsonData := ReadData(resp)
// 	var ParsedData info
// 	err := json.Unmarshal(jsonData, &ParsedData)
// 	if err != nil{
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	//fmt.Println(string(jsonData))
// 	return &ParsedData

// }

func ParsedData(jsonData []byte) (*result){
  
	var ParsedData result
	err := json.Unmarshal(jsonData, &ParsedData)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return &ParsedData
}

func ReadData(resp* http.Response)[]byte{
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return jsonData

}

func GETresponse(s string)(resp* http.Response){   

	resp , err := http.Get(s)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	return resp
}

func linkGenerator() string {
	//count := getInfo()
	//totalCount := count.Info["count"]
	//fmt.Println(totalCount)
	//fmt.Println(string(count.Info[]))
   return "https://rickandmortyapi.com/api/character/"+strconv.Itoa(random(826))
}