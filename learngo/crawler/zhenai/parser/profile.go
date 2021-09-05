package parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"brucego.com/learngo/crawler/engine"
	"brucego.com/learngo/crawler/model"
)

// <div data-v-8b1eac0c="" class="purple-btns">
// <div data-v-8b1eac0c="" class="m-btn purple">未婚</div>
// <div data-v-8b1eac0c="" class="m-btn purple">34岁</div>
// <div data-v-8b1eac0c="" class="m-btn purple">魔羯座(12.22-01.19)</div>
// <div data-v-8b1eac0c="" class="m-btn purple">182cm</div>
// <div data-v-8b1eac0c="" class="m-btn purple">79kg</div>
// <div data-v-8b1eac0c="" class="m-btn purple">工作地:大连庄河市</div>
// <div data-v-8b1eac0c="" class="m-btn purple">月收入:5-8千</div>
// <div data-v-8b1eac0c="" class="m-btn purple">财务经理</div>
// <div data-v-8b1eac0c="" class="m-btn purple">大学本科</div>
// </div>
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

var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[^>]*>[^>男女]*(.)士征婚</a>`)
var ageRe = regexp.MustCompile(`<div [^>]+>(\d+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div [^>]+>([^>]+)</div><div [^>]+>\d+岁</div>`)
var heightRe = regexp.MustCompile(`<div [^>]+>(\d+)cm</div>`)
var weightRe = regexp.MustCompile(`<div [^>]+>(\d+)kg</div>`)
var LocationRe = regexp.MustCompile(`<div [^>]+>工作地:([^>]*)</div>`)
var incomeRe = regexp.MustCompile(`<div [^>]+>[^<>]*收入:([^<>]*)</div>`)
var occupationRe = regexp.MustCompile(`<div [^>]+>[^<>]*收入:[^<>]*</div><div [^>]+>([^>]*)</div>`)
var educationRe = regexp.MustCompile(`<div [^>]+>[^<>]*收入:[^>]+</div><div [^>]+>[^><]*</div><div [^>]+>([^>]*)</div>`)
var houseRe = regexp.MustCompile(`<div [^>]+>([^>]*房)</div>`)
var carRe = regexp.MustCompile(`<div [^>]+>([^>]*车)</div>`)
var horoscopeRe = regexp.MustCompile(`<div [^>]+>\d+岁</div><div [^>]+>([^>]*)</div>`)
var hometownRe = regexp.MustCompile(`<div [^>]+>籍贯:([^>]*)</div>`)

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
	log.Printf("Profile %v", results.Items)

	return results

}

func compileProfileField(re *regexp.Regexp, content []byte) string {
	result := re.FindSubmatch(content)
	if result != nil {
		return strings.Trim(string(result[1]), " ")

	}
	return ""
}
