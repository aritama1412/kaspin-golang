package models

type Diterima struct {
	ResultCd     string `json:"resultCd"`
	ResultMsg    string `json:"resultMsg"`
	TXid         string `json:"tXid"`
	ReferenceNo  string `json:"referenceNo"`
	PayMethod    string `json:"payMethod"`
	Amt          string `json:"amt"`
	TransDt      string `json:"transDt"`
	TransTm      string `json:"transTm"`
	Description  string `json:"description"`
	BankCd       int    `json:"bankCd"`
	VacctNo      int    `json:"vacctNo"`
	MitraCd      int    `json:"mitraCd"`
	PayNo        int    `json:"payNo"`
	Currency     int    `json:"currency"`
	GoodsNm      int    `json:"goodsNm"`
	BillingNm    int    `json:"billingNm"`
	VacctValidDt int    `json:"vacctValidDt"`
	VacctValidTm int    `json:"vacctValidTm"`
	PayValidDt   int    `json:"payValidDt"`
	PayValidTm   int    `json:"payValidTm"`
	RequestURL   int    `json:"requestURL"`
	PaymentExpDt int    `json:"paymentExpDt"`
	PaymentExpTm int    `json:"paymentExpTm"`
	QrContent    int    `json:"qrContent"`
	QrUrl        int    `json:"qrUrl"`
}

// ResultCd     string `json:"resultCd"`
// ResultMsg    string `json:"resultMsg"`
// TXid         string `json:"tXid"`
// ReferenceNo  string `json:"referenceNo"`
// PayMethod    int    `json:"payMethod"`
// Amt          int    `json:"amt"`
// TransDt      string `json:"transDt"`
// TransTm      string `json:"transTm"`
// Description  string `json:"description"`
// BankCd       string `json:"bankCd"`
// VacctNo      string `json:"vacctNo"`
// MitraCd      string `json:"mitraCd"`
// PayNo        string `json:"payNo"`
// Currency     string `json:"currency"`
// GoodsNm      string `json:"goodsNm"`
// BillingNm    string `json:"billingNm"`
// VacctValidDt string `json:"vacctValidDt"`
// VacctValidTm string `json:"vacctValidTm"`
// PayValidDt   string `json:"payValidDt"`
// PayValidTm   string `json:"payValidTm"`
// RequestURL   string `json:"requestURL"`
// PaymentExpDt string `json:"paymentExpDt"`
// PaymentExpTm string `json:"paymentExpTm"`
// QrContent    string `json:"qrContent"`
// QrUrl        string `json:"qrUrl"`
