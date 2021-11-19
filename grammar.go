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

type Grammar struct {
	rules []Rule
}

type Rule struct {
	handler func()
	tokens  []TokenType
}

func NewGrammar() *Grammar {
	grammar := &Grammar{}
	grammar.rules = make([]Rule, 0)
	return grammar
}

func (grammar *Grammar) RegisterRule(handler func(), tokens ...interface{}) {
	rule := Rule{}
	rule.handler = handler
	rule.tokens = make([]TokenType, 0)

	for _, token := range tokens {
		tokenType := TokenType(token.(string))
		rule.tokens = append(rule.tokens, tokenType)
	}

	grammar.rules = append(grammar.rules, rule)
}

func (grammar *Grammar) Match(tokens []Token) bool {
	matched := false
	for _, rule := range grammar.rules {
		if len(rule.tokens) != len(tokens) {
			continue
		}

		for offset, token := range tokens {
			if token.tokenType != rule.tokens[offset] {
				break
			}
			if offset == len(tokens)-1 {
				matched = true
				break
			}
		}
		if matched {
			rule.handler()
			return true
		}
	}
	return false
}
