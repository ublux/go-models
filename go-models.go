package gomodels

import (
	"fmt"
	"time"

	"github.com/ublux/go-models/models"
)

// TODO do not compile backend content
// https://stackoverflow.com/q/36703867/637142
// Remove properties like IdAccount for client

// rm -r /root/repos/go-models/models ; cp -r /tmp/output/models /root/repos/go-models/
// rm -r /root/repos/go-models/enums ; cp -r /tmp/output/enums /root/repos/go-models/

func Add(a, b int) int {
	return a + b
}

func foo() {
	var c1 models.Contact
	fmt.Println(c1)

	// Calling Date() method
	// with all its parameters
	tm := time.Date(2020, time.April, 11, 21, 34, 01, 0, time.UTC)

	fmt.Println(tm)
}
