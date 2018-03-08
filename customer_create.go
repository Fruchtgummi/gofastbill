package gofastbill

func (s *Initialization) Customer_Create() (*FBAPI, error) {

	xmlbody := `<?xml version="1.0" encoding="utf-8"?>
				<FBAPI>
				    <SERVICE>customer.create</SERVICE>
				    <DATA>
				        <CUSTOMER_TYPE>business</CUSTOMER_TYPE>
				        <ORGANIZATION>Musterfirma</ORGANIZATION>
				        <LAST_NAME>Mustermann</LAST_NAME>
				    </DATA>
				</FBAPI>`

	resp, err := s.FastbillRequest(xmlbody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
