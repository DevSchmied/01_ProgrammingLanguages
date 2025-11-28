package main

import "testing"

func TestAnalyzeAttendance_InvalidID(t *testing.T) {
	data := []string{"2", "0 5", "1 7"}
	r := AnalyzeAttendance(data)
	if r.Valid != -1 {
		t.Errorf("expected invalid result")
	}
}
