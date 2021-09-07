package parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"brucego.com/learngo/crawler/engine"
	"brucego.com/learngo/crawler/model"
)

// {
// 	"data": {
// 	  "age": 30,
// 	  "avatarURL": "https://photo.zastatic.com/images/photo/297633/1190530908/63183738655096038.png",
// 	  "basicInfo": [
// 		"未婚",
// 		"30岁",
// 		"魔羯座(12.22-01.19)",
// 		"165cm",
// 		"52kg",
// 		"工作地:长春朝阳区",
// 		"月收入:5-8千",
// 		"运营管理",
// 		"大专"
// 	  ],
// 	  "detailInfo": [
// 		"汉族",
// 		"籍贯:吉林长春",
// 		"体型:丰满",
// 		"不吸烟",
// 		"不喝酒",
// 		"已买车",
// 		"没有小孩",
// 		"是否想要孩子:想要孩子",
// 		"何时结婚:时机成熟就结婚"
// 	  ],
// 	  "educationString": "大专",
// 	  "genderString": "女士",
// 	  "heightString": "165cm",
// 	  "introduceContent": "喜欢高个子的暖男，不喜欢瘦的男生！非诚勿扰！",
// 	  "marriageString": "未婚",
// 	  "nickname": "少女心",
// 	  "salaryString": "5001-8000元",
// 	  "workCityString": "长春",
// 	  "workProvinceCityString": "长春朝阳区"
// 	},

//   }

var genderRe = regexp.MustCompile(`"genderString":"([^,]*)"`)
var ageRe = regexp.MustCompile(`"age":(\d+)`)
var marriageRe = regexp.MustCompile(`"marriageString":"([^,]*)"`)
var heightRe = regexp.MustCompile(`"heightString":"([^,]*)cm"`)
var weightRe = regexp.MustCompile(`"(\d+)kg"`)
var LocationRe = regexp.MustCompile(`"workProvinceCityString":"([^,]*)"`)
var incomeRe = regexp.MustCompile(`"salaryString":"([^,"]*)"`)
var occupationRe = regexp.MustCompile(`"月收入:[^,"]*","([^,"]*)"[^\]]`)
var educationRe = regexp.MustCompile(`"educationString":"([^,]*)"`)
var houseRe = regexp.MustCompile(`"([^,"]*房)"`)
var carRe = regexp.MustCompile(`"([^,"]*车)"`)
var horoscopeRe = regexp.MustCompile(`"([^",]*座\(\d+.\d+-\d+.\d+\))"`)
var hometownRe = regexp.MustCompile(`"籍贯:([^",]*)"`)

func ParseProfile(content []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	gender := compileProfileField(genderRe, content)
	profile.Gender = gender

	age, err := strconv.Atoi(compileProfileField(ageRe, content))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(compileProfileField(heightRe, content))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(compileProfileField(weightRe, content))
	if err == nil {
		profile.Weight = weight
	}

	income := compileProfileField(incomeRe, content)
	profile.Income = income

	marriage := compileProfileField(marriageRe, content)
	profile.Marriage = marriage

	education := compileProfileField(educationRe, content)
	profile.Education = education
	occupation := compileProfileField(occupationRe, content)
	profile.Occupation = occupation

	// occupation is not filled
	if education == "" {
		profile.Education = occupation
		profile.Occupation = education

	}

	location := compileProfileField(LocationRe, content)
	profile.Location = location

	hometown := compileProfileField(hometownRe, content)
	profile.Hometown = hometown

	horoscope := compileProfileField(horoscopeRe, content)
	profile.Horoscope = horoscope

	house := compileProfileField(houseRe, content)
	profile.House = house

	car := compileProfileField(carRe, content)
	profile.Car = car

	results := engine.ParseResult{}
	results.Items = append(results.Items, profile)
	log.Printf("Profile %+v", results.Items)

	return results

}

func compileProfileField(re *regexp.Regexp, content []byte) string {
	result := re.FindSubmatch(content)
	if result != nil {
		return strings.Trim(string(result[1]), " ")

	}
	return ""
}
