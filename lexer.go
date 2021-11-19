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
	"strconv"
	"strings"
)

type TokenType string

type Lexer struct {
	tokens      map[TokenType]func(string) bool
	tokensMatch map[TokenType]func(string) bool
}

type Token struct {
	tokenType TokenType
	buffer    string
}

func NewLexer() *Lexer {
	return &Lexer{tokens: make(map[TokenType]func(string) bool), tokensMatch: make(map[TokenType]func(string) bool)}
}

func (lexer *Lexer) RegisterToken(token TokenType) error {
	if _, exists := lexer.tokens[token]; !exists {
		lexer.tokens[token] = func(tokenName string) bool { return TokenType(tokenName) == token }
	}
	return nil
}

func (lexer *Lexer) RegisterTokenMatch(token TokenType, detector func(string) bool) error {
	if _, exists := lexer.tokensMatch[token]; !exists {
		lexer.tokensMatch[token] = detector
	}
	return nil
}

func (lexer *Lexer) GetTokenType(buffer string) TokenType {
	if _, exists := lexer.tokens[TokenType(buffer)]; exists {
		return TokenType(buffer)
	}

	for tokenType, tokenValidator := range lexer.tokensMatch {
		if tokenValidator(buffer) {
			return tokenType
		}
	}
	return ""
}

func (lexer *Lexer) Tokenize(buffer string) ([]Token, error) {
	tokens := make([]Token, 0)

	var token *Token
	skipComment := false
	skipSpaces := true

	for _, character := range buffer {
		if skipComment {
			continue
		}
		if skipSpaces && character == ' ' || character == '\t' {
			continue
		}
		skipSpaces = false

		if token == nil {
			token = &Token{}
		}

		switch character {
		case ' ':
			fallthrough
		case '\t':
			if token != nil {
				token.tokenType = lexer.GetTokenType(token.buffer)
				tokens = append(tokens, *token)
				token = nil
			}
			skipSpaces = true
			continue

		default:
			token.buffer += string(character)
		}

	}
	if token != nil {
		token.tokenType = lexer.GetTokenType(token.buffer)
		tokens = append(tokens, *token)
	}

	return tokens, nil
}

func (lexer *Lexer) IsString(buffer string) bool {
	_, err := strconv.Atoi(buffer)
	return err != nil
}

func (lexer *Lexer) IsNumber(buffer string) bool {
	_, err := strconv.Atoi(buffer)
	return err == nil
}

func (lexer *Lexer) IsFloat(buffer string) bool {
	_, err := strconv.Atoi(buffer)
	return err == nil && strings.Contains(buffer, ".")
}
