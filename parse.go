// Copyright 2014, Truveris Inc. All Rights Reserved.
// Use of this source code is governed by the ISC license in the LICENSE file.

package main

import (
	"strings"
)

// Extract the voice from the URL. It should be a single word after the slash.
func parseVoiceFromPath(path string) string {
	tokens := strings.SplitN(strings.TrimLeft(path, "/"), "/", 2)
	return tokens[0]
}
