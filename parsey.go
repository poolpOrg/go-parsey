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

package parsey

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Configuration struct {
	lexer   *Lexer
	grammar *Grammar
}

func NewConfiguration(lexer *Lexer, grammar *Grammar) *Configuration {
	return &Configuration{lexer: lexer, grammar: grammar}
}

func (config *Configuration) ParseFile(filename string) (bool, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer fp.Close()

	return config.ParseReader(bufio.NewReader(fp))
}

func (config *Configuration) ParseReader(rd *bufio.Reader) (bool, error) {
	i := 0
	success := true
	for {
		i++
		line, err := readline(rd)
		if err != nil {
			if err == io.EOF {
				return success, nil
			}
			return false, err
		}

		tokens, err := config.lexer.Tokenize(line)
		if err != nil {
			return false, err
		}

		if len(tokens) == 0 {
			continue
		}

		if !config.grammar.Match(tokens) {
			fmt.Println("error on line", i, ":", line)
			success = false
			continue
		}
	}
}

func readline(rd *bufio.Reader) (string, error) {
	cur := ""
	for {
		line, isPrefix, err := rd.ReadLine()
		if err != nil {
			return cur, err
		}
		if len(cur) != 0 && cur[len(cur)-1] == '\\' {
			cur = cur[:len(cur)-1]
		}
		cur += string(line)
		if !isPrefix || len(line) == 0 || line[len(line)-1] != '\\' {
			return cur, nil
		}
	}
}
