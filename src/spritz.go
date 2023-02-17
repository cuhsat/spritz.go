// spritz.go
package spritz

const N = 256

var a, i, j, k, w, z byte
var sbox [N]byte

func Encrypt(message, key []byte) {
	keySetup(key)
	for v := range message {
		message[v] += drip()
	}
}

func Decrypt(message, key []byte) {
	keySetup(key)
	for v := range message {
		message[v] -= drip()
	}
}

func Hash(message, digest []byte) {
	r := []byte{byte(len(digest))}
	initializeState()
	absorb(message)
	absorbStop()
	absorb(r)
	squeeze(digest)
}

func keySetup(k []byte) {
	initializeState()
	absorb(k)
}

func initializeState() {
	a, i, j, k, z = 0, 0, 0, 0, 0
	w = 1
	for v := 0; v < N; v++ {
		sbox[v] = byte(v)
	}
}

func absorb(i []byte) {
	for _, v := range i {
		absorbByte(v)
	}
}

func absorbByte(b byte) {
	absorbNibble(byte(b & 0x0F))
	absorbNibble(byte(b >> 4))
}

func absorbStop() {
	if a == (N / 2) {
		shuffle()
	}
	a++
}

func absorbNibble(x byte) {
	if a == (N / 2) {
		shuffle()
	}
	var t byte = (N / 2) + x
	sbox[a], sbox[t] = sbox[t], sbox[a]
	a++
}

func squeeze(r []byte) {
	if a > 0 {
		shuffle()
	}
	for v := 0; v < min(len(r), int(N)); v++ {
		r[v] = drip()
	}
}

func drip() byte {
	if a > 0 {
		shuffle()
	}
	update()
	return output()
}

func shuffle() {
	whip()
	crush()
	whip()
	crush()
	whip()
	a = 0
}

func output() byte {
	z = sbox[j+sbox[i+sbox[(z+k)]]]
	return z
}

func whip() {
	for v := 0; v < N*2; v++ {
		update()
	}
	w += 2
}

func crush() {
	for v := byte(0); v < (N / 2); v++ {
		var t byte = N - 1 - v
		if sbox[v] > sbox[t] {
			sbox[v], sbox[t] = sbox[t], sbox[v]
		}
	}
}

func update() {
	i += w
	j = k + sbox[j+sbox[i]]
	k = i + k + sbox[j]
	sbox[i], sbox[j] = sbox[j], sbox[i]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
