package model

type ResourceNamesResponse struct {
	Names     []string `json:"names"`
	Namespace string   `json:"namespace"`
	Kind      string   `json:"kind"`
}
