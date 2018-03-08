gofastbill
======
~~~
go get github.com/Fruchtgummi/gofastbill
~~~
gofastbill a API for Fastbill

Here are the things you can do until now:

  * Authentication : Login => Email, API-Key | Login for Mobile and Add-Ons with => Email, Password | Type xml or json
  * Invoice : get, create, update, delete, complete, cancel, sign, sendbyemail, sendbypost, setpaid
  * Template : get
  * Webhooks : sendto

Additional Features
-------------------
  * Minimal Models with hooks
  * Authentication ( https://gowalker.org/github.com/Fruchtgummi/gofastbill#Authentication, Connect with Email & API-Key or as Mobile & Add-On with Email & Passwort  )

Any ideas for the API or bug fixes please feel free to create a issue

Documentation
-------------

https://gowalker.org/github.com/Fruchtgummi/gofastbill

Basic Usage
-----------
~~~~

	import "github.com/Fruchtgummi/gofastbill"
	
	
	you find out all Structs for request and response on "gofastbill"-object


~~~~

Authentication and GET a Customer
-----------------------------------
~~~
    //Initialisierung
   
    fastbill, err = gofastbill.Init("email@domain.com", "xxxxxf2f39ebe33418e4f173cxxxx05IdihIGYAxSiABUfXgMeMBExxxpBC4xxx", false, true, "json")
	if err != nil {
		log.Println(err)
	}

	var x gofastbill.CustomerGet_Request
	
	//is CUSTOMER_ID empty you get all customers
	x.CUSTOMER_ID = "123456"
	
	fbapi, err := fastbill.Customer_Get(x)

	if err != nil {
		log.Println(err)
	}

	log.Println(fbapi)

	...


~~~~

Create a New Customer
-----------------------------------
~~~
	
    var x gofastbill.CustomerCreate_Request

	x.CUSTOMER_TYPE = "consumer" // or business
	x.LAST_NAME = "Mustermann" // Required by consumer
	x.ORGANIZATION = "Mustermann GmbH" //Required by business

	fbapi, err := fastbill.Customer_Create(x)

	if err != nil {
		log.Println(err)
	}

	log.Println(fbapi)


~~~~