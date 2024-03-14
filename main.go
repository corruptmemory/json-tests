package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
)

const idString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const sampleCount = 1000000

var currencies = []string{
	"AFN",
	"EUR",
	"ALL",
	"DZD",
	"USD",
	"EUR",
	"AOA",
	"XCD",
	"XCD",
	"ARS",
	"AMD",
	"AWG",
	"AUD",
	"EUR",
	"AZN",
	"BSD",
	"BHD",
	"BDT",
	"BBD",
	"BYN",
	"EUR",
	"BZD",
	"XOF",
	"BMD",
	"INR",
	"BTN",
	"BOB",
	"BOV",
	"USD",
	"BAM",
	"BWP",
	"NOK",
	"BRL",
	"USD",
	"BND",
	"BGN",
	"XOF",
	"BIF",
	"CVE",
	"KHR",
	"XAF",
	"CAD",
	"KYD",
	"XAF",
	"XAF",
	"CLP",
	"CLF",
	"CNY",
	"AUD",
	"AUD",
	"COP",
	"COU",
	"KMF",
	"CDF",
	"XAF",
	"NZD",
	"CRC",
	"XOF",
	"EUR",
	"CUP",
	"CUC",
	"ANG",
	"EUR",
	"CZK",
	"DKK",
	"DJF",
	"XCD",
	"DOP",
	"USD",
	"EGP",
	"SVC",
	"USD",
	"XAF",
	"ERN",
	"EUR",
	"SZL",
	"ETB",
	"EUR",
	"FKP",
	"DKK",
	"FJD",
	"EUR",
	"EUR",
	"EUR",
	"XPF",
	"EUR",
	"XAF",
	"GMD",
	"GEL",
	"EUR",
	"GHS",
	"GIP",
	"EUR",
	"DKK",
	"XCD",
	"EUR",
	"USD",
	"GTQ",
	"GBP",
	"GNF",
	"XOF",
	"GYD",
	"HTG",
	"USD",
	"AUD",
	"EUR",
	"HNL",
	"HKD",
	"HUF",
	"ISK",
	"INR",
	"IDR",
	"XDR",
	"IRR",
	"IQD",
	"EUR",
	"GBP",
	"ILS",
	"EUR",
	"JMD",
	"JPY",
	"GBP",
	"JOD",
	"KZT",
	"KES",
	"AUD",
	"KPW",
	"KRW",
	"KWD",
	"KGS",
	"LAK",
	"EUR",
	"LBP",
	"LSL",
	"ZAR",
	"LRD",
	"LYD",
	"CHF",
	"EUR",
	"EUR",
	"MOP",
	"MKD",
	"MGA",
	"MWK",
	"MYR",
	"MVR",
	"XOF",
	"EUR",
	"USD",
	"EUR",
	"MRU",
	"MUR",
	"EUR",
	"XUA",
	"MXN",
	"MXV",
	"USD",
	"MDL",
	"EUR",
	"MNT",
	"EUR",
	"XCD",
	"MAD",
	"MZN",
	"MMK",
	"NAD",
	"ZAR",
	"AUD",
	"NPR",
	"EUR",
	"XPF",
	"NZD",
	"NIO",
	"XOF",
	"NGN",
	"NZD",
	"AUD",
	"USD",
	"NOK",
	"OMR",
	"PKR",
	"USD",
	"PAB",
	"USD",
	"PGK",
	"PYG",
	"PEN",
	"PHP",
	"NZD",
	"PLN",
	"EUR",
	"USD",
	"QAR",
	"EUR",
	"RON",
	"RUB",
	"RWF",
	"EUR",
	"SHP",
	"XCD",
	"XCD",
	"EUR",
	"EUR",
	"XCD",
	"WST",
	"EUR",
	"STN",
	"SAR",
	"XOF",
	"RSD",
	"SCR",
	"SLE",
	"SGD",
	"ANG",
	"XSU",
	"EUR",
	"EUR",
	"SBD",
	"SOS",
	"ZAR",
	"SSP",
	"EUR",
	"LKR",
	"SDG",
	"SRD",
	"NOK",
	"SEK",
	"CHF",
	"CHE",
	"CHW",
	"SYP",
	"TWD",
	"TJS",
	"TZS",
	"THB",
	"USD",
	"XOF",
	"NZD",
	"TOP",
	"TTD",
	"TND",
	"TRY",
	"TMT",
	"USD",
	"AUD",
	"UGX",
	"UAH",
	"AED",
	"GBP",
	"USD",
	"USD",
	"USN",
	"UYU",
	"UYI",
	"UYW",
	"UZS",
	"VUV",
	"VES",
	"VED",
	"VND",
	"USD",
	"USD",
	"XPF",
	"MAD",
	"YER",
	"ZMW",
	"ZWL",
	"XBA",
	"XBB",
	"XBC",
	"XBD",
	"XTS",
	"XXX",
	"XAU",
	"XPD",
	"XPT",
	"XAG",
}

var samples []Transaction

type MessageContext struct {
	ID uuid.UUID `json:"id"`
}

type Transaction struct {
	Context         MessageContext `json:"context"`
	Time            time.Time      `json:"time"`
	SourceAccountID string         `json:"source_account_id"`
	TargetAccountID string         `json:"target_account_id"`
	Amount          *big.Float     `json:"amount"`
	Currency        string         `json:"currency"`
	FillerField1    string         `json:"filler_field_1"`
	FillerField2    string         `json:"filler_field_2"`
	FillerField3    int            `json:"filler_field_3"`
	FillerField4    string         `json:"filler_field_4"`
	FillerField5    string         `json:"filler_field_5"`
	FillerField6    string         `json:"filler_field_6"`
	FillerField7    int            `json:"filler_field_7"`
	FillerField8    string         `json:"filler_field_8"`
	FillerField9    string         `json:"filler_field_9"`
	FillerField10   string         `json:"filler_field_10"`
}

func generateRandomCurrency(rng *rand.Rand) string {
	return currencies[rng.Intn(len(currencies))]
}

func generateRandomContext() MessageContext {
	return MessageContext{
		ID: uuid.New(),
	}
}

func generateRandomAccountID(rng *rand.Rand) string {
	length := rng.Intn(5) + 10
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = idString[rng.Intn(len(idString))]
	}
	return string(id)
}

func generateRandomAmount(rng *rand.Rand) *big.Float {
	return big.NewFloat(rng.Float64() * 1000.0)
}

func generateRandomTransaction(rng *rand.Rand) Transaction {
	return Transaction{
		Context:         generateRandomContext(),
		Time:            time.Now(),
		SourceAccountID: generateRandomAccountID(rng),
		TargetAccountID: generateRandomAccountID(rng),
		Amount:          generateRandomAmount(rng),
		Currency:        generateRandomCurrency(rng),
		FillerField1:    generateRandomAccountID(rng),
		FillerField2:    generateRandomAccountID(rng),
		FillerField3:    rng.Intn(1000),
		FillerField4:    generateRandomAccountID(rng),
		FillerField5:    generateRandomAccountID(rng),
		FillerField6:    generateRandomAccountID(rng),
		FillerField7:    rng.Intn(1000),
		FillerField8:    generateRandomAccountID(rng),
		FillerField9:    generateRandomAccountID(rng),
		FillerField10:   generateRandomAccountID(rng),
	}
}

func singleThreadedTest() {
	encoded := make([][]byte, sampleCount)

	encodeStart := time.Now()
	for i := 0; i < sampleCount; i++ {
		encoded[i], _ = json.Marshal(samples[i])
	}
	encodeEnd := time.Now()
	encodeDuration := encodeEnd.Sub(encodeStart)
	fmt.Printf("Encoding %d samples took %v or %f per second\n", sampleCount, encodeDuration, float64(sampleCount)/encodeDuration.Seconds())

	decodedSamples := make([]Transaction, sampleCount)
	decodeStart := time.Now()
	for i := 0; i < sampleCount; i++ {
		json.Unmarshal(encoded[i], &decodedSamples[i])
	}
	decodeEnd := time.Now()
	decodeDuration := decodeEnd.Sub(decodeStart)
	fmt.Printf("Decoding %d samples took %v or %f per second\n", sampleCount, decodeDuration, float64(sampleCount)/decodeDuration.Seconds())
	fmt.Printf("Encoding/Decoding %d samples took %v or %f per second\n", sampleCount, encodeDuration+decodeDuration, float64(sampleCount)/(encodeDuration+decodeDuration).Seconds())
}

func multiThreadedBody() {
	encoded := make([][]byte, sampleCount)

	for i := 0; i < sampleCount; i++ {
		encoded[i], _ = json.Marshal(samples[i])
	}
	decodedSamples := make([]Transaction, sampleCount)
	for i := 0; i < sampleCount; i++ {
		json.Unmarshal(encoded[i], &decodedSamples[i])
	}
}

func multiThreadedTest() {
	wg := &sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			multiThreadedBody()
			wg.Done()
		}()
	}
	startTime := time.Now()
	wg.Wait()
	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)
	totalSamples := sampleCount * runtime.NumCPU()
	fmt.Printf("Multi-threaded test took %v\n", endTime.Sub(startTime))
	fmt.Printf("Encoding/Decoding %d samples took %v or %f per second\n", totalSamples, endTime.Sub(startTime), float64(totalSamples)/totalDuration.Seconds())
}

func main() {
	fmt.Printf("Generating %d samples\n", sampleCount)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	samples = make([]Transaction, sampleCount)
	for i := 0; i < sampleCount; i++ {
		samples[i] = generateRandomTransaction(rng)
	}
	fmt.Println("*** Starting single threaded test ***")
	singleThreadedTest()
	fmt.Println("*** Starting multi threaded test ***")
	multiThreadedTest()
}
