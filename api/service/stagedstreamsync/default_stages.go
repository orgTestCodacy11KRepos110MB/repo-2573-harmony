package stagedstreamsync

import (
	"context"
)

type ForwardOrder []SyncStageID
type RevertOrder []SyncStageID
type CleanUpOrder []SyncStageID

var DefaultForwardOrder = ForwardOrder{
	Heads,
	ShortRange,
	SyncEpoch,
	BlockBodies,
	// Stages below don't use Internet
	States,
	Finish,
}

var DefaultRevertOrder = RevertOrder{
	Finish,
	States,
	BlockBodies,
	SyncEpoch,
	ShortRange,
	Heads,
}

var DefaultCleanUpOrder = CleanUpOrder{
	Finish,
	States,
	BlockBodies,
	SyncEpoch,
	ShortRange,
	Heads,
}

func DefaultStages(ctx context.Context,
	headsCfg StageHeadsCfg,
	srCfg StageShortRangeCfg,
	seCfg StageEpochCfg,
	bodiesCfg StageBodiesCfg,
	statesCfg StageStatesCfg,
	finishCfg StageFinishCfg,
) []*Stage {

	handlerStageHeads := NewStageHeads(headsCfg)
	handlerStageShortRange := NewStageShortRange(srCfg)
	handlerStageEpochSync := NewStageEpoch(seCfg)
	handlerStageBodies := NewStageBodies(bodiesCfg)
	handlerStageStates := NewStageStates(statesCfg)
	handlerStageFinish := NewStageFinish(finishCfg)

	return []*Stage{
		{
			ID:          Heads,
			Description: "Retrieve Chain Heads",
			Handler:     handlerStageHeads,
		},
		{
			ID:          ShortRange,
			Description: "Short Range Sync",
			Handler:     handlerStageShortRange,
		},
		{
			ID:          SyncEpoch,
			Description: "Sync only Last Block of Epoch",
			Handler:     handlerStageEpochSync,
		},
		{
			ID:          BlockBodies,
			Description: "Retrieve Block Bodies",
			Handler:     handlerStageBodies,
		},
		{
			ID:          States,
			Description: "Update Blockchain State",
			Handler:     handlerStageStates,
		},
		{
			ID:          Finish,
			Description: "Finalize Changes",
			Handler:     handlerStageFinish,
		},
	}
}
