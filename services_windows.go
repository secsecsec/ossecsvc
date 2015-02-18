package ossecsvc

import (
	"fmt"

	"code.google.com/p/winsvc/mgr"
	"code.google.com/p/winsvc/svc"
)

const (
	ServiceName = "OSSEC HIDS"
)

func Start() error {
	ctrl, err := mgr.Connect()
	if err != nil {
		ctrl.Disconnect()
		return err
	}

	defer ctrl.Disconnect()

	service, err := ctrl.OpenService(ServiceName)
	if err != nil {
		return fmt.Errorf("error creating handle to %s: %v", ServiceName, err)
	}
	defer service.Close()

	if err := service.Start(nil); err != nil {
		return fmt.Errorf("error starting service %s: %v", ServiceName, err)
	}

	return nil
}

func Stop() error {
	ctrl, err := mgr.Connect()
	if err != nil {
		ctrl.Disconnect()
		return err
	}

	defer ctrl.Disconnect()

	service, err := ctrl.OpenService(ServiceName)
	if err != nil {
		return fmt.Errorf("error creating handle to %s: %v", ServiceName, err)
	}
	defer service.Close()

	if _, err := service.Control(svc.Stop); err != nil {
		return fmt.Errorf("error stopping service %s: %v", ServiceName, err)
	}

	return nil
}
