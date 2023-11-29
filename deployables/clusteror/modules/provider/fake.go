package provider

import "errors"

type Fake struct {
	Code ProviderCode
}

func (fake *Fake) CreateMachine() error {
	return errors.New("not implemented yet")
}

func (fake *Fake) DestroyMachine() error {
	return errors.New("not implemented yet")
}

func (fake *Fake) Identify() *ProviderCommonData {
	return &ProviderCommonData{Code: fake.Code}
}
