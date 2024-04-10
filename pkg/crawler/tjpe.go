package crawler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var TJPEbaseUrl = "https://www2.tjpe.jus.br/dje/djeletronico?visaoId=tjdf.djeletronico.comum.internet.apresentacao.VisaoDiarioEletronicoInternetPorData"

func TJPE() (map[string]string, error) {

	dioUrl, err := GetDIOUrl(TJPEGetDIOUrl)
	if err != nil {
		return nil, err
	}

	processes, err := DownloadDIOAndExtractProcesses(dioUrl)
	if err != nil {
		return nil, err
	}

	return processes, nil
}

func TJPEGetDIOUrl(client *http.Client) (string, error) {
	res, err := client.Get(TJPEbaseUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("error when accessing TJPEcrawler url. Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	val, exists := doc.Find("#diariosConsultados_numero_0").Attr("value")

	if !exists {
		return "", errors.New("TJPE_DIO number not found not found")
	}

	dioUrl := fmt.Sprintf("https://www2.tjpe.jus.br/dje/DownloadServlet?dj=DJ%s_%s-ASSINADO.PDF&statusDoDiario=ASSINADO", val, time.Now().Format("2006"))

	return dioUrl, nil
}
