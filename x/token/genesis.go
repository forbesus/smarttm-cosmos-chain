package token

// GenesisState defines the module's initial state
type GenesisState struct {
	Tokens   []Token   `json:"tokens"`   // Define token types
	Balances []Balance `json:"balances"` // Define account balances
}

// Token represents a token in the genesis state
type Token struct {
	Denom       string `json:"denom"`        // Token denomination (e.g., smartTM)
	TotalSupply uint64 `json:"total_supply"` // Total supply of the token
}

// Balance represents a balance entry in genesis
type Balance struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

// DefaultGenesis sets up initial tokens and balances
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Tokens: []Token{
			{Denom: "smartTM", TotalSupply: 100000}, // Set total supply
		},
		Balances: []Balance{
			{Address: "cosmos1q7ev25sze0g6pua0sdvqpnkjge86f6elpqyxmn", Amount: 10000},
			{Address: "cosmos1ml95mze9j50nre38w7aqkatssq3fxyuc3y3j2y", Amount: 10000},
			{Address: "cosmos1cxq2ynzy7y96s8thpggfnupr2c4yesjclszteu", Amount: 10000},
			{Address: "cosmos1cq9s4kyrepxxy3vp7c6uqzypclvrwtx75atysu", Amount: 10000},
			{Address: "cosmos1e9r2fhg234xvh8fywfh4g5ygr6g7jrkm3ut792", Amount: 10000},
			{Address: "cosmos1ardueyuwhwuxm0x4gn7tvrukv8w6jx200y08k4", Amount: 10000},
			{Address: "cosmos1zhfk7pf6q7zrdzmtk5kfjgkjpacnyc6cfwesjc", Amount: 10000},
			{Address: "cosmos19cc5agyhh40cp4vzm0wqek8xhe9nf0sxvyc2tn", Amount: 10000},
			{Address: "cosmos1vafkych49fzcqqjxz90tp3km38zy8h5hmsgwpg", Amount: 10000},
			{Address: "cosmos1qrd52sxtj9aywle7n8vaethnyjy4typu0p4gj0", Amount: 10000},
		},
	}
}
