package model

import (
	"github.com/astaxie/beego"
	"unicontract/src/common"
	"unicontract/src/core/protos"
)

// table [contract]
type ContractModel struct {
	Id         string `json:"id"`          //合约唯一标识ID，对合约主体信息计算hash
	Version    int8   `json:"version"`     //合约描述结构版本号
	MainPubkey string `json:"main_pubkey"` //合约处理主节点公钥
	NodePubkey string `json:"node_pubkey"` //合约运行节点公钥
	Signature  string `json:"signature"`   //合约运行节点签名
	Timestamp  string `json:"timestamp"`   //合约运行跟踪时间戳（以合约执行层输出结果时间为准）

	//合约运行节点投票公钥环
	//“voters”:[voter_node1_pubkey, voter_node2_pubkey, voter_node3_pubkey],
	Voters   []string        `json:"voters"`
	Contract protos.Contract `json:"contract"` //合约描述集合, (引用contract描述 for proto3)
}

// validate the contract
func (c *ContractModel) Validate() bool {
	first_valid := c.validateContract()
	if !first_valid {
		return false
	}
	content_valid := c.validateContractContent()
	if !content_valid {
		return false
	}
	return true
}

//Create a signature for the contract
func (c *ContractModel) Sign(private_key string) string {
	contract_serialized := c.ToString()
	signatureContract := common.Sign(private_key, contract_serialized)
	return signatureContract
}

// Check the validity of a Contract's signature
func (c *ContractModel) IsSignatureValid() bool {
	contract_serialized := c.ToString()
	node_pubkey := c.NodePubkey
	signatureContract := c.Signature
	return common.Verify(node_pubkey, contract_serialized, signatureContract)
}

// TODO return new Contract with attach info
func (c *ContractModel) ToDict() *ContractModel {
	contract := &c.Contract
	if contract == nil {
		panic("Empty contract creation is not allowed")
	}
	// hash the contract in [contractModel]
	c.Id = c.GetId()
	// todo NodePubkey
	c.NodePubkey = ""
	// todo MainPubkey
	c.MainPubkey = ""
	// todo Signature
	c.Signature = ""
	// todo voters
	c.Voters = []string{}
	c.Timestamp = common.GenTimestamp()
	// todo version
	c.Version = 1
	c.Contract = *contract

	return c
}

func (c *ContractModel) ToString() string {
	return common.Serialize(c)
}

// return the  id (hash generate)
func (c *ContractModel) GetId() string {
	contract_serialized := common.Serialize(c.Contract)
	return common.HashData(contract_serialized)
}

//Validate the contract
func (c *ContractModel) validateContract() bool {
	federation := c.Voters
	nodePubkey := c.NodePubkey
	flag := false
	for _, vote := range federation {
		if vote == nodePubkey {
			flag = true
			break
		}
	}

	if !flag {
		beego.Error("Only federation nodes can create contract")
		//panic("Only federation nodes can create contract")
		return false
	}

	if !c.IsSignatureValid() {
		beego.Error("Invalid contract signature")
		//panic("Invalid contract signature")
		return false
	}
	return true
}

//Validate the contract content
func (c *ContractModel) validateContractContent() bool {
	contract := &c.Contract
	if contract == nil {
		beego.Error("Empty contract is not allowed")
		//panic("Empty contract is not allowed")
		return false
	}

	if contract.Operation == "CREATE" {
		beego.Error("missing validate the contract [creator_pubkey]")
		beego.Error("missing validate the contract [create_timestamp]")
		beego.Error("missing validate the contract [contract_attributes]")
		beego.Error("missing validate the contract [contract_owners]")
		beego.Error("missing validate the contract [contract_signatures]")
		beego.Error("missing validate the contract [contract_asserts]")
		beego.Error("missing validate the contract [contract_components]")
	}

	return true
}