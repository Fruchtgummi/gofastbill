package gofastbill

type DocumentGet_Request struct {
	FOLDER_ID string `xml:"FOLDER_ID,omitempty" json:"FOLDER_ID,omitempty"` //	Rechnungs-ID
}

//GET Documents
//FILTER => FOLDER_ID (Buggy: does not matter if you specify an FOLDER_ID, the RESPONSE always shows everyone)
//RESPONSE: you get a map[string]interface{}, because the RESPONSE use the same name of a Struct but there are different ones Methodes.
//Example:
//	var x gofastbill.DocumentGet_Request
//	fbapi, err := fastbill.Document_get(x)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(fbapi.RESPONSE.ITEMS.FOLDERS)
//	for k, v := range fbapi.RESPONSE.ITEMS.FOLDERS {
//		log.Println(k, v)
//	}
//All values for the map[] FOLDER_ID, NAME, PARENTFOLDER_ID, CREATED, CONTENT_COUNT; The key of the rows is a string and indiscriminately
func (s *Initialization) Document_get(req DocumentGet_Request) (*FBAPI, error) {

	var fastbillbody string

	var r FBAPI

	r.SERVICE = "document.get"

	r.FILTER.FOLDER_ID = req.FOLDER_ID

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
