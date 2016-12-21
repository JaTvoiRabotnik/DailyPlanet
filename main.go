package main

import (
	"io"
	"net/http"
	"os"

	"google.golang.org/api/storagetransfer/v1"
)

// Copies a file from the URL provided into a local file.
func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Transfer the file from the local storage into the Google storage bucket
func transferToStorage(filepath string, projectID string) (err error) {
	service_account, err := storagetransfer.Get(projectID)
	if err != nil {
		return err
	}
	return nil
}
