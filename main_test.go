package main

import "testing"

func TestFieldProductivity1(t *testing.T) {
	field := Field{Name: "1", Area: 200, Productions: []int{200, 450, 120, 30}}
	expectedProductivity := 4.0
	productivity := field.Productivity()
	if productivity == expectedProductivity {
		t.Logf("PASSED! Result: %.2f", productivity)
	} else {
		t.Errorf("ERROR! Expected: %.2f  Result: %.2f", expectedProductivity, productivity)
	}
}

func TestFieldProductivity2(t *testing.T) {
	field := Field{Name: "2", Area: 50, Productions: []int{500, 210, 720}}
	expectedProductivity := 28.6
	productivity := field.Productivity()
	if productivity == expectedProductivity {
		t.Logf("PASSED! Result: %.2f", productivity)
	} else {
		t.Errorf("ERROR! Expected: %.2f  Result: %.2f", expectedProductivity, productivity)
	}
}

func TestFarmProductivity(t *testing.T) {
	fieldsFazenda := []Field{
		{Name: "1", Area: 200, Productions: []int{200, 450, 120, 30}},
		{Name: "2", Area: 50, Productions: []int{500, 210, 720}},
	}
	farm := &Farm{Name: "Fazenda", Fields: fieldsFazenda}
	expectedProductivity := 8.92
	productivity := farm.Productivity()
	if productivity == expectedProductivity {
		t.Logf("PASSED! Result: %.2f", productivity)
	} else {
		t.Errorf("ERROR! Expected: %.2f  Result: %.2f", expectedProductivity, productivity)
	}
}
