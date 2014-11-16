package lingo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLingo(t *testing.T) {
	l := New("en_US", "translations")
	t1 := l.TranslationsForLocale("sr_RS")

	r1 := t1.Value("lingo.example.1")
	if r1 != "Srbija value 1" {
		t.Error(r1)
		t.Fail()
	}
	r2 := t1.Value("lingo.example.2")
	if r2 != "Nested Srbija" {
		t.Error(r2)
		t.Fail()
	}
	r3 := t1.Value("lingo.example.4.inception")
	if r3 != "Double nested?" {
		t.Error(r3)
		t.Fail()
	}
	r4 := t1.Value("lingo.example.5", "Arg1", "Milk", "4")
	if r4 != "Arguments? Arg1 Milk 4" {
		t.Error(r4)
		t.Fail()
	}
}

func TestLingoHttp(t *testing.T) {
	l := New("en_US", "translations")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := l.TranslationsForRequest(r)
		r1 := t1.Value("lingo.example.1")
		if r1 != "Srbija value 1" {
			t.Error(r1)
			t.Fail()
		}
	}))
	defer srv.Close()
	url, _ := url.Parse(srv.URL)
	req := &http.Request{
		Method: "GET",
		Header: map[string][]string{
			"Accept-Language": {"sr, en-gb;q=0.8, en;q=0.7"},
		},
		URL: url,
	}
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

}
