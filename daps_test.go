package backendgis

import (
	"fmt"
	"testing"
)

var dbname = "datajson"
var collname = "get"

// func TestGeoIntersects(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", dbname)
// 	coordinates := Polygon{
// 		Coordinates: [][][]float64{
// 			{
// 				{107.54190844499595, -6.873854824661734},
// 				{107.54184261920568, -6.873975211086574},
// 				{107.54201584497002, -6.874061201371745},
// 				{107.54208686753464, -6.8739494139982185},
// 				{107.54190844499595, -6.873854824661734},
// 			},
// 		},
// 	}
// 	name := GeoIntersects(mconn, collname, coordinates)
// 	fmt.Println(name)
// }

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{95.31123456789012, 5.553210987654321},
				{95.31133456789011, 5.553210987654321},
				{95.31133456789011, 5.553310987654321},
				{95.31123456789012, 5.553310987654321},
				{95.31123456789012, 5.553210987654321},
			},
		},
	}
	name := GeoWithin(mconn, collname, coordinates)
	fmt.Println(name)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "geojson", "get")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	name := Near(mconn, "get", coordinates)
	fmt.Println(name)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "geojson")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	name := NearSphere(mconn, "get", coordinates)
	fmt.Println(name)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "geojson")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{107.54165179627012, -6.873347393494967},
			{107.54253039713798, -6.873937966575383},
		},
	}
	name := Box(mconn, collname, coordinates)
	fmt.Println(name)
}
