package devicedetector

import "github.com/mallsbill/device-detector-go/parser"

const (
	UNKNOW = "UNK"
)

type DeviceDetector struct {
	userAgent string
	device    string
	brand     string
	model     string
	os        map[string]string
	client    map[string]string
}

func New(userAgent string) DeviceDetector {
	devicedetector := DeviceDetector{
		userAgent: userAgent,
	}

	return devicedetector
}

func (d *DeviceDetector) GetUserAgent() string {
	return d.userAgent
}

func (d *DeviceDetector) GetDevice() string {
	return d.device
}

func (d *DeviceDetector) GetBrand() string {
	return d.brand
}

func (d *DeviceDetector) GetModel() string {
	return d.model
}

func (d *DeviceDetector) GetOs(attr interface{}) interface{} {

	switch a := attr.(type) {
	case string:
		if os, isset := d.os[a]; isset {
			return os
		}
		return UNKNOW
	default:
		return d.os
	}
}

func (d *DeviceDetector) GetClient(attr interface{}) interface{} {

	switch a := attr.(type) {
	case string:
		if client, isset := d.client[a]; isset {
			return client
		}
		return UNKNOW
	default:
		return d.client
	}
}

func (d *DeviceDetector) Parse() {
	d.parseOs()
}

func (d *DeviceDetector) parseOs() {
	operating_system := parser.NewOperatingSystem(d.userAgent)
	operating_system.Parse()
}

func (d *DeviceDetector) parseClient() {
}

func (d *DeviceDetector) parseDevice() {
}
