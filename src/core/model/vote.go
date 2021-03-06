package model

import (
	"unicontract/src/common"
	"unicontract/src/config"
)

type VoteBody struct {
	IsValid       bool   //合约、合约交易投票结果，如true,false
	InvalidReason string //合约、合约交易投无效票原因
	VoteFor       string //contract.id
	VoteType      string //投票对象的类型，如CONTRACT，TRANSACTION等
	Timestamp     string //节点投票时间戳
}

// table [Votes]
type Vote struct {
	Id         string   `json:"id"` //投票唯一标识ID，UUID
	NodePubkey string   //投票节点的公钥
	VoteBody   VoteBody `json:"Vote"` //投票信息
	Signature  string   //投票节点签名
}

// Calculate the election status of a contract.
func (v *Vote) ContractElection() {

}

//  Filter votes from unknown nodes or nodes that are not listed on
//	block. This is the primary Sybill protection.
func PartitionEligibleVotes(votes []Vote, eligible_voters []string) ([]Vote, []Vote) {
	//eligible := make([]Votes, len(votes))
	//ineligible := make([]Votes, len(votes))
	var eligible []Vote
	var ineligible []Vote
	for _, _votes := range votes {
		voter_eligible := make([]string, len(votes))
		for _, _voter_eligible := range eligible_voters {
			if _votes.NodePubkey == _voter_eligible {
				voter_eligible = append(voter_eligible[:], _voter_eligible)
			}
		}
		if voter_eligible != nil {
			if _votes.VerifyVoteSignature() {
				eligible = append(eligible[:], _votes)
				continue
			}
		}
		ineligible = append(ineligible[:], _votes)
	}
	return eligible, ineligible
}

func CountVotes(eligible_votes []Vote) map[string]interface{} {

	// Group by pubkey to detect duplicate voting
	by_voter := make(map[Vote]bool)
	for _, votes := range eligible_votes {
		if !by_voter[votes] {
			by_voter[votes] = true
		}
	}
	n_valid := 0
	cheat := make([]Vote, len(by_voter))
	for votes, _ := range by_voter {
		cheat = append(cheat, votes)
		vote := votes.VoteBody
		if vote.IsValid {
			n_valid += 1
		}
	}

	resultMap := make(map[string]interface{})
	counts := make(map[string]int)
	counts["n_valid"] = n_valid
	counts["n_invalid"] = len(by_voter) - n_valid
	resultMap["cheat"] = cheat
	resultMap["counts"] = counts

	return resultMap
}

// TODO Decide on votes.
// TODO To return VALID there must be a clear majority that say VALID.
// TODO A tie on an even number of votes counts as INVALID.
func (v *Vote) DecideVotes(n_voters int, n_valid int, n_invalid int) string {
	if n_invalid*2 >= n_voters {
		return "INVALID"
	}

	if n_valid*2 > n_voters {
		return "VALID"
	}

	return "UNDECIDED"
}

func (v *Vote) VerifyVoteSignature() bool {
	signature := v.Signature
	pub := v.NodePubkey
	body := common.StructSerialize(v.VoteBody)
	return common.Verify(pub, body, signature)
}

func (v *Vote) SignVote() string {
	priv_key := config.Config.Keypair.PrivateKey
	msg := common.StructSerialize(v.VoteBody)
	sig := common.Sign(priv_key, msg)
	return sig
}

func (v *Vote) ToString() string {
	return common.StructSerialize(v)
}

// generate UUID
func (v *Vote) GenerateId() string {
	return common.GenerateUUID()
}
