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
			{Address: "wallet1_address", Amount: 10000},
			{Address: "wallet2_address", Amount: 10000},
			{Address: "wallet3_address", Amount: 10000},
			{Address: "wallet4_address", Amount: 10000},
			{Address: "wallet5_address", Amount: 10000},
			{Address: "wallet6_address", Amount: 10000},
			{Address: "wallet7_address", Amount: 10000},
			{Address: "wallet8_address", Amount: 10000},
			{Address: "wallet9_address", Amount: 10000},
			{Address: "wallet10_address", Amount: 10000},
		},
	}
}
