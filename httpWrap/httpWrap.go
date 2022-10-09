package httpWrap

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func MapToUrlValue(params map[string]string) url.Values {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}
	return values
}

func Get(hostUrl string, params url.Values, headers map[string]string) ([]byte, error) {
	return httpRequest("GET", hostUrl, params, headers, false, "")
}

func Delete(hostUrl string, params url.Values, headers map[string]string) ([]byte, error) {
	return httpRequest("DELETE", hostUrl, params, headers, false, "")
}

func Post(hostUrl string, params url.Values, headers map[string]string) ([]byte, error) {
	return httpRequest("POST", hostUrl, params, headers, false, "")
}

func PostToBody(hostUrl string, body string, headers map[string]string) ([]byte, error) {
	return httpRequest("POST", hostUrl, nil, headers, true, body)
}

func Put(hostUrl string, body string, headers map[string]string) ([]byte, error) {
	return httpRequest("PUT", hostUrl, nil, headers, true, body)
}

func httpRequest(method string, hostUrl string, params url.Values, headers map[string]string, isPostToBody bool, body string) ([]byte, error) {
	var request *http.Request
	var err error

	if method == "GET" || method == "DELETE" {
		if strings.Contains(hostUrl, "?") {
			hostUrl += params.Encode()
		} else {
			hostUrl += "?" + params.Encode()
		}
		request, err = http.NewRequest(method, hostUrl, nil)
	} else {
		if isPostToBody {
			request, err = http.NewRequest(method, hostUrl, strings.NewReader(body))
		} else {
			request, err = http.NewRequest(method, hostUrl, strings.NewReader(params.Encode()))
			if err == nil {
				request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			}
		}
	}

	if err != nil {
		return nil, err
	}

	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 404 {
		var body []byte
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			body, _ = ioutil.ReadAll(reader)
		default:
			body, _ = ioutil.ReadAll(response.Body)
		}

		return body, nil
	}

	return nil, errors.New("response status code " + strconv.Itoa(response.StatusCode))
}

func PostWithFile(url string, param url.Values, files map[string]string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range param {
		for _, item := range v {
			_ = writer.WriteField(k, item)
		}
	}

	for k, v := range files {
		file, err := os.Open(v)
		if err != nil {
			return nil, err
		}

		part, err := writer.CreateFormFile(k, v)
		if err == nil {
			_, err = io.Copy(part, file)
		}

		file.Close()
	}

	err := writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, body)

	formType := writer.FormDataContentType()
	request.Header.Set("Content-Type", formType)

	client := &http.Client{}
	response, err := client.Do(request)
	if nil != err {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func Download(filename string, url string) error {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	return err
}
