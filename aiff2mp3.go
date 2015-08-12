// Copyright 2015, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// Using ffmpeg, convert the given AIFF file to MP3.
func aiff2mp3(aiffpath, mp3path string) error {
	log.Printf("aiff2mp3(%s, %s)", aiffpath, mp3path)

	cmd := exec.Command("ffmpeg", "-i", aiffpath, mp3path)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	errdata, err := ioutil.ReadAll(stderr)
	if err != nil {
		return err
	}

	if strings.Contains(string(errdata), "error") {
		return errors.New(string(errdata))
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
