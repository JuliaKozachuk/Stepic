package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://instagram188.p.rapidapi.com/userinfo/mega6obep3000"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "998eac06a4msh716839cbbb7ec92p1f6e79jsnd5f140a472d2")
	req.Header.Add("X-RapidAPI-Host", "instagram188.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer w.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	user_info_data := string(body)
	js := json.Unmarshal([]byte(user_info_data), &user_info_data)

	fmt.Println(js)

}
