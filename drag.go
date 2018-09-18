package pencere

var globalDragContext *DragContext

type DragContext struct {
	Properties   Properties
	DraggingIcon *Pencere
	IconOffsetX  int
	IconOffsetY  int
}

func NewDragContext(source *Pencere) *DragContext {
	dc := &DragContext{
		Properties: make(map[string]interface{}),
	}

	return dc
}
