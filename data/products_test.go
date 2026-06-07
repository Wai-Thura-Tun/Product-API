package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 3,
		SKU:   "ahc-sur-rufe",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
