package utils

import (
	"fmt"
	"regexp"
)

// Ref: https://www.cnj.jus.br/programas-e-acoes/numeracao-unica/ and https://atos.cnj.jus.br/atos/detalhar/atos-normativos?documento=119
// Example: 0000001-12.2014.2.00.0000
var (
	ProcessNumberPattern            = `\d{7}\-\d{2}\.\d{4}\.\d\.\d{2}\.\d{4}`
	ProcessNumberWithContextPattern = `\d{7}\-\d{2}\.\d{4}\.\d\.\d{2}\.\d{4}.{0,500}`
)

var courtCodes = map[string]string{
	"01": "TJAC",
	"02": "TJAL",
	"03": "TJAP",
	"04": "TJAM",
	"05": "TJBA",
	"06": "TJCE",
	"07": "TJDF",
	"08": "TJES",
	"09": "TJGO",
	"10": "TJMA",
	"11": "TJMT",
	"12": "TJMS",
	"13": "TJMG",
	"14": "TJPA",
	"15": "TJPB",
	"16": "TJPR",
	"17": "TJPE",
	"18": "TJPI",
	"19": "TJRJ",
	"20": "TJRN",
	"21": "TJRS",
	"22": "TJRO",
	"23": "TJRR",
	"24": "TJSC",
	"25": "TJSP",
	"26": "TJSE",
	"27": "TJTO",
}

func GetCourtByProcessNumber(processNumber string) (*string, error) {
	pattern := `\d{7}-\d{2}\.\d{4}\.\d\.(?P<part02>\d{2})\.\d{4}`

	matches := regexp.MustCompile(pattern).FindStringSubmatch(processNumber)
	if len(matches) < 2 {
		return nil, fmt.Errorf("Process number %s is invalid", processNumber)
	}

	courtCode := matches[1]
	court, ok := courtCodes[courtCode]
	if !ok {
		unknownCourt := "UNKNOWN"

		return &unknownCourt, nil
	}

	return &court, nil
}
