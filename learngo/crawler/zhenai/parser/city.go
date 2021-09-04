package parser

import (
	"regexp"

	"brucego.com/brucego.com/learngo/crawler/engine"
	"brucego.com/brucego.com/learngo/crawler/model"
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

var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[^>]*>[^>]+(男|女)士征婚</a>`)

func ParseCity(content []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	return engine.ParseResult{}

}
