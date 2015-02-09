package test

import (
	"oncemsg/utility"
	"testing"
)

func TestCryAse(t *testing.T) {
	origin := "AAA"
	deed, err := utility.CryAse(origin)
	if err != nil {
		t.Error(err)
	}
	t.Log(deed)
	originC, _ := utility.DeCryAse(deed)
	if originC == origin {
		t.Log("PASS")
	} else {
		t.Error("FAIL" + origin + " " + deed + " " + originC)
	}
}
