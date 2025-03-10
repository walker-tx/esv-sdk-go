// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PassageResponse struct {
	// The passage reference that was requested
	Query *string `json:"query,omitempty"`
	// The canonical version of the passage reference
	Canonical *string `json:"canonical,omitempty"`
	// Array of parsed passage references
	Parsed      [][]int64     `json:"parsed,omitempty"`
	PassageMeta []PassageMeta `json:"passage_meta,omitempty"`
	// Array of formatted passage text
	Passages []string `json:"passages,omitempty"`
}

func (o *PassageResponse) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *PassageResponse) GetCanonical() *string {
	if o == nil {
		return nil
	}
	return o.Canonical
}

func (o *PassageResponse) GetParsed() [][]int64 {
	if o == nil {
		return nil
	}
	return o.Parsed
}

func (o *PassageResponse) GetPassageMeta() []PassageMeta {
	if o == nil {
		return nil
	}
	return o.PassageMeta
}

func (o *PassageResponse) GetPassages() []string {
	if o == nil {
		return nil
	}
	return o.Passages
}
