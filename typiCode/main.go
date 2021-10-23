package typiCode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
var todos []Todo

func HandleMain(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	todos = nil
	fmt.Println("starting up...")
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		in := strconv.Itoa(i)
		fmt.Printf("Running %s\n", in)
		go callTypicode(in, &wg)
	}
	fmt.Println("Finished, waiting...")
	wg.Wait()

	apiResponse, err := json.Marshal(todos)
	if nil != err {
		log.Fatalf("Can't encode json. Error: %s", err)
	}
	_, _ = w.Write(apiResponse)
	fmt.Println("Returning response")
	return
}

func callTypicode(endpoint string, wg *sync.WaitGroup) {
	httpResponse, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + endpoint)
	response, err := ioutil.ReadAll(httpResponse.Body)
	if nil != err {
		log.Fatalf("Can't call typicode, error: %s", err)
		return
	}
	var todo Todo
	err = json.Unmarshal(response, &todo)
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatalf("Can't decode json response, error: %s", err)
	}
	todos = append(todos, todo)
	wg.Done()
}

