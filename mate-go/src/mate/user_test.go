package mate

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	r, err := http.NewRequest("GET", "http://localhost:8083/Hello/Get", nil)

	if err != nil {
		log.Println("http.NewRequest: ", err.Error())
	}

	resp, err := http.DefaultClient.Do(r)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		log.Println("http.DefaultClient.Do: ", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("resp.StatusCode!=http.StatusOK: ", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		log.Println("ioutil.ReadAll: ", err.Error())
	}

	log.Println(string(data))
}
