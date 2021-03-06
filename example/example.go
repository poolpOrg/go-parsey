/*
 * Copyright (c) 2021 Gilles Chehade <gilles@poolp.org>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/poolpOrg/go-parsey"
)

func configBuilderRule1(config *parsey.Configuration, tokens []parsey.Token) error {
	return nil
}

func configBuilderRule2(config *parsey.Configuration, tokens []parsey.Token) error {
	return nil
}

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "example.conf", "configuration file")
	flag.Parse()

	lexer := parsey.NewLexer()
	lexer.RegisterToken("listen")
	lexer.RegisterToken("on")
	lexer.RegisterToken("match")
	lexer.RegisterToken("=>")
	lexer.RegisterTokenMatch("STRING", lexer.IsString)
	lexer.RegisterTokenMatch("NUMBER", lexer.IsNumber)
	lexer.RegisterTokenMatch("FLOAT", lexer.IsFloat)

	grammar := parsey.NewGrammar()
	grammar.RegisterRule(configBuilderRule1, "listen", "on", "STRING")
	grammar.RegisterRule(configBuilderRule2, "match", "=>", "STRING")

	config := parsey.NewConfiguration(lexer, grammar)

	success, err := config.ParseFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if success {
		fmt.Println(configFile, "parsed successfully")
	}
}
