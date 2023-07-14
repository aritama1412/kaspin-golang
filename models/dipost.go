package models

type Dipost struct {
	TimeStamp      string `json:"timeStamp"`
	Imid           string `json:"iMid"`
	ReferenceNo    string `json:"referenceNo"`
	Amt            int    `json:"amt"`
	Currency       string `json:"currency"`
	MerchantToken  string `json:"merchantToken"`
	PayMethod      string `json:"payMethod"`
	InstmntMon     string `json:"instmntMon"`
	Description    string `json:"description"`
	GoodsNm        string `json:"goodsNm"`
	BillingNm      string `json:"billingNm"`
	BillingPhone   string `json:"billingPhone"`
	BillingCity    string `json:"billingCity"`
	BillingState   string `json:"billingState"`
	BillingPostCd  string `json:"billingPostCd"`
	BillingCountry string `json:"billingCountry"`
	BillingEmail   string `json:"billingEmail"`
	DbProcessUrl   string `json:"dbProcessUrl"`
}
