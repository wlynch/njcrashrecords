package njcrash

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	Zip   string `json:"zip,omitempty"`
}
