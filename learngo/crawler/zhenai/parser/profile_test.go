package parser

import (
	"io/ioutil"
	"reflect"
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
var profile1 = model.Profile{
	Name:       "美丽心晴",
	Gender:     "女",
	Age:        34,
	Height:     157,
	Weight:     0,
	Income:     "3-5千",
	Marriage:   "离异",
	Education:  "中专",
	Occupation: "",
	Location:   "安康汉阴",
	Hometown:   "陕西安康",
	Horoscope:  "魔羯座(12.22-01.19)",
	House:      "租房",
	Car:        "已买车",
}
var profile2 = model.Profile{
	Name:       "与Ta相濡以沫 ",
	Gender:     "女",
	Age:        40,
	Height:     160,
	Weight:     65,
	Income:     "3千以下",
	Marriage:   "离异",
	Education:  "高中及以下",
	Occupation: "自由职业",
	Location:   "安康汉滨区",
	Hometown:   "四川攀枝花",
	Horoscope:  "天蝎座(10.23-11.21)",
	House:      "租房",
	Car:        "未买车",
}
var profile3 = model.Profile{
	Name:       "大树 ",
	Gender:     "男",
	Age:        40,
	Height:     160,
	Weight:     65,
	Income:     "5-8千",
	Marriage:   "离异",
	Education:  "大学本科",
	Occupation: "人事总监",
	Location:   "安康市辖区",
	Hometown:   "陕西安康",
	Horoscope:  "魔羯座(12.22-01.19)",
	House:      "已购房",
	Car:        "未买车",
}

var profileMap = map[model.Profile]string{
	profile1: "test_fixture/profile.html",
	profile2: "test_fixture/profile2.html",
	profile3: "test_fixture/profile3.html",
}

func TestParseProfile(t *testing.T) {

	for key, value := range profileMap {
		content, err := ioutil.ReadFile(value)
		if err != nil {
			panic(err)
		}
		results := ParseProfile(content, key.Name)
		profileReal := results.Items[0].(model.Profile)

		fields := reflect.TypeOf(key)
		values := reflect.ValueOf(key)
		resultsValues := reflect.ValueOf(profileReal)
		num := fields.NumField()

		for i := 0; i < num; i++ {
			field := fields.Field(i)
			value := values.Field(i)
			resultsValue := resultsValues.Field(i)
			if value.String() != resultsValue.String() {
				t.Errorf("%s %s should be %s:%s, instead of %s", key.Name, field.Name, field.Name, resultsValue, value)
			}
		}

	}
}
