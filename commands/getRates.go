package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/xml"
	"io/ioutil"
	"github.com/mike-neutron/go_currency_app/src/initializers"
	"github.com/mike-neutron/go_currency_app/src/models"
	"time"
	"strings"
	"strconv"
	"gorm.io/datatypes"
	"golang.org/x/exp/slices"
)

type RateXML struct {
	XMLName xml.Name `xml:"ValCurs"`
	RateXMLItem []struct {
		NumCode    string  `xml:"NumCode"`
		CharCode   string  `xml:"CharCode"`
		Nominal    uint    `xml:"Nominal"`
		Name       string  `xml:"Name"`
		Value      string  `xml:"Value"`
	} `xml:"Valute"`
}

func init() {
	config, err := initializers.LoadConfig("../.env")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	// Get response from cbr
	response, error := http.Get("https://www.cbr-xml-daily.ru/daily_utf8.xml")
	if error != nil {
	    fmt.Println(error)
	}

	// Parse request
	body, error := ioutil.ReadAll(response.Body); 
	if error != nil {
		fmt.Println(error)
	}

	// close response body
	defer response.Body.Close()

	// Make rates structs
	rates := new(RateXML)
	err := xml.Unmarshal([]byte(body), rates)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Get existed values
	now := datatypes.Date(time.Now())
	var existed []models.Rate
	initializers.DB.Where("date = ?", now).Find(&existed)

	// Make existed array
	var existedKeys []string
	for _, existedRate := range existed {
		existedKeys = append(existedKeys, existedRate.Code)
	}

	for _, rate := range rates.RateXMLItem {
		if slices.Contains(existedKeys, rate.CharCode) {
			continue
		}

		value := strings.Replace(rate.Value, ",", ".", -1)
		valueFloat64, err := strconv.ParseFloat(value, 32)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		valueFloat32 := float32(valueFloat64)
		newRate := models.Rate{
			Code: rate.CharCode,
			Rate: valueFloat32,
			Date: now,
		}
		initializers.DB.Create(&newRate)
	}
}