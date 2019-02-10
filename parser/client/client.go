package client

import "github.com/mallsbill/device-detector-go/parser"

type Client struct {
	parser  parser.Parser
	name    string
	version string
}

func NewClient(userAgent string, fixtureFile string, parserName string) Client {
	parser := parser.NewParser(userAgent, fixtureFile, parserName)

	client := Client{
		parser:  parser,
		name:    "",
		version: "",
	}

	return client
}

func (c *Client) Parse() bool {

	regexes := c.parser.GetRegexes()

	if c.parser.PreMatchOverall() != nil {
		for _, element := range regexes {
			clientRegex, _ := element.(map[interface{}]interface{})

			matches := c.parser.MatchUserAgent(clientRegex["regex"].(string))

			if matches != nil && len(matches) > 0 {
				c.name = c.parser.BuildByMatch(clientRegex["name"].(string), matches)
				c.version = c.parser.BuildVersion(clientRegex["version"].(string), matches)

				break
			}
		}
	}

	if c.name == "" {
		return false
	}

	return true
}
