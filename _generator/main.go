package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"golang.org/x/tools/imports"
)

var fn = flag.String("o", "generated.go", "Name of the generated file")

const url = "https://raw.githubusercontent.com/muan/emojilib/master/emojis.json"

type Emojis map[string]Emoji

type Emoji struct {
	Keywords []string `json:"keywords"`
	Char     string   `json:"char"`
	Category string   `json:"category"`
}

func main() {
	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var emojis Emojis

	json.NewDecoder(resp.Body).Decode(&emojis)

	src := `// DO NOT EDIT THIS FILE!
					//
					// Update it by running: go generate
					//
					`

	generatedAt := time.Now().UTC().Format("2006-01-02 15:04:05 -0700")

	src += "// Generated at: " + generatedAt + `

					package emojilib

					var emojis = Emojis{
					`
	var (
		emojiNames     = []string{}
		keywordNames   = []string{}
		keywordBuckets = map[string][]string{}
	)

	for n, _ := range emojis {
		// Skip the keys
		if n == "keys" {
			continue
		}

		emojiNames = append(emojiNames, n)
	}

	sort.Strings(emojiNames)

	for _, n := range emojiNames {
		e := emojis[n]

		var ks string

		if len(e.Keywords) > 0 {
			ks = `"` + strings.Join(e.Keywords, `","`) + `"`

			for _, k := range e.Keywords {
				keywordBuckets[k] = append(keywordBuckets[k], n)
			}
		}

		src += `"` + n + `": Emoji{` + "\n" +
			`Keywords: []string{` + ks + "},\n" +
			`Char: "` + e.Char + "\",\n" +
			`Category: "` + e.Category + "\",\n" +
			"},\n"
	}
	src += "}\n\n"

	for kn, _ := range keywordBuckets {
		keywordNames = append(keywordNames, kn)
	}

	sort.Strings(keywordNames)

	src += "var keywordLookup = map[string][]string{\n"

	for _, n := range keywordNames {
		src += `"` + n + "\": []string{\n \"" + strings.Join(keywordBuckets[n], "\",\n \"") + "\",\n},\n"
	}
	src += "}\n\n"

	src += "var emojiReplacer = strings.NewReplacer(\n"

	for _, n := range emojiNames {
		src += `":` + n + `:", "` + emojis[n].Char + "\",\n"
	}
	src += ")\n\n"

	src += "var emojiPaddedReplacer = strings.NewReplacer(\n"

	for _, n := range emojiNames {
		src += `":` + n + `:", "` + emojis[n].Char + " \",\n"
	}
	src += ")\n\n"

	res, err := imports.Process(*fn, []byte(src), nil)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*fn, res, 0644)
	if err != nil {
		panic(err)
	}
}
