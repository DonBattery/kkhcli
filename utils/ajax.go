package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/pkg/errors"
)

// HTTPHeaderJSON is Content-Type : application/json
var (
	HTTPHeaderJSON = HTTPHeader{
		key:   "Content-Type",
		value: "application/json",
	}
)

// sendRequest creates a HTTP request with the given Method, to the given URL,
// with (optional) body and HTTP headers.
// The function returns the response body's byte-representation (JSON), and on optional error.
func sendRequest(method string, url string, body []byte, headers ...HTTPHeader) ([]byte, error) {
	var response []byte

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		err = errors.Wrapf(err, "Cannot create %v request to %v\n", method, url)
		return response, err
	}
	for _, header := range headers {
		req.Header.Set(header.key, header.value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "Cannot send request\n")
		return response, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			closeErr = errors.Wrap(closeErr, "Error closing response body\n")
			log.Fatal(closeErr)
		}
	}()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "Error reading response body\n")
		return response, err
	}
	if resp.StatusCode != 200 {
		return response, errors.Errorf("Server answered with a non-200 status: %v\n", resp.StatusCode)
	}
	return response, nil
}

// SendCommand Issues a command to the KKHC Server
func SendCommand(command string, body []byte) ([]byte, error) {
	adminHeader := HTTPHeader{"adminpassword", viper.GetString("adminPass")}
	commandHeader := HTTPHeader{"command", command}
	return sendRequest("POST", viper.GetString("apiURL"), body, HTTPHeaderJSON, adminHeader, commandHeader)
}

// Creates a new file upload http request
func newfileUploadRequest(url, fileName, path string) (*http.Request, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot open path %s", path)
	}
	defer file.Close()
	if _, err := file.Stat(); err != nil {
		return nil, errors.Wrapf(err, "Cannot read path %s", path)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("adminpassword", viper.GetString("adminPass"))
	req.Header.Set("command", "addAvatars")
	return req, err

}

// SendAvatar sends an avatar pic to the KKHC server
func SendAvatar(path string) ([]byte, error) {
	var response []byte

	req, err := newfileUploadRequest(viper.GetString("apiURL"), "avatar", path)
	if err != nil {
		return response, errors.Wrap(err, "Cannot create request\n")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "Cannot send request\n")
		return response, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			closeErr = errors.Wrap(closeErr, "Error closing response body\n")
			log.Fatal(closeErr)
		}
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "Error reading response body\n")
		return response, err
	}
	if resp.StatusCode != 200 {
		return response, errors.Errorf("Server answered with a non-200 status: %v\n", resp.StatusCode)
	}
	return response, nil

}
