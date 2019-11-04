package go_shadowsocks

type Cipher struct {
	encodePassword *Password
	decodePassword *Password
}

func (cipher Cipher) encode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.encodePassword[v]
	}
}

func (cipher Cipher) decode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.decodePassword[v]
	}
}

func New(password *Password) *Cipher {
	var decodePassword *Password
	encodePassword := password
	for i, v := range encodePassword {
		decodePassword[v] = byte(i)
	}

	return &Cipher{
		encodePassword,
		decodePassword,
	}
}