package transaction

import (
	"testing"
	"github.com/astaxie/beego/logs"
	"unicontract/src/common"
	"unicontract/src/core/model"
	"unicontract/src/core/protos"
	"unicontract/src/config"
	"fmt"
	"unicontract/src/core/db/rethinkdb"
)

func Test_createTx(t *testing.T) {
	config.Init()
	tx_signers := []string{
		"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL",
	}
	recipients := [][2]interface{}{
		{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL", 500},
	}

	tempMap := make(map[string]interface{})
	tempMap["a"] = "1"
	tempMap["c"] = "3"
	tempMap["b"] = "2"
	tempMap["A"] = "4"
	tempMap["6"] = map[string]string{"QQQQ": "9999"}
	metadata := model.Metadata{
		Id:   "meta-data-id",
		Data: tempMap,
	}
	asset := model.Asset{}

	//--------------------contract-------------------------
	contract := model.ContractModel{}
	contractAsset := []*protos.ContractAsset{}
	contractComponent := []*protos.ContractComponent{}
	contractBody := &protos.ContractBody{
		ContractId:         "feca0672-4ad7-4d9a-ad57-83d48db2269b",
		Cname:              "test contract output",
		Ctype:              "CREATE",
		Caption:            "购智能手机返话费合约产品协议",
		Description:        "移动用户A花费500元购买移动运营商B的提供的合约智能手机C后",
		ContractState:      "",
		Creator:            common.GenTimestamp(),
		CreatorTime:        "1493111926720",
		StartTime:          "1493111926730",
		EndTime:            "1493111926740",
		ContractOwners:     []string{"BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443", "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy", "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ", },
		ContractSignatures: nil,
		ContractAssets:     contractAsset,
		ContractComponents: contractComponent,
	}
	contractSignatures := []*protos.ContractSignature{
		{
			OwnerPubkey:   "BtS4rHnMvhJELuP5PKKrdjN7Mp1rqerx6iuEz3diW443",
			Signature:     contract.Sign("hg6uXBjkcpn6kmeBthETonH66c26GyAcasGdBMaYTbC"),
			SignTimestamp: "1493111926751",
		},
		{
			OwnerPubkey:   "4tBAt7QjZE8Eub58UFNVg6DSAcH3uY4rftZJZb5ngPMy",
			Signature:     contract.Sign("AnV4aa3KCpsNF8bEqQ8qF8T97iW4KnhMmPKwaFWgKhRo"),
			SignTimestamp: "1493111926752",
		},
		{
			OwnerPubkey:   "9cEcV6CywjZSed8AC2zUFUYC94KXbn4Fe7DnqBQgYpwQ",
			Signature:     contract.Sign("9647UfPdDSwBf5kw7tUrSe7cmYY5RvVX47GrGqSh4XVi"),
			SignTimestamp: "1493111926753",
		},
	}
	contractBody.ContractSignatures = contractSignatures
	contract.ContractHead = &protos.ContractHead{config.Config.Keypair.PublicKey, 1}
	contract.ContractBody = contractBody
	contract.Id = common.HashData(common.StructSerialize(contractBody))

	relation := model.Relation{
		ContractId: contract.Id,
		TaskId:     "task-id-123456789",
		Voters: []string{
			config.Config.Keypair.PublicKey,
		},
	}

	output,_ := Create(tx_signers, recipients, metadata, asset, relation, contract)
	output = NodeSign(output)
	logs.Info(common.StructSerialize(output))
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	fmt.Println(b)
}

func Test_FreezeTx(t *testing.T) {
	config.Init()
	ownerbefore := "5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL"
	recipients := [][2]interface{}{
		[2]interface{}{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL", 300},
	}
	metadata := model.Metadata{}
	asset := GetAsset(ownerbefore)

	contract := GetContractFromUnichain("feca0672-4ad7-4d9a-ad57-83d48db2269b")

	relation := model.Relation{}
	relation.GenerateRelation("feca0672-4ad7-4d9a-ad57-83d48db2269b", "taskId")

	output,_ := Transfer("FREEZE", ownerbefore, recipients, metadata, asset, relation, contract)
	output = NodeSign(output)
	logs.Info(common.StructSerialize(output))
	b := rethinkdb.InsertContractOutput(common.StructSerialize(output))
	fmt.Println(b)
}

func Test_GetUnspent(t *testing.T) {
	config.Init()
	//pubkey := config.Config.Keypair.PublicKey
	inps, bal := GetUnspent("5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL")
	logs.Info(inps)
	logs.Info(bal)
}

func Test_GetFreezespent(t *testing.T) {
	config.Init()
	//pubkey := config.Config.Keypair.PublicKey
	inps, bal := GetFreezeUnspent("5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL")
	logs.Info(inps)
	logs.Info(bal)
}

func Test_GetContractFromUnichain(t *testing.T) {
	contract := GetContractFromUnichain("feca0672-4ad7-4d9a-ad57-83d48db2269b")
	logs.Info(common.StructSerialize(contract))
}
