package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {

	fileUrl := "https://www.youtube.com/embed/P_z9LhH-924?autoplay=0&mute=0&controls=0&origin=https%3A%2F%2Frecord-playback-2.kynaforkids.vn&playsinline=1&showinfo=0&rel=0&iv_load_policy=3&modestbranding=1&enablejsapi=1&widgetid=1"

	// Download the file, params:
	// 1) name of file to save as
	// 2) URL to download FROM
	err := DownloadFile(fileUrl)
	if err != nil {
		fmt.Println("Error downloading file: ", err)
		return
	}

	fmt.Println("Downloaded: " + fileUrl)
}

func DownloadFile(url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	filepath := path.Base(resp.Request.URL.String())
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DownloadFile will download from a given url to a file. It will
// write as it downloads (useful for large files).
func DownloadFileStaticName(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
