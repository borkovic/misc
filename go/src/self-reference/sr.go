//  https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

package main

import "fmt"

//**************************************************************
type Foo struct {
    verbosity int
}

type option func(f *Foo) option


//**************************************************************
func (f *Foo) Option(opts ...option) (previous option) {
    for _, opt := range opts {
        previous = opt(f)
    }
    return previous
}

//**************************************************************
// Verbosity sets Foo's verbosity level to v.
func Verbosity(v int) option {
    return func(f *Foo) option {
        previous := f.verbosity
        fmt.Println("Setting verbosity to", v, ", Old verbosity", f.verbosity)
        f.verbosity = v
        return Verbosity(previous)
    }
}


//**************************************************************
func main() {
    var f Foo 
    f.verbosity = 14

    prevVerbosity := f.Option(Verbosity(3))
    //DoSomeDebugging()
    f.Option(prevVerbosity)
}

