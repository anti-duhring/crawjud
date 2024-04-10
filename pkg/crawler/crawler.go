package crawler

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/anti-duhring/crawjud/pkg/utils"
	"github.com/anti-duhring/pdf"
)

func GetDIOUrl(fn func(client *http.Client) (string, error)) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return fn(client)
}

func DownloadDIOAndExtractProcesses(dioUrl string) (map[string]string, error) {
	file, err := DownloadContent(dioUrl)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := ExtractBytes(file)
	if err != nil {
		return nil, err
	}

	return ParseProcesses(content)
}

func DownloadContent(url string) (io.ReadCloser, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error when accessing TJPEcrawler url. Status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}

func ExtractText(file io.ReadCloser) (string, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return "", err
	}

	f, err := pdf.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return "", err
	}

	b, err := f.GetPlainText()
	if err != nil {
		return "", err
	}

	_, err = buf.ReadFrom(b)
	if err != nil {
		return "", err
	}
	text := buf.String()

	return text, nil
}

func ExtractBytes(file io.ReadCloser) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}

	f, err := pdf.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return nil, err
	}

	b, err := f.GetPlainText()
	if err != nil {
		return nil, err
	}

	_, err = buf.ReadFrom(b)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func ParseProcesses(content []byte) (map[string]string, error) {

	re, err := regexp.Compile(utils.ProcessNumberWithContextPattern)

	if err != nil {
		return nil, err
	}

	matches := re.FindAll(content, -1)
	processes := make(map[string]string)

	for _, match := range matches {
		re := regexp.MustCompile(utils.ProcessNumberPattern)
		processId := re.Find(match)
		key := string(processId)

		processes[key] = string(match)
	}

	return processes, nil
}
