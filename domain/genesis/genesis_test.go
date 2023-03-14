package genesis

import (
	"fmt"
	"testing"
	"time"

	"github.com/steve-care-software/libs/cryptography/hash"
)

func TestGenesis_Success(t *testing.T) {
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
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if ins.Difficulty() != difficulty {
		t.Errorf("the difficulty was expected to be %d, %d returned", ins.Difficulty(), difficulty)
		return
	}

	if ins.Reward() != reward {
		t.Errorf("the reward was expected to be %d, %d returned", ins.Reward(), reward)
		return
	}

	if ins.Halving() != halving {
		t.Errorf("the halving was expected to be %d, %d returned", ins.Halving(), halving)
		return
	}

	if ins.MiningValue() != miningValue {
		t.Errorf("the miningValue was expected to be %d, %d returned", ins.MiningValue(), miningValue)
		return
	}

	if ins.MaxTrxPerBlock() != maxTrxPerBlock {
		t.Errorf("the maxTrxPerBlock was expected to be %d, %d returned", ins.MaxTrxPerBlock(), maxTrxPerBlock)
		return
	}

	if ins.BlockDuration() != blockDuration {
		t.Errorf("the blockDuration was expected to be %d, %d returned", ins.BlockDuration(), blockDuration)
		return
	}

	if !ins.CreatedOn().Equal(createdOn) {
		t.Errorf("the creation time was expected to be %s, %s returned", createdOn.String(), ins.CreatedOn().String())
		return
	}

	pHash, err := hash.NewAdapter().FromMultiBytes([][]byte{
		[]byte(fmt.Sprintf("%d", difficulty)),
		[]byte(fmt.Sprintf("%d", reward)),
		[]byte(fmt.Sprintf("%d", halving)),
		[]byte(fmt.Sprintf("%d", miningValue)),
		[]byte(fmt.Sprintf("%d", maxTrxPerBlock)),
		[]byte(fmt.Sprintf("%d", blockDuration)),
		[]byte(fmt.Sprintf("%d", createdOn.Unix())),
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !ins.Hash().Compare(*pHash) {
		t.Errorf("the returned hash is invalid")
		return
	}
}

func TestGenesis_withZeroDifficulty_ReturnsError(t *testing.T) {
	difficulty := uint(0)
	reward := uint(20)
	halving := uint(200)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutDifficulty_ReturnsError(t *testing.T) {
	reward := uint(20)
	halving := uint(200)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withZeroReward_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(0)
	halving := uint(200)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutReward_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	halving := uint(200)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withZeroHalving_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(0)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withMiningValueTooBig_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	miningValue := uint8(10)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutMiningValueTooBig_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutHalving_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withZeroMaxTrxPerBlock_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(0)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutMaxTrxPerBlock_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	miningValue := uint8(0)
	blockDuration := time.Duration(time.Second * 60 * 10)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withZeroBlockDuration_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(0)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutBlockDuration_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	createdOn := time.Now().UTC()
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		CreatedOn(createdOn).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestGenesis_withoutCreationTime_ReturnsError(t *testing.T) {
	difficulty := uint(1)
	reward := uint(1)
	halving := uint(20)
	miningValue := uint8(0)
	maxTrxPerBlock := uint(20)
	blockDuration := time.Duration(time.Second * 60 * 10)
	_, err := NewBuilder().Create().
		WithDifficulty(difficulty).
		WithReward(reward).
		WithHalving(halving).
		WithMiningValue(miningValue).
		WithMaxTrxPerBlock(maxTrxPerBlock).
		WithBlockDuration(blockDuration).
		Now()

	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
