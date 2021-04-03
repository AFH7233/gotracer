package utils

type Octree struct {
	visibleObjects []int
	nodes          [8]*Octree
	box            Box
}

func (octree *Octree) Intersect(ray *Ray, objects []VisibleObject) (int, Ray, float64, bool) {
	if octree.box.Intersect(*ray) {
		for _, node := range octree.nodes {
			node.Intersect(ray, objects)
		}
	}
	var r Ray
	return 1, r, 0.0, true
}
