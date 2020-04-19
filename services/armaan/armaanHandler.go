package armaan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pranjal2209/armaan/config"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}
	response["error"] = ""
	if r.Method != "POST" {
		response["error"] = "invalid request!!"
		data, _ := json.Marshal(response)
		w.Write(data)
		return
	}
	a, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(a))
	fmt.Println(err)
	responses := map[string]string{}
	err = json.Unmarshal(a, &responses)
	fmt.Println(responses)
	//a, err := r.GetBody()
	//fmt.Println(string(a))
	fmt.Println(err)
	response["data"] = map[string]string{}
	data, _ := json.Marshal(response)
	w.Write(data)
	return
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetCountries(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	response := make(map[string][]string)
	for key := range config.Config {
		response["country"] = append(response["country"], key)
	}
	data, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(`{"country":["India"]}`))
		return
	}
	w.Write(data)
	return
}
