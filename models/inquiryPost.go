package models

type InquiryPost struct {
	TimeStamp     string `json:"timeStamp"`
	MerchantToken string `json:"merchantToken"`
	ReferenceNo   string `json:"referenceNo"`
	TXid          string `json:"tXid"`
	Imt           int    `json:"amt"`
	IMid          string `json:"iMid"`
}
