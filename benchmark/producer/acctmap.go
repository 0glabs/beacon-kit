package producer

type AccountMap struct {
	total      uint32
	accounts   []Account
	faucetAcct *Account
}

func NewAccountMap(total uint32, faucetPrivateKey string) *AccountMap {
	am := &AccountMap{
		total:      total,
		accounts:   make([]Account, 0, total),
		faucetAcct: CreateFaucetAccount(faucetPrivateKey),
	}

	for i := uint32(0); i < total; i++ {
		am.accounts = append(am.accounts, NewAccount(i))
	}

	return am
}

func (am AccountMap) GetAccount(index uint32) *Account {
	if index < am.total {
		return &am.accounts[index]
	}
	return nil
}

func (am AccountMap) GetAccountCount() uint32 {
	return am.total
}

func (am AccountMap) GetFaucetAccount() *Account {
	return am.faucetAcct
}
