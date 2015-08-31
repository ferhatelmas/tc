package tc

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// No - models TC Number
type No struct {
	ID        string
	FirstName string
	LastName  string
	BirthYear int
}

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

// IsValid checks given TC No is valid or not for given
// first, last name and birth year.
func IsValid(id, firstName, lastName string, birthYear int) (bool, error) {
	if len(id) != 11 {
		return false, fmt.Errorf("TC No must be 11 characters")
	}
	r := strings.NewReader(fmt.Sprintf(reqBody, id, strings.ToUpper(firstName), strings.ToUpper(lastName), birthYear))
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
