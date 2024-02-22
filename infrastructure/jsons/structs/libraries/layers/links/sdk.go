package links

// Link represents a link
type Link struct {
	Origin   Origin    `json:"origin"`
	Elements []Element `json:"elements"`
}

// Element represents an element
type Element struct {
	Layer     string     `json:"layer"`
	Condition *Condition `json:"condition"`
}

// Condition represents a condition
type Condition struct {
	Resource ConditionResource `json:"resource"`
	Next     *ConditionValue   `json:"next"`
}

// ConditionValue represents a condition value
type ConditionValue struct {
	Resource  *ConditionResource `json:"resource"`
	Condition *Condition         `json:"condition"`
}

// ConditionResource represents a condition resource
type ConditionResource struct {
	Code            uint `json:"code"`
	IsRaisedInLayer bool `json:"is_raised_in_layer"`
}

// Origin represents an origin
type Origin struct {
	Resource OriginResource `json:"resource"`
	Operator Operator       `json:"operator"`
	Next     OriginValue    `json:"next"`
}

// OriginValue represents an origin value
type OriginValue struct {
	Resource *OriginResource `json:"resource"`
	Origin   *Origin         `json:"origin"`
}

// OriginResource represents an origin resource
type OriginResource struct {
	Layer       string `json:"layer"`
	IsMandatory bool   `json:"is_mandatory"`
}

// Operator represents an operator
type Operator struct {
	IsAnd bool `json:"is_and"`
	IsOr  bool `json:"is_or"`
	IsXor bool `json:"is_xor"`
}
