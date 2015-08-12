// Copyright 2014-2015, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		post(w, r)
	default:
		errorHandler(w, nil, "unsupported method")
	}
}

// The GET method for this service expects the voice to be used as only path in
// the URL and the sentence to be in the query string, for example:
//
//   /alex?hello%20world!
//
func get(w http.ResponseWriter, r *http.Request) {
	voice, format := parseVoiceFromPath(r.URL.Path)
	sentence, err := url.QueryUnescape(r.URL.RawQuery)
	if err != nil {
		errorHandler(w, err, "bad query")
		return
	}

	responseHandler(w, sentence, voice, format)
}

// The POST method expects the voice to be used as the only path and the
// sentence to be passed as data.
func post(w http.ResponseWriter, r *http.Request) {
	voice, format := parseVoiceFromPath(r.URL.Path)

	sentence, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, err, "bad request body")
		return
	}

	responseHandler(w, string(sentence), voice, format)
}

func responseHandler(w http.ResponseWriter, sentence, voice, format string) {
	dirpath, err := ioutil.TempDir("", "sayd")
	if err != nil {
		errorHandler(w, err, "failed to run say")
		return
	}

	mimetype := "audio/aiff"
	filepath := dirpath + "/output.aiff"

	err = say(sentence, voice, filepath)
	if err != nil {
		errorHandler(w, err, "failed to run say")
		return
	}

	if format == "mp3" {
		mimetype = "audio/mpeg"
		mp3path := dirpath + "/output.mp3"
		aiff2mp3(filepath, mp3path)
		filepath = mp3path
	}

	f, err := os.Open(filepath)
	if err != nil {
		errorHandler(w, err, "failed to run say")
		return
	}

	info, err := f.Stat()
	if err != nil {
		errorHandler(w, err, "failed to run say")
		return
	}

	w.Header().Set("Content-type", mimetype)
	w.Header().Set("Content-length", fmt.Sprintf("%d", info.Size()))
	_, err = io.Copy(w, f)
	if err != nil {
		errorHandler(w, err, "failed to run say")
		return
	}

	f.Close()

	os.RemoveAll(dirpath)
}
