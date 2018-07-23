package common

import (
	"sort"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
)

func SortInts(iter []int) []int {
	sort.Ints(iter)
	return iter
}

//Convert slice of interface to ints
func AsInts(iter []interface{}) []int {
	var ints = make([]int, len(iter))
	for i, o := range iter {
		ints[i] = o.(int)
	}
	return ints
}

//node.Nodes from Rtree nodes
func NodesFromObjects(iter []*rtree.Obj) []*node.Node {
	var nodes = make([]*node.Node, 0, len(iter))
	for i := range iter {
		nodes = append(nodes, iter[i].Object.(*node.Node))
	}
	return nodes
}

//hull geom
func HullGeom(coordinates []geom.Point) geom.Geometry {
	var g geom.Geometry
	if len(coordinates) > 2 {
		g = geom.NewPolygon(coordinates)
	} else if len(coordinates) == 2 {
		g = geom.NewLineString(coordinates)
	} else {
		var pt = coordinates[0]
		g = &pt
	}
	return g
}

func LinearCoords(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func CreateHulls(indxs [][]int, coords []geom.Point) []*node.Node {
	var poly    = pln.New(coords)
	var hulls   = make([]*node.Node, 0)
	for _, o := range indxs {
		hulls = append(hulls, newNodeFromPolyline(poly, rng.Range(o[0], o[1]), HullGeom))
	}
	return hulls
}

//New Node
func newNodeFromPolyline(polyline *pln.Polyline, rng rng.Rng, geomFn geom.GeometryFn) *node.Node {
	return node.New(polyline.SubCoordinates(rng), rng, geomFn)
}
