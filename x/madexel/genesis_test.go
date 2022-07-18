package madexel_test

import (
	"testing"

	keepertest "github.com/MadeXeL/madexel/testutil/keeper"
	"github.com/MadeXeL/madexel/testutil/nullify"
	"github.com/MadeXeL/madexel/x/madexel"
	"github.com/MadeXeL/madexel/x/madexel/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MadexelKeeper(t)
	madexel.InitGenesis(ctx, *k, genesisState)
	got := madexel.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
