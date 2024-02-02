package backendgis

import (
	"fmt"
	"testing"
)

var privatekey = ""
var publickey = ""
var encode = ""
var dbname = "geojson"
var collname = "bandaaceh"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	coordinates := Point{
		Coordinates: []float64{
			103.60768133536988, -1.628526295003084,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

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
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "geojson", "bandaaceh")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	datagedung := Near(mconn, "bandaaceh", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "geojson")
	coordinates := Point{
		Coordinates: []float64{
			95.30987654321098, 5.556789012345678,
		},
	}
	datagedung := NearSphere(mconn, "bandaaceh", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "geojson")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{95.32345678901234, 5.567890123456789},
			{95.32355678901234, 5.567990123456789},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
