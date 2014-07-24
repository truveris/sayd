// Copyright 2014, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/jessevdk/go-flags"
)

type Cmd struct {
	ConfigFile string `short:"c" description:"Configuration file" default:"/etc/sayd.conf"`
}

type Cfg struct {
	// If defined, start a web server to list the aliases (e.g. :8989)
	HTTPServerAddress string
}

var (
	cfg = Cfg{}
	cmd = Cmd{}
)

// Look in the current directory for an config.json file.
func ParseConfigFile() error {
	file, err := os.Open(cmd.ConfigFile)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	if cfg.HTTPServerAddress == "" {
		return errors.New("'HTTPServerAddress' is not defined")
	}

	return nil
}

// Parse the command line arguments and populate the global cmd struct.
func ParseCommandLine() {
	flagParser := flags.NewParser(&cmd, flags.PassDoubleDash)
	_, err := flagParser.Parse()
	if err != nil {
		println("command line error: " + err.Error())
		flagParser.WriteHelp(os.Stderr)
		os.Exit(1)
	}
}
