package main

import (
	"fmt"

	"gogeo/gpkg"
	// "gogeo/ops"

	"github.com/slawler/gdal"
)

func main() {

	greeting := fmt.Sprintf("%s\n", "Welcome!")
	fmt.Println(greeting)

	ds := gdal.OpenDataSource(gpkg.TestData, gpkg.UpdateModeFalse)
	defer ds.Destroy()

	dsLayers := gpkg.GetLayerNames(&ds)
	fmt.Println(dsLayers)

	layerName := dsLayers[1]
	featureFields := gpkg.GetFeatureFields(layerName, &ds)
	fmt.Println(featureFields)

	layer := gpkg.GetLayer(layerName, &ds)

	// for i:=1;i<layer.FeatureCount(){
	for i := int64(1); i < 2; i++ {
		feature := layer.Feature(i)
		geometry := feature.Geometry()

		fmt.Println(geometry.Length(), geometry.PointCount(), geometry.X(10))

	}

	// feature :=

	// newGPKG := "data/Test.gpkg"
	// gogeo.NewGPKG(newGPKG)
}
