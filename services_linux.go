package ossecsvc

import "fmt"

const (
	ServiceName = "ossec-agent"
)

func Start() error {
	return fmt.Errorf("Service Start() not implemented on linux")
}

func Stop() error {
	return fmt.Errorf("Service Stop() not implemented on linux")
}
