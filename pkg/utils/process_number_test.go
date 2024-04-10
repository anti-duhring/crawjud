package utils_test

import (
	"testing"

	"github.com/anti-duhring/crawjud/pkg/utils"
)

func TestGetCourtByProcessNumber(t *testing.T) {
	t.Run("Should return TJAC when process number is 0000000-01.0000.00.0000", func(t *testing.T) {
		processNumber := "0000000-00.0000.0.01.0000"
		expected := "TJAC"
		result, err := utils.GetCourtByProcessNumber(processNumber)

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			return
		}

		if *result != expected {
			t.Errorf("Expected %s, got %s", expected, *result)
		}
	})

	t.Run("Should return TJAL when process number is 0000000-02.0000.00.0000", func(t *testing.T) {
		processNumber := "0000000-00.0000.0.02.0000"
		expected := "TJAL"
		result, err := utils.GetCourtByProcessNumber(processNumber)

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			return
		}

		if *result != expected {
			t.Errorf("Expected %s, got %s", expected, *result)
		}
	})

	t.Run("Should return TJAP when process number is 0000000-03.0000.00.0000", func(t *testing.T) {
		processNumber := "0000000-00.0000.0.03.0000"
		expected := "TJAP"
		result, err := utils.GetCourtByProcessNumber(processNumber)

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			return
		}

		if *result != expected {
			t.Errorf("Expected %s, got %s", expected, *result)
		}
	})

	t.Run("Should return TJPE when process number is 0001459-18.2017.8.17.3130", func(t *testing.T) {
		processNumber := "0001459-18.2017.8.17.3130"
		expected := "TJPE"
		result, err := utils.GetCourtByProcessNumber(processNumber)

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			return
		}

		if *result != expected {
			t.Errorf("Expected %s, got %s", expected, *result)
		}
	})
}
