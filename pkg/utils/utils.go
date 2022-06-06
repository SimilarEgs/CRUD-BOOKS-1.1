package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// this function тainly will be used with POST method, when creating a book entity,
// which help us easily unmarshal requst json data
func ParseBody(r *http.Request, x interface{}) {

	//reading body, and if there is no error -> unmarshal that body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
