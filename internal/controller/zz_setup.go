/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	segment "github.com/dkb-bank/provider-zscaler-zpa/internal/controller/application/segment"
	server "github.com/dkb-bank/provider-zscaler-zpa/internal/controller/application/server"
	providerconfig "github.com/dkb-bank/provider-zscaler-zpa/internal/controller/providerconfig"
	group "github.com/dkb-bank/provider-zscaler-zpa/internal/controller/segment/group"
	groupserver "github.com/dkb-bank/provider-zscaler-zpa/internal/controller/server/group"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		segment.Setup,
		server.Setup,
		providerconfig.Setup,
		group.Setup,
		groupserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
