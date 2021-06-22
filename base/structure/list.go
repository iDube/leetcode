package structure

// list

type List struct {
	heard, tail *ListItem
}

type ListItem struct {
	item interface{}
	next *ListItem
}

func (this *List) Add(i *interface{}) {
	if this.heard == nil {
		this.heard = &ListItem{
			item: i,
			next: nil,
		}
	} else {
		this.tail.next = &ListItem{
			item: i,
			next: nil,
		}
	}
}
