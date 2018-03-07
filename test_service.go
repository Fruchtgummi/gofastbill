package gofastbill

import (
	"log"
)

var fastbill *Initialization

func main() {

	fastbill, err := Init("roman.koch@schwaebisch-hall.de", "99a0cf2f39ebe33418e4f173c93bd005IdihIGYAxSiABUfXgMeMBETL6pBC4hbH", false)
	if err != nil {
		log.Println(err)
	}

	log.Println(fastbill)

}
