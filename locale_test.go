package lingo

import (
	"testing"
)

func TestLocale(t *testing.T) {
	l1 := supportedLocales("en,de-AT; q=0.8,de;q=0.6,bg; q=0.4,en-US;q=0.2,sr;q=0.2")
	if len(l1) != 6 {
		t.Error(l1)
		t.Fail()
	}
	l2 := supportedLocales("en")
	if len(l2) != 1 {
		t.Error(l2)
		t.Fail()
	}
	l3 := supportedLocales("")
	if len(l3) != 0 {
		t.Error(l3)
		t.Fail()
	}
	l4 := ParseLocale("en_US")
	if l4.Lang != "en" || l4.Country != "US" {
		t.Error(l4)
		t.Fail()
	}
	l5 := ParseLocale("en")
	if l5.Lang != "en" || l5.Country != "" {
		t.Error(l4)
		t.Fail()
	}

}
