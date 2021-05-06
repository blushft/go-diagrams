package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type iotContainer struct {
	path  string
	attrs []attr.Attribute
}

var Iot = &iotContainer{path: "assets/aws/iot"}

func (c *iotContainer) IotAlexaEcho(opts ...attr.Attribute) *node.Node {
	return node.New("iot-alexa-echo", attr.AssetImage("assets/aws/iot/iot-alexa-echo.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotCamera(opts ...attr.Attribute) *node.Node {
	return node.New("iot-camera", attr.AssetImage("assets/aws/iot/iot-camera.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotHardwareBoard(opts ...attr.Attribute) *node.Node {
	return node.New("iot-hardware-board", attr.AssetImage("assets/aws/iot/iot-hardware-board.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotTopic(opts ...attr.Attribute) *node.Node {
	return node.New("iot-topic", attr.AssetImage("assets/aws/iot/iot-topic.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotGreengrass(opts ...attr.Attribute) *node.Node {
	return node.New("iot-greengrass", attr.AssetImage("assets/aws/iot/iot-greengrass.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotRule(opts ...attr.Attribute) *node.Node {
	return node.New("iot-rule", attr.AssetImage("assets/aws/iot/iot-rule.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("iot-analytics", attr.AssetImage("assets/aws/iot/iot-analytics.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotDeviceDefender(opts ...attr.Attribute) *node.Node {
	return node.New("iot-device-defender", attr.AssetImage("assets/aws/iot/iot-device-defender.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotGreengrassConnector(opts ...attr.Attribute) *node.Node {
	return node.New("iot-greengrass-connector", attr.AssetImage("assets/aws/iot/iot-greengrass-connector.png"), attr.Shape(attr.None))
}

func (c *iotContainer) Freertos(opts ...attr.Attribute) *node.Node {
	return node.New("freertos", attr.AssetImage("assets/aws/iot/freertos.png"), attr.Shape(attr.None))
}

func (c *iotContainer) InternetOfThings(opts ...attr.Attribute) *node.Node {
	return node.New("internet-of-things", attr.AssetImage("assets/aws/iot/internet-of-things.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotEvents(opts ...attr.Attribute) *node.Node {
	return node.New("iot-events", attr.AssetImage("assets/aws/iot/iot-events.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotJobs(opts ...attr.Attribute) *node.Node {
	return node.New("iot-jobs", attr.AssetImage("assets/aws/iot/iot-jobs.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotLambda(opts ...attr.Attribute) *node.Node {
	return node.New("iot-lambda", attr.AssetImage("assets/aws/iot/iot-lambda.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotPolicyEmergency(opts ...attr.Attribute) *node.Node {
	return node.New("iot-policy-emergency", attr.AssetImage("assets/aws/iot/iot-policy-emergency.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotThingsGraph(opts ...attr.Attribute) *node.Node {
	return node.New("iot-things-graph", attr.AssetImage("assets/aws/iot/iot-things-graph.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotCertificate(opts ...attr.Attribute) *node.Node {
	return node.New("iot-certificate", attr.AssetImage("assets/aws/iot/iot-certificate.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotDeviceManagement(opts ...attr.Attribute) *node.Node {
	return node.New("iot-device-management", attr.AssetImage("assets/aws/iot/iot-device-management.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotSitewise(opts ...attr.Attribute) *node.Node {
	return node.New("iot-sitewise", attr.AssetImage("assets/aws/iot/iot-sitewise.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotButton(opts ...attr.Attribute) *node.Node {
	return node.New("iot-button", attr.AssetImage("assets/aws/iot/iot-button.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotHttp(opts ...attr.Attribute) *node.Node {
	return node.New("iot-http", attr.AssetImage("assets/aws/iot/iot-http.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotPolicy(opts ...attr.Attribute) *node.Node {
	return node.New("iot-policy", attr.AssetImage("assets/aws/iot/iot-policy.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotShadow(opts ...attr.Attribute) *node.Node {
	return node.New("iot-shadow", attr.AssetImage("assets/aws/iot/iot-shadow.png"), attr.Shape(attr.None))
}

func (c *iotContainer) Iot1Click(opts ...attr.Attribute) *node.Node {
	return node.New("iot-1-click", attr.AssetImage("assets/aws/iot/iot-1-click.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotAction(opts ...attr.Attribute) *node.Node {
	return node.New("iot-action", attr.AssetImage("assets/aws/iot/iot-action.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotHttp2(opts ...attr.Attribute) *node.Node {
	return node.New("iot-http2", attr.AssetImage("assets/aws/iot/iot-http2.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotMqtt(opts ...attr.Attribute) *node.Node {
	return node.New("iot-mqtt", attr.AssetImage("assets/aws/iot/iot-mqtt.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotAlexaSkill(opts ...attr.Attribute) *node.Node {
	return node.New("iot-alexa-skill", attr.AssetImage("assets/aws/iot/iot-alexa-skill.png"), attr.Shape(attr.None))
}

func (c *iotContainer) IotCore(opts ...attr.Attribute) *node.Node {
	return node.New("iot-core", attr.AssetImage("assets/aws/iot/iot-core.png"), attr.Shape(attr.None))
}
