package gofastbill

type RevenueGet_Request struct {
	INVOICE_ID     string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`         //Rechnungs-ID
	INVOICE_NUMBER string `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"` //Rechnungsnummer
	INVOICE_TITLE  string `xml:"INVOICE_TITLE,omitempty" json:"INVOICE_TITLE,omitempty"`   //Rechnungstitel
	CUSTOMER_ID    string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`       //Eine bestimmte Kundennummer
	MONTH          string `xml:"MONTH,omitempty" json:"MONTH,omitempty"`                   //Monat
	YEAR           string `xml:"YEAR,omitempty" json:"YEAR,omitempty"`                     //Jahr
	START_DUE_DATE string `xml:"START_DUE_DATE,omitempty" json:"START_DUE_DATE,omitempty"` //Rechnungen die nach einem bestimmten Datum fällig werden
	END_DUE_DATE   string `xml:"END_DUE_DATE,omitempty" json:"END_DUE_DATE,omitempty"`     //Rechnungen die vor einem bestimmten Datum fällig werden
	TYPE           string `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                     //Zahlungsart
}

//GET Revenue
//FILTER => INVOICE_ID, INVOICE_NUMBER, INVOICE_TITLE, CUSTOMER_ID, MONTH, YEAR,START_DUE_DATE, END_DUE_DATE, TYPE
//Querying the details of one or more invoices. If no filter is set, 10 invoices will be returned. When LIMIT is set, up to 100 invoices are possible
func (s *Initialization) Revenue_get(req RevenueGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "revenue.get"

	r.FILTER.INVOICE_ID = req.INVOICE_ID
	r.FILTER.INVOICE_NUMBER = req.INVOICE_NUMBER
	r.FILTER.INVOICE_TITLE = req.INVOICE_TITLE
	r.FILTER.CUSTOMER_ID = req.CUSTOMER_ID
	r.FILTER.MONTH = req.MONTH
	r.FILTER.YEAR = req.YEAR
	r.FILTER.START_DUE_DATE = req.START_DUE_DATE
	r.FILTER.END_DUE_DATE = req.END_DUE_DATE
	r.FILTER.TYPE = req.TYPE

	fastbillbody, err := s.GenerateRequest(r)
	if err != nil {
		return nil, err
	}

	resp, err := s.FastbillRequest(fastbillbody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
