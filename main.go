package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pgschk/namegen/pkg/namegen"
)

var numEmployeeRecords = 6

type KSCEmployee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Position  string `json:"position"`
}

type Status struct {
	Status int   `json:"statusCode"`
	Time   int64 `json:"timestamp"`
}

type Response struct {
	Status Status         `json:"status"`
	Data   []*KSCEmployee `json:"data"`
}

var Jobs = [...]string{
	"Pilot",
	"Scientist",
	"Engineer",
}

func newKSCEmployee() *KSCEmployee {
	k := KSCEmployee{}
	k.FirstName = namegen.GenerateName()
	k.LastName = "Kerman"
	k.Age = rand.Intn(17) + 22
	k.Position = Jobs[rand.Intn(3)]
	return &k
}

func getEmployees(w http.ResponseWriter, req *http.Request) {
	response := Response{}
	response.Status.Status = 200
	response.Status.Time = time.Now().Unix()
	for i := 0; i < numEmployeeRecords; i++ {
		response.Data = append(response.Data, newKSCEmployee())
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprint(w, string(jsonResponse))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // initialize RNG seed
	http.HandleFunc("/v1/employees", getEmployees)
	http.HandleFunc("/headers", headers)

	fmt.Println("KSC Employee Record API mock server starting. Don't expect more logging.")
	http.ListenAndServe(":8080", nil)
}
