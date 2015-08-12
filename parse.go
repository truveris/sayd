// Copyright 2014-2015, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"strings"
)

// Extract the voice and format from the URL. It should be a single word after the slash.
func parseVoiceFromPath(path string) (string, string) {
	// First extract the first part of the path, the only one that matters.
	tokens := strings.SplitN(strings.TrimLeft(path, "/"), "/", 2)
	// Try to find a format, only aiff and mp3 should be valid.
	tokens = strings.Split(tokens[0], ".")

	voice := tokens[0]
	format := ""

	if len(tokens) > 1 {
		format = tokens[1]
	} else {
		format = "aiff"
	}

	return voice, format
}
