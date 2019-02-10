package client

import (
	"fmt"
	"strings"
)

var AVAILABLE_BROWSERS = map[string]string{
	"36": "360 Phone Browser",
	"3B": "360 Browser",
	"AA": "Avant Browser",
	"AB": "ABrowse",
	"AF": "ANT Fresco",
	"AG": "ANTGalio",
	"AL": "Aloha Browser",
	"AM": "Amaya",
	"AO": "Amigo",
	"AN": "Android Browser",
	"AR": "Arora",
	"AV": "Amiga Voyager",
	"AW": "Amiga Aweb",
	"AT": "Atomic Web Browser",
	"AS": "Avast Secure Browser",
	"BA": "Beaker Browser",
	"BB": "BlackBerry Browser",
	"BD": "Baidu Browser",
	"BS": "Baidu Spark",
	"BE": "Beonex",
	"BJ": "Bunjalloo",
	"BL": "B-Line",
	"BR": "Brave",
	"BK": "BriskBard",
	"BX": "BrowseX",
	"CA": "Camino",
	"CC": "Coc Coc",
	"CD": "Comodo Dragon",
	"C1": "Coast",
	"CX": "Charon",
	"CF": "Chrome Frame",
	"HC": "Headless Chrome",
	"CH": "Chrome",
	"CI": "Chrome Mobile iOS",
	"CK": "Conkeror",
	"CM": "Chrome Mobile",
	"CN": "CoolNovo",
	"CO": "CometBird",
	"CP": "ChromePlus",
	"CR": "Chromium",
	"CY": "Cyberfox",
	"CS": "Cheshire",
	"CU": "Cunaguaro",
	"DB": "dbrowser",
	"DE": "Deepnet Explorer",
	"DF": "Dolphin",
	"DO": "Dorado",
	"DL": "Dooble",
	"DI": "Dillo",
	"EI": "Epic",
	"EL": "Elinks",
	"EB": "Element Browser",
	"EP": "GNOME Web",
	"ES": "Espial TV Browser",
	"FB": "Firebird",
	"FD": "Fluid",
	"FE": "Fennec",
	"FF": "Firefox",
	"FK": "Firefox Focus",
	"FL": "Flock",
	"FM": "Firefox Mobile",
	"FW": "Fireweb",
	"FN": "Fireweb Navigator",
	"GA": "Galeon",
	"GE": "Google Earth",
	"HJ": "HotJava",
	"IA": "Iceape",
	"IB": "IBrowse",
	"IC": "iCab",
	"I2": "iCab Mobile",
	"I1": "Iridium",
	"ID": "IceDragon",
	"IV": "Isivioo",
	"IW": "Iceweasel",
	"IE": "Internet Explorer",
	"IM": "IE Mobile",
	"IR": "Iron",
	"JS": "Jasmine",
	"JI": "Jig Browser",
	"KI": "Kindle Browser",
	"KM": "K-meleon",
	"KO": "Konqueror",
	"KP": "Kapiko",
	"KY": "Kylo",
	"KZ": "Kazehakase",
	"LB": "Liebao",
	"LG": "LG Browser",
	"LI": "Links",
	"LU": "LuaKit",
	"LS": "Lunascape",
	"LX": "Lynx",
	"MB": "MicroB",
	"MC": "NCSA Mosaic",
	"ME": "Mercury",
	"MF": "Mobile Safari",
	"MI": "Midori",
	"MU": "MIUI Browser",
	"MS": "Mobile Silk",
	"MX": "Maxthon",
	"NB": "Nokia Browser",
	"NO": "Nokia OSS Browser",
	"NV": "Nokia Ovi Browser",
	"NE": "NetSurf",
	"NF": "NetFront",
	"NL": "NetFront Life",
	"NP": "NetPositive",
	"NS": "Netscape",
	"NT": "NTENT Browser",
	"OB": "Obigo",
	"OD": "Odyssey Web Browser",
	"OF": "Off By One",
	"OE": "ONE Browser",
	"OI": "Opera Mini",
	"OM": "Opera Mobile",
	"OP": "Opera",
	"ON": "Opera Next",
	"OO": "Opera Touch",
	"OR": "Oregano",
	"OV": "Openwave Mobile Browser",
	"OW": "OmniWeb",
	"OT": "Otter Browser",
	"PL": "Palm Blazer",
	"PM": "Pale Moon",
	"PP": "Oppo Browser",
	"PR": "Palm Pre",
	"PU": "Puffin",
	"PW": "Palm WebPro",
	"PA": "Palmscape",
	"PX": "Phoenix",
	"PO": "Polaris",
	"PT": "Polarity",
	"PS": "Microsoft Edge",
	"QQ": "QQ Browser",
	"QT": "Qutebrowser",
	"QZ": "QupZilla",
	"RK": "Rekonq",
	"RM": "RockMelt",
	"SB": "Samsung Browser",
	"SA": "Sailfish Browser",
	"SC": "SEMC-Browser",
	"SE": "Sogou Explorer",
	"SF": "Safari",
	"SH": "Shiira",
	"SK": "Skyfire",
	"SS": "Seraphic Sraf",
	"SL": "Sleipnir",
	"SM": "SeaMonkey",
	"SN": "Snowshoe",
	"SR": "Sunrise",
	"SP": "SuperBird",
	"ST": "Streamy",
	"SX": "Swiftfox",
	"TB": "Tenta Browser",
	"TZ": "Tizen Browser",
	"TS": "TweakStyle",
	"UC": "UC Browser",
	"VI": "Vivaldi",
	"VB": "Vision Mobile Browser",
	"WE": "WebPositive",
	"WF": "Waterfox",
	"WO": "wOSBrowser",
	"WT": "WeTab Browser",
	"YA": "Yandex Browser",
	"XI": "Xiino",
}

var BROWSER_FAMILIES = map[string][]string{
	"Android Browser":    []string{"AN", "MU"},
	"BlackBerry Browser": []string{"BB"},
	"Baidu":              []string{"BD", "BS"},
	"Amiga":              []string{"AV", "AW"},
	"Chrome":             []string{"CH", "BA", "BR", "CC", "CD", "CM", "CI", "CF", "CN", "CR", "CP", "IR", "RM", "AO", "TS", "VI", "PT", "AS", "TB"},
	"Firefox":            []string{"FF", "FE", "FM", "SX", "FB", "PX", "MB", "EI", "WF", "CU"},
	"Internet Explorer":  []string{"IE", "IM", "PS"},
	"Konqueror":          []string{"KO"},
	"NetFront":           []string{"NF"},
	"NetSurf":            []string{"NE"},
	"Nokia Browser":      []string{"NB", "NO", "NV", "DO"},
	"Opera":              []string{"OP", "OM", "OI", "ON", "OO"},
	"Safari":             []string{"SF", "MF"},
	"Sailfish Browser":   []string{"SA"},
}

type Browser struct {
	client         Client
	name           string
	short_name     string
	version        string
	engine         string
	engine_version string
	family         string
}

func NewBrowser(userAgent string) Browser {
	client := NewClient(userAgent, "client/browsers.yml", "browser")

	browser := Browser{
		client: client,
	}

	return browser
}

func (b *Browser) Parse() bool {

	regexes := b.client.parser.GetRegexes()

	for _, element := range regexes {
		browserRegex, _ := element.(map[interface{}]interface{})

		matches := b.client.parser.MatchUserAgent(browserRegex["regex"].(string))

		if matches != nil && len(matches) > 0 {
			b.name = b.client.parser.BuildByMatch(browserRegex["name"].(string), matches)
			b.version = b.client.parser.BuildVersion(browserRegex["version"].(string), matches)
			break
		}
	}

	if b.name == "" {
		return false
	}

	for browserShort, browserName := range AVAILABLE_BROWSERS {
		if strings.ToLower(b.name) == strings.ToLower(browserName) {
			b.name = browserName
			b.short_name = browserShort
		}
	}

	b.family = b.GetBrowserFamily(b.short_name)

	fmt.Println(b.name)
	fmt.Println(b.short_name)
	fmt.Println(b.family)

	return true
}

func (b *Browser) GetBrowserFamily(browserLabel string) string {

	for family, labels := range BROWSER_FAMILIES {
		for _, label := range labels {
			if label == browserLabel {
				return family
			}
		}
	}

	return "Unknown"
}
