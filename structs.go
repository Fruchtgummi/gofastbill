package gofastbill

type CUSTOMER struct {
	CUSTOMER_ID                    string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CUSTOMER_NUMBER                string `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`
	DAYS_FOR_PAYMENT               string `xml:"DAYS_FOR_PAYMENT,omitempty" json:"DAYS_FOR_PAYMENT,omitempty"`
	CREATED                        string `xml:"CREATED,omitempty" json:"CREATED,omitempty"`
	PAYMENT_TYPE                   string `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`
	BANK_NAME                      string `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`
	BANK_ACCOUNT_NUMBER            string `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`
	BANK_CODE                      string `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`
	BANK_ACCOUNT_OWNER             string `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`
	BANK_IBAN                      string `xml:"BANK_IBAN,omitempty" json:"BANK_IBAN,omitempty"`
	BANK_BIC                       string `xml:"BANK_BIC,omitempty" json:"BANK_BIC,omitempty"`
	BANK_ACCOUNT_MANDATE_REFERENCE string `xml:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty" json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty"`
	SHOW_PAYMENT_NOTICE            string `xml:"SHOW_PAYMENT_NOTICE,omitempty" json:"SHOW_PAYMENT_NOTICE,omitempty"`
	ACCOUNT_RECEIVABLE             string `xml:"ACCOUNT_RECEIVABLE,omitempty" json:"ACCOUNT_RECEIVABLE,omitempty"`
	CUSTOMER_TYPE                  string `xml:"CUSTOMER_TYPE,omitempty" json:"CUSTOMER_TYPE,omitempty"`
	TOP                            string `xml:"TOP,omitempty" json:"TOP,omitempty"`
	NEWSLETTER_OPTIN               string `xml:"NEWSLETTER_OPTIN,omitempty" json:"NEWSLETTER_OPTIN,omitempty"`
	ORGANIZATION                   string `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`
	POSITION                       string `xml:"POSITION,omitempty" json:"POSITION,omitempty"`
	ACADEMIC_DEGREE                string `xml:"ACADEMIC_DEGREE,omitempty" json:"ACADEMIC_DEGREE,omitempty"`
	SALUTATION                     string `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`
	FIRST_NAME                     string `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`
	LAST_NAME                      string `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`
	ADDRESS                        string `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`
	ADDRESS_2                      string `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`
	ZIPCODE                        string `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`
	CITY                           string `xml:"CITY,omitempty" json:"CITY,omitempty"`
	COUNTRY_CODE                   string `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`
	SECONDARY_ADDRESS              string `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"`
	PHONE                          string `xml:"PHONE,omitempty" json:"PHONE,omitempty"`
	PHONE_2                        string `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`
	FAX                            string `xml:"FAX,omitempty" json:"FAX,omitempty"`
	MOBILE                         string `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`
	EMAIL                          string `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`
	WEBSITE                        string `xml:"WEBSITE,omitempty" json:"WEBSITE,omitempty"`
	VAT_ID                         string `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`
	CURRENCY_CODE                  string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`
	LASTUPDATE                     string `xml:"LASTUPDATE,omitempty" json:"LASTUPDATE,omitempty"`
	TAGS                           string `xml:"TAGS,omitempty" json:"TAGS,omitempty"`
}

type CONTACT struct {
	CONTACT_ID        string `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`
	CUSTOMER_ID       string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	ORGANIZATION      string `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`
	POSITION          string `xml:"POSITION,omitempty" json:"POSITION,omitempty"`
	SALUTATION        string `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`
	FIRST_NAME        string `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`
	LAST_NAME         string `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`
	ADDRESS           string `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`
	ADDRESS_2         string `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`
	ZIPCODE           string `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`
	CITY              string `xml:"CITY,omitempty" json:"CITY,omitempty"`
	COUNTRY_CODE      string `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`
	SECONDARY_ADDRESS string `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"`
	PHONE             string `xml:"PHONE,omitempty" json:"PHONE,omitempty"`
	PHONE_2           string `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`
	FAX               string `xml:"FAX,omitempty" json:"FAX,omitempty"`
	MOBILE            string `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`
	EMAIL             string `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`
	WEBSITE           string `xml:"WEBSITE,omitempty" json:"WEBSITE,omitempty"`
	VAT_ID            string `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`
	CURRENCY_CODE     string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`
	CREATED           string `xml:"CREATED,omitempty" json:"CREATED,omitempty"`
	TAGS              string `xml:"TAGS,omitempty" json:"TAGS,omitempty"`
}

type DATA struct {
	CUSTOMER_ID                    string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CUSTOMER_NUMBER                string `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`                               //	Eigene Kundennummer
	CUSTOMER_TYPE                  string `xml:"CUSTOMER_TYPE,omitempty" json:"CUSTOMER_TYPE,omitempty"`                                   // 	Required	Kundentyp: business = Geschäftskunde | consumer = Privatperson
	ORGANIZATION                   string `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                                     // 	Required	Firmenname [REQUIRED] wenn CUSTOMER_TYPE = business
	POSITION                       string `xml:"POSITION,omitempty" json:"POSITION,omitempty"`                                             // 	Position in der Firma
	ACADEMIC_DEGREE                string `xml:"ACADEMIC_DEGREE,omitempty" json:"ACADEMIC_DEGREE,omitempty"`                               //	Akademischer Grad
	SALUTATION                     string `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                                         //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME                     string `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                                         //	Vorname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	LAST_NAME                      string `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                                           // 	Required	Nachname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	ADDRESS                        string `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                                               //	Adresszeile 1
	ADDRESS_2                      string `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                                           //	Adresszeile 2
	SECONDARY_ADDRESS              string `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"`                           //	Lieferadresse
	ZIPCODE                        string `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                                               //	Postleitzahl
	CITY                           string `xml:"CITY,omitempty" json:"CITY,omitempty"`                                                     //	Stadt
	STATE                          string `xml:"STATE,omitempty" json:"STATE,omitempty"`                                                   //	Bundestaat
	COUNTRY_CODE                   string `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                                     //	Länder-Code (ISO 3166 ALPHA-2)
	PHONE                          string `xml:"PHONE,omitempty" json:"PHONE,omitempty"`                                                   //	Telefon
	PHONE_2                        string `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`                                               //	Telefon 2
	FAX                            string `xml:"FAX,omitempty" json:"FAX,omitempty"`                                                       //	Faxnummer
	MOBILE                         string `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`                                                 //	Mobiltelefonnummer
	EMAIL                          string `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`                                                   //	E-Mailadresse
	WEBSITE                        string `xml:"WEBSITE,omitempty" json:"WEBSITE,omitempty"`                                               //	Webseite
	ACCOUNT_RECEIVABLE             string `xml:"ACCOUNT_RECEIVABLE,omitempty" json:"ACCOUNT_RECEIVABLE,omitempty"`                         //	Datev Debitoren-Kontonummer
	CURRENCY_CODE                  string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                                   //	Standardwährung
	VAT_ID                         string `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                                 //	USt-IdNr.
	DAYS_FOR_PAYMENT               string `xml:"DAYS_FOR_PAYMENT,omitempty" json:"DAYS_FOR_PAYMENT,omitempty"`                             //	Tage bis zum Zahlungsziel
	PAYMENT_TYPE                   string `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                                     //	1 = send | 2 = lastS | 3 = bar | 4 = paypal | 5 = vorK | 6 = kreditK
	SHOW_PAYMENT_NOTICE            string `xml:"SHOW_PAYMENT_NOTICE,omitempty" json:"SHOW_PAYMENT_NOTICE,omitempty"`                       //	Zahlungshinweis anzeigen
	BANK_NAME                      string `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                                           //	Bankname
	BANK_CODE                      string `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                                           //	Bankleitzahl
	BANK_ACCOUNT_NUMBER            string `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`                       //	Kontonummer
	BANK_ACCOUNT_OWNER             string `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`                         //	Kontoinhaber
	BANK_ACCOUNT_MANDATE_REFERENCE string `xml:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty" json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty"` //	Mandatsrefernznummer
	TAGS                           string `xml:"TAGS,omitempty" json:"TAGS,omitempty"`                                                     //	Tag halt
}
