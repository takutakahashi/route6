package iface

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	v1 "github.com/takutakahashi/route6/api/v1"
	"github.com/takutakahashi/route6/pkg/util"
	"github.com/vishvananda/netlink"
	corev1 "k8s.io/api/core/v1"
)

func Apply(node *corev1.Node, i *v1.Interface) error {
	ownDevice := v1.Device{}
	for _, device := range i.Spec.Devices {
		if util.MatchSelector(node.Labels, device.NodeSelector) {
			ownDevice = device
			break
		}
	}
	if ownDevice.Name == "" {
		return errors.New("no device found on this node")
	}
	dev, err := netlink.LinkByName("lo")
	if err != nil {
		return errors.Wrap(err, "failed to find device")
	}
	logrus.Info(dev)
	return nil
}
