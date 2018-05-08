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

//node.Nodes from Rtree boxes
func NodesFromBoxes(iter []rtree.BoxObj) []*node.Node {
	var nodes = make([]*node.Node, len(iter))
	for i, h := range iter {
		nodes[i] = h.(*node.Node)
	}
	return nodes
}

//node.Nodes from Rtree nodes
func NodesFromRtreeNodes(iter []*rtree.Node) []*node.Node {
	var nodes = make([]*node.Node, len(iter))
	for i, h := range iter {
		nodes[i] = h.GetItem().(*node.Node)
	}
	return nodes
}

//hull geom
func HullGeom(coords []*geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0].Clone()
	}
	return g
}

func LinearCoords(wkt string) []*geom.Point{
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func CreateHulls(indxs [][]int, coords []*geom.Point) []*node.Node {
	poly := pln.New(coords)
	hulls := make([]*node.Node, 0)
	for _, o := range indxs {
		hulls = append(hulls, newNodeFromPolyline(poly, rng.NewRange(o[0], o[1]), HullGeom))
	}
	return hulls
}



//New Node
func newNodeFromPolyline(polyline *pln.Polyline, rng *rng.Range, gfn geom.GeometryFn) *node.Node {
	return node.New(polyline.SubCoordinates(rng), rng, gfn)
}


