package main

/*
// The minimum size of stack used by Go code
_StackMin = 2048
*/

/*
// Max stack size is 1 GB on 64-bit, 250 MB on 32-bit.
// Using decimal instead of binary GB and MB because
// they look nicer in the stack overflow failure message.
if goarch.PtrSize == 8 {
	maxstacksize = 1000000000
} else {
	maxstacksize = 250000000
}
*/

//go:noinline
func allocateBuffer(_ [1500]byte) {
}

func main() {
	var buffer [1500]byte
	allocateBuffer(buffer)
}
