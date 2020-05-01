package main

import "testing"

//the thest func will be called automatically by the go test runner with that argument
//in cod am pus 4 carti si 4 culori => 16 carti in total
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 16, but got ", len(d))
	}
}

// to make a test, create a new file ending in  ceva_test.go
// then run go test.
// fara numele fisierului, doar:
// go test

// ai functia x pe care vrei sa o testezi, in fisierul _test.go creezi functia TestX, in care compui testul.
// (daca functia initiala incepea cu litera mica, e ok sa inceapa cu litera mare TestCeva )
// Test cu T mare, pt ca vrei sa fie available in afara..
