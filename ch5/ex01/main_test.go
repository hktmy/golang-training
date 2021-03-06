package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args string
		want []string
	}{
		{`<a href="foo">Foo</a>`, []string{"foo"}},
		{`<ul><li><a href="/foo">Foo</a></li><li><a href="/bar">Bar</a></li></ul>`, []string{"/foo", "/bar"}},
		{`<ul><li><a href="/foo">Foo</a></li><li><a href="/bar">Bar</a></li></ul><ul><li><a href="/hoge">Hoge</a></li><li><a href="/piyo">Piyo</a></li></ul>`, []string{"/foo", "/bar", "/hoge", "/piyo"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("visit(%q)", test.args)
		doc, err := html.Parse(strings.NewReader(test.args))
		if err != nil {
			log.Fatal(err)
		}
		links := visit(nil, doc)
		if !reflect.DeepEqual(links, test.want) {
			t.Errorf("%s", descr)
			t.Errorf("got")
			for _, link := range links {
				t.Errorf("%s", link)
			}
			t.Errorf("expect")
			for _, link := range test.want {
				t.Errorf("%s", link)
			}
		}
	}
}
