package base

func ListPtrDemo() {
	fill1 := Fill{Name:"tom", Heigth:184}
	fill2 := Fill{Name:"jesse", Heigth:172}
	list := []Fill{fill1, fill2}
	setFillList(list)

	for _, v := range list {
		println(v.Heigth)
	}
}

type Fill struct {
	Name string
	Heigth int
}

func setFillList(list []Fill) {
	for _, v := range list {
		v.Heigth = 1
	}
}