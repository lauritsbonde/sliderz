package filehandler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FileUpload(w http.ResponseWriter, req *http.Request){
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum of 10MB of memory
	req.ParseMultipartForm(10 << 20)

	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename, the Header and the size of the file
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Ensure the temp-files directory exists
	if _, err := os.Stat("./temp-files"); os.IsNotExist(err) {
		os.Mkdir("./temp-files", 0755)
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	// We prefix the filename with "upload-" and use the current Unix timestamp
	tempFile, err := os.CreateTemp("./temp-files", "upload-*"+handler.Filename)
	if err != nil {
		fmt.Println("Error Creating Temp File")
		fmt.Println(err)
	}
	defer tempFile.Close()

	// Read all of the contents of our uploaded file into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// Write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// Return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}