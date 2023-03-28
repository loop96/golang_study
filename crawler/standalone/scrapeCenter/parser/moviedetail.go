package parser

import (
	"golang_study/crawler/standalone/engines"
	"golang_study/crawler/standalone/model"
	"log"
	"regexp"
	"strconv"
)

var nameReg = regexp.MustCompile(`<a\s+[a-zA-Z0-9\-]+=""\s+class="[a-zA-Z0-9\-]+[^>]*">\s*<[^<]*h2\s+[a-zA-Z0-9\-]+=""\s+class="[a-zA-Z0-9\-]+">(.+)</h2>\s*</a>`)
var tagReg = regexp.MustCompile(`<button\s+[a-zA-Z0-9\-]+=""\s+type="button"\s+class="[a-zA-Z0-9\-\s]+">\s*<span>(.+)</span>\s*</button>`)
var CountryAndTimeMinutesReg = regexp.MustCompile(`<div[0-9a-zA-Z\s-]+=""\s+class="[0-9a-zA-Z\-\s]+">\s*<span[0-9a-zA-Z\s-]+="">(.+)</span>\s*<span[0-9a-zA-Z\s-]+="">(.+)</span>\s*<span[0-9a-zA-Z\s-]+="">(\d+)\s*分钟</span>\s*</div>`)
var DateReg = regexp.MustCompile(`<div[0-9a-zA-Z\s-]+=""\s+class="[0-9a-zA-Z\-\s]+">\s*<span[0-9a-zA-Z\s-]+="">(.+)\s上映</span>\s*</div>`)
var StarsReg = regexp.MustCompile(`<p[0-9a-zA-Z\s-]+=""\s+class="[0-9a-zA-Z\s-]+">\s*(\d\.\d)</p>`)
var IntroduceReg = regexp.MustCompile(`<div[\s-\w]+=""\s+class="\w+"><h3[\w-\s]+="">剧情简介</h3>\s*<p[\w\s-]+="">\s*(.+)\s*</p></div>`)
var CastReg = regexp.MustCompile(`<div\s+class="[\w-]+">\s*<img[\w\s-]+=""\s+src="(https://.+)@[0-9a-z_]+"\s+class="\w+">\s*<p[\s-\w]+=""\s+class="[\s\w-]+"\s+aria-describedby="[\w-]+"\s+tabindex="\d+">(.+)</p>\s*<p[\s\w-]+=""\s+class="[\w\s-]+"\s+aria-describedby="[\w-]+"\s+tabindex="\d">饰：(.+)</p></div>`)

// ParseMovieDetail 解析详情
func ParseMovieDetail(htmlByte []byte) engines.ParserResult {
	nameMatch := nameReg.FindSubmatch(htmlByte)
	tagMatch := tagReg.FindAllSubmatch(htmlByte, -1)
	countryAndTimeMinutesMatch := CountryAndTimeMinutesReg.FindSubmatch(htmlByte)
	dateMatch := DateReg.FindSubmatch(htmlByte)
	starsMatch := StarsReg.FindSubmatch(htmlByte)
	introduceMatch := IntroduceReg.FindSubmatch(htmlByte)
	castMatch := CastReg.FindAllSubmatch(htmlByte, -1)

	name := string(nameMatch[1])

	tag := make([]string, len(tagMatch))
	for _, m := range tagMatch {
		tag = append(tag, string(m[1]))
	}

	timeMinutes, err := strconv.ParseFloat(string(countryAndTimeMinutesMatch[3]), 64)
	if err != nil {
		log.Printf("%s has Err TimeMinutes , skip", name)
		return engines.ParserResult{}
	}

	stars, err := strconv.ParseFloat(string(starsMatch[1]), 64)
	if err != nil {
		log.Printf("%s has Err Stars , skip", name)
		return engines.ParserResult{}
	}

	cast := make([]model.Cast, len(castMatch))
	for _, m := range castMatch {
		cast = append(cast, model.Cast{
			ActorName: string(m[2]),
			PlayName:  string(m[3]),
			PicUrl:    string(m[1]),
		})
	}

	movie := model.Movie{
		Name:        name,
		Tag:         tag,
		Country:     string(countryAndTimeMinutesMatch[1]),
		TimeMinutes: timeMinutes,
		Date: func() string {
			if len(dateMatch) <= 1 {
				return ""
			}
			return string(dateMatch[1])
		}(),
		Stars: stars,
		Introduce: func() string {
			if len(introduceMatch) <= 1 {
				return ""
			}
			return string(introduceMatch[1])
		}(),
		Casts: cast,
	}
	return engines.ParserResult{
		Items: []interface{}{movie},
	}
}
