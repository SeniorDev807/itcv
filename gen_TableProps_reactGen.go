// Code generated by reactGen. DO NOT EDIT.

package react

// TableProps are the props for a <table> component
type TableProps struct {
	AriaSet
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

func (t *TableProps) assign(_v *_TableProps) {

	if t.AriaSet != nil {
		for dk, dv := range t.AriaSet {
			_v.o.Set("aria-"+dk, dv)
		}
	}

	_v.ClassName = t.ClassName

	_v.DangerouslySetInnerHTML = t.DangerouslySetInnerHTML

	if t.DataSet != nil {
		for dk, dv := range t.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if t.ID != "" {
		_v.ID = t.ID
	}

	if t.Key != "" {
		_v.Key = t.Key
	}

	if t.OnChange != nil {
		_v.o.Set("onChange", t.OnChange.OnChange)
	}

	if t.OnClick != nil {
		_v.o.Set("onClick", t.OnClick.OnClick)
	}

	if t.Ref != nil {
		_v.o.Set("ref", t.Ref.Ref)
	}

	_v.Role = t.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = t.Style.hack()

}