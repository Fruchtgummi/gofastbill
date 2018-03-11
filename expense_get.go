package gofastbill

type ExpenseGet_Request struct {
	INVOICE_ID     string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`       //	Rechnungs-ID
	INVOICE_NUMBER string `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"` //	Rechnungsnummer
	MONTH          string `xml:"MONTH,omitempty" json:"MONTH,omitempty"`                   //	Monat
	YEAR           string `xml:"YEAR,omitempty" json:"YEAR,omitempty"`                     //	Jahr
}

//GET Expense
//FILTER => All fields from Struct: gofastbill.ExpenseGet_Request
//RESPONSE:All fields from Struct: gofastbill.EXPENSE
func (s *Initialization) Expense_get(req ExpenseGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "expense.get"

	r.FILTER.INVOICE_ID = req.INVOICE_ID
	r.FILTER.INVOICE_NUMBER = req.INVOICE_NUMBER
	r.FILTER.MONTH = req.MONTH
	r.FILTER.YEAR = req.YEAR

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
