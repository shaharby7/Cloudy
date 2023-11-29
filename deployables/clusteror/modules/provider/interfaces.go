package provider

type Provider interface {
	Identify() *ProviderCommonData
	CreateMachine() error
	DestroyMachine() error
}

type ProviderCommonData struct {
	Code ProviderCode `json:"code"`
}
