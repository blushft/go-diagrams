package aws

import "github.com/blushft/go-diagrams/diagram"

type iotContainer struct {
	path string
	opts []diagram.NodeOption
}

var Iot = &iotContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/iot",
}

func (c *iotContainer) Freertos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/freertos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotLambda(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-lambda.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotTopic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-topic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotShadow(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-shadow.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotAction(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-action.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotAlexaSkill(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-alexa-skill.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotEvents(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-events.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotHardwareBoard(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-hardware-board.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotJobs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-jobs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotPolicy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-policy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotPolicyEmergency(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-policy-emergency.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotSitewise(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-sitewise.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) InternetOfThings(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/internet-of-things.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotAlexaEcho(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-alexa-echo.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotCore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-core.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotDeviceManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-device-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotGreengrassConnector(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-greengrass-connector.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotHttp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-http.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotAnalytics(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-analytics.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotCamera(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-camera.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotHttp2(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-http2.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotThingsGraph(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-things-graph.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotDeviceDefender(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-device-defender.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) Iot1Click(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-1-click.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotButton(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-button.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotCertificate(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-certificate.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotGreengrass(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-greengrass.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotMqtt(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-mqtt.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *iotContainer) IotRule(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/iot/iot-rule.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
