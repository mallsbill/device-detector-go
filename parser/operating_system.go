package parser

import (
	"strings"
)

var OPERATING_SYSTEMS = map[string]string{
	"AIX": "AIX",
	"AND": "Android",
	"AMG": "AmigaOS",
	"ATV": "Apple TV",
	"ARL": "Arch Linux",
	"BTR": "BackTrack",
	"SBA": "Bada",
	"BEO": "BeOS",
	"BLB": "BlackBerry OS",
	"QNX": "BlackBerry Tablet OS",
	"BMP": "Brew",
	"CES": "CentOS",
	"COS": "Chrome OS",
	"CYN": "CyanogenMod",
	"DEB": "Debian",
	"DFB": "DragonFly",
	"FED": "Fedora",
	"FIR": "Fire OS",
	"FOS": "Firefox OS",
	"BSD": "FreeBSD",
	"GNT": "Gentoo",
	"GTV": "Google TV",
	"HPX": "HP-UX",
	"HAI": "Haiku OS",
	"IRI": "IRIX",
	"INF": "Inferno",
	"KOS": "KaiOS",
	"KNO": "Knoppix",
	"KBT": "Kubuntu",
	"LIN": "GNU/Linux",
	"LBT": "Lubuntu",
	"VLN": "VectorLinux",
	"MAC": "Mac",
	"MAE": "Maemo",
	"MDR": "Mandriva",
	"SMG": "MeeGo",
	"MCD": "MocorDroid",
	"MIN": "Mint",
	"MLD": "MildWild",
	"MOR": "MorphOS",
	"NBS": "NetBSD",
	"MTK": "MTK / Nucleus",
	"WII": "Nintendo",
	"NDS": "Nintendo Mobile",
	"OS2": "OS/2",
	"T64": "OSF1",
	"OBS": "OpenBSD",
	"PSP": "PlayStation Portable",
	"PS3": "PlayStation",
	"RHT": "Red Hat",
	"ROS": "RISC OS",
	"REM": "Remix OS",
	"RZD": "RazoDroiD",
	"SAB": "Sabayon",
	"SSE": "SUSE",
	"SAF": "Sailfish OS",
	"SLW": "Slackware",
	"SOS": "Solaris",
	"SYL": "Syllable",
	"SYM": "Symbian",
	"SYS": "Symbian OS",
	"S40": "Symbian OS Series 40",
	"S60": "Symbian OS Series 60",
	"SY3": "Symbian^3",
	"TDX": "ThreadX",
	"TIZ": "Tizen",
	"UBT": "Ubuntu",
	"WTV": "WebTV",
	"WIN": "Windows",
	"WCE": "Windows CE",
	"WMO": "Windows Mobile",
	"WPH": "Windows Phone",
	"WRT": "Windows RT",
	"XBX": "Xbox",
	"XBT": "Xubuntu",
	"YNS": "YunOs",
	"IOS": "iOS",
	"POS": "palmOS",
	"WOS": "webOS",
	"WIO": "Windows IoT",
	"UNK": "Unknown",
}

var OS_FAMILIES = map[string][]string{
	"Android": []string{
		"AND", "CYN", "REM", "RZD", "MLD", "MCD", "YNS", "FIR",
	},
	"AmigaOS":        []string{"AMG", "MOR"},
	"Apple TV":       []string{"ATV"},
	"BlackBerry":     []string{"BLB", "QNX"},
	"Brew":           []string{"BMP"},
	"BeOS":           []string{"BEO", "HAI"},
	"Chrome OS":      []string{"COS"},
	"Firefox OS":     []string{"FOS", "KOS"},
	"Gaming Console": []string{"WII", "PS3"},
	"Google TV":      []string{"GTV"},
	"IBM":            []string{"OS2"},
	"iOS":            []string{"IOS"},
	"RISC OS":        []string{"ROS"},
	"GNU/Linux": []string{
		"LIN", "ARL", "DEB", "KNO", "MIN", "UBT", "KBT", "XBT", "LBT", "FED",
		"RHT", "VLN", "MDR", "GNT", "SAB", "SLW", "SSE", "CES", "BTR", "SAF",
	},
	"Mac":                   []string{"MAC"},
	"Mobile Gaming Console": []string{"PSP", "NDS", "XBX"},
	"Real-time OS":          []string{"MTK", "TDX"},
	"Other Mobile":          []string{"WOS", "POS", "SBA", "TIZ", "SMG", "MAE"},
	"Symbian":               []string{"SYM", "SYS", "SY3", "S60", "S40"},
	"Unix": []string{
		"SOS", "AIX", "HPX", "BSD", "NBS", "OBS", "DFB", "SYL", "IRI", "T64", "INF",
	},
	"WebTV":          []string{"WTV"},
	"Windows":        []string{"WIN"},
	"Windows IoT":    []string{"WIO"},
	"Windows Mobile": []string{"WPH", "WMO", "WCE", "WRT"},
	"Unknown":        []string{"UNK"},
}

const FIXTURE_FILE string = "oss.yml"

type OperatingSystem struct {
	parser   Parser
	name     string
	short    string
	version  string
	platform string
	family   string
}

func NewOperatingSystem(userAgent string) OperatingSystem {
	parser := NewParser(userAgent, FIXTURE_FILE, "os")

	operatingSystem := OperatingSystem{
		parser:   parser,
		name:     "",
		short:    "UNK",
		version:  "",
		platform: "",
		family:   "",
	}

	return operatingSystem
}

func (o *OperatingSystem) Parse() bool {

	regexes := o.parser.GetRegexes()

	for _, element := range regexes {
		osRegex, _ := element.(map[interface{}]interface{})

		matches := o.parser.MatchUserAgent(osRegex["regex"].(string))

		if matches != nil && len(matches) > 0 {
			o.name = o.parser.BuildByMatch(osRegex["name"].(string), matches)
			o.version = o.parser.BuildVersion(osRegex["version"].(string), matches)
			o.platform = o.parsePlatform()
			break
		}
	}

	if o.name == "" {
		return false
	}

	for osShort, osName := range OPERATING_SYSTEMS {
		if strings.ToLower(o.name) == strings.ToLower(osName) {
			o.name = osName
			o.short = osShort
		}
	}

	o.family = o.GetOsFamily(o.short)
	return true
}

func (o *OperatingSystem) parsePlatform() string {
	if len(o.parser.MatchUserAgent("arm")) > 0 {
		return "ARM"
	} else if len(o.parser.MatchUserAgent("WOW64|x64|win64|amd64|x86_64")) > 0 {
		return "x64"
	} else if len(o.parser.MatchUserAgent("i[0-9]86|i86pc")) > 0 {
		return "x86"
	}

	return ""
}

func (o *OperatingSystem) GetOsFamily(osLabel string) string {

	for family, labels := range OS_FAMILIES {
		for _, label := range labels {
			if label == osLabel {
				return family
			}
		}
	}

	return "Unknown"
}

func (o *OperatingSystem) GetName() string {
	return o.name
}

func (o *OperatingSystem) GetShort() string {
	return o.short
}

func (o *OperatingSystem) GetVersion() string {
	return o.version
}

func (o *OperatingSystem) GetPlatform() string {
	return o.platform
}

func (o *OperatingSystem) GetFamily() string {
	return o.family
}
