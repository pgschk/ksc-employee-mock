package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/pgschk/namegen/pkg/namegen"
)

// NumEmployeeRecords defines how many employee records will be returned per call
const NumEmployeeRecords = 6

/* struct of a KSC Employee containing:
 * FirstName: A generated first name for the Kerbal.
 * LastName: Kerman. Always Kerman. Don't think too much about it.
 * Age: A random age between 22 and 39. Kerbals in the prime of their life.
 * Position: One of the available jobs at KSC: Pilot, Scientist or Engineer.
 */
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

/* List of available jobs at KSC */
var Jobs = [3]string{
	"Pilot",
	"Scientist",
	"Engineer",
}

/* newKSCEmployee returns a KSC employee with a randomly generated first name */
func newKSCEmployee() *KSCEmployee {
	k := KSCEmployee{}
	k.FirstName = namegen.GenerateName()
	k.LastName = "Kerman"
	k.Age = rand.Intn(17) + 22
	k.Position = Jobs[rand.Intn(3)]
	return &k
}

/* getEmployees handles a HTTP call and writes a JSON array of KSC employees */
func getEmployees(w http.ResponseWriter, req *http.Request) {
	response := Response{}
	response.Status.Status = 200
	response.Status.Time = time.Now().Unix()

	// fill array with multiple random KSC employees
	for i := 0; i < NumEmployeeRecords; i++ {
		response.Data = append(response.Data, newKSCEmployee())
	}

	// Marshal JSON response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// output JSON
	fmt.Fprint(w, string(jsonResponse))
}

/* main set up simple HTTP server to return a JSON array of KSC employees */
func main() {
	rand.Seed(time.Now().UnixNano()) // initialize RNG seed
	http.HandleFunc("/v1/employees", getEmployees)

	fmt.Println("KSC Employee Record API mock server starting. Don't expect more logging.")
	http.ListenAndServe(":8080", nil)
}
