package parser

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/smallfish/simpleyaml"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Parser struct {
	fixtureFile string
	parserName  string
	userAgent   string
	regexList   []interface{}
}

func NewParser(userAgent string, fixtureFile string) Parser {
	parser := Parser{
		userAgent:   userAgent,
		fixtureFile: fixtureFile,
	}

	return parser
}

func (p *Parser) GetRegexes() []interface{} {
	var regexFile = path.Join(getRegexesDirectory(), p.fixtureFile)
	yaml, err := ioutil.ReadFile(regexFile)
	check(err)

	y, err := simpleyaml.NewYaml(yaml)
	check(err)

	if y.IsArray() {
		regexList, err := y.Array()
		check(err)
		p.regexList = regexList
	} else {
		panic("parse result is not an array")
	}

	return p.regexList
}

func getRegexesDirectory() string {
	regexDirAbs, _ := filepath.Abs("regexes")
	regexDir := strings.Replace(regexDirAbs, "test", "", 1)

	return regexDir
}

func (p *Parser) MatchUserAgent(regex string) []string {
	fullregex := "(?i)(?:^|[^A-Z0-9\\-_]|[^A-Z0-9\\-]_|sprd-)(?:" + strings.Replace(regex, "/", "\\/", -1) + ")"
	re := regexp.MustCompile(fullregex)
	matches := re.FindStringSubmatch(p.userAgent)

	return matches
}

func (p *Parser) BuildByMatch(item string, matches []string) string {

	for nb := 1; nb <= 3; nb++ {

		if strings.Contains(item, "$"+strconv.Itoa(nb)) == false {
			continue
		}

		replace := ""

		if len(matches) > nb && matches[nb] != "" {
			replace = matches[nb]
		}

		item = strings.Replace("$"+strconv.Itoa(nb), item, replace, -1)
	}

	return item
}
