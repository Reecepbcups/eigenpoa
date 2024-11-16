package keeper

import (
	"encoding/json"
	"errors"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AVSProposalHandler struct {
	keeper *Keeper

	prepareProposalHandler sdk.PrepareProposalHandler
	processProposalHandler sdk.ProcessProposalHandler
}

func NewAVSProposalHandler(
	k *Keeper,
	prepareProposalHandler sdk.PrepareProposalHandler,
	processProposalHandler sdk.ProcessProposalHandler,
) *AVSProposalHandler {
	fmt.Println("NewAVSProposalHandler registered")
	return &AVSProposalHandler{
		keeper:                 k,
		prepareProposalHandler: prepareProposalHandler,
		processProposalHandler: processProposalHandler,
	}
}

type (
	// TODO: make this a proto type or keep as is? (iirc the sdk team recommends not making it a proper tx)
	VoteExtension struct {
		Height    int64
		EthHeight uint64
		// TODO: instead of actually doing this, I remember some hash mechanism was used instead? this way if no updates no update.
		Operators [][]byte
	}

	InjectedVoteExtension struct {
		Operators          [][]byte
		ExtendedCommitInfo abci.ExtendedCommitInfo
	}
)

func (a AVSProposalHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (resp *abci.ResponseExtendVote, err error) {
		defer func() {
			// catch panics if possible
			if r := recover(); r != nil {
				// h.logger.Error(
				// 	"recovered from panic in ExtendVoteHandler",
				// 	"err", r,
				// )
				fmt.Println("recovered from panic in ExtendVoteHandler", "err", r)

				resp, err = &abci.ResponseExtendVote{VoteExtension: []byte{}},
					fmt.Errorf("recovered application panic in ExtendVote: %v", r)
			}
		}()

		if req == nil {
			err := fmt.Errorf("extend vote handler received a nil request")
			// h.logger.Error(err.Error())
			fmt.Println(err.Error())
			return nil, err
		}

		ethBlockHeight, err := a.keeper.Eth.BlockNumber(ctx)
		if err != nil {
			return nil, err
		}
		fmt.Println("ethBlockHeight", ethBlockHeight)

		var operators = make([][]byte, 0)
		if ethBlockHeight > 0 {
			operators, err = a.keeper.GetOperators(ctx, ethBlockHeight)
			if err != nil {
				return nil, fmt.Errorf("failed to get operator addresses: %w", err)
			}
		}

		if len(operators) == 0 {
			// TODO: error, or just continue? I feel error but then the chain would halt with PoA
			return nil, errors.New("no operators found")
		}

		// ve := VoteExtension{
		// 	Height:    req.Height,
		// 	EthHeight: ethBlockHeight,
		// 	Operators: operators,
		// }
		// fmt.Println("ExtendVoteHandler VoteExtension", ve)

		// bz, err := json.Marshal(ve)
		// if err != nil {
		// 	return nil, fmt.Errorf("failed to marshal vote extension: %w", err)
		// }

		// return &abci.ResponseExtendVote{VoteExtension: bz}, nil
		return &abci.ResponseExtendVote{VoteExtension: nil}, nil
	}
}

func (a AVSProposalHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
		var ext VoteExtension
		err := json.Unmarshal(req.VoteExtension, &ext)
		if err != nil {
			// NOTE: It is safe to return an error as the Cosmos SDK will capture all
			// errors, log them, and reject the proposal.
			return nil, fmt.Errorf("failed to unmarshal vote extension: %w", err)
		}

		if ext.Height != req.Height {
			return nil, fmt.Errorf("vote extension height does not match request height; expected: %d, got: %d", req.Height, ext.Height)
		}

		// switch {
		// case !bytes.Equal(req.Hash, ve.Hash):
		// 	return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_REJECT}, nil
		// case len(ve.Data) != 1024:
		// 	return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_REJECT}, nil
		// }

		// TODO: verify the data here (MUST be deterministic)

		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}

func (a AVSProposalHandler) PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	proposalTxs := req.Txs
	fmt.Printf("PrepareProposal: %v\n", req)

	if req.Height >= ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight && ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight != 0 {
		err := baseapp.ValidateVoteExtensions(ctx, a.keeper.valStore, 0, "", abci.ExtendedCommitInfo{
			Round: req.LocalLastCommit.Round,
			Votes: req.LocalLastCommit.Votes,
		})
		if err != nil {
			return nil, err
		}

		finalizedOperators, err := computeFinalizedOperators(ctx, req.LocalLastCommit)
		if err != nil {
			return nil, errors.New("failed to compute stake-weighted oracle prices")
		}

		injectedVoteExtTx := InjectedVoteExtension{
			Operators:          finalizedOperators,
			ExtendedCommitInfo: req.LocalLastCommit,
		}

		// NOTE: We use stdlib JSON encoding, but an application may choose to use
		// a performant mechanism. This is for demo purposes only.
		bz, err := json.Marshal(injectedVoteExtTx)
		if err != nil {
			fmt.Println("failed to encode injected vote extension tx", err)
			return nil, errors.New("failed to encode injected vote extension tx")
		}

		// Inject a "fake" tx into the proposal s.t. validators can decode, verify,
		// and store the canonical stake-weighted average prices.
		proposalTxs = append(proposalTxs, bz)
	} else {
		fmt.Println("PrepareProposal: not injecting vote extension tx because of VoteExtensionsEnableHeight")
	}

	// proceed with normal block proposal construction, e.g. POB, normal txs, etc...

	return &abci.ResponsePrepareProposal{
		Txs: proposalTxs,
	}, nil
}

func (a AVSProposalHandler) ProcessProposal(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	// if len(req.Txs) == 0 {
	// return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	// }

	// get injected data, pop it off at 1: if it is not nil, processe

	// return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	return a.processProposalHandler(ctx, req)
}

func computeFinalizedOperators(_ sdk.Context, ci abci.ExtendedCommitInfo) ([][]byte, error) {
	// TODO: some check to see how many validators voted to have an operator added. If someone votes for an operator that should NOT be in here, punish in some way?
	operators := make(map[string]struct{})
	for _, v := range ci.Votes {
		if v.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue
		}

		if len(v.VoteExtension) == 0 {
			fmt.Println("vote extension is empty skipping", "cosmos validator", fmt.Sprintf("%x", v.Validator.Address), "with power", v.Validator.Power)
			continue
		}

		var voteExt VoteExtension
		fmt.Println("v.VoteExtension", v.VoteExtension)
		if err := json.Unmarshal(v.VoteExtension, &voteExt); err != nil {
			fmt.Println("failed to decode vote extension", "err", err, "cosmos validator", fmt.Sprintf("%x", v.Validator.Address), "with power", v.Validator.Power)
			return nil, err
		}

		for _, operator := range voteExt.Operators {
			operators[string(operator)] = struct{}{}
		}
	}

	// convert back to bytes after map string sorting for uniques
	operatorsBz := make([][]byte, 0, len(operators))
	for operator := range operators {
		operatorsBz = append(operatorsBz, []byte(operator))
	}

	return operatorsBz, nil
}
