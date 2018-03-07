gofastbill
======
~~~
go get github.com/Fruchtgummi/gofastbill
~~~
gofastbill a API for Fastbill

Here are the things you can do until now:

  * Authentication : Login => Email, API-Key | Login for Mobile and Add-Ons with => Email, Password
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

STURCT



~~~~

Authentication and Create a Invoice
-----------------------------------
~~~
    //write here a comment
   
    c, err := gofastbill.Authentication("https://my.fastbill.com/api/1.0/api.php","email@domain.com","AKI-Key", false) 
    if err != nil{
        panic(err)
    }

	...


~~~~

