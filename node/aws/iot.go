package aws

import "github.com/blushft/go-diagrams/node"

type iotContainer struct {
	path string
	opts []node.Option
}

var Iot = &iotContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/iot",
}

func (c *iotContainer) IotCertificate(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-certificate.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotHttp2(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-http2.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotRule(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-rule.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotTopic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-topic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotAlexaSkill(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-alexa-skill.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotButton(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-button.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotHardwareBoard(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-hardware-board.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotMqtt(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-mqtt.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotPolicy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-policy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotDeviceManagement(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-device-management.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotSitewise(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-sitewise.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) Freertos(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/freertos.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) InternetOfThings(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/internet-of-things.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotCore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-core.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotEvents(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-events.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotGreengrassConnector(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-greengrass-connector.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotGreengrass(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-greengrass.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotPolicyEmergency(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-policy-emergency.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotAction(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-action.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotLambda(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-lambda.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotAlexaEcho(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-alexa-echo.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) Iot1Click(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-1-click.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotAnalytics(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-analytics.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotCamera(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-camera.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotDeviceDefender(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-device-defender.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotHttp(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-http.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotJobs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-jobs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotShadow(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-shadow.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *iotContainer) IotThingsGraph(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/iot/iot-things-graph.png")}, c.opts, opts)
	return node.New(nopts...)
}
