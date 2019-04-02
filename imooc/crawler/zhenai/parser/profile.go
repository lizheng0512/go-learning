package parser

import (
	"github.com/lizheng0512/go-learning/imooc/crawler/model"
	"regexp"
	"strconv"
)

const ageRe = `<div class="m-btn purple" data-v-bff6f798>(.+)</div><div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`
const xingzuoRe = `<div class="m-btn purple" data-v-bff6f798>(.+座).+</div>`
const heightRe = `<div class="m-btn purple" data-v-bff6f798>([\d]+cm)</div>`
const weightRe = `<div class="m-btn purple" data-v-bff6f798>([\d]+kg)</div>`
const gongzuodiRe = `<div class="m-btn purple" data-v-bff6f798>工作地:([]+)</div>`
const incomeRe = `<div class="m-btn purple" data-v-bff6f798>月收入:(.+)</div><div class="m-btn purple" data-v-bff6f798>(.+)</div><div class="m-btn purple" data-v-bff6f798>(.+)</div>`

func ParseProfile(contents []byte) {
	profile := model.Profile{}
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)

	if match != nil {
		age, err := strconv.Atoi(string(match[1])
		if err == nil {
			profile.Age = age
		}
	}
}