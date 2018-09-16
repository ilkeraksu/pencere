package pencere

import (
	"sync/atomic"
)

var focusedPencere *Pencere
var isRenderDirty bool
var handleId uint64

var Bus = NewEventBus()

func getHandleId() uint64 {
	return atomic.AddUint64(&handleId, 1)
}
