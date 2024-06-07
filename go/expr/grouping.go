package expr

type Grouping struct{}

func NewGrouping() Grouping {
	return Grouping{}
}
func (g *Grouping) Accept() {

}
