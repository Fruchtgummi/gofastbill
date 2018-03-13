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
	STATUS                         string `xml:"STATUS,omitempty" json:"STATUS,omitempty"` //Status
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
	STATUS            string `xml:"STATUS,omitempty" json:"STATUS,omitempty"` //Status
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
	SUB_TOTAL              float64   `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`                           //Nettobetrag
	VAT_TOTAL              float64   `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`                           //Vorsteuerbetrag
	VAT_ITEMS              VAT_ITEMS `xml:"VAT_ITEMS,omitempty" json:"VAT_ITEMS,omitempty"`                           //Liste aller Artikel zu einem Datensatz
	ITEMS                  []ITEMS   `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //Liste der Artikel
	TOTAL                  float64   `xml:"TOTAL,omitempty" json:"TOTAL,omitempty"`                                   //Gesamtmenge
	PAYMENTS               PAYMENTS  `xml:"PAYMENTS,omitempty" json:"PAYMENTS,omitempty"`                             //Liste aller Zahlungen zu einem Datensatz
	PAYMENT_INFO           string    `xml:"PAYMENT_INFO,omitempty" json:"PAYMENT_INFO,omitempty"`                     //Informationen zu Zahlung
	DOCUMENT_URL           string    `xml:"DOCUMENT_URL,omitempty" json:"DOCUMENT_URL,omitempty"`                     //URL eines Dokuments
	STATUS                 string    `xml:"STATUS,omitempty" json:"STATUS,omitempty"`                                 //Status
}

type VAT_ITEMS struct {
	VAT_PERCENT  string  `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`   //MwSt in Prozent
	COMPLETE_NET float64 `xml:"COMPLETE_NET,omitempty" json:"COMPLETE_NET,omitempty"` //Gesamtnettobetrag
	VAT_VALUE    float64 `xml:"VAT_VALUE,omitempty" json:"VAT_VALUE,omitempty"`       //Mehrwertsteuerwert
}

type ITEMS struct {
	ARTICLE_NUMBER string  `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"` //Artikelnummer
	DESCRIPTION    string  `xml:"DESCRIPTION,omitempty" json:"DESCRIPTION,omitempty"`       //Beschreibung
	QUANTITY       string  `xml:"QUANTITY,omitempty" json:"QUANTITY,omitempty"`             //Anzahl
	UNIT           string  `xml:"UNIT,omitempty" json:"UNIT,omitempty"`                     //Produkt Einheitstype
	UNIT_PRICE     string  `xml:"UNIT_PRICE,omitempty" json:"UNIT_PRICE,omitempty"`         //Einzelpreis
	VAT_PERCENT    string  `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`       //MwSt in Prozent
	VAT_VALUE      float64 `xml:"VAT_VALUE,omitempty" json:"VAT_VALUE,omitempty"`           //
	IS_GROSS       string  `xml:"IS_GROSS,omitempty" json:"IS_GROSS,omitempty"`             //Flag ob auf Brutto umgestellt werden soll: 0 = no | 1= yes
	COMPLETE_GROSS float64 `xml:"COMPLETE_GROSS,omitempty" json:"COMPLETE_GROSS,omitempty"` //Brutto
	SORT_ORDER     float64 `xml:"SORT_ORDER,omitempty" json:"SORT_ORDER,omitempty"`         //Sortierungschlüssel der Rechnungsposition
	COMPLETE_NET   float64 `xml:"COMPLETE_NET,omitempty" json:"COMPLETE_NET,omitempty"`     //Sortierungschlüssel der Rechnungsposition
}

type ITEM struct {
	INVOICE_ITEM_ID string                 `xml:"INVOICE_ITEM_ID,omitempty" json:"INVOICE_ITEM_ID,omitempty"` //	ID der Rechnungsposition
	INVOICE_ID      string                 `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`           //	Rechnungs-ID
	CUSTOMER_ID     string                 `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`         //	Eine bestimmte Kundennummer
	ARTICLE_NUMBER  string                 `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"`   //	Artikelnummer
	DESCRIPTION     string                 `xml:"DESCRIPTION,omitempty" json:"DESCRIPTION,omitempty"`         //	Beschreibung
	QUANTITY        string                 `xml:"QUANTITY,omitempty" json:"QUANTITY,omitempty"`               //	Anzahl
	UNIT_PRICE      string                 `xml:"UNIT_PRICE,omitempty" json:"UNIT_PRICE,omitempty"`           //	Einzelpreis
	VAT_PERCENT     string                 `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`         //	MwSt in Prozent
	VAT_VALUE       float64                `xml:"VAT_VALUE,omitempty" json:"VAT_VALUE,omitempty"`             //	Mehrwertsteuerwert
	COMPLETE_NET    float64                `xml:"COMPLETE_NET,omitempty" json:"COMPLETE_NET,omitempty"`       //	Gesamtnettobetrag
	COMPLETE_GROSS  float64                `xml:"COMPLETE_GROSS,omitempty" json:"COMPLETE_GROSS,omitempty"`   //
	CURRENCY_CODE   string                 `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`     //	Standardwährung
	SORT_ORDER      float64                `xml:"SORT_ORDER,omitempty" json:"SORT_ORDER,omitempty"`           //	Sortierschlüssel der Rechnungsposition
	DOCUMENTS       []DOCUMENT             `xml:"DOCUMENTS,omitempty" json:"ITEMS,omitempty"`
	FOLDERS         map[string]interface{} `xml:"FOLDERS,omitempty" json:"FOLDERS,omitempty"`
}

type PAYMENTS struct {
	PAYMENT_ID    string `xml:"PAYMENT_ID,omitempty" json:"PAYMENT_ID,omitempty"`       //Zahlungs_id
	DATE          string `xml:"DATE,omitempty" json:"DATE,omitempty"`                   //Datum
	AMOUNT        string `xml:"AMOUNT,omitempty" json:"AMOUNT,omitempty"`               //Betrag
	CURRENCY_CODE string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"` //Standardwährung
	TYPE          string `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                   //Zahlungsart
	NOTE          string `xml:"NOTE,omitempty" json:"NOTE,omitempty"`                   //Notiz

}

type WEBHOOK struct {
	WEBHOOK_ID string `xml:"WEBHOOK_ID,omitempty" json:"WEBHOOK_ID,omitempty"` //Webhook ID
	ENDPOINT   string `xml:"ENDPOINT,omitempty" json:"ENDPOINT,omitempty"`     //Endpoint
	TYPE       string `xml:"TYPE,omitempty" json:"TYPE,omitempty"`             //Typ des Endpoints: "url" | "email". (Derzeit ist nur "url" verfügbar)
	EVENTS     string `xml:"EVENTS,omitempty" json:"EVENTS,omitempty"`         //Events: bei welchen Ereignissen wir er aktiviert. ("customer.created,customer.updated")
}

//Daten (Comma-separated values) der Events die dem Webhook-Endpoint zugeordnet werden können:
//customer.created
//customer.updated
//customer.deleted
//invoice.created
//invoice.completed
//invoice.canceled
//estimate.created
//estimate.updated
//contact.created
//contact.updated
//contact.deleted

type TEMPLATE struct {
	TEMPLATE_ID   int    `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`     //Tempalte ID
	TEMPLATE_NAME string `xml:"TEMPLATE_NAME,omitempty" json:"TEMPLATE_NAME,omitempty"` //Vorlagenname
	TEMPLATE_HASH string `xml:"TEMPLATE_HASH,omitempty" json:"TEMPLATE_HASH,omitempty"` //Eindeutige ID des Templates
}

type ARTICLE struct {
	ARTICLE_ID     string `xml:"ARTICLE_ID,omitempty" json:"ARTICLE_ID,omitempty"`         //	Artikel ID
	ARTICLE_NUMBER string `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"` //	Artikelnummer
	TITLE          string `xml:"TITLE,omitempty" json:"TITLE,omitempty"`                   //	Titel
	DESCRIPTION    string `xml:"DESCRIPTION,omitempty" json:"DESCRIPTION,omitempty"`       //	Beschreibung
	UNIT           string `xml:"UNIT,omitempty" json:"UNIT,omitempty"`                     //	Produkt Einheitstyp
	UNIT_PRICE     string `xml:"UNIT_PRICE,omitempty" json:"UNIT_PRICE,omitempty"`         //	Einzelpreis
	CURRENCY_CODE  string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`   //	Standardwährung
	VAT_PERCENT    string `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`       //	MwSt in Prozent
	TAGS           string `xml:"TAGS,omitempty" json:"TAGS,omitempty"`                     //
}

type FOLDER struct {
	FOLDER_ID       string `xml:"FOLDER_ID,omitempty" json:"FOLDER_ID,omitempty"`             // ID des Ordners
	NAME            string `xml:"NAME,omitempty" json:"NAME,omitempty"`                       // Name des Ordners
	PARENTFOLDER_ID string `xml:"PARENTFOLDER_ID,omitempty" json:"PARENTFOLDER_ID,omitempty"` // ID des Parent-Odners
	CREATED         string `xml:"CREATED,omitempty" json:"CREATED,omitempty"`                 // Datum der Erstellung
	CONTENT_COUNT   string `xml:"CONTENT_COUNT,omitempty" json:"CONTENT_COUNT,omitempty"`     // Menge der enthaltenden Dateien
}

type DOCUMENT struct {
	DOCUMENT_ID string `xml:"DOCUMENT_ID,omitempty" json:"DOCUMENT_ID,omitempty"` // ID des Dokuments
	TYPE        string `xml:"TYPE,omitempty" json:"TYPE,omitempty"`               // Dokumententype
	TITLE       string `xml:"TITLE,omitempty" json:"TITLE,omitempty"`             // Titel
	DATE        string `xml:"DATE,omitempty" json:"DATE,omitempty"`               // Datum
	NOTE        string `xml:"NOTE,omitempty" json:"NOTE,omitempty"`               // Bemerkung
}

type ESTIMATE struct {
	ESTIMATE_ID            string    `xml:"ESTIMATE_ID,omitempty" json:"ESTIMATE_ID,omitempty"`                       //	Angebots-ID
	STATE                  string    `xml:"STATE,omitempty" json:"STATE,omitempty"`                                   //	Status: a = ueberweisung | b = lastschrift | c = bar | d = paypal | e = vorkasse
	CUSTOMER_ID            string    `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //	Eine bestimmte Kundennummer
	CUSTOMER_NUMBER        string    `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`               //	Eigene Kundennummer
	CUSTOMER_COSTCENTER_ID string    `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //	ID der Kostenstelle
	PROJECT_ID             string    `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`                         //	Eine bestimmte Projekt ID
	ORGANIZATION           string    `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                     //	Firmenname [REQUIRED] wenn customer_type = business
	SALUTATION             string    `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                         //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME             string    `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                         //	Vorname
	LAST_NAME              string    `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                           //	Nachname
	ADDRESS                string    `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                               //	Adresszeile 1
	ADDRESS_2              string    `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                           //	Adresszeile 2
	ZIPCODE                string    `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                               //	Postleitzahl
	CITY                   string    `xml:"CITY,omitempty" json:"CITY,omitempty"`                                     //	Stadt
	PAYMENT_TYPE           string    `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                     //	1 = ueberweisung | 2 = lastschrift | 3 = bar | 4 = paypal | 5 = vorkasse | 6 = kreditkarte
	BANK_NAME              string    `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                           //	Bankname
	BANK_ACCOUNT_NUMBER    string    `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`       //	Kontonummer
	BANK_CODE              string    `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                           //	Bankleitzahl
	BANK_ACCOUNT_OWNER     string    `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`         //	Kontoinhaber
	BANK_IBAN              string    `xml:"BANK_IBAN,omitempty" json:"BANK_IBAN,omitempty"`                           //	IBAN
	BANK_BIC               string    `xml:"BANK_BIC,omitempty" json:"BANK_BIC,omitempty"`                             //	BIC
	COUNTRY_CODE           string    `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                     //	Länder-Code (ISO 3166 ALPHA-2)
	VAT_ID                 string    `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                 //	USt-IdNr.
	CURRENCY_CODE          string    `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //	Standardwährung
	TEMPLATE_ID            string    `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //	Vorlage-ID
	ESTIMATE_NUMBER        string    `xml:"ESTIMATE_NUMBER,omitempty" json:"ESTIMATE_NUMBER,omitempty"`               //	Angebotsnummer
	INTROTEXT              string    `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //	Einleitungstext
	ESTIMATE_DATE          string    `xml:"ESTIMATE_DATE,omitempty" json:"ESTIMATE_DATE,omitempty"`                   //	Angebote Datum
	DUE_DATE               string    `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`                             //	Fälligkeitsdatum
	SUB_TOTAL              float64   `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`                           //	Nettobetrag
	VAT_TOTAL              float64   `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`                           //	Vorsteuerbetrag
	VAT_ITEMS              VAT_ITEMS `xml:"VAT_ITEMS,omitempty" json:"VAT_ITEMS,omitempty"`                           //	Liste aller Artikel zu einem Datensatz
	ITEMS                  string    `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //	Liste der Artikel
	TOTAL                  string    `xml:"TOTAL,omitempty" json:"TOTAL,omitempty"`                                   //	Gesamtmenge
	DOCUMENT_URL           string    `xml:"DOCUMENT_URL,omitempty" json:"DOCUMENT_URL,omitempty"`                     //	URL eines Dokuments
	START_ESTIMATE_DATE    string    `xml:"START_ESTIMATE_DATE,omitempty" json:"START_ESTIMATE_DATE,omitempty"`       //	Angebote ab einem bestimmten Datum
	END_ESTIMATE_DATE      string    `xml:"END_ESTIMATE_DATE,omitempty" json:"END_ESTIMATE_DATE,omitempty"`           //	Angebote bis zu einem bestimmten Datum
}

type EXPENSE struct {
	INVOICE_ID     string    `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`         //	Rechnungs-ID
	ORGANIZATION   string    `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`     //	Firmenname
	INVOICE_NUMBER string    `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"` //	Rechnungsnummer
	INVOICE_DATE   string    `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`     //	Rechnungsdatum
	DUE_DATE       string    `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`             //	Fälligkeitsdatum
	PROJECT_ID     string    `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`         //	Eine bestimmte Projekt ID
	CUSTOMER_ID    string    `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`       //	Eine bestimmte Kundennummer
	SUB_TOTAL      float64   `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`           //	Nettobetrag
	VAT_TOTAL      float64   `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`           //	Vorsteuerbetrag
	TOTAL          float64   `xml:"TOTAL,omitempty" json:"TOTAL,omitempty"`                   //	Gesamtmenge
	PAID_DATE      string    `xml:"PAID_DATE,omitempty" json:"PAID_DATE,omitempty"`           //	Datum der Zahlung
	CURRENCY_CODE  string    `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`   //	Standardwährung
	COMMENT        string    `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`               //	Bemerkung
	VAT_ITEMS      VAT_ITEMS `xml:"VAT_ITEMS,omitempty" json:"VAT_ITEMS,omitempty"`           //	Liste aller Artikel zu einem Datensatz
	PAYMENT_INFO   string    `xml:"PAYMENT_INFO,omitempty" json:"PAYMENT_INFO,omitempty"`     //	Informationen zu Zahlung
}

type TIME struct {
	TIME_ID          string `xml:"TIME_ID,omitempty" json:"TIME_ID,omitempty"`                   //ID eines bestimmten Zeiteintrags
	CUSTOMER_ID      string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`           //Eine bestimmte Kundennummer
	PROJECT_ID       string `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`             //Eine bestimmte Projekt ID
	INVOICE_ID       string `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`             //Rechnungs-ID
	DATE             string `xml:"DATE,omitempty" json:"DATE,omitempty"`                         //Datum
	START_TIME       string `xml:"START_TIME,omitempty" json:"START_TIME,omitempty"`             //Startzeit
	END_TIME         string `xml:"END_TIME,omitempty" json:"END_TIME,omitempty"`                 //Endzeit
	MINUTES          string `xml:"MINUTES,omitempty" json:"MINUTES,omitempty"`                   //Minuten
	BILLABLE_MINUTES string `xml:"BILLABLE_MINUTES,omitempty" json:"BILLABLE_MINUTES,omitempty"` //Abrechenbare Minuten
	COMMENT          string `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`                   //Bemerkung
}

type RECURRING struct {
	INVOICE_ID             string    `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`                         //	Rechnungs-ID
	TYPE                   string    `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                                     //	Rechnungen eines bestimmten Typs: outgoing = Rechnungen draft = Entwürfe | credit = Gutschriften
	CUSTOMER_ID            string    `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //	Eine bestimmte Kundennummer
	CUSTOMER_COSTCENTER_ID string    `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //	ID der Kostenstelle
	CURRENCY_CODE          string    `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //	Standardwährung
	TEMPLATE_ID            string    `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //	Entwurfsnummer
	INTROTEXT              string    `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //	Einleitungstext
	INVOICE_NUMBER         string    `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"`                 //	Rechnungsnummer
	PAID_DATE              string    `xml:"PAID_DATE,omitempty" json:"PAID_DATE,omitempty"`                           //	Datum der Zahlung
	IS_CANCELED            string    `xml:"IS_CANCELED,omitempty" json:"IS_CANCELED,omitempty"`                       //	Flag für Stonierungsstatus: 0 = nein | 1 = ja
	INVOICE_DATE           string    `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`                     //	Rechnungsdatum
	DUE_DATE               string    `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`                             //	Fälligkeitsdatum
	DELIVERY_DATE          string    `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                   //	Lieferdatum
	CASH_DISCOUNT_PERCENT  string    `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`   //	Skonto in Prozent
	CASH_DISCOUNT_DAYS     string    `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`         //	Skonto-Zeitraum in Tagen
	SUB_TOTAL              float64   `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`                           //	Nettobetrag
	VAT_TOTAL              float64   `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`                           //	Vorsteuerbetrag
	TOTAL                  float64   `xml:"TOTAL,omitempty" json:"TOTAL,omitempty"`                                   //	Gesamtmenge
	VAT_ITEMS              VAT_ITEMS `xml:"VAT_ITEMS,omitempty" json:"VAT_ITEMS,omitempty"`                           //	Liste aller Artikel zu einem Datensatz
	ITEMS                  []ITEMS   `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //	Liste der Artikel
}

type REVENUE struct {
	INVOICE_ID             string    `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`                         //	Rechnungs-ID
	TYPE                   string    `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                                     //	Rechnungen eines bestimmten Typs: outgoing = Rechnungen draft = Entwürfe | credit = Gutschriften
	CUSTOMER_ID            string    `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //	Eine bestimmte Kundennummer
	CUSTOMER_NUMBER        string    `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`               //	Eigene Kundennummer
	CUSTOMER_COSTCENTER_ID string    `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //	ID der Kostenstelle
	PROJECT_ID             string    `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`                         //	Eine bestimmte Projekt ID
	CURRENCY_CODE          string    `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //	Standardwährung
	DELIVERY_DATE          string    `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                   //	Lieferdatum
	INVOICE_TITLE          string    `xml:"INVOICE_TITLE,omitempty" json:"INVOICE_TITLE,omitempty"`                   //	Rechnungstitel
	CASH_DISCOUNT_PERCENT  string    `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`   //	Skonto in Prozent
	CASH_DISCOUNT_DAYS     string    `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`         //	Skonto-Zeitraum in Tagen
	SUB_TOTAL              float64   `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`                           //	Nettobetrag
	VAT_TOTAL              float64   `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`                           //	Vorsteuerbetrag
	VAT_ITEMS              VAT_ITEMS `xml:"VAT_ITEMS,omitempty" json:"VAT_ITEMS,omitempty"`                           //	Liste aller Artikel zu einem Datensatz
	ITEMS                  []ITEMS   `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                   //	Liste der Artikel
	TOTAL                  float64   `xml:"TOTAL,omitempty" json:"TOTAL,omitempty"`                                   //	Gesamtmenge
	ORGANIZATION           string    `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                     //	Firmenname [REQUIRED] wenn customer_type = business
	SALUTATION             string    `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                         //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME             string    `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                         //	Vorname
	LAST_NAME              string    `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                           //	Nachname
	ADDRESS                string    `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                               //	Adresszeile 1
	ADDRESS_2              string    `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                           //	Adresszeile 2
	ZIPCODE                string    `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                               //	Postleitzahl
	CITY                   string    `xml:"CITY,omitempty" json:"CITY,omitempty"`                                     //	Stadt
	PAYMENT_TYPE           string    `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                     //	1 = ueberweisung | 2 = lastschrift | 3 = bar | 4 = paypal | 5 = vorkasse | 6 = kreditkarte
	BANK_NAME              string    `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                           //	Bankname
	BANK_ACCOUNT_NUMBER    string    `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`       //	Kontonummer
	BANK_CODE              string    `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                           //	Bankleitzahl
	BANK_ACCOUNT_OWNER     string    `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`         //	Kontoinhaber
	BANK_IBAN              string    `xml:"BANK_IBAN,omitempty" json:"BANK_IBAN,omitempty"`                           //	IBAN
	BANK_BIC               string    `xml:"BANK_BIC,omitempty" json:"BANK_BIC,omitempty"`                             //	BIC
	COMMENTS               []COMMENT `xml:"COMMENTS,omitempty" json:"COMMENTS,omitempty"`                             //	Liste aller Kommentare zu einem Datensatz
	COUNTRY_CODE           string    `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                     //	Länder-Code (ISO 3166 ALPHA-2)
	VAT_ID                 string    `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                 //	USt-IdNr.
	TEMPLATE_ID            string    `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                       //	Entwurfsnummer
	INVOICE_NUMBER         string    `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"`                 //	Rechnungsnummer
	INTROTEXT              string    `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                           //	Einleitungstext
	PAID_DATE              string    `xml:"PAID_DATE,omitempty" json:"PAID_DATE,omitempty"`                           //	Datum der Zahlung
	IS_CANCELED            string    `xml:"IS_CANCELED,omitempty" json:"IS_CANCELED,omitempty"`                       //	Flag für Stonierungsstatus: 0 = nein | 1 = ja
	INVOICE_DATE           string    `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`                     //	Rechnungsdatum
	DUE_DATE               string    `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`                             //	Fälligkeitsdatum
	PAYMENT_INFO           string    `xml:"PAYMENT_INFO,omitempty" json:"PAYMENT_INFO,omitempty"`                     //	Informationen zu Zahlung
	LASTUPDATE             string    `xml:"LASTUPDATE,omitempty" json:"LASTUPDATE,omitempty"`                         //	Datum der letzten Bearbeitung
	DOCUMENT_URL           string    `xml:"DOCUMENT_URL,omitempty" json:"DOCUMENT_URL,omitempty"`                     //	URL eines Dokuments
}

type PROJECT struct {
	PROJECT_ID             string `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`                         //	Eine bestimmte Projekt ID
	PROJECT_NAME           string `xml:"PROJECT_NAME,omitempty" json:"PROJECT_NAME,omitempty"`                     //	Projektname
	CUSTOMER_ID            string `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`                       //	Eine bestimmte Kundennummer
	CUSTOMER_COSTCENTER_ID string `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"` //	ID der Kostenstelle
	HOUR_PRICE             string `xml:"HOUR_PRICE,omitempty" json:"HOUR_PRICE,omitempty"`                         //	Stundensatz
	CURRENCY_CODE          string `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                   //	Standardwährung
	VAT_PERCENT            string `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`                       //	MwSt in Prozent
	START_DATE             string `xml:"START_DATE,omitempty" json:"START_DATE,omitempty"`                         //	Datum des ersten Rechnungslaufs
	END_DATE               string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"`                             //	Enddatum
	TASKS                  []TASK `xml:"DOCUMENT_URL,omitempty" json:"DOCUMENT_URL,omitempty"`                     //	Liste der Aufgaben
}

type TASK struct {
	TASK_ID       string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` //ID einer bestimmten Aufagbe
	TASK_NUMBER   string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` //
	TASK_NAME     string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` //
	DESCRIPTION   string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` // Beschreibung
	STATUS        string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` // Ergebnis einer Aktion
	PRIORITY      string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` //
	HOUR_PRICE    string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` // Stundensatz
	CURRENCY_CODE string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` // Standardwährung
	VAR_PERCENT   string `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"` // MwSt in %
}

type COMMENT struct {
	COMMENT string
}

//Recipient

type RECIPIENT struct {
	TO  string `xml:"TO,omitempty" json:"TO,omitempty"`   //
	CC  string `xml:"CC,omitempty" json:"CC,omitempty"`   //
	BCC string `xml:"BCC,omitempty" json:"BCC,omitempty"` //
}

//Datastruct include all tags
type DATA struct {
	CUSTOMER_ID                    string    `xml:"CUSTOMER_ID,omitempty" json:"CUSTOMER_ID,omitempty"`
	CONTACT_ID                     string    `xml:"CONTACT_ID,omitempty" json:"CONTACT_ID,omitempty"`
	WEBHOOK_ID                     string    `xml:"WEBHOOK_ID,omitempty" json:"WEBHOOK_ID,omitempty"`
	CUSTOMER_NUMBER                string    `xml:"CUSTOMER_NUMBER,omitempty" json:"CUSTOMER_NUMBER,omitempty"`                               //	Eigene Kundennummer
	CUSTOMER_TYPE                  string    `xml:"CUSTOMER_TYPE,omitempty" json:"CUSTOMER_TYPE,omitempty"`                                   // 	Required	Kundentyp: business = Geschäftskunde | consumer = Privatperson
	ORGANIZATION                   string    `xml:"ORGANIZATION,omitempty" json:"ORGANIZATION,omitempty"`                                     // 	Required	Firmenname [REQUIRED] wenn CUSTOMER_TYPE = business
	POSITION                       string    `xml:"POSITION,omitempty" json:"POSITION,omitempty"`                                             // 	Position in der Firma
	ACADEMIC_DEGREE                string    `xml:"ACADEMIC_DEGREE,omitempty" json:"ACADEMIC_DEGREE,omitempty"`                               //	Akademischer Grad
	SALUTATION                     string    `xml:"SALUTATION,omitempty" json:"SALUTATION,omitempty"`                                         //	Anrede: mr = Herr | mrs = Frau | family = Familie | "empty" = ohne Titel
	FIRST_NAME                     string    `xml:"FIRST_NAME,omitempty" json:"FIRST_NAME,omitempty"`                                         //	Vorname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	LAST_NAME                      string    `xml:"LAST_NAME,omitempty" json:"LAST_NAME,omitempty"`                                           // 	Required	Nachname [REQUIRED] wenn CUSTOMER_TYPE = consumer
	ADDRESS                        string    `xml:"ADDRESS,omitempty" json:"ADDRESS,omitempty"`                                               //	Adresszeile 1
	ADDRESS_2                      string    `xml:"ADDRESS_2,omitempty" json:"ADDRESS_2,omitempty"`                                           //	Adresszeile 2
	SECONDARY_ADDRESS              string    `xml:"SECONDARY_ADDRESS,omitempty" json:"SECONDARY_ADDRESS,omitempty"`                           //	Lieferadresse
	ZIPCODE                        string    `xml:"ZIPCODE,omitempty" json:"ZIPCODE,omitempty"`                                               //	Postleitzahl
	CITY                           string    `xml:"CITY,omitempty" json:"CITY,omitempty"`                                                     //	Stadt
	STATE                          string    `xml:"STATE,omitempty" json:"STATE,omitempty"`                                                   //	Bundestaat
	COUNTRY_CODE                   string    `xml:"COUNTRY_CODE,omitempty" json:"COUNTRY_CODE,omitempty"`                                     //	Länder-Code (ISO 3166 ALPHA-2)
	PHONE                          string    `xml:"PHONE,omitempty" json:"PHONE,omitempty"`                                                   //	Telefon
	PHONE_2                        string    `xml:"PHONE_2,omitempty" json:"PHONE_2,omitempty"`                                               //	Telefon 2
	FAX                            string    `xml:"FAX,omitempty" json:"FAX,omitempty"`                                                       //	Faxnummer
	MOBILE                         string    `xml:"MOBILE,omitempty" json:"MOBILE,omitempty"`                                                 //	Mobiltelefonnummer
	EMAIL                          string    `xml:"EMAIL,omitempty" json:"EMAIL,omitempty"`                                                   //	E-Mailadresse
	WEBSITE                        string    `xml:"WEBSITE,omitempty" json:"WEBSITE,omitempty"`                                               //	Webseite
	ACCOUNT_RECEIVABLE             string    `xml:"ACCOUNT_RECEIVABLE,omitempty" json:"ACCOUNT_RECEIVABLE,omitempty"`                         //	Datev Debitoren-Kontonummer
	CURRENCY_CODE                  string    `xml:"CURRENCY_CODE,omitempty" json:"CURRENCY_CODE,omitempty"`                                   //	Standardwährung
	VAT_ID                         string    `xml:"VAT_ID,omitempty" json:"VAT_ID,omitempty"`                                                 //	USt-IdNr.
	DAYS_FOR_PAYMENT               string    `xml:"DAYS_FOR_PAYMENT,omitempty" json:"DAYS_FOR_PAYMENT,omitempty"`                             //	Tage bis zum Zahlungsziel
	PAYMENT_TYPE                   string    `xml:"PAYMENT_TYPE,omitempty" json:"PAYMENT_TYPE,omitempty"`                                     //	1 = send | 2 = lastS | 3 = bar | 4 = paypal | 5 = vorK | 6 = kreditK
	SHOW_PAYMENT_NOTICE            string    `xml:"SHOW_PAYMENT_NOTICE,omitempty" json:"SHOW_PAYMENT_NOTICE,omitempty"`                       //	Zahlungshinweis anzeigen
	BANK_NAME                      string    `xml:"BANK_NAME,omitempty" json:"BANK_NAME,omitempty"`                                           //	Bankname
	BANK_CODE                      string    `xml:"BANK_CODE,omitempty" json:"BANK_CODE,omitempty"`                                           //	Bankleitzahl
	BANK_ACCOUNT_NUMBER            string    `xml:"BANK_ACCOUNT_NUMBER,omitempty" json:"BANK_ACCOUNT_NUMBER,omitempty"`                       //	Kontonummer
	BANK_ACCOUNT_OWNER             string    `xml:"BANK_ACCOUNT_OWNER,omitempty" json:"BANK_ACCOUNT_OWNER,omitempty"`                         //	Kontoinhaber
	BANK_ACCOUNT_MANDATE_REFERENCE string    `xml:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty" json:"BANK_ACCOUNT_MANDATE_REFERENCE,omitempty"` //	Mandatsrefernznummer
	CREATED                        string    `xml:"CREATED,omitempty" json:"CREATED,omitempty"`                                               //
	TAGS                           string    `xml:"TAGS,omitempty" json:"TAGS,omitempty"`                                                     //	Tag halt
	CUSTOMER_COSTCENTER_ID         string    `xml:"CUSTOMER_COSTCENTER_ID,omitempty" json:"CUSTOMER_COSTCENTER_ID,omitempty"`                 //
	TEMPLATE_ID                    string    `xml:"TEMPLATE_ID,omitempty" json:"TEMPLATE_ID,omitempty"`                                       //
	INTROTEXT                      string    `xml:"INTROTEXT,omitempty" json:"INTROTEXT,omitempty"`                                           //
	INVOICE_TITLE                  string    `xml:"INVOICE_TITLE,omitempty" json:"INVOICE_TITLE,omitempty"`                                   //
	INVOICE_DATE                   string    `xml:"INVOICE_DATE,omitempty" json:"INVOICE_DATE,omitempty"`                                     //
	DELIVERY_DATE                  string    `xml:"DELIVERY_DATE,omitempty" json:"DELIVERY_DATE,omitempty"`                                   //
	CASH_DISCOUNT_PERCENT          string    `xml:"CASH_DISCOUNT_PERCENT,omitempty" json:"CASH_DISCOUNT_PERCENT,omitempty"`                   //
	CASH_DISCOUNT_DAYS             string    `xml:"CASH_DISCOUNT_DAYS,omitempty" json:"CASH_DISCOUNT_DAYS,omitempty"`                         //
	EU_DELIVERY                    string    `xml:"EU_DELIVERY,omitempty" json:"EU_DELIVERY,omitempty"`                                       //
	IS_GROSS                       string    `xml:"IS_GROSS,omitempty" json:"IS_GROSS,omitempty"`                                             //
	TEMPLATE_HASH                  string    `xml:"TEMPLATE_HASH,omitempty" json:"TEMPLATE_HASH,omitempty"`                                   //
	ITEMS                          []ITEMS   `xml:"ITEMS,omitempty" json:"ITEMS,omitempty"`                                                   //
	TYPE                           string    `xml:"TYPE,omitempty" json:"TYPE,omitempty"`                                                     //
	ENDPOINT                       string    `xml:"ENDPOINT,omitempty" json:"ENDPOINT,omitempty"`                                             //
	EVENTS                         string    `xml:"EVENTS,omitempty" json:"EVENTS,omitempty"`                                                 //
	INVOICE_ID                     string    `xml:"INVOICE_ID,omitempty" json:"INVOICE_ID,omitempty"`                                         //Required	Rechnungs-ID
	DELETE_EXISTING_ITEMS          string    `xml:"DELETE_EXISTING_ITEMS,omitempty" json:"DELETE_EXISTING_ITEMS,omitempty"`                   //Flag zur Löschung aller bestehenden Rechnungsposten: 0 = nein | 1 = ja
	RECIPIENT                      RECIPIENT `xml:"RECIPIENT,omitempty" json:"RECIPIENT,omitempty"`                                           //Required	Empfänger
	SUBJECT                        string    `xml:"SUBJECT,omitempty" json:"SUBJECT,omitempty"`                                               //Betreff
	MESSAGE                        string    `xml:"MESSAGE,omitempty" json:"MESSAGE,omitempty"`                                               //e-Mail Text
	RECEIPT_CONFIRMATION           string    `xml:"RECEIPT_CONFIRMATION,omitempty" json:"RECEIPT_CONFIRMATION,omitempty"`                     //
	PAID_DATE                      string    `xml:"PAID_DATE,omitempty" json:"PAID_DATE,omitempty"`                                           //
	ARTICLE_ID                     string    `xml:"ARTICLE_ID,omitempty" json:"ARTICLE_ID,omitempty"`                                         //	Artikel ID
	ARTICLE_NUMBER                 string    `xml:"ARTICLE_NUMBER,omitempty" json:"ARTICLE_NUMBER,omitempty"`                                 //	Artikelnummer
	TITLE                          string    `xml:"TITLE,omitempty" json:"TITLE,omitempty"`                                                   //	Titel
	DESCRIPTION                    string    `xml:"DESCRIPTION,omitempty" json:"DESCRIPTION,omitempty"`                                       //	Beschreibung
	UNIT                           string    `xml:"UNIT,omitempty" json:"UNIT,omitempty"`                                                     //	Produkt Einheitstyp
	UNIT_PRICE                     string    `xml:"UNIT_PRICE,omitempty" json:"UNIT_PRICE,omitempty"`                                         //	Einzelpreis
	VAT_PERCENT                    string    `xml:"VAT_PERCENT,omitempty" json:"VAT_PERCENT,omitempty"`                                       //	MwSt in Prozent
	TEMPLATE_NAME                  string    `xml:"TEMPLATE_NAME,omitempty" json:"TEMPLATE_NAME,omitempty"`                                   //Vorlagenname
	ESTIMATE_ID                    string    `xml:"ESTIMATE_ID,omitempty" json:"ESTIMATE_ID,omitempty"`                                       //	Angebots-ID
	ESTIMATE_NUMBER                string    `xml:"ESTIMATE_NUMBER,omitempty" json:"ESTIMATE_NUMBER,omitempty"`                               //	Angebotsnummer
	START_ESTIMATE_DATE            string    `xml:"START_ESTIMATE_DATE,omitempty" json:"START_ESTIMATE_DATE,omitempty"`                       //	Angebote ab einem bestimmten Datum
	END_ESTIMATE_DATE              string    `xml:"END_ESTIMATE_DATE,omitempty" json:"END_ESTIMATE_DATE,omitempty"`                           //	Angebote bis zu einem bestimmten Datum
	INVOICE_ITEM_ID                string    `xml:"INVOICE_ITEM_ID,omitempty" json:"INVOICE_ITEM_ID,omitempty"`                               //
	INVOICE_NUMBER                 string    `xml:"INVOICE_NUMBER,omitempty" json:"INVOICE_NUMBER,omitempty"`                                 //	Rechnungsnumm
	MONTH                          string    `xml:"MONTH,omitempty" json:"MONTH,omitempty"`                                                   //	Monat
	YEAR                           string    `xml:"YEAR,omitempty" json:"YEAR,omitempty"`                                                     //
	DUE_DATE                       string    `xml:"DUE_DATE,omitempty" json:"DUE_DATE,omitempty"`                                             //	Fälligkeitsdatum
	PROJECT_ID                     string    `xml:"PROJECT_ID,omitempty" json:"PROJECT_ID,omitempty"`                                         //	Eine bestimmte Projekt ID
	COMMENT                        string    `xml:"COMMENT,omitempty" json:"COMMENT,omitempty"`                                               //	Bemerkung
	SUB_TOTAL                      float64   `xml:"SUB_TOTAL,omitempty" json:"SUB_TOTAL,omitempty"`                                           // Required	Nettobetrag
	VAT_TOTAL                      float64   `xml:"VAT_TOTAL,omitempty" json:"VAT_TOTAL,omitempty"`                                           //	Vorsteuerbetrag
	DATE                           string    `xml:"DATE,omitempty" json:"DATE,omitempty"`                                                     //	Datum
	TASK_ID                        string    `xml:"TASK_ID,omitempty" json:"TASK_ID,omitempty"`                                               //	ID einer bestimmten Aufgabe
	START_TIME                     string    `xml:"START_TIME,omitempty" json:"START_TIME,omitempty"`                                         // Required	Startzeit
	END_TIME                       string    `xml:"END_TIME,omitempty" json:"END_TIME,omitempty"`                                             //	Endzeit
	MINUTES                        string    `xml:"MINUTES,omitempty" json:"MINUTES,omitempty"`                                               //	Minuten
	BILLABLE_MINUTES               string    `xml:"BILLABLE_MINUTES,omitempty" json:"BILLABLE_MINUTES,omitempty"`                             //	Abrechenbare Minuten
	TIME_ID                        string    `xml:"TIME_ID,omitempty" json:"TIME_ID,omitempty"`                                               //
	START_DATE                     string    `xml:"START_DATE,omitempty" json:"START_DATE,omitempty"`                                         // Required	Datum des ersten Rechnungslaufs
	FREQUENCY                      string    `xml:"FREQUENCY,omitempty" json:"FREQUENCY,omitempty"`                                           // Required	Wiederholrate des Rechnungslauf: Weekly = wöchentlich | 2 weeks = alle 2 Wochen | 4 weeks = alle 4 Wochen | monthly = monatlich | 2 months = alle 2 Monate | 3 months = vierteljährlich | 6 months = alle 6 Monate| yearly = jährlich| 2 years = alle 2 Jahre
	OCCURENCES                     string    `xml:"OCCURENCES,omitempty" json:"OCCURENCES,omitempty"`                                         //	Anzahl der wiederkehrenden Rechnungen: 0 = unbegrenzt
	OUTPUT_TYPE                    string    `xml:"OUTPUT_TYPE,omitempty" json:"OUTPUT_TYPE,omitempty"`                                       // Required	Typ der automatisch zu erzeugenden Rechnung: draft = Entwurf | Outgoing = Rechnung
	EMAIL_NOTIFY                   string    `xml:"EMAIL_NOTIFY,omitempty" json:"EMAIL_NOTIFY,omitempty"`                                     //	Flag für den Versand an die eigene e-Mailadresse: 0 = nein | 1 = ja
	PROJECT_NAME                   string    `xml:"EMAIL_NOTIFY,omitempty" json:"EMAIL_NOTIFY,omitempty"`                                     //Name für ein Projekt
	HOUR_PRICE                     string    `xml:"HOUR_PRICE,omitempty" json:"HOUR_PRICE,omitempty"`                                         //	Stundensatz
	END_DATE                       string    `xml:"END_DATE,omitempty" json:"END_DATE,omitempty"`                                             //	Enddatum
}
