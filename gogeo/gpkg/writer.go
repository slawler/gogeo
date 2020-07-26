package gpkg

import (
	"fmt"
	"io/ioutil"

	"github.com/slawler/gdal"
)

const (
	// TempalteGPKG is an empty, gpkg formatted sqllite db
	TempalteGPKG string = "data/empty.gpkg"
	// TestData available in the data directory
	TestData string = "data/example.gpkg"
	// UpdateModeFalse option for opening gpkg
	UpdateModeFalse int = 0
	// UpdateModeTrue option for opening gpkg
	UpdateModeTrue int = 1
)

// NewGPKG Creates an empty GeoPackage
func NewGPKG(newEmptyGPKG string) {
	input, err := ioutil.ReadFile(TempalteGPKG)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(newEmptyGPKG, input, 0644)
	if err != nil {
		fmt.Println("Error creating", newEmptyGPKG)
		fmt.Println(err)
		return
	}
}

// AddVectorLayer creates a new feature dataset in a geopackage
func (v GoLayer) AddVectorLayer(ds *gdal.DataSource) {
	spRef := gdal.CreateSpatialReference(v.SpatialRef)
	newLayer := ds.CreateLayer(v.FeatureName, spRef, v.GeometryType, v.Options)

	// Add Fields
	for _, field := range v.Fields {
		fiedlDef := gdal.CreateFieldDefinition(field.FieldName, field.FieldType)
		newLayer.CreateField(fiedlDef, true)
	}
}

// AddFeature creates a new feature dataset in a geopackage
func (v GoLayer) AddFeature(layerName string, ds *gdal.DataSource) {

	layer := ds.LayerByName(layerName)
	featureDef := layer.Definition()
	newFeature := featureDef.Create()

	for _, field := range v.Fields {

		fieldNameID := featureDef.FieldIndex(field.FieldName)

		switch field.FieldType {

		case gdal.FT_String:
			value := field.FieldValue.(string)
			newFeature.SetFieldString(fieldNameID, value)

		case gdal.FT_Integer64:
			value := field.FieldValue.(int64)
			newFeature.SetFieldInteger64(fieldNameID, value)

		case gdal.FT_Real:
			value := field.FieldValue.(float64)
			newFeature.SetFieldFloat64(fieldNameID, value)
		}

	}
	newFeature.SetGeometryDirectly(v.Geometry)
	layer.Create(newFeature)
	newFeature.Destroy()

}
