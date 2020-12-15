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
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-alexa-echo.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-alexa-echo", opts...)
}

func (c *awsIoTContainer) IotDeviceManagement(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-device-management.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-device-management", opts...)
}

func (c *awsIoTContainer) IotEvents(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-events.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-events", opts...)
}

func (c *awsIoTContainer) IotTopic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-topic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-topic", opts...)
}

func (c *awsIoTContainer) IotGreengrassConnector(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-greengrass-connector.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-greengrass-connector", opts...)
}

func (c *awsIoTContainer) IotHardwareBoard(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-hardware-board.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-hardware-board", opts...)
}

func (c *awsIoTContainer) IotLambda(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-lambda.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-lambda", opts...)
}

func (c *awsIoTContainer) IotPolicyEmergency(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-policy-emergency.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-policy-emergency", opts...)
}

func (c *awsIoTContainer) IotThingsGraph(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-things-graph.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-things-graph", opts...)
}

func (c *awsIoTContainer) IotCertificate(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-certificate.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-certificate", opts...)
}

func (c *awsIoTContainer) IotShadow(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-shadow.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-shadow", opts...)
}

func (c *awsIoTContainer) IotHttp2(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-http2.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-http2", opts...)
}

func (c *awsIoTContainer) InternetOfThings(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/internet-of-things.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("internet-of-things", opts...)
}

func (c *awsIoTContainer) IotAlexaSkill(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-alexa-skill.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-alexa-skill", opts...)
}

func (c *awsIoTContainer) IotCamera(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-camera.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-camera", opts...)
}

func (c *awsIoTContainer) IotCore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-core.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-core", opts...)
}

func (c *awsIoTContainer) IotDeviceDefender(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-device-defender.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-device-defender", opts...)
}

func (c *awsIoTContainer) IotHttp(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-http.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-http", opts...)
}

func (c *awsIoTContainer) Freertos(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/freertos.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("freertos", opts...)
}

func (c *awsIoTContainer) Iot1Click(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-1-click.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-1-click", opts...)
}

func (c *awsIoTContainer) IotRule(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-rule.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-rule", opts...)
}

func (c *awsIoTContainer) IotSitewise(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-sitewise.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-sitewise", opts...)
}

func (c *awsIoTContainer) IotButton(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-button.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-button", opts...)
}

func (c *awsIoTContainer) IotJobs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-jobs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-jobs", opts...)
}

func (c *awsIoTContainer) IotPolicy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-policy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-policy", opts...)
}

func (c *awsIoTContainer) IotAction(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-action.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-action", opts...)
}

func (c *awsIoTContainer) IotAnalytics(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-analytics.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-analytics", opts...)
}

func (c *awsIoTContainer) IotMqtt(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-mqtt.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-mqtt", opts...)
}

func (c *awsIoTContainer) IotGreengrass(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/iot/iot-greengrass.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-greengrass", opts...)
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
