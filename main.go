package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.GET("/", printM)
	router.Run("localhost:8083")
	fmt.Println("Started at 8083")
	

	
}

func printer(p *result){
	fmt.Println("Name-> ",p.Name)
	fmt.Println("Status-> ",p.Status)
	fmt.Println("Species-> ",p.Species)
	fmt.Println("Type-> ",p.Type)
	fmt.Println("Gender-> ",p.Gender)
	fmt.Println("Origin-> ",p.Origin["name"])
	fmt.Println("Location-> ",p.Location["name"])
}

func printM(c *gin.Context){
	resp := GETresponse(linkGenerator()) 

	defer resp.Body.Close()

	jsonData := ReadData(resp)

	ParsedData := ParsedData(jsonData)

	printer(ParsedData)

	htmlTemplate := `<div style="display: flex; justify-content: center; align-items: center; height: 100vh; background-color: #C4CB24;"><div style="text-align: center; background-color: #fff; padding: 20px;"><h1>Daily Dose of Rick and Morty</h1><img src="%s" alt="Rick and Morty Image"><p>Name: %s</p><p>Status: %s</p><p>Species: %s</p><p>Gender: %s</p><p>Origin: %s</p><p>Last Location: %s</p></div></div>`
    formattedHTML := fmt.Sprintf(htmlTemplate, ParsedData.Image, ParsedData.Name, ParsedData.Status, ParsedData.Species,ParsedData.Gender,ParsedData.Origin["name"],ParsedData.Location["name"])

	fmt.Fprintln(c.Writer,formattedHTML)
	
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
