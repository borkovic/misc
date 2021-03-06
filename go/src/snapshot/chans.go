package snapshot

type HorizBidirChan chan Data
type HorizInChan <-chan Data
type HorizOutChan chan<- Data

type HorizChanPair struct {
	in   HorizInChan
	out  HorizOutChan
	from ProcIdx
	to   ProcIdx
}

type VertBidirChan chan Data
type VertInChan <-chan Data
type VertOutChan chan<- Data
type VertChanPair struct {
	in  VertInChan
	out VertOutChan
}

type NeighborChans []HorizChanPair

func HorizBidir2InChan(c HorizBidirChan) HorizInChan {
	var c2 chan Data = (chan Data)(c)
	var c3 <-chan Data = (<-chan Data)(c2)
	var c4 HorizInChan = HorizInChan(c3)
	return c4
}

func HorizBidir2OutChan(c HorizBidirChan) HorizOutChan {
	return (HorizOutChan)((chan<- Data)((chan Data)(c)))
}

func VertBidir2InChan(c VertBidirChan) VertInChan {
	return (VertInChan)((<-chan Data)((chan Data)(c)))
}

func VertBidir2OutChan(c VertBidirChan) VertOutChan {
	return (VertOutChan)((chan<- Data)((chan Data)(c)))
}
