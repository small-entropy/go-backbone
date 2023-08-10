package meta

type Meta struct {
	Limit  int64 `json:"Limit,omitempty"`
	Skip   int64 `json:"Skip,omitempty"`
	Count  int   `json:"Count,omitempty"`
	Total  int64 `json:"Total,omitempty"`
	Filter map[string]interface{}
}
