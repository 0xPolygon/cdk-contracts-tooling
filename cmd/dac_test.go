package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygondatacommittee"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestSetupDAC(t *testing.T) {
	// Deploy DAC
	out, err := exec.Command(
		"go", "run", "./cmd", "deploy-dac",
		"-l1", "local",
		"-w", walletAddr,
		"-wp", walletPass,
		"-skip-confirmation",
	).CombinedOutput()
	require.NoError(t, err, string(out))
	_, msg, _ := strings.Cut(string(out), "DAC proxy will be deployed at ")
	dacAddr, _, _ := strings.Cut(msg, " with the tx ")
	time.Sleep(3 * time.Second) // TODO: wait for tx to be mined
	fmt.Println(string(out))
	client, err := getClient()
	require.NoError(t, err)
	dac, err := polygondatacommittee.NewPolygondatacommittee(common.HexToAddress(dacAddr), client)
	require.NoError(t, err)
	owner, err := dac.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(walletAddr), owner)

	// Setup DAC
	setup := DACSetup{
		RequiredSingatures: 1,
		Members: []DACMember{
			{
				Address: common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"),
				URL:     "http://foo.bar",
			},
			{
				Address: common.HexToAddress("0xFFFFd3AccA5a467e9e704C703E8D87F634fB0Fc9"),
				URL:     "http://bar.foo",
			},
		},
	}
	setupFile, err := os.CreateTemp("", "")
	require.NoError(t, err)
	setupData, err := json.Marshal(setup)
	require.NoError(t, err)
	err = os.WriteFile(setupFile.Name(), setupData, 0644)
	require.NoError(t, err)
	out, err = exec.Command(
		"go", "run", "./cmd", "setup-dac",
		"-l1", "local",
		"-w", walletAddr,
		"-wp", walletPass,
		"-a", dacAddr,
		"-f", setupFile.Name(),
		"-skip-confirmation",
	).CombinedOutput()
	require.NoError(t, err, string(out))

	reqSigsSC, err := dac.RequiredAmountOfSignatures(nil)
	require.NoError(t, err)
	require.Equal(t, setup.RequiredSingatures, reqSigsSC.Int64())
	nMembers, err := dac.GetAmountOfMembers(nil)
	require.NoError(t, err)
	require.Equal(t, len(setup.Members), int(nMembers.Int64()))
	for i, member := range setup.Members {
		memberSC, err := dac.Members(nil, big.NewInt(int64(i)))
		require.NoError(t, err)
		require.Equal(t, member.URL, memberSC.Url)
		require.Equal(t, member.Address, memberSC.Addr)
	}
}
