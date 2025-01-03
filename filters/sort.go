package filters

type treeNode struct {
	pixel                 *Pixel
	leftChild, rightChild *treeNode
}

func sortXByZone(zoneAmount int, pg PixelGrid) PixelGrid {
	buckets := make([]*treeNode, zoneAmount)

	for y := 0; y < pg.MaxY; y++ {

		for i := 0; i < zoneAmount; i++ {
			buckets[i] = &treeNode{}
		}

		for x := 0; x < pg.MaxX; x++ {
			buckets[pg.Grid[x][y].Zone].addNode(pg.Grid[x][y])
		}

		sortedSlice := make([][]*Pixel, zoneAmount)
		for i := 0; i < zoneAmount; i++ {
			sortedSlice[i] = buckets[i].toSlice()
		}

		for x := 0; x < pg.MaxX; x++ {
			flag := pg.Grid[x][y].Zone
			pg.Grid[x][y] = sortedSlice[flag][0]
			sortedSlice[flag] = sortedSlice[flag][1:]
		}
	}

	return pg
}

func sortYByZone(zoneAmount int, pg PixelGrid) PixelGrid {
	buckets := make([]*treeNode, zoneAmount)

	for x := 0; x < pg.MaxX; x++ {

		for i := 0; i < zoneAmount; i++ {
			buckets[i] = &treeNode{}
		}

		for y := 0; y < pg.MaxY; y++ {
			buckets[pg.Grid[x][y].Zone].addNode(pg.Grid[x][y])
		}

		sortedSlice := make([][]*Pixel, zoneAmount)
		for i := 0; i < zoneAmount; i++ {
			sortedSlice[i] = buckets[i].toSlice()
		}

		for y := 0; y < pg.MaxY; y++ {
			flag := pg.Grid[x][y].Zone
			pg.Grid[x][y] = sortedSlice[flag][0]
			sortedSlice[flag] = sortedSlice[flag][1:]
		}
	}

	return pg
}

func (node *treeNode) addNode(p *Pixel) {
	if node == nil || node.pixel == nil {
		node.pixel = p
		return
	}

	if p.Lightness < node.pixel.Lightness {
		if node.leftChild == nil {
			node.leftChild = &treeNode{pixel: p}
		} else {
			node.leftChild.addNode(p)
		}
	} else {
		if node.rightChild == nil {
			node.rightChild = &treeNode{pixel: p}
		} else {
			node.rightChild.addNode(p)
		}
	}
}

func (n *treeNode) toSlice() []*Pixel {
	var result []*Pixel
	n.inOrder(&result)
	return result
}

func (node *treeNode) inOrder(result *[]*Pixel) {
	if node == nil || node.pixel == nil {
		return
	}

	if node.leftChild != nil {
		node.leftChild.inOrder(result)
	}

	*result = append(*result, node.pixel)

	if node.rightChild != nil {
		node.rightChild.inOrder(result)
	}
}
