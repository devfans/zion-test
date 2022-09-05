package base

import "math/big"

var ZionPrecision = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

const (
	MethodCancelValidator = "cancelValidator"

	MethodChangeEpoch = "changeEpoch"

	MethodCreateValidator = "createValidator"

	MethodEndBlock = "endBlock"

	MethodStake = "stake"

	MethodUnStake = "unStake"

	MethodUpdateCommission = "updateCommission"

	MethodUpdateValidator = "updateValidator"

	MethodWithdraw = "withdraw"

	MethodWithdrawCommission = "withdrawCommission"

	MethodWithdrawStakeRewards = "withdrawStakeRewards"

	MethodWithdrawValidator = "withdrawValidator"

	MethodGetAccumulatedCommission = "getAccumulatedCommission"

	MethodGetAllValidators = "getAllValidators"

	MethodGetCommunityInfo = "getCommunityInfo"

	MethodGetCurrentEpochInfo = "getCurrentEpochInfo"

	MethodGetEpochInfo = "getEpochInfo"

	MethodGetGlobalConfig = "getGlobalConfig"

	MethodGetOutstandingRewards = "getOutstandingRewards"

	MethodGetStakeInfo = "getStakeInfo"

	MethodGetStakeRewards = "getStakeRewards"

	MethodGetStakeStartingInfo = "getStakeStartingInfo"

	MethodGetTotalPool = "getTotalPool"

	MethodGetUnlockingInfo = "getUnlockingInfo"

	MethodGetValidator = "getValidator"

	MethodGetValidatorAccumulatedRewards = "getValidatorAccumulatedRewards"

	MethodGetValidatorOutstandingRewards = "getValidatorOutstandingRewards"

	MethodGetValidatorSnapshotRewards = "getValidatorSnapshotRewards"

	EventCancelValidator = "CancelValidator"

	EventChangeEpoch = "ChangeEpoch"

	EventCreateValidator = "CreateValidator"

	EventStake = "Stake"

	EventUnStake = "UnStake"

	EventUpdateCommission = "UpdateCommission"

	EventUpdateValidator = "UpdateValidator"

	EventWithdraw = "Withdraw"

	EventWithdrawCommission = "WithdrawCommission"

	EventWithdrawStakeRewards = "WithdrawStakeRewards"

	EventWithdrawValidator = "WithdrawValidator"
)

const (
	MethodPropose = "propose"

	MethodProposeCommunity = "proposeCommunity"

	MethodProposeConfig = "proposeConfig"

	MethodVoteProposal = "voteProposal"

	MethodGetCommunityProposalList = "getCommunityProposalList"

	MethodGetConfigProposalList = "getConfigProposalList"

	MethodGetProposal = "getProposal"

	MethodGetProposalList = "getProposalList"

	EventPropose = "Propose"

	EventProposeCommunity = "ProposeCommunity"

	EventProposeConfig = "ProposeConfig"

	EventVoteProposal = "VoteProposal"
)

var InitialBalanceMap = map[string]string{
	"0x02484cb50AAC86Eae85610D6f4Bf026f30f6627D": "1000000000000000000000000",
	"0x08135Da0A343E492FA2d4282F2AE34c6c5CC1BbE": "1000000000000000000000000",
	"0x08A2DE6F3528319123b25935C92888B16db8913E": "1000000000000000000000000",
	"0x09DB0a93B389bEF724429898f539AEB7ac2Dd55f": "1000000000000000000000000",
	"0x1003ff39d25F2Ab16dBCc18EcE05a9B6154f65F4": "1000000000000000000000000",
	"0x11e8F3eA3C6FcF12EcfF2722d75CEFC539c51a1C": "1000000000000000000000000",
	"0x14dC79964da2C08b23698B3D3cc7Ca32193d9955": "1000000000000000000000000",
	"0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65": "1000000000000000000000000",
	"0x1BcB8e569EedAb4668e55145Cfeaf190902d3CF2": "1000000000000000000000000",
	"0x1CBd3b2770909D4e10f157cABC84C7264073C9Ec": "1000000000000000000000000",
	"0x1aac82773CB722166D7dA0d5b0FA35B0307dD99D": "1000000000000000000000000",
	"0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f": "1000000000000000000000000",
	"0x2546BcD3c84621e976D8185a91A922aE77ECEc30": "1000000000000000000000000",
	"0x2f4f06d218E426344CFE1A83D53dAd806994D325": "1000000000000000000000000",
	"0x30Bf53315437B47AeB9f6576F0f9094226342a58": "2600000000000000000000000",
	"0x35304262b9E87C00c430149f28dD154995d01207": "1000000000000000000000000",
	"0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC": "1000000000000000000000000",
	"0x3c3E2E178C69D4baD964568415a0f0c84fd6320A": "1000000000000000000000000",
	"0x40Fc963A729c542424cD800349a7E4Ecc4896624": "1000000000000000000000000",
	"0x49675089329c637c59b7DA465E00c7A8fe4c3247": "2000000000000000000000000",
	"0x4b23D303D9e3719D6CDf8d172Ea030F80509ea15": "1000000000000000000000000",
	"0x553BC17A05702530097c3677091C5BB47a3a7931": "1000000000000000000000000",
	"0x58621F440dA6cdf1A79C47E75f2dc3c804789A22": "2000000000000000000000000",
	"0x5E661B79FE2D3F6cE70F5AAC07d8Cd9abb2743F1": "1000000000000000000000000",
	"0x5F9805B99eCb2a9C3C23f9FDafF022efeeFD79b6": "2000000000000000000000000",
	"0x5eb15C0992734B5e77c888D713b4FC67b3D679A2": "1000000000000000000000000",
	"0x61097BA76cD906d2ba4FD106E757f7Eb455fc295": "1000000000000000000000000",
	"0x70997970C51812dc3A010C7d01b50e0d17dc79C8": "1000000000000000000000000",
	"0x71bE63f3384f5fb98995898A86B02Fb2426c5788": "1000000000000000000000000",
	"0x7D86687F980A56b832e9378952B738b614A99dc6": "1000000000000000000000000",
	"0x7Ebb637fd68c523613bE51aad27C35C4DB199B9c": "1000000000000000000000000",
	"0x7f8BB57E811B35783a125c5D14Fc2B820acA4C6B": "2000000000000000000000000",
	"0x8263Fce86B1b78F95Ab4dae11907d8AF88f841e7": "1000000000000000000000000",
	"0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199": "1000000000000000000000000",
	"0x86c53Eb85D0B7548fea5C4B4F82b4205C8f6Ac18": "1000000000000000000000000",
	"0x87BdCE72c06C21cd96219BD8521bDF1F42C78b5e": "1000000000000000000000000",
	"0x90F79bf6EB2c4f870365E785982E1f101E93b906": "1000000000000000000000000",
	"0x976EA74026E726554dB657fA54763abd0C3a0aa9": "1000000000000000000000000",
	"0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc": "1000000000000000000000000",
	"0x9CDED4a330682cB88093214d1DE56D0B7aE525BF": "2000000000000000000000000",
	"0x9DCCe783B6464611f38631e6C851bf441907c710": "1000000000000000000000000",
	"0x9eAF5590f2c84912A08de97FA28d0529361Deb9E": "1000000000000000000000000",
	"0x9eF6c02FB2ECc446146E05F1fF687a788a8BF76d": "1000000000000000000000000",
	"0xBcd4042DE499D14e55001CcbB24a551F3b954096": "1000000000000000000000000",
	"0xC004e69C5C04A223463Ff32042dd36DabF63A25a": "1000000000000000000000000",
	"0xC39d4b38577b1d7a465A99FCd7Bd56ce1F821A5c": "2000000000000000000000000",
	"0xD308a07F97db36C338e8FE2AfB09267781d00811": "2000000000000000000000000",
	"0xD3A01Ed4e1D0554a179Daebf95508b668767D441": "2000000000000000000000000",
	"0xD4A1E660C916855229e1712090CcfD8a424A2E33": "1000000000000000000000000",
	"0xDe224dd66aeC53E5Ee234c94F9928601777dEac7": "2000000000000000000000000",
	"0xDf37F81dAAD2b0327A0A50003740e1C935C70913": "1000000000000000000000000",
	"0xFABB0ac9d68B0B445fB7357272Ff202C5651694a": "1000000000000000000000000",
	"0xa0Ee7A142d267C1f36714E4a8F75612F20a79720": "1000000000000000000000000",
	"0xa7580f28d5304b55594CfC1907F36D91b3D77cE5": "2000000000000000000000000",
	"0xbDA5747bFD65F08deb54cb465eB87D40e51B197E": "1000000000000000000000000",
	"0xcF2d5b3cBb4D7bF04e3F7bFa8e27081B52191f91": "1000000000000000000000000",
	"0xcaBd7634D99020996c887AeE2B536f3fF1B71Fb6": "2000000000000000000000000",
	"0xcd3B766CCDd6AE721141F452C550Ca635964ce71": "1000000000000000000000000",
	"0xd98c495FE343dEDb29fB3Ed6b79834dA84f23631": "2000000000000000000000000",
	"0xdD2FD4581271e230360230F9337D5c0430Bf44C0": "1000000000000000000000000",
	"0xdF3e18d64BC6A983f673Ab319CCaE4f1a57C7097": "1000000000000000000000000",
	"0xe141C82D99D85098e03E1a1cC1CdE676556fDdE0": "1000000000000000000000000",
	"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266": "1000000000000000000000000",
}
