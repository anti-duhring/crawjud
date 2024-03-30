package crawler_test

import (
	"testing"

	"github.com/anti-duhring/crawjud/internal/crawler"
)

func TestTJPE(t *testing.T) {
	result, err := crawler.TJPE()
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result == nil {
		t.Errorf("Error: result is empty")
	}

	if len(result) == 0 {
		t.Errorf("Error: result is empty")
	}

}
