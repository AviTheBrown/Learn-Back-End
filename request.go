package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
)

func headers(writer http.ResponseWriter, req *http.Request) {
	// this retreives information from the Head of the request
	// there should be no whitespace in the string arg
	fmt.Fprintf(writer, "this is the header\n")
	//	head := req.Header["Accept-Encoding"]
	// or
	head2 := req.Header.Get("Accept-Language")

	// prints a map of string
	//    fmt.Fprintln(writer, head)
	// prints a comma-delimited list of values
	writer.Write([]byte(head2))
}

func body(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "this is the body\n")

	// creates a var that represents the length of the message body in bytes,
	// not the header of the request.
	len := req.ContentLength
	// creates a slice of bytes that is equal to the length of the body message
	// of the request
	body := make([]byte, len)
	// reads all the body message into the body slice
	req.Body.Read(body)
	defer req.Body.Close()
	writer.Write([]byte(body))
}
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Printf("The handler %T was invoked\n", name)
		h(writer, req)
	}
}
func process(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "this is the process handler\n")
	file, header, err := req.FormFile("uploaded")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("File path:", header.Filename)
	fmt.Fprintln(writer, "Uploaded file %+v\n", header.Filename)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(writer, string(data))
	// this only parses and retireves the key-value pair of the form and omits the
	// URL key-value pairs.
	fmt.Fprintln(writer, req.PostForm)

	// prints both the URL and form KV pairs
	fmt.Fprintln(writer, req.Form)

	// MultiparseForm only contains the form KV pairs
	// only with "multipart/form-data" enctype
	// if you use application/x-www-form-urlencoded enctype it will return nil
	fmt.Fprintln(writer, req.MultipartForm)

	// access the KV pair from the form without have to parse the form
	// it does it by itself
	// with "application/x-www-form-urlencoded" :
	// it retrieves the first value associated with the given key from the request body,
	// and it ignores other values with the same key.
	// --------------
	// with  "multipart/form-data"
	// it  retrieves the first value associated with the given key from the request body,
	// and it ignores other values with the same key.
	// However, it also looks for the value in the URL query parameters
	// and the path parameters, and it returns the first value found.

	//	// this will display to the client the value that is found in the URL and not the form
	//    fmt.Fprintf(writer, req.FormValue("hello"))
	//	// this will display the values from the key "hello" of the form to the client
	//	fmt.Fprintln(writer, "(1)", req.FormValue("hello"))
	//	// this will display the value of the key "hello" from the Post to the client
	//	fmt.Fprintln(writer, "(2)", req.PostFormValue("hello"))
	//	// this will display the mapping of the KV pairs if the form using the "application-wwww"
	//	fmt.Fprintln(writer, "(3)", req.PostForm)
	//	// the is will diplay to the client the mapping to the KV pairs of the form
	//        // this can only be used when using the multipart-form data enctype
	//	fmt.Fprintln(writer, "(4)", req.MultipartForm)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//	http.HandleFunc("/header", headers)
	//    http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	fmt.Printf("Starting server.....\n")
	server.ListenAndServe()
}
