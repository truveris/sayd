// Copyright 2014, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	ParseCommandLine()

	err := ParseConfigFile()
	if err != nil {
		log.Fatal("config error: ", err.Error())
	}

	log.Printf("starting sayd on %s", cfg.HTTPServerAddress)
	s := &http.Server{
		Addr:         cfg.HTTPServerAddress,
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("error: ", err)
	}
}
