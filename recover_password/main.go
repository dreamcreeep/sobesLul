package main

import (
	"crypto/md5"
	"testing"
)

var alphabet = []rune{'a', 'b', 'c', 'd', '1', '2', '3'}

func RecoverPassword(h []byte) string {

	return "" // рекурсия
}

func TestRecoverPassword(t *testing.T) {
	for _, exp := range []string{
		"a",
		"12",
		"abc333d",
	} {
		t.Run(exp, func(t *testing.T) {
			act := RecoverPassword(hashPassword(exp))
			if act != exp {
				t.Error("recovered:", act, "expected:", exp)
			}
		})
	}
}

func hashPassword(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
