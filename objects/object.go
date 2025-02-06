package objects

type Object interface {
	Draw()
}

type ObjectGroup = []Object

func NewObjectGroup() *ObjectGroup {
	objGr := make(ObjectGroup, 0)

	registerObjects(&objGr)

	return &objGr
}

func registerObjects(objGr *ObjectGroup) {
	*objGr = append(*objGr,
		NewMaze(70, 30, 50, 50, 15, 5433),
	)
}
