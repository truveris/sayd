// Copyright 2014, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, err error, msg string) {
	if err == nil {
		log.Printf("error: %s", msg)
	} else {
		log.Printf("error: %s", err.Error())
	}
	http.Error(w, msg+" (see logs for more details)", 500)
}
