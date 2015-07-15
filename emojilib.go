/*

Package emojilib is a port of the Emoji keyword library to Go

Installation

Just go get the package:

    go get -u github.com/peterhellberg/emojilib

Usage

A small usage example

		package main

		import (
			"fmt"

			"github.com/peterhellberg/emojilib"
		)

		func main() {
			fmt.Println(emojilib.ReplaceWithPadding("I :green_heart: You!"))
		}

*/
package emojilib

import "errors"

//go:generate go run _generator/main.go

// Emojis contain emojis keyed on their name
type Emojis map[string]Emoji

// Emoji contains the keywords, char and category for an emoji
type Emoji struct {
	Keywords []string `json:"keywords"`
	Char     string   `json:"char"`
	Category string   `json:"category"`
}

// ErrUnknownEmoji is returned from Find if provided with a unknown emoji name
var ErrUnknownEmoji = errors.New("unknown emoji")

// Find returns an Emoji if provided with a known name
func Find(n string) (Emoji, error) {
	if e, ok := emojis[n]; ok {
		return e, nil
	}

	return Emoji{}, ErrUnknownEmoji
}

// All returns all the emojis
func All() Emojis {
	return emojis
}

// Replace takes a string and replaces all emoji names with their emoji character
func Replace(s string) string {
	return emojiReplacer.Replace(s)
}

// ReplaceWithPadding takes a string and replaces all emoji names with their
// emoji character and a space in order to display better in terminals
func ReplaceWithPadding(s string) string {
	return emojiPaddedReplacer.Replace(s)
}
