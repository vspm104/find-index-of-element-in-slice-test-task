package main

import (
	"net/http"
	"testing"
)

type searchTest struct {
	arg1     []int
	arg2     int
	expected int
}

var fixture = []int{0, 10, 1000, 1000000}

func TestBelowRange(t *testing.T) {

	got := indexSearch(fixture, -1)
	want := http.StatusBadRequest

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLowBorderValue(t *testing.T) {

	got := indexSearch(fixture, 0)
	want := http.StatusOK

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestExactMatch(t *testing.T) {

	got := indexSearch(fixture, 10)
	want := http.StatusOK

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestTenLevelMatch(t *testing.T) {

	got := indexSearch(fixture, 950)
	want := http.StatusOK

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIndexNotFound(t *testing.T) {

	got := indexSearch(fixture, 2000)
	want := http.StatusBadRequest

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestHighBorderValue(t *testing.T) {

	got := indexSearch(fixture, 1000000)
	want := http.StatusOK

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAboveBorder(t *testing.T) {

	got := indexSearch(fixture, 1500000)
	want := http.StatusBadRequest

	if got.Code != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
