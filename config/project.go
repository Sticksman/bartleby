package config

const defaultMetadataDir = ".metadata"
const defaultHeadingDepth = 1
const defaultSeparator = ""
const defaultShowTitle = true
const defaultAutoIncrementSections = true
const defaultShowSectionClass = true
const defaultShowSectionName = true

var defaultTwoSectionClassNames = []string{"chapter", "section"}
var defaultThreeSectionClassNames = []string{"part", "chapter", "section"}

// Project defines a configuration of a long-form text project
type Project struct {
	Name                  string   `json:"name"`
	MetadataDir           string   `json:"metadataDir"`           // A directory that will be ignored. Useful for keeping notes and front matter
	ShowTitle             bool     `json:"showTitle"`             // If true, will show the title first
	ShowHeadingDepth      int      `json:"showHeadingDepth"`      // Show the section headings up to some level of depth. 0 will not show headings.
	Separator             string   `json:"separator"`             // Define a separator string between subsections. Defaults to empty
	SectionClassNames     []string `json:"sectionNames"`          // Define a set of names for sections, (e.g. [part, chapter, section]). These will be auto-generated at the lowest level
	AutoIncrementSections bool     `json:"autoIncrementSections"` // If true automatically labels each section with a heading that is its sectionclass + number (e.g. section 1)
	ShowSectionClass      bool     `json:"showSectionClass"`      // if true, will show the section class name and section number
	ShowSectionName       bool     `json:"showSectionName"`       // If true, will use the block's name as the sectionname. Combines with Section class (e.g. Chatper 1\nThe Fire Starts)
	Structure             []*Block `json:"structure"`             // A set of sections that will build into the final document
}

// NewProjectConfig creates a project config at the path
func NewProjectConfig(name string, path string) error {
	return nil
}

func loadProjectConfig(path string) (*Project, error) {
	return nil, nil
}
