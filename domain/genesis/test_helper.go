package genesis

import "time"

// NewGenesisForTests creates a new genesis for tests
func NewGenesisForTests() Genesis {
	difficulty := uint(1)
	reward := uint(20)
	halving := uint(200)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	ins, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
