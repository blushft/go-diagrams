package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsIoTContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSIot = &awsIoTContainer{path: "assets/aws/iot"}

func (c *awsIoTContainer) IotAlexaEcho(opts ...attr.Attribute) *node.Node {
	return node.New("iot-alexa-echo", attr.AssetImage("assets/aws/iot/iot-alexa-echo.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotDeviceManagement(opts ...attr.Attribute) *node.Node {
	return node.New("iot-device-management", attr.AssetImage("assets/aws/iot/iot-device-management.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotEvents(opts ...attr.Attribute) *node.Node {
	return node.New("iot-events", attr.AssetImage("assets/aws/iot/iot-events.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotTopic(opts ...attr.Attribute) *node.Node {
	return node.New("iot-topic", attr.AssetImage("assets/aws/iot/iot-topic.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotGreengrassConnector(opts ...attr.Attribute) *node.Node {
	return node.New("iot-greengrass-connector", attr.AssetImage("assets/aws/iot/iot-greengrass-connector.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotHardwareBoard(opts ...attr.Attribute) *node.Node {
	return node.New("iot-hardware-board", attr.AssetImage("assets/aws/iot/iot-hardware-board.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotLambda(opts ...attr.Attribute) *node.Node {
	return node.New("iot-lambda", attr.AssetImage("assets/aws/iot/iot-lambda.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotPolicyEmergency(opts ...attr.Attribute) *node.Node {
	return node.New("iot-policy-emergency", attr.AssetImage("assets/aws/iot/iot-policy-emergency.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotThingsGraph(opts ...attr.Attribute) *node.Node {
	return node.New("iot-things-graph", attr.AssetImage("assets/aws/iot/iot-things-graph.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotCertificate(opts ...attr.Attribute) *node.Node {
	return node.New("iot-certificate", attr.AssetImage("assets/aws/iot/iot-certificate.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotShadow(opts ...attr.Attribute) *node.Node {
	return node.New("iot-shadow", attr.AssetImage("assets/aws/iot/iot-shadow.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotHttp2(opts ...attr.Attribute) *node.Node {
	return node.New("iot-http2", attr.AssetImage("assets/aws/iot/iot-http2.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) InternetOfThings(opts ...attr.Attribute) *node.Node {
	return node.New("internet-of-things", attr.AssetImage("assets/aws/iot/internet-of-things.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotAlexaSkill(opts ...attr.Attribute) *node.Node {
	return node.New("iot-alexa-skill", attr.AssetImage("assets/aws/iot/iot-alexa-skill.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotCamera(opts ...attr.Attribute) *node.Node {
	return node.New("iot-camera", attr.AssetImage("assets/aws/iot/iot-camera.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotCore(opts ...attr.Attribute) *node.Node {
	return node.New("iot-core", attr.AssetImage("assets/aws/iot/iot-core.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotDeviceDefender(opts ...attr.Attribute) *node.Node {
	return node.New("iot-device-defender", attr.AssetImage("assets/aws/iot/iot-device-defender.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotHttp(opts ...attr.Attribute) *node.Node {
	return node.New("iot-http", attr.AssetImage("assets/aws/iot/iot-http.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) Freertos(opts ...attr.Attribute) *node.Node {
	return node.New("freertos", attr.AssetImage("assets/aws/iot/freertos.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) Iot1Click(opts ...attr.Attribute) *node.Node {
	return node.New("iot-1-click", attr.AssetImage("assets/aws/iot/iot-1-click.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotRule(opts ...attr.Attribute) *node.Node {
	return node.New("iot-rule", attr.AssetImage("assets/aws/iot/iot-rule.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotSitewise(opts ...attr.Attribute) *node.Node {
	return node.New("iot-sitewise", attr.AssetImage("assets/aws/iot/iot-sitewise.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotButton(opts ...attr.Attribute) *node.Node {
	return node.New("iot-button", attr.AssetImage("assets/aws/iot/iot-button.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotJobs(opts ...attr.Attribute) *node.Node {
	return node.New("iot-jobs", attr.AssetImage("assets/aws/iot/iot-jobs.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotPolicy(opts ...attr.Attribute) *node.Node {
	return node.New("iot-policy", attr.AssetImage("assets/aws/iot/iot-policy.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotAction(opts ...attr.Attribute) *node.Node {
	return node.New("iot-action", attr.AssetImage("assets/aws/iot/iot-action.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotAnalytics(opts ...attr.Attribute) *node.Node {
	return node.New("iot-analytics", attr.AssetImage("assets/aws/iot/iot-analytics.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotMqtt(opts ...attr.Attribute) *node.Node {
	return node.New("iot-mqtt", attr.AssetImage("assets/aws/iot/iot-mqtt.png"), attr.Shape(attr.None))
}

func (c *awsIoTContainer) IotGreengrass(opts ...attr.Attribute) *node.Node {
	return node.New("iot-greengrass", attr.AssetImage("assets/aws/iot/iot-greengrass.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "aws.iot"
  Register(namespace, "IotAlexaEcho", AWSIot.IotAlexaEcho)
  Register(namespace, "IotDeviceManagement", AWSIot.IotDeviceManagement)
  Register(namespace, "IotEvents", AWSIot.IotEvents)
  Register(namespace, "IotTopic", AWSIot.IotTopic)
  Register(namespace, "IotGreengrassConnector", AWSIot.IotGreengrassConnector)
  Register(namespace, "IotHardwareBoard", AWSIot.IotHardwareBoard)
  Register(namespace, "IotLambda", AWSIot.IotLambda)
  Register(namespace, "IotPolicyEmergency", AWSIot.IotPolicyEmergency)
  Register(namespace, "IotThingsGraph", AWSIot.IotThingsGraph)
  Register(namespace, "IotCertificate", AWSIot.IotCertificate)
  Register(namespace, "IotShadow", AWSIot.IotShadow)
  Register(namespace, "IotHttp2", AWSIot.IotHttp2)
  Register(namespace, "InternetOfThings", AWSIot.InternetOfThings)
  Register(namespace, "IotAlexaSkill", AWSIot.IotAlexaSkill)
  Register(namespace, "IotCamera", AWSIot.IotCamera)
  Register(namespace, "IotCore", AWSIot.IotCore)
  Register(namespace, "IotDeviceDefender", AWSIot.IotDeviceDefender)
  Register(namespace, "IotHttp", AWSIot.IotHttp)
  Register(namespace, "Freertos", AWSIot.Freertos)
  Register(namespace, "Iot1Click", AWSIot.Iot1Click)
  Register(namespace, "IotRule", AWSIot.IotRule)
  Register(namespace, "IotSitewise", AWSIot.IotSitewise)
  Register(namespace, "IotButton", AWSIot.IotButton)
  Register(namespace, "IotJobs", AWSIot.IotJobs)
  Register(namespace, "IotPolicy", AWSIot.IotPolicy)
  Register(namespace, "IotAction", AWSIot.IotAction)
  Register(namespace, "IotAnalytics", AWSIot.IotAnalytics)
  Register(namespace, "IotMqtt", AWSIot.IotMqtt)
  Register(namespace, "IotGreengrass", AWSIot.IotGreengrass)
}
