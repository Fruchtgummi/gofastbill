package gofastbill

type ArticleGet_Request struct {
	ARTICLE_NUMBER string `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"` //Artikelnummer
}

//GET Article
//FILTER => All fields from Struct: gofastbill.ArticleGet_Request
//Querying the details of one or more invoices. If no filter is set, 10 invoices will be returned.
func (s *Initialization) Article_get(req ArticleGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "article.get"
	r.FILTER.ARTICLE_NUMBER = req.ARTICLE_NUMBER

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
