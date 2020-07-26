package main

import (
	"fmt"

	gogeo "gogeo/gpkg"

	"github.com/slawler/gdal"
)

func main() {

	greeting := fmt.Sprintf("%s\n", "Welcome!")
	fmt.Println(greeting)

	ds := gdal.OpenDataSource(gogeo.TestData, gogeo.UpdateModeFalse)
	defer ds.Destroy()

	dsLayers := gogeo.GetLayerNames(&ds)
	fmt.Println(dsLayers)

	layerName := dsLayers[1]
	featureFields := gogeo.GetFeatureFields(layerName, &ds)
	fmt.Println(featureFields)

	// newGPKG := "data/Test.gpkg"
	// gogeo.NewGPKG(newGPKG)
}
