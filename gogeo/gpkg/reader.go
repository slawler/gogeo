package gpkg

import (
	"github.com/slawler/gdal"
)

// GoLayer is a generic container for vector data
type GoLayer struct {
	FilePath     string            `json:"file_path"`
	SpatialRef   string            `json:"spatial_reference"`
	FeatureName  string            `json:"feature_name"`
	Fields       []GoLayerField    `json:"fields"`
	Geometry     gdal.Geometry     `json:"geometry"`
	GeometryType gdal.GeometryType `json:"geometry_type"`
	Options      []string          `json:"creation_options"`
}

// GoLayerField include mappings for vector data in a geopackage
type GoLayerField struct {
	FieldName  string         `json:"field_name"`
	FieldType  gdal.FieldType `json:"field_type"`
	FieldValue interface{}    `json:"field_value"`
}

// GetLayerNames avaliable in open gpkg
func GetLayerNames(ds *gdal.DataSource) []string {
	var layers []string

	for i := 0; i < ds.LayerCount(); i++ {
		layers = append(layers, ds.LayerByIndex(i).Name())
	}
	return layers
}

// GetFeatureFields avaliable in open gpkg
// TODO: Map fieldTypes to names and save as struct for reference
func GetFeatureFields(layerName string, ds *gdal.DataSource) []GoLayerField {
	var features = make([]GoLayerField, 0)

	layer := ds.LayerByName(layerName)
	fieldDef := layer.Definition()

	for i := 0; i < fieldDef.FieldCount(); i++ {
		fieldName := fieldDef.FieldDefinition(i).Name()
		fieldType := fieldDef.FieldDefinition(i).Type()
		features = append(features, GoLayerField{fieldName, fieldType, nil})
	}

	return features
}
