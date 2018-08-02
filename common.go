package common

import (
	"sort"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/intdxdt/iter"
)

func SortInts(iter []int) []int {
	sort.Ints(iter)
	return iter
}

//Convert slice of interface to ints
func AsInts(iter []interface{}) []int {
	var ints = make([]int, len(iter))
	for i := range iter {
		ints[i] = iter[i].(int)
	}
	return ints
}

//hull geom
func HullGeom(coordinates geom.Coords) geom.Geometry {
	var g geom.Geometry
	if coordinates.Len() > 2 {
		g = geom.NewPolygon(coordinates)
	} else if coordinates.Len() == 2 {
		g = geom.NewLineString(coordinates)
	} else {
		g = coordinates.Pt(0)
	}
	return g
}

func LinearCoords(wkt string) geom.Coords {
	return geom.NewLineStringFromWKT(wkt).Coordinates
}

func CreateHulls(id *iter.Igen, indices [][]int, coords geom.Coords) []node.Node {
	var poly = pln.New(coords)
	var hulls = make([]node.Node, 0)
	for _, o := range indices {
		hulls = append(hulls, nodeFromPolyline(id, poly, rng.Range(o[0], o[1]), HullGeom))
	}
	return hulls
}

//New Node
func nodeFromPolyline(id *iter.Igen, polyline *pln.Polyline, rng rng.Rng, geomFn geom.GeometryFn) node.Node {
	return node.CreateNode(id, polyline.SubCoordinates(rng), rng, geomFn)
}
