package xsd

import (
	"encoding/xml"
	"path/filepath"
)

type Include struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2001/XMLSchema include"`
	Namespace      string   `xml:"namespace,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	IncludedSchema *Schema  `xml:"-"`
}

func (i *Include) load(ws *Workspace, baseDir string) (err error) {
	if i.SchemaLocation != "" {
		i.IncludedSchema, err =
			ws.loadXsd(filepath.Join(baseDir, i.SchemaLocation), false)
	}
	return
}
