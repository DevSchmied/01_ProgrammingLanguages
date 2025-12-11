package main

import "testing"

func TestAnalyzeAttendance_InvalidID(t *testing.T) {
	data := []string{"2", "0 5", "1 7"}
	r := AnalyzeAttendance(data)
	if r.Valid != -1 {
		t.Errorf("expected invalid result")
	}
}

func TestAnalyzeAttendance_InvalidNFormat(t *testing.T) {
	data := []string{"abc", "1", "2"}
	r := AnalyzeAttendance(data)
	if r.Valid != -1 {
		t.Errorf("expected invalid result for non-numeric N")
	}
}

func TestAnalyzeAttendance_NotEnoughLines(t *testing.T) {
	data := []string{"3", "1 2", "2 3"}
	r := AnalyzeAttendance(data)
	if r.Valid != -1 {
		t.Errorf("expected invalid result when data lines are fewer than N")
	}
}

func TestAnalyzeAttendance_InvalidFieldsCount(t *testing.T) {
	data := []string{"2", "1", "2 3"}
	r := AnalyzeAttendance(data)
	if r.Valid != -1 {
		t.Errorf("expected invalid result for malformed student record with missing fields")
	}
}
