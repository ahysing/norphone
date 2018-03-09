package main

import (
	"fmt"
	"testing"
)

func Test_Empty(t *testing.T) {
	assertSame(t, "", "")
}

func Test_Aarrestad(t *testing.T) {
	assertSame(t, "Aarrestad", "Årrestad")
}

func Test_Andreassen(t *testing.T) {
	assertSame(t, "Andreasen", "Andreassen")
}

func Test_Arntsen(t *testing.T) {
	assertSame(t, "Arntsen", "Arntzen")
}

func Test_Bache(t *testing.T) {
	assertSame(t, "Bache", "Bakke")
}

func Test_Franck(t *testing.T) {
	assertSame(t, "Frank", "Franck")
}

func Test_Christian(t *testing.T) {
	assertSame(t, "Christian", "Kristian")
}

func Test_Kielland(t *testing.T) {
	assertSame(t, "Kielland", "Kjelland")
}

func Test_Krogh(t *testing.T) {
	assertSame(t, "Krogh", "Krog")
}

func Test_Krohg(t *testing.T) {
	assertSame(t, "Krog", "Krohg")
}

func Test_Jendahl(t *testing.T) {
	assertSame(t, "Jendal", "Jendahl")
}

func Test_Hjendal(t *testing.T) {
	assertSame(t, "Jendal", "Hjendal")
}

func Test_Gjendal(t *testing.T) {
	assertSame(t, "Jendal", "Gjendal")
}

func Test_Wold(t *testing.T) {
	assertSame(t, "Vold", "Wold")
}

func Test_Thomas(t *testing.T) {
	assertSame(t, "Thomas", "Tomas")
}

func Test_Aamodt(t *testing.T) {
	assertSame(t, "Aamodt", "Aamot")
}

func Test_Aksel(t *testing.T) {
	assertSame(t, "Aksel", "Axel")
}

func Test_Christophersen(t *testing.T) {
	assertSame(t, "Kristoffersen", "Christophersen")
}

func Test_Vold(t *testing.T) {
	assertSame(t, "Voll", "Vold")
}

func Test_Granlid(t *testing.T) {
	assertSame(t, "Granli", "Granlid")
}

func Test_Giever(t *testing.T) {
	assertSame(t, "Gjever", "Giever")
}

func Test_Sanderhaugen(t *testing.T) {
	assertSame(t, "Sannerhaugen", "Sanderhaugen")
}

func Test_Jahren(t *testing.T) {
	assertSame(t, "Jahren", "Jaren")
}

func Test_Amundsroed(t *testing.T) {
	assertSame(t, "Amundsrud", "Amundsrød")
}

func Test_Carlson(t *testing.T) {
	assertSame(t, "Karlson", "Carlson")
}

func assertSame(t *testing.T, key1 string, key2 string) {
	result1 := Norphone(key1)
	result2 := Norphone(key2)
	if result1 != result2 {
		msg := fmt.Sprintf("wrong key '%s' != '%s'", key1, key2)
		t.Errorf(msg)
	}
}
