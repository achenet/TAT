package request_reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Request struct {
	Name     string   `json:"name"`
	Endpoint string   `json:"endpoint"`
	Method   string   `json:"method"`
	Data     string   `json:"data"`
	Headers  []string `json:"headers'`
}

func main() {
	fmt.Println("RequestReader")
}

func NewRequestFromJson(filename string) (Request, error) {
	var request Request
	requestFile, err := os.Open(filename)
	if err != nil {
		return request, err

	}
	defer requestFile.Close()
	// Read file data into a bytes buffer, and then unmarshal into struct.
	// Could simply use io.ReadAll
	// Get file size
	buffer, err := io.ReadAll(requestFile)
	if err != nil {
		return request, err
	}
	fmt.Println("Buffer:", string(buffer))
	// Unmarshal into JSON
	err = json.Unmarshal(buffer, &request)
	if err != nil {
		return request, err
	}
	return request, nil
}
