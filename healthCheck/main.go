package healthCheck

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const StatusUp = "UP"
const StatusDown = "DOWN"

type Health struct {
	Status string
	Checks
}

type Checks []struct {
	Name string
	Status string
}

var HealthEntity Health

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	HealthEntity.Status = StatusUp
	HealthEntity.Checks = append(HealthEntity.Checks, struct {
		Name   string
		Status string
	}{Name: "HTTP", Status: StatusUp})

	dependency, err := http.Get("https://jsonplaceholder.typicode.com/health")
	if nil != err {
		HealthEntity.Status = StatusDown
		HealthEntity.Checks = append(HealthEntity.Checks, struct {
			Name   string
			Status string
		}{Name: "Typicode", Status: StatusDown})
	} else {
		httpResponse, err := ioutil.ReadAll(dependency.Body)
		if nil == err && json.Valid(httpResponse) {
			HealthEntity.Checks = append(HealthEntity.Checks, struct {
				Name   string
				Status string
			}{Name: "Typicode", Status: StatusUp})
		} else {
			HealthEntity.Status = StatusDown
			HealthEntity.Checks = append(HealthEntity.Checks, struct {
				Name   string
				Status string
			}{Name: "TypicodeJson", Status: StatusDown})
		}
	}
	apiResponse, err := json.Marshal(HealthEntity)
	if nil != err {
		log.Fatalf("Can't encode json. Error: %s", err)
	}
	_, _ = w.Write(apiResponse)
	return
}