package config

// Block refers to a subtree in a long-form text structure.
type Block struct {
	Name              string   `json:"name"`
	CustomHeading     string   `json:"customHeading"`     // Will override section class and name (e.g. prologue)
	HideSectionClass  bool     `json:"hideSectionClass"`  // Will hide section class
	HideSectionName   bool     `json:"hideSectionName"`   // Will hide section name
	HideSubHeadings   bool     `json:"hideSubheadings"`   // Will hide subheadings for a particular section
	HideSubSeparators bool     `json:"hideSubSeparators"` // Will hide subseparators for a particular section
	Substructure      []*Block `json:"substructure"`
}
