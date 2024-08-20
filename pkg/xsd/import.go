package xsd

import (
	"encoding/xml"
)

type Import struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2001/XMLSchema import"`
	Namespace      string   `xml:"namespace,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	ImportedSchema *Schema  `xml:"-"`
}

func (i *Import) load(ws *Workspace, baseDir string) (err error) {

	if i.SchemaLocation != "" {
		i.ImportedSchema, err =
			ws.loadXsd(i.SchemaLocation, true)
	}
	return
}
