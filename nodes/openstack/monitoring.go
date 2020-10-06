package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type monitoringContainer struct {
	path  string
	attrs []attr.Attribute
}

var Monitoring = &monitoringContainer{path: "assets/openstack/monitoring"}

func (c *monitoringContainer) Monasca(opts ...attr.Attribute) *node.Node {
	return node.New("monasca", attr.AssetImage("assets/openstack/monitoring/monasca.png"), attr.Shape(attr.None))
}

func (c *monitoringContainer) Telemetry(opts ...attr.Attribute) *node.Node {
	return node.New("telemetry", attr.AssetImage("assets/openstack/monitoring/telemetry.png"), attr.Shape(attr.None))
}
