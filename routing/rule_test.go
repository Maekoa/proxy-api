package routing

import (
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestMatch(t *testing.T) {
	dict := Set{}
	for _, p := range []string{
		"baidu.com",
		"google.com",
		"v2ex.asia",
		"sedkiuyer",
		"awiuskiefurg",
		"awdiusgyef",
	} {
		dict.Add(p)
	}
	if dict.MatchSuffix("baidu.com") != true {
		t.Error()
	}
	if dict.MatchSuffix("baidu.cn") != false {
		t.Error()
	}

}

func BenchmarkAdd(b *testing.B) {
	dict := Set{}
	for i := 0; i < b.N; i++ {
		dict.Add(randSeq(1))
	}
}

func BenchmarkAddInline(b *testing.B) {
	dict := Set{}
	for i := 0; i < b.N; i++ {
		dict[randSeq(1)] = Null
	}
}
