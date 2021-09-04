package parser

import (
	"io/ioutil"
	"testing"

	"brucego.com/learngo/crawler/model"
)

// Name       string
// Gender     string
// Age        int
// Height     int
// Weight     int
// Income     string
// Marriage   string
// Education  string
// Occupation string
// Location   string
// Hometown   string
// Horoscope  string
// House      string
// Car        string
func TestParseProfile(t *testing.T) {

	content, err := ioutil.ReadFile("profile.html")
	if err != nil {
		panic(err)
	}
	results := ParseProfile(content, "test")
	profile := results.Items[0].(model.Profile)
	if profile.Name != "test" {
		t.Errorf("Gender should be %v instead of test", profile.Name)

	}
	if profile.Gender != "女" {
		t.Errorf("Gender should be %v instead of 女", profile.Gender)

	}

	if profile.Age != 34 {
		t.Errorf("Age should be %v ", profile.Age)

	}
	if profile.Height != 157 {
		t.Errorf("Height should be %v ", profile.Height)

	}
	if profile.Weight != 0 {
		t.Errorf("Weight should be %v ", profile.Weight)

	}
	if profile.Income != "3-5千" {
		t.Errorf("Income should be %v ", profile.Income)

	}
	if profile.Marriage != "离异" {
		t.Errorf("Marriage should be %v", profile.Marriage)

	}

	if profile.Education != "中专" {
		t.Errorf("Education should be %v ", profile.Education)

	}
	if profile.Occupation != "" {
		t.Errorf("Occupation should be %v ", profile.Occupation)

	}
	if profile.Location != "安康汉阴" {
		t.Errorf("Location should be %v ", profile.Location)

	}
	if profile.Hometown != "陕西安康" {
		t.Errorf("Hometown should be %v ", profile.Hometown)

	}
	if profile.Horoscope != "魔羯座(12.22-01.19)" {
		t.Errorf("Horoscope should be %v ", profile.Horoscope)

	}
	if profile.Car != "已买车" {
		t.Errorf("Car should be %v ", profile.Car)

	}
	if profile.House != "租房" {
		t.Errorf("House should be %v ", profile.House)

	}
}
