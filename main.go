package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kaspin-golang/models"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/registration", registrationHandler)
	http.HandleFunc("/inquiry", inquiryHandler)
	http.HandleFunc("/payment", paymentHandler)

	// Start the HTTP server
	fmt.Println("server started at: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("data: ", err)
		return
	}

	// Parse the request body into a struct
	var req models.Dipost
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("Error parsing request body:", err)
		return
	}

	post := models.Dipost{
		TimeStamp:      req.TimeStamp,
		Imid:           req.Imid,
		ReferenceNo:    req.ReferenceNo,
		Amt:            req.Amt,
		Currency:       req.Currency,
		MerchantToken:  req.MerchantToken,
		PayMethod:      req.PayMethod,
		InstmntMon:     req.InstmntMon,
		Description:    req.Description,
		GoodsNm:        req.GoodsNm,
		BillingNm:      req.BillingNm,
		BillingPhone:   req.BillingPhone,
		BillingCity:    req.BillingCity,
		BillingState:   req.BillingState,
		BillingPostCd:  req.BillingPostCd,
		BillingCountry: req.BillingCountry,
		BillingEmail:   req.BillingEmail,
		DbProcessUrl:   req.DbProcessUrl,
	}

	json_data, err := json.Marshal(post)
	if err != nil {
		fmt.Print("er c: ", err)
	}
	response, err := http.Post("https://dev.nicepay.co.id/nicepay/direct/v2/registration", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Print("er b: ", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("er a: ", err)
	}

	var responseObject models.Diterima
	json.Unmarshal(responseData, &responseObject)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func inquiryHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("data: ", err)
		return
	}

	// Parse the request body into a struct
	var req models.InquiryPost
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("Error parsing request body:", err)
		return
	}

	post := models.InquiryPost{
		TimeStamp:     req.TimeStamp,
		MerchantToken: req.MerchantToken,
		ReferenceNo:   req.ReferenceNo,
		TXid:          req.TXid,
		Imt:           req.Imt,
		IMid:          req.IMid,
	}

	json_data, err := json.Marshal(post)
	if err != nil {
		fmt.Print("er c: ", err)
	}
	response, err := http.Post("https://dev.nicepay.co.id/nicepay/direct/v2/inquiry", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Print("er b: ", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("er a: ", err)
	}

	var responseObject models.InquiryReceived
	json.Unmarshal(responseData, &responseObject)

	// Check if the status is 9 and update the struct accordingly
	if responseObject.Status == "9" {
		responseObject.AdditionalInfo = "Payment - Initialization / Reversal"
	} else if responseObject.Status == "0" {
		responseObject.AdditionalInfo = "Payment - Success"
	} else if responseObject.Status == "1" {
		responseObject.AdditionalInfo = "Payment - Failed"
	} else if responseObject.Status == "2" {
		responseObject.AdditionalInfo = "Payment - Void/Refund"
	}

	// Marshal the updated struct to JSON
	updatedResponseData, err := json.Marshal(responseObject)
	if err != nil {
		fmt.Println("Error marshaling updated response data:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(updatedResponseData)
}

type ResponseData struct {
	ResultMsg string `json:"resultMsg"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	paramValues := r.URL.Query()

	// Convert the parameters to URL-encoded form data
	formData := url.Values{}
	for key, values := range paramValues {
		for _, value := range values {
			formData.Add(key, value)
		}
	}
	body := strings.NewReader(formData.Encode())

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", "https://dev.nicepay.co.id/nicepay/direct/v2/payment", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	bodyx, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the value of resultMsg using a regular expression
	resultMsgRegex := regexp.MustCompile(`<input type="hidden" name="resultMsg" value="([^"]+)"`)
	match := resultMsgRegex.FindStringSubmatch(string(bodyx))
	var resultMsg string
	if len(match) >= 2 {
		resultMsg = match[1]
	} else {
		resultMsg = "Result Message not found"
	}

	// Create a ResponseData struct with the resultMsg value
	responseData := ResponseData{
		ResultMsg: resultMsg,
	}

	// Convert the struct to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		log.Fatal(err)
	}

	// Print the JSON data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
