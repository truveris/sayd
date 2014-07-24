// Copyright 2014, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

// Create a sound file from the given sentence and voice.
func say(sentence, voice, filepath string) error {
	log.Printf("say(%s, %s, %s)", sentence, voice, filepath)

	cmd := exec.Command("say", "-v", voice, "-o", filepath, sentence)
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

	if strings.Contains(string(errdata), "Voice") {
		return errors.New(string(errdata))
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
