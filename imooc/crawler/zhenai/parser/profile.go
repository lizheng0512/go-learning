package parser

import (
	"github.com/lizheng0512/go-learning/imooc/crawler/engine"
	"github.com/lizheng0512/go-learning/imooc/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(.+座)[^<]+</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+cm)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+kg)</div>`)
var gongzuodiRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
var genderRe = regexp.MustCompile(`"genderString":"([^"]+)",`)
var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([未买车][已买车])</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([和家人同住][租房][已购房])</div>`)
var jiguanRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>籍贯:([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	// 名称
	match := nameRe.FindSubmatch(contents)
	if match != nil {
		name := string(match[1])
		profile.Name = name
	}
	// 籍贯
	match = jiguanRe.FindSubmatch(contents)
	if match != nil {
		jiguan := string(match[1])
		profile.Jiguan = jiguan
	}
	// 年龄婚况
	match = ageRe.FindSubmatch(contents)
	if match != nil {
		marriage := string(match[1])
		profile.Marriage = marriage
		age, err := strconv.Atoi(string(match[2]))
		if err == nil {
			profile.Age = age
		}
	}
	// 性别
	match = genderRe.FindSubmatch(contents)
	if match != nil {
		gender := string(match[1])
		profile.Gender = gender
	}
	// 星座
	match = xingzuoRe.FindSubmatch(contents)
	if match != nil {
		marriage := string(match[1])
		profile.Xingzuo = marriage
	}
	// 身高
	match = heightRe.FindSubmatch(contents)
	if match != nil {
		height, err := strconv.Atoi(string(match[1]))
		if err == nil {
			profile.Height = height
		}
	}
	// 体重
	match = weightRe.FindSubmatch(contents)
	if match != nil {
		weight, err := strconv.Atoi(string(match[1]))
		if err == nil {
			profile.Weight = weight
		}
	}
	// 工作地
	match = gongzuodiRe.FindSubmatch(contents)
	if match != nil {
		gongzuodi := string(match[1])
		profile.Gongzuodi = gongzuodi
	}
	// 月收入 工作 学历
	match = incomeRe.FindSubmatch(contents)
	if match != nil {
		income := string(match[1])
		profile.Income = income
		occupation := string(match[2])
		profile.Occupation = occupation
		education := string(match[3])
		profile.Education = education
	}
	// 购车情况
	match = carRe.FindSubmatch(contents)
	if match != nil {
		car := string(match[1])
		profile.Car = car
	}
	// 购房情况
	match = houseRe.FindSubmatch(contents)
	if match != nil {
		house := string(match[1])
		profile.House = house
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
