import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	counter2            int    = 0
	firstInstanceHost2  string = "http://localhost:8080"
	secondInstanceHost2 string = "http://localhost:8081"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleProxy2)
	http.ListenAndServe(":9000", mux)
}

func handleProxy2(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	text := string(textBytes)

	if counter2 == 0 {
		resp, err := http.Post(firstInstanceHost2, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatal(err)
		}
		counter2++

		textBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		fmt.Println(string(textBytes))

		return
	}

	resp, err := http.Post(secondInstanceHost2, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatal(err)
	}
	textBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(textBytes)
	counter2--
}