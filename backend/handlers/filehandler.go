package filehandler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func readFile(fileName string, startPage int, endPage int) {
	fmt.Println("\nReading file: " + fileName, "Start Page:", startPage, "End Page:", endPage)

	args := []string{
		"-layout",              // Maintain (as best as possible) the original physical layout of the text.
		"-nopgbrk",             // Don't insert page breaks (form feed characters) between pages.
	}
	
	if startPage > 0 {
		args = append(args, fmt.Sprintf("-f %d", startPage))
	}

	if endPage > 0 {
		if endPage < startPage {
			fmt.Println("End page cannot be less than start page")
			return
		}
		args = append(args, fmt.Sprintf("-l %d", endPage))
	}

	endArgs := []string{
		fileName, // The input file.
		"-",                    // Send the output to stdout.
	}

	args = append(args, endArgs...)

	cmd := exec.CommandContext(context.Background(), "pdftotext", args...)

	fmt.Printf("Command: %v\n", cmd)

	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running pdftotext")
		fmt.Println(err)
		fmt.Println("Stderr:", errBuf.String())
		return
	}

	fmt.Println(outBuf.String())
}

func deleteFile(fileName string) {
	// delete file
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted file: " + fileName)
}

func FileUpload(w http.ResponseWriter, req *http.Request){
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum of 10MB of memory
	req.ParseMultipartForm(10 << 20)

	var startPage, startErr = strconv.Atoi(req.FormValue("startPage"))
	var endPage, endErr = strconv.Atoi(req.FormValue("endPage"))
	if startErr != nil || endErr != nil {
		fmt.Println("Error Parsing Start and End Page")
		fmt.Println(startErr)
		fmt.Println(endErr)
		return
	}

	// FormFile returns the first file for the given key `file`
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
	tempFile, err := os.CreateTemp("./temp-files", "upload-*"+handler.Filename)
	if err != nil {
		fmt.Println("Error Creating Temp File")
		fmt.Println(err)
	}
	defer tempFile.Close()
	// defer deleteFile(tempFile.Name())

	// Read all of the contents of our uploaded file into a byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// Write this byte array to our temporary file
	tempFile.Write(fileBytes)

	// Return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	readFile(tempFile.Name(), startPage, endPage)
}