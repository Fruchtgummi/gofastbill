gofastbill
======

*gofastbill* is API for FastBill

FastBill is a finance-app specifically for small businesses and the self-employed. 
With FastBill have your finances easy under control, for example, easily and easily create invoices or collect receipts. 

https://www.fastbill.com/en/

~~~
go get github.com/Fruchtgummi/gofastbill
~~~


Here are the things you can do until now:

  * Authentication : Login => Email, API-Key | Login for Mobile and Add-Ons with => Email, Password | Type xml or json
  * Customer : get, create, update, delete
  * Contact : get, create, update, delete
  * Invoice : get, create, update, delete, complete, cancel, sign, sendbyemail, sendbypost, setpaid
  * Article : get, create, update, delete
  * Estimate : get, create, sendbyemail, creatinvoice, delete
  * Template : get
  * Item : get, delete
  * Time : get, create, update, delete
  * Expense : get, create **(TODO: Upload for paid document)**
  * Document : get, create **(TODO: Upload for document)**
  * Recurring : get, create, update, delete
  * Revenue : get, create, setpaid, delete
  * Project : get, create, update, delete

cooming soon:
  
  * Webhooks : get, create, delete (BETA)
  * GO-Routines (threads)
  * Chan (Stats) 

Features
-------------------
  * Minimal Models with hooks
  * Authentication (https://gowalker.org/github.com/Fruchtgummi/gofastbill#Init, Connect with Email & API-Key or as Mobile & Add-On with Email & Passwort)

Any ideas for the API or bug fixes please feel free to create a issue

Documentation
-------------

https://gowalker.org/github.com/Fruchtgummi/gofastbill

Basic Usage
-----------
~~~~
	import "github.com/Fruchtgummi/gofastbill"
~~~~


you find out all Structs for request and response on "gofastbill"-object


Authentication and GET a Customer
-----------------------------------
~~~
	//Initialisierung
	
	fastbill, err = gofastbill.Init("email@domain.com", "xxxxxf2f39ebe33418e4f173cxxxx05IdihIGYAxSiABUfXgMeMBExxxpBC4xxx", false, true, "json")
	
	if err != nil {
		log.Println(err)
	}
	
	var x gofastbill.CustomerGet_Request
	
	
	x.CUSTOMER_ID = "123456" //is CUSTOMER_ID empty you get all customers
	
	fbapi, err := fastbill.Customer_get(x)
	
	if err != nil {
		log.Println(err)
	}
	
	log.Println(fbapi)
	
	...


~~~~

Create a NEW Customer
-----------------------------------
~~~	
	var x gofastbill.CustomerCreate_Request
	
	x.CUSTOMER_TYPE = "consumer" 					// or business
	x.LAST_NAME = "Mustermann" 						// Required by consumer
	x.ORGANIZATION = "Mustermann GmbH" 				//Required by business
	
	fbapi, err := fastbill.Customer_create(x)
	
	if err != nil {
		log.Println(err)
	}
	
	log.Println(fbapi)

~~~~


Create a Invoice
-----------------------------------
~~~	
	var x gofastbill.InvoiceCreate_Request
	
	x.CUSTOMER_ID = "123456"
	x.TEMPLATE_ID = "123456"
	x.INTROTEXT = "I'm a example"
	x.INVOICE_TITLE = "Title of Invoice"
	
	var y gofastbill.ITEMS
	
	//loop	
	
		y.ARTICLE_NUMBER = "1234567"
		y.DESCRIPTION = "I'm a description of ITEM"
		y.QUANTITY = "4"
		y.UNIT_PRICE = "25.00"
		y.VAT_PERCENT = "19"
		
		x.ITEMS = append(x.ITEMS, y)
	
	//loop end
	
	fbapi, err := fastbill.Invoice_create(x)
	
	if err != nil {
		log.Println(err)
	}
~~~~