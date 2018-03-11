package gofastbill

import (
	"errors"
)

type ArticleCreate_Request struct {
	ARTICLE_NUMBER string `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"` //	Artikelnummer
	TITLE          string `xml:"TITLE,omitempty" json:"TITLE,omitempty"`                   //	Titel
	DESCRIPTION    string `xml:"DESCRIPTION,omitempty" json:"DESCRIPTION,omitempty"`       //	Beschreibung
	UNIT           string `xml:"UNIT,omitempty" json:"UNIT,omitempty"`                     //	Produkt Einheitstyp
	UNIT_PRICE     string `xml:"UNIT_PRICE,omitempty" json:"UNIT_PRICE,omitempty"`         //	Einzelpreis
	CURRENCY_CODE  string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`   //	Standardwährung
	VAT_PERCENT    string `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`       //	MwSt in Prozent
}

//CREATE Article
//FILTER => All fields from Struct: gofastbill.ArticleCreate_Request
//Required:
//ARTICLE_NUMBER, TITLE, UNIT_PRICE
//RESPONSE: STATUS, ARTICLE_ID
func (s *Initialization) Article_create(req ArticleCreate_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "article.create"

	if req.ARTICLE_NUMBER == "" {
		return nil, errors.New(s.Typ + ": ARTICLE_NUMBER must not be empty")
	}

	if req.TITLE == "" {
		return nil, errors.New(s.Typ + ": TITLE must not be empty")
	}
	if req.UNIT_PRICE == "" {
		return nil, errors.New(s.Typ + ": UNIT_PRICE must not be empty")
	}

	r.DATA.ARTICLE_NUMBER = req.ARTICLE_NUMBER
	r.DATA.TITLE = req.TITLE
	r.DATA.DESCRIPTION = req.DESCRIPTION
	r.DATA.UNIT = req.UNIT
	r.DATA.UNIT_PRICE = req.UNIT_PRICE
	r.DATA.CURRENCY_CODE = req.CURRENCY_CODE
	r.DATA.VAT_PERCENT = req.VAT_PERCENT

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
