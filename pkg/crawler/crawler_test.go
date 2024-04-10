package crawler_test

import (
	"testing"

	"github.com/anti-duhring/crawjud/pkg/crawler"
)

func TestGetDIOUrl(t *testing.T) {
	result, err := crawler.GetDIOUrl(crawler.TJPEGetDIOUrl)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result == "" {
		t.Errorf("Error: result is empty")
	}

	t.Logf("Result: %s", result)
}

func TestDownloadContent(t *testing.T) {
	url, err := crawler.GetDIOUrl(crawler.TJPEGetDIOUrl)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	result, err := crawler.DownloadContent(url)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result == nil {
		t.Errorf("Error: result is empty")
	}

}

func TestExtractText(t *testing.T) {
	url, err := crawler.GetDIOUrl(crawler.TJPEGetDIOUrl)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	result, err := crawler.DownloadContent(url)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	text, err := crawler.ExtractText(result)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if text == "" {
		t.Errorf("Error: text is empty")
	}

}

func TestExtractBytes(t *testing.T) {
	url, err := crawler.GetDIOUrl(crawler.TJPEGetDIOUrl)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	result, err := crawler.DownloadContent(url)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	text, err := crawler.ExtractBytes(result)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if text == nil {
		t.Errorf("Error: text is empty")
	}
}

func TestParseProcesses(t *testing.T) {
	url, err := crawler.GetDIOUrl(crawler.TJPEGetDIOUrl)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	result, err := crawler.DownloadContent(url)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	content, err := crawler.ExtractBytes(result)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	processes, err := crawler.ParseProcesses(content)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if processes == nil {
		t.Errorf("Error: processes is empty")
	}

	for k, v := range processes {
		if k == "" {
			t.Errorf("Error: key is empty")
		}
		if v == "" {
			t.Errorf("Error: value is empty")
		}
	}

	t.Logf("Processes found: %d", len(processes))
}
