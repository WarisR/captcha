package captcha_test

import (
	"testing"
	"github.com/WarisR/captcha"
)

func TestCaptcha(t *testing.T) {
	result := captcha.Generate("1111")

	expected := "1 + one"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptcha2(t *testing.T) {
	result := captcha.Generate("1211")

	expected := "2 + one"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptcha3(t *testing.T) {
	result := captcha.Generate("1311")

	expected := "3 + one"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

func TestCaptcha4(t *testing.T) {
	result := captcha.Generate("2321")

	expected := "three - 1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}


func TestCaptcha5(t *testing.T) {
	result := captcha.Generate("2511")

	expected := "five + 1"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}


func TestCaptcha6(t *testing.T) {
	result := captcha.Generate("2732")

	expected := "seven x 2"
	if expected != result {
		t.Errorf("it should say %q but get %q", expected, result)
	}
}

