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
	Name:       "lindy",
	Gender:     "女士",
	Age:        32,
	Height:     160,
	Weight:     44,
	Income:     "8001-12000元",
	Marriage:   "未婚",
	Education:  "大专",
	Occupation: "其他职业",
	Location:   "东莞东城区",
	Hometown:   "广东清远",
	Horoscope:  "天秤座(09.23-10.22)",
	House:      "已购房",
	Car:        "未买车",
}
var profile2 = model.Profile{
	Name:       "空旷的夜",
	Gender:     "男士",
	Age:        28,
	Height:     175,
	Weight:     0,
	Income:     "8001-12000元",
	Marriage:   "未婚",
	Education:  "大学本科",
	Occupation: "其他职业",
	Location:   "武汉武昌区",
	Hometown:   "湖北孝感",
	Horoscope:  "魔羯座(12.22-01.19)",
	House:      "已购房",
	Car:        "未买车",
}

var profileMap = map[model.Profile]string{
	profile1: "test_fixture/profile.json",
	profile2: "test_fixture/profile2.json",
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
			valueType := values.Field(i).Type()
			if valueType.String() == reflect.Int.String() {
				if value.Int() != resultsValue.Int() {
					t.Errorf("%s %v should be %s:%v, instead of %v", key.Name, field.Name, field.Name, resultsValue, value)
				}
			}

			if value.String() != resultsValue.String() {
				t.Errorf("%s %v should be %s:%v, instead of %v", key.Name, field.Name, field.Name, resultsValue, value)
			}
		}

	}
}
