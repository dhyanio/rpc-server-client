package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args holds arguments passed to JSON-RPC service
type Args struct {
	ID string
}

// Book struct holds Book JSON structure
type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}
type JSONServer struct{}

// GiveBookDetail is RPC implementation
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args,
	reply *Book) error {
	var books []Book
	// Read JSON file and load data
	absPath, _ := filepath.Abs("chapter3/books.json")
	raw, readerr := ioutil.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	// Unmarshal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over each book to find the given book
	for _, book := range books {
		if book.ID == args.ID {
			// If book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// Create a new RPC server
	s := rpc.NewServer()
	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}

// A slight difference here is we have to register codec type using
// the RegisterCodec method. That is JSON codec in this case. Then, we can
// register the service using the RegisterService method and start a normal
// HTTP server. If you have noticed well, we used the alias jsonparse for
// the encoding/json package because it can conflict with another
// package, github.com/gorilla/rpc/json .

// Now, do we have to develop a client? Not necessarily, because a client can be a
// curl program since the RPC server is serving requests over HTTP, we need to
// post JSON with a book ID to get the details. So, fire up another shell and execute
// this curl request:
// curl -X POST \
// http://localhost:1234/rpc \
// -H 'cache-control: no-cache' \
// -H 'content-type: application/json' \
// -d '{
// "method": "JSONServer.GiveBookDetail",
// "params": [{
// "ID": "1234"
// }],
// "id": "1"
// }'