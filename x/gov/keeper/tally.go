package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// DefaultContractAddr is the wasm contract address generated by code ID 1 and instance ID 1.
//
// In other words, the first ever contract to be deployed on this chain will necessarily have this address.
var DefaultContractAddr = wasmkeeper.BuildContractAddress(1, 1)

// Tally iterates over the votes and updates the tally of a proposal based on the voting power of the voters
func (k Keeper) Tally(ctx sdk.Context, proposal govtypes.Proposal) (passes bool, burnDeposits bool, tallyResults govtypes.TallyResult) {
	results := make(map[govtypes.VoteOption]sdk.Dec)
	results[govtypes.OptionYes] = sdk.ZeroDec()
	results[govtypes.OptionAbstain] = sdk.ZeroDec()
	results[govtypes.OptionNo] = sdk.ZeroDec()
	results[govtypes.OptionNoWithVeto] = sdk.ZeroDec()

	// fetch all tokens locked in the DAO contract
	//
	// NOTE: for now we simply use the default contract address. later it may be a better idea to use
	// a configurable parameter
	tokensLocked, totalTokensLocked := MustGetTokensInVesting(ctx, k.wasmKeeper, DefaultContractAddr)

	// total amount of tokens that have voted in this poll; used to determine whether the poll reaches
	// quorum and the pass threshold
	totalTokensVoted := sdk.ZeroDec()

	// iterate through votes
	k.IterateVotes(ctx, proposal.ProposalId, func(vote govtypes.Vote) bool {
		voterAddr := sdk.MustAccAddressFromBech32(vote.Voter)

		votingPower := sdk.ZeroDec()

		// if the voter has tokens locked in DAO contract, add that to the voting power
		if votingPowerInDao, ok := tokensLocked[vote.Voter]; ok {
			votingPower = votingPower.Add(votingPowerInDao.ToDec())
		}

		incrementTallyResult(votingPower, vote.Options, results, &totalTokensVoted)
		k.deleteVote(ctx, vote.ProposalId, voterAddr)

		return false
	})

	tallyParams := k.GetTallyParams(ctx)
	tallyResults = govtypes.NewTallyResultFromMap(results)

	// if there is no staked coins, the proposal fails
	if totalTokensLocked.IsZero() {
		return false, false, tallyResults
	}

	// if there is not enough quorum of votes, the proposal fails, and deposit burned
	//
	// NOTE: should the deposit really be burned here?
	if totalTokensVoted.Quo(totalTokensLocked.ToDec()).LT(tallyParams.Quorum) {
		return false, true, tallyResults
	}

	// if everyone abstains, proposal fails
	if totalTokensVoted.Sub(results[govtypes.OptionAbstain]).IsZero() {
		return false, false, tallyResults
	}

	// if more than 1/3 of voters veto, proposal fails, and deposit burned
	//
	// NOTE: here 1/3 is defined as 1/3 *of all votes*, including abstaining votes. could it make more
	// sense to instead define it as 1/3 *of all non-abstaining votes*?
	if results[govtypes.OptionNoWithVeto].Quo(totalTokensVoted).GT(tallyParams.VetoThreshold) {
		return false, true, tallyResults
	}

	// if no less than 1/2 of non-abstaining voters vote No, proposal fails
	if results[govtypes.OptionNo].Quo(totalTokensVoted.Sub(results[govtypes.OptionAbstain])).GTE(tallyParams.Threshold) {
		return false, false, tallyResults
	}

	// otherwise, meaning more than 1/2 of non-abstaining voters vote Yes, proposal passes
	return true, false, tallyResults
}

func incrementTallyResult(votingPower sdk.Dec, options []govtypes.WeightedVoteOption, results map[govtypes.VoteOption]sdk.Dec, totalTokensVoted *sdk.Dec) {
	for _, option := range options {
		subPower := votingPower.Mul(option.Weight)
		results[option.Option] = results[option.Option].Add(subPower)
	}

	*totalTokensVoted = totalTokensVoted.Add(votingPower)
}
