package xlsxer

// TagName defines key in the struct field's tag to scan
var TagName = "xlsx"

// TagSeparator defines seperator string for multiple xlsx tags in struct fields
var TagSeparator = ","

// FieldSeperator defines how to combine parent struct with child struct
var FieldsCombiner = "."

// Normalizer is a function that takes and returns a string. It is applied to
// struct and header field values before they are compared. It can be used to alter
// names for comparison. For instance, you could allow case insensitive matching
// or convert '-' to '_'.
type Normalizer func(string) string

// normalizeName function initially set to a nop Normalizer.
var normalizeName = DefaultNameNormalizer()

// DefaultNameNormalizer is a nop Normalizer.
func DefaultNameNormalizer() Normalizer { return func(s string) string { return s } }
