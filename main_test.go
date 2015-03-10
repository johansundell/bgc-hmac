package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFall_1A(t *testing.T) {
	in := "00000000"
	expect := "FF365893D899291C3BF505FB3175E880"

	result, err := getHmacString(in)
	if err != nil {
		t.Error(err)
	}
	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}

	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "testfall-1a.txt")
	result, _ = getHmacByte(dat)
	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestFall_1B(t *testing.T) {
	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "testfall-1b.txt")
	expect := "9BA94363739D45256DF4B6FA3B9DE1CD"

	result, err := getHmacByte(dat)
	if err != nil {
		t.Error(err)
	}

	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestFall_1C(t *testing.T) {
	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "testfall-1c.txt")
	expect := "826CA6CBA33F7E1D3CC9161A0956A35B"

	result, err := getHmacByte(dat)
	if err != nil {
		t.Error(err)
	}

	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestFall_2A(t *testing.T) {
	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "hmac_testfall-2a.txt")
	expect := "9473EDFCAA8CD2434D6D76ABFFC991BD"

	result, err := getHmacByte(dat)
	if err != nil {
		t.Error(err)
	}

	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestFall_2B(t *testing.T) {
	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "hmac_testfall-2b.txt")
	expect := "20956E44B404C4085446139B2B952D77"

	result, err := getHmacByte(dat)
	if err != nil {
		t.Error(err)
	}

	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestFall_2C(t *testing.T) {
	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "hmac_testfall-2c.txt")
	expect := "515704694958361678194D51850FF157"

	result, err := getHmacByte(dat)
	if err != nil {
		t.Error(err)
	}

	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestFall_2D(t *testing.T) {
	dat, _ := ioutil.ReadFile("files" + string(os.PathSeparator) + "hmac_testfall-2d.txt")
	expect := "68C5954818010DDDCB1EC02963E8C123"

	result, err := getHmacByte(dat)
	if err != nil {
		t.Error(err)
	}

	if result != expect {
		t.Errorf("Expeded %s, got %s", expect, result)
	}
}

func TestWrongSecretLenght(t *testing.T) {
	tmp := secret
	secret = "123"
	_, err := getHmacString("123456")
	if err == nil {
		t.Error(err)
	}
	secret = tmp
}
