package key

type Mint struct {
	MINT_KEY_PAIR       Key
	MINT_PUBLIC_ADDRESS string
}

func Init() *Mint {
	var (
		MINT_KEY_PAIR       = *GenKeyPairT()
		MINT_PUBLIC_ADDRESS = MINT_KEY_PAIR.Public
	)

	return &Mint{
		MINT_KEY_PAIR:       MINT_KEY_PAIR,
		MINT_PUBLIC_ADDRESS: MINT_PUBLIC_ADDRESS,
	}
}

func (m *Mint) GET_MINT_KEY_PAIR() Key {
	return m.MINT_KEY_PAIR
}

func (m *Mint) GET_MINT_PUBLIC_ADDRESS() string {
	return m.MINT_PUBLIC_ADDRESS
}
