package gofastbill

import (
	"errors"
)

type ArticleDelete_Request struct {
	ARTICLE_ID string `xml:"ARTICLE_ID,omitempty" json:"ARTICLE_ID,omitempty"` //	Artikel ID
}

//DELETE Article
//FILTER => All fields from Struct: gofastbill.ArticleDelete_Request
//Required:
//ARTICLE_ID
//RESPONSE: STATUS
func (s *Initialization) Article_delete(req ArticleDelete_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "article.update"

	if req.ARTICLE_ID == "" {
		return nil, errors.New(s.Typ + ": ARTICLE_ID must not be empty")
	}

	r.DATA.ARTICLE_ID = req.ARTICLE_ID

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
