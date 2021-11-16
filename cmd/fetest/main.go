package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const PORT = 8000

func main() {
	fmt.Println("Sending a dummy request to port", PORT)
	url := fmt.Sprintf("http://localhost:%d/formula", PORT)
	formulaStr := []byte(`(p' <-> (p | q)) & (q | q') & !(q' & q) & !(q & r & r') & (r' -> (p | q | r))`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(formulaStr))
	if err != nil {
		fmt.Printf("[ERROR] Malformed request! %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "test/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] Request resulted in error %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[RESP] Response status:", resp.Status)
	fmt.Println("[RESP] Response header:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("[RESP] Response:", string(body))
}
