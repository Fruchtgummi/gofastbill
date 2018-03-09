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

type INVOICE struct {
	INVOICE_ID             string    `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`                         //Rechnungs-ID
	TYPE                   string    `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                                     //
	CUSTOMER_ID            string    `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //Eine bestimmte Kundennummer
	CUSTOMER_NUMBER        string    `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`               //Eigene Kundennummer
	CUSTOMER_COSTCENTER_ID string    `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //ID der Kostenstelle
	ORGANIZATION           string    `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                     //Firmenname [REQUIRED] wenn customer_type = business
	SALUTATION             string    `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                         //Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME             string    `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                         //Vorname
	LAST_NAME              string    `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                           //Nachname
	ADDRESS                string    `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                               //Adresszeile 1
	ADDRESS_2              string    `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                           //Adresszeile 2
	ZIPCODE                string    `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                               //Postleitzahl
	CITY                   string    `xml:"CITY,omitempty" json:"CITY,omitempty"`                                     //Stadt
	COMMENT                string    `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`                               //Bemerkung
	PAYMENT_TYPE           string    `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                     //1 = ueberweisung | 2 = lastschrift | 3 = bar | 4 = paypal | 5 = vorkasse | 6 = kreditkarte
	DAYS_FOR_PAYMENT       string    `xml:"DAYS_FOR_PAYMENT,omitempty" json:"DAYS_FOR_PAYMENT,omitempty"`             //Tage bis zum Zahlungsziel
	BANK_NAME              string    `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                           //Bankname
	BANK_ACCOUNT_NUMBER    string    `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`       //Kontonummer
	BANK_CODE              string    `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                           //Bankleitzahl
	BANK_ACCOUNT_OWNER     string    `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`         //Kontoinhaber
	BANK_IBAN              string    `xml:"BANK_IBAN,omitempty" json:"BANK_IBAN,omitempty"`                           //IBAN
	BANK_BIC               string    `xml:"BANK_BIC,omitempty" json:"BANK_BIC,omitempty"`                             //BIC
	AFFILIATE              string    `xml:"AFFILIATE,omitempty" json:"AFFILIATE,omitempty"`                           //Vertriebspartner
	COUNTRY_CODE           string    `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                     //Länder-Code (ISO 3166 ALPHA-2)
	VAT_ID                 string    `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                 //USt-IdNr.
	CURRENCY_CODE          string    `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //Standardwährung
	TEMPLATE_ID            string    `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //Entwurfsnummer
	CONTACT_ID             string    `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`                         //Kunden ID
	SUBSCRIPTION_ID        string    `xml:"SUBSCRIPTION_ID,omitempty" json:"SUBSCRIPTION_ID,omitempty"`               //Abonnement ID
	INVOICE_NUMBER         string    `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"`                 //Rechnungsnummer
	INVOICE_TITLE          string    `xml:"INVOICE_TITLE,omitempty" json:"INVOICE_TITLE,omitempty"`                   //Rechnungstitel
	INTROTEXT              string    `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //Einleitungstext
	PAID_DATE              string    `xml:"PAID_DATE,omitempty" json:"PAID_DATE,omitempty"`                           //Datum der Zahlung
	IS_CANCELED            string    `xml:"IS_CANCELED,omitempty" json:"IS_CANCELED,omitempty"`                       //Flag für Stonierungsstatus: 0 = nein | 1 = ja
	INVOICE_DATE           string    `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`                     //Rechnungsdatum
	DUE_DATE               string    `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`                             //Fälligkeitsdatum
	DELIVERY_DATE          string    `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                   //Lieferdatum
	CASH_DISCOUNT_PERCENT  string    `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`   //Skonto in Prozent
	CASH_DISCOUNT_DAYS     string    `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`         //Skonto-Zeitraum in Tagen
	SUB_TOTAL              string    `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`                           //Nettobetrag
	VAT_TOTAL              string    `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`                           //Vorsteuerbetrag
	VAT_ITEMS              VAT_ITEMS `xml:"VAT_ITEMS,omitempty" json:"VAT_ITEMS,omitempty"`                           //Liste aller Artikel zu einem Datensatz
	ITEMS                  []ITEMS   `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //Liste der Artikel
	TOTAL                  string    `xml:"TOTAL,omitempty" json:"TOTAL,omitempty"`                                   //Gesamtmenge
	PAYMENTS               PAYMENTS  `xml:"PAYMENTS,omitempty" json:"PAYMENTS,omitempty"`                             //Liste aller Zahlungen zu einem Datensatz
	PAYMENT_INFO           string    `xml:"PAYMENT_INFO,omitempty" json:"PAYMENT_INFO,omitempty"`                     //Informationen zu Zahlung
	DOCUMENT_URL           string    `xml:"DOCUMENT_URL,omitempty" json:"DOCUMENT_URL,omitempty"`                     //URL eines Dokuments

}

type VAT_ITEMS struct {
	VAT_PERCENT  string `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`   //MwSt in Prozent
	COMPLETE_NET string `xml:"COMPLETE_NET,omitempty" json:"COMPLETE_NET,omitempty"` //Gesamtnettobetrag
	VAT_VALUE    string `xml:"VAT_VALUE,omitempty" json:"VAT_VALUE,omitempty"`       //Mehrwertsteuerwert
}

type ITEMS struct {
	ARTICLE_NUMBER string `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"` //Artikelnummer
	DESCRIPTION    string `xml:"DESCRIPTION,omitempty" json:"DESCRIPTION,omitempty"`       //Beschreibung
	QUANTITY       string `xml:"QUANTITY,omitempty" json:"QUANTITY,omitempty"`             //Anzahl
	UNIT           string `xml:"UNIT,omitempty" json:"UNIT,omitempty"`                     //Produkt Einheitstype
	UNIT_PRICE     string `xml:"UNIT_PRICE,omitempty" json:"UNIT_PRICE,omitempty"`         //Einzelpreis
	VAT_PERCENT    string `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`       //MwSt in Prozent
	IS_GROSS       string `xml:"IS_GROSS,omitempty" json:"IS_GROSS,omitempty"`             //Flag ob auf Brutto umgestellt werden soll: 0 = no | 1= yes
	SORT_ORDER     string `xml:"SORT_ORDER,omitempty" json:"SORT_ORDER,omitempty"`         //Sortierungschlüssel der Rechnungsposition
}

type PAYMENTS struct {
	PAYMENT_ID    string `xml:"PAYMENT_ID,omitempty" json:"PAYMENT_ID,omitempty"`       //Zahlungs_id
	DATE          string `xml:"DATE,omitempty" json:"DATE,omitempty"`                   //Datum
	AMOUNT        string `xml:"AMOUNT,omitempty" json:"AMOUNT,omitempty"`               //Betrag
	CURRENCY_CODE string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"` //Standardwährung
	TYPE          string `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                   //Zahlungsart
	NOTE          string `xml:"NOTE,omitempty" json:"NOTE,omitempty"`                   //Notiz

}

//Datastruct include all tags
type DATA struct {
	CUSTOMER_ID                    string  `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CONTACT_ID                     string  `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`
	CUSTOMER_NUMBER                string  `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`                               //	Eigene Kundennummer
	CUSTOMER_TYPE                  string  `xml:"CUSTOMER_TYPE,omitempty" json:"CUSTOMER_TYPE,omitempty"`                                   // 	Required	Kundentyp: business = Geschäftskunde | consumer = Privatperson
	ORGANIZATION                   string  `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                                     // 	Required	Firmenname [REQUIRED] wenn CUSTOMER_TYPE = business
	POSITION                       string  `xml:"POSITION,omitempty" json:"POSITION,omitempty"`                                             // 	Position in der Firma
	ACADEMIC_DEGREE                string  `xml:"ACADEMIC_DEGREE,omitempty" json:"ACADEMIC_DEGREE,omitempty"`                               //	Akademischer Grad
	SALUTATION                     string  `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                                         //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME                     string  `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                                         //	Vorname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	LAST_NAME                      string  `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                                           // 	Required	Nachname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	ADDRESS                        string  `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                                               //	Adresszeile 1
	ADDRESS_2                      string  `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                                           //	Adresszeile 2
	SECONDARY_ADDRESS              string  `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"`                           //	Lieferadresse
	ZIPCODE                        string  `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                                               //	Postleitzahl
	CITY                           string  `xml:"CITY,omitempty" json:"CITY,omitempty"`                                                     //	Stadt
	STATE                          string  `xml:"STATE,omitempty" json:"STATE,omitempty"`                                                   //	Bundestaat
	COUNTRY_CODE                   string  `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                                     //	Länder-Code (ISO 3166 ALPHA-2)
	PHONE                          string  `xml:"PHONE,omitempty" json:"PHONE,omitempty"`                                                   //	Telefon
	PHONE_2                        string  `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`                                               //	Telefon 2
	FAX                            string  `xml:"FAX,omitempty" json:"FAX,omitempty"`                                                       //	Faxnummer
	MOBILE                         string  `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`                                                 //	Mobiltelefonnummer
	EMAIL                          string  `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`                                                   //	E-Mailadresse
	WEBSITE                        string  `xml:"WEBSITE,omitempty" json:"WEBSITE,omitempty"`                                               //	Webseite
	ACCOUNT_RECEIVABLE             string  `xml:"ACCOUNT_RECEIVABLE,omitempty" json:"ACCOUNT_RECEIVABLE,omitempty"`                         //	Datev Debitoren-Kontonummer
	CURRENCY_CODE                  string  `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                                   //	Standardwährung
	VAT_ID                         string  `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                                 //	USt-IdNr.
	DAYS_FOR_PAYMENT               string  `xml:"DAYS_FOR_PAYMENT,omitempty" json:"DAYS_FOR_PAYMENT,omitempty"`                             //	Tage bis zum Zahlungsziel
	PAYMENT_TYPE                   string  `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                                     //	1 = send | 2 = lastS | 3 = bar | 4 = paypal | 5 = vorK | 6 = kreditK
	SHOW_PAYMENT_NOTICE            string  `xml:"SHOW_PAYMENT_NOTICE,omitempty" json:"SHOW_PAYMENT_NOTICE,omitempty"`                       //	Zahlungshinweis anzeigen
	BANK_NAME                      string  `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                                           //	Bankname
	BANK_CODE                      string  `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                                           //	Bankleitzahl
	BANK_ACCOUNT_NUMBER            string  `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`                       //	Kontonummer
	BANK_ACCOUNT_OWNER             string  `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`                         //	Kontoinhaber
	BANK_ACCOUNT_MANDATE_REFERENCE string  `xml:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty" json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty"` //	Mandatsrefernznummer
	CREATED                        string  `xml:"CREATED,omitempty" json:"CREATED,omitempty"`
	TAGS                           string  `xml:"TAGS,omitempty" json:"TAGS,omitempty"`                                     //	Tag halt
	CUSTOMER_COSTCENTER_ID         string  `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //
	TEMPLATE_ID                    string  `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //
	INTROTEXT                      string  `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //
	INVOICE_TITLE                  string  `xml:"INVOICE_TITLE,omitempty" json:"INVOICE_TITLE,omitempty"`                   //
	INVOICE_DATE                   string  `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`                     //
	DELIVERY_DATE                  string  `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                   //
	CASH_DISCOUNT_PERCENT          string  `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`   //
	CASH_DISCOUNT_DAYS             string  `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`         //
	EU_DELIVERY                    string  `xml:"EU_DELIVERY,omitempty" json:"EU_DELIVERY,omitempty"`                       //
	IS_GROSS                       string  `xml:"IS_GROSS,omitempty" json:"IS_GROSS,omitempty"`                             //
	TEMPLATE_HASH                  string  `xml:"TEMPLATE_HASH,omitempty" json:"TEMPLATE_HASH,omitempty"`                   //
	ITEMS                          []ITEMS `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //

}
