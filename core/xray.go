package core

import "github.com/lmmqxyx404/my_core/common"

// Server is an instance of Xray. At any time, there must be at most one Server instance running.
type Server interface {
	common.Runnable
}
