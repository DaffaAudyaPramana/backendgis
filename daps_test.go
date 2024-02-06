package backendgis

import (
	"fmt"
	"testing"
)

var privatekey = ""
var publickey = ""
var encode = ""
var dbname = "datajson"
var collname = "get"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{100.5918079639361, -0.45832970470556234},
				{100.59180519679398, -0.4584272433564678},
				{100.5918916699963, -0.4584327774636705},
				{100.59189305356739, -0.4583317799957598},
				{100.5918079639361, -0.45832970470556234},
			},
		},
	}
	name := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(name)
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
			{95.32345678901234, 5.567890123456789},
			{95.32355678901234, 5.567990123456789},
		},
	}
	name := Box(mconn, collname, coordinates)
	fmt.Println(name)
}