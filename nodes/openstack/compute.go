package openstack

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type computeContainer struct {
	path  string
	attrs []attr.Attribute
}

var Compute = &computeContainer{path: "assets/openstack/compute"}

func (c *computeContainer) Nova(opts ...attr.Attribute) *node.Node {
	return node.New("nova", attr.AssetImage("assets/openstack/compute/nova.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Qinling(opts ...attr.Attribute) *node.Node {
	return node.New("qinling", attr.AssetImage("assets/openstack/compute/qinling.png"), attr.Shape(attr.None))
}

func (c *computeContainer) Zun(opts ...attr.Attribute) *node.Node {
	return node.New("zun", attr.AssetImage("assets/openstack/compute/zun.png"), attr.Shape(attr.None))
}
