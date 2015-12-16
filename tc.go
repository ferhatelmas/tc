package tc

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

type response struct {
	XMLName xml.Name
	Body    body
}

type body struct {
	XMLName  xml.Name
	Response result `xml:"TCKimlikNoDogrulaResponse"`
}

type result struct {
	XMLName xml.Name `xml:"TCKimlikNoDogrulaResponse"`
	Result  bool     `xml:"TCKimlikNoDogrulaResult"`
}

var reqBody = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
      <TCKimlikNoDogrula xmlns="http://tckimlik.nvi.gov.tr/WS">
	<TCKimlikNo>%s</TCKimlikNo>
	<Ad>%s</Ad>
	<Soyad>%s</Soyad>
	<DogumYili>%d</DogumYili>
      </TCKimlikNoDogrula>
  </soap:Body>
</soap:Envelope>
`

// IsValidFor checks given TC No is valid or not for given
// first, last name and birth year.
func IsValidFor(id, firstName, lastName string, birthYear int) (bool, error) {
	if len(id) != 11 {
		return false, fmt.Errorf("TC No must be 11 characters")
	}
	r := strings.NewReader(fmt.Sprintf(reqBody, id, toUpper(firstName), toUpper(lastName), birthYear))
	req, err := http.NewRequest("POST", "https://tckimlik.nvi.gov.tr/Service/KPSPublic.asmx?WSDL", r)
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPAction", "http://tckimlik.nvi.gov.tr/WS/TCKimlikNoDogrula")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var b response
	if err = xml.Unmarshal(res, &b); err != nil {
		return false, err
	}
	return b.Body.Response.Result, nil
}

func toUpper(s string) string {
	return strings.ToUpperSpecial(unicode.TurkishCase, s)
}

// IsValid verifies given identification is a valid possible id.
func IsValid(id string) bool {
	return len(id) == 11 && validChars(id) && tenthDigit(id) && eleventhDigit(id)
}

func validChars(id string) bool {
	for _, c := range id {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func tenthDigit(id string) bool {
	return int(id[9])-'0' == mode(3*sum(id, 0, 9, 2)+sum(id, 1, 8, 2))
}

func eleventhDigit(id string) bool {
	return int(id[10])-'0' == mode(sum(id, 0, 9, 2)+3*sum(id, 1, 10, 2))
}

func mode(x int) int {
	x = (10 - x) % 10
	if x < 0 {
		x += 10
	}
	return x
}

func sum(id string, s, e, t int) int {
	var n int
	for ; s < e; s += t {
		n += int(id[s]) - '0'
	}
	return n
}
