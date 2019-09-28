package snapshot

/*************************************************************
 * This process is root.
 * It sends its neg. value to all children.
 * Collects sums from the children
*************************************************************/
func (proc *Proc) runRoot(neighbors NeighborChans) Data {
	for _, n := range neighbors {
		n.out <- (-proc.m_MyVal)
	}
	sum := proc.m_MyVal
	for i := 0; i < NumNeighbors; i++ {
		select {
		case x := <-(*neighbors)[0].in:
			sum += x
		case x := <-(*neighbors)[1].in:
			sum += x
		case x := <-(*neighbors)[2].in:
			sum += x
		}
	}
	return sum
}

/*************************************************************
 * Child
 * 1. receive value from parent
 * 2. send its neg. value to others (children and siblings)
 * 3. receive values from children(positive) and siblings(negative)
 * 4. sums values from children (positive)`
 * 5. send sum to parent
*************************************************************/
func (proc *Proc) runChild(neighbors NeighborChans) {
	var parIdx int = -1

	// 1. read from parent in tree
	select {
	case <-neighbors[0].in:
		parIdx = 0
	case <-neighbors[1].in:
		parIdx = 1
	case <-neighbors[2].in:
		parIdx = 2
	}

	// 2. send to all but parent (children and siblings)
	for i := 0; i < NumNeighbors; i++ {
		if i == parIdx {
			continue
		}
		neighbors[i].out <- (-proc.m_MyVal)
	}

	// 3. receive from children and siblings
	// from children positive, from siblings negative
	// 4. sum from children
	sum := proc.m_MyVal
	for i := 0; i < NumNeighbors; i++ {
		if i == parIdx {
			continue
		}
		v := <-neighbors[i].in
		if v > 0 { // this is child, siblings send negative
			sum += v
		}
	}

	// 5. send sum to parent
	neighbors[parIdx].out <- sum
}

/*************************************************************
 * This process recieves value from top.
 * If negative, that's root, otherwise it's child
*************************************************************/
func (proc *Proc) run(topChan IoChan, neighbors NeighborChans) {
	proc.m_MyVal = <-topChan.in
	if proc.m_MyVal < 0 { // master
		proc.m_MyVal = -proc.m_MyVal
		sum := proc.runRoot(neighbors)
		topChan.out <- sum
	} else { // slave
		proc.runChild(neighbors)
	}
}
