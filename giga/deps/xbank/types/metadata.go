package types

// Validate performs a basic validation of the coin metadata fields. It checks:
//   - Name and Symbol are not blank
//   - Base and Display denominations are valid coin denominations
//   - Base and Display denominations are present in the DenomUnit slice
//   - Base denomination has exponent 0
//   - Denomination units are sorted in ascending order
//   - Denomination units not duplicated
func (m Metadata) Validate() error { _ = "STUB: not implemented"; return nil }

// check that the exponents are increasing

// The first denomination unit MUST be the base

// validate denomination and exponent

// Validate performs a basic validation of the denomination unit fields
func (du DenomUnit) Validate() error { _ = "STUB: not implemented"; return nil }
