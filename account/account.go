package account

type Account struct  {
	ID int
	Balance	int
	Blocked bool
}

var Accounts = make(map[int]Account)

func FindAccountById(id int) (Account, error) {
	return Accounts[id], nil
}

func (a *Account) Deposit(amount int) {
	acc := Accounts[a.ID]
	acc.Balance += amount
	Accounts[a.ID] = acc
}

func (a *Account) Withdraw(amount int) {
	acc := Accounts[a.ID]
	acc.Balance -= amount
	Accounts[a.ID] = acc
}

func (a Account) HasAvailableFunds(amount int) bool{
	if a.Balance >= amount {
		return true
	}
	return false
}