package parser

type Parser struct {
	fixtureFile string
	parserName  string
	userAgent   string
	regexList   map[string]string
}

func NewParser(userAgent string) Parser {
	parser := Parser{
		userAgent: userAgent,
	}

	return parser
}
