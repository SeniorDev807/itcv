// Code generated by reactGen. DO NOT EDIT.

package react

// H1Props defines the properties for the <h1> element
type H1Props struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (h *H1Props) assign(v *_H1Props) {

	v.ClassName = h.ClassName

	v.DangerouslySetInnerHTML = h.DangerouslySetInnerHTML

	if h.DataSet != nil {
		for dk, dv := range h.DataSet {
			v.o.Set("data-"+dk, dv)
		}
	}

	if h.ID != "" {
		v.ID = h.ID
	}

	if h.Key != "" {
		v.Key = h.Key
	}

	if h.OnChange != nil {
		v.o.Set("onChange", h.OnChange.OnChange)
	}

	if h.OnClick != nil {
		v.o.Set("onClick", h.OnClick.OnClick)
	}

	if h.Ref != nil {
		v.o.Set("ref", h.Ref.Ref)
	}

	v.Role = h.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = h.Style.hack()

}
