package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	// data -> is a slice of bytes
	data := []byte("Hello world\n")
	// err -> will handle any errors that might occur as the result of this call
		// the first parameter of the WriteFile() mehod is the name of the file
		// that will written to 
		// the second parameter is what that data which will be written to the file.
		// the third is the permissions on the file
			// 0 at the beginning is to indicate an octal number.
			// 6 sets the read-write permissions for the owner of the file.
			// 4 sets read-only permissions for the group that owns the file.
			// 4 sets read-only permissions for all other users who don't own the file or aren't in the group that owns the file.
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read, err := ioutil.ReadeFile("data1")
	if err != nil {
		fmt.Println("there is not file to read from.")
	}
	fmt.Println(string(read))

	// create a new file named "data2"
	file1, err := os.Create("data2")
	if err != nil {
		fmt.Println("uable to create the file")
	}
	// after the file is done being read close it
	defere file1.Close()

	// writer the data that is in bytes and writes it 
	// file1 which is called "data2"
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to files\n", bytes)

	// this will open up data2 that was created previously 
	file2, _ := os.Open("data2")
	// close onced open
	defer file2.Close()

	// creates a slice of bytes that has the length of the bytes repesentation
	// of the data variable
	read2 := make([]byte, len(data))
	// bytes -> count the number of bytes in the read2 slice
	bytes,_ := file2.Read(read2)
	// prints the number of bytes
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
