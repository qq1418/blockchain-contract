package transaction

import (
	"encoding/json"
	"errors"
	"strconv"

	"unicontract/src/chain"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"

	"github.com/astaxie/beego/logs"
)

const (
	_GENESIS  = "GENESIS"
	_CREATE   = "CREATE"
	_TRANSFER = "TRANSFER"
	_INTERIM  = "INTERIM"
	_CONTRACT = "CONTRACT"
	_FREEZE   = "FREEZE"
	_UNFREEZ  = "UNFREEZE"
	_VERSION  = 2
)

//var ALLOWED_OPERATIONS = [4]string{_GENESIS, _CREATE, _TRANSFER, _CONTRACT}

//create asset
func Create(tx_signers []string, recipients [][2]interface{}, metadata *model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) (model.ContractOutput, error) {

	isFeeze := false
	operation := _CREATE
	version := _VERSION
	timestamp := common.GenTimestamp()
	//generate outputs
	outputs := []*model.ConditionsItem{}
	for index, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		output := &model.ConditionsItem{}
		output.GenerateOutput(index, isFeeze, pubkey, amount)
		outputs = append(outputs, output)
	}
	//generate inputs
	inputs := []*model.Fulfillment{}
	input := &model.Fulfillment{}
	input.GenerateInput(tx_signers)
	inputs = append(inputs, input)

	contractOutput := model.ContractOutput{}
	contractOutput.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput, nil
}

//transfer asset include:transfer/freeze/unfreeze
func Transfer(operation string, ownerbefore string, recipients [][2]interface{}, metadata *model.Metadata, asset model.Asset,
	relation model.Relation, contract model.ContractModel) (model.ContractOutput, error) {
	contractId := relation.ContractId
	taskId := relation.TaskId
	taskExecuteIdx := relation.TaskExecuteIdx
	isFeeze := false
	//generate inputs
	var inputs = []*model.Fulfillment{}
	var balance int = 0
	var spentFlag float64 = -1 //0:no asset was frozen;  1:the asset was frozen; 2:the frozen asset had transfer
	logs.Info(operation)
	if operation == _FREEZE {
		isFeeze = true
		inputs, balance, spentFlag = GetFrozenUnspent(ownerbefore, contractId, taskId, taskExecuteIdx)
		if spentFlag == 1 || spentFlag == 3 || spentFlag == 4 {
			err := errors.New("there had one/mutil frozen asset , step over 'FREEZE' ")
			return model.ContractOutput{}, err
		}
		//generate inputs
		inputs, balance = GetUnfreezeUnspent(ownerbefore)
		if len(recipients) != 1 || recipients[0][0].(string) != ownerbefore {
			err := errors.New("The opertion `FREEZE` should has one ownerafter = ownerbefore !")
			return model.ContractOutput{}, err
		}
	} else if operation == _UNFREEZ {
		if len(recipients) > 0 {
			err := errors.New("The opertion `UNFREEZE` should not has any ownner-afters !")
			return model.ContractOutput{}, err
		}

		inputs, balance, spentFlag = GetFrozenUnspent(ownerbefore, contractId, taskId, taskExecuteIdx)
		//NOTE  not sure whether I need to check the inputs is it has some frozen asset or not

	} else if operation == _TRANSFER {
		//generate inputs
		inputs, balance, spentFlag = GetFrozenUnspent(ownerbefore, contractId, taskId, taskExecuteIdx)

		if spentFlag == 0 {
			err := errors.New("Can not find any frozen asset !")
			return model.ContractOutput{}, err
		} else if spentFlag == 2 {
			err := errors.New("The frozen asset had be unfreezed !")
			return model.ContractOutput{}, err
		} else if spentFlag == 3 {
			err := errors.New("The frozen asset had be transfered !")
			return model.ContractOutput{}, err
		} else if spentFlag == 4 {
			err := errors.New("There is muti-frozen asset ,please check on !")
			return model.ContractOutput{}, err
		}
	}
	if len(inputs) == 0 {
		err := errors.New("Can not find any asset to do this operation!")
		return model.ContractOutput{}, err
	}
	//the operation in DB needs to be 'TRANSFER'
	operation = _TRANSFER
	version := _VERSION
	timestamp := common.GenTimestamp()

	//generate outputs
	outputs := []*model.ConditionsItem{}
	amounts := 0
	for index, recipient := range recipients {
		pubkey := recipient[0].(string)
		amount := recipient[1].(int)
		output := &model.ConditionsItem{}
		output.GenerateOutput(index, isFeeze, pubkey, amount)
		outputs = append(outputs, output)
		amounts += amount
	}
	if balance < amounts {
		err := errors.New("not enough asset to do the operation !")
		return model.ContractOutput{}, err
	} else if balance > amounts {
		pubkey := ownerbefore
		amount := balance - amounts
		output := &model.ConditionsItem{}
		output.GenerateOutput(len(outputs), false, pubkey, amount)
		outputs = append(outputs, output)
	}

	contractOutput := model.ContractOutput{}
	contractOutput.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	//logs.Info("down---",common.StructSerialize(contractOutput))
	return contractOutput, nil
}

func Interim(metadata *model.Metadata,
	relation model.Relation, contract model.ContractModel) (model.ContractOutput, error) {

	operation := _INTERIM
	version := _VERSION
	timestamp := common.GenTimestamp()
	//generate outputs
	outputs := []*model.ConditionsItem{}
	inputs := []*model.Fulfillment{}
	contractOutput := model.ContractOutput{}
	asset := model.Asset{}
	contractOutput.GenerateConOutput(operation, asset, inputs, outputs, metadata, timestamp, version, relation, contract)
	return contractOutput, nil
}

// the all unspent asset include 'freeze'/'unfreeze'
func GetAllUnspent(pubkey string, contractId string) {
	//TODO when it needed
}

// the unspent asset only include 'unfreeze'
func GetUnfreezeUnspent(pubkey string) (inps []*model.Fulfillment, bal int) {
	param := "unspent=true&public_key=" + pubkey
	result, err := chain.GetUnspentTxs(param)
	if err != nil {
		logs.Error(err.Error())
		return nil, 0
	}
	inputs := []*model.Fulfillment{}
	var balance int
	for index, unspend := range result.Data.([]interface{}) {
		//logs.Info("unspend-map====",unspend)
		unspenStruct := model.UnSpentOutput{}
		mapObjBytes, _ := json.Marshal(unspend)
		json.Unmarshal(mapObjBytes, &unspenStruct)
		//logs.Info("unspend====", common.StructSerialize(unspenStruct))
		//generate input
		inoutLink := model.ContractOutputLink{
			Cid:  unspenStruct.Cid,
			Txid: unspenStruct.Txid,
		}
		ownerbefore := []string{pubkey}
		input := model.Fulfillment{
			Fid:          index,
			OwnersBefore: ownerbefore,
			Fulfillment:  "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD",
			Input:        &inoutLink,
		}
		inputs = append(inputs, &input)
		balance += unspenStruct.Amount
		logs.Info("input====", common.StructSerialize(input))
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		return nil, 0
	}
	return inputs, balance
}

func GetAsset(ownerbefore string) model.Asset {
	asset := model.Asset{}
	//TODO  get asset
	//asset.Id = common.GenerateUUID()
	asset.Id = "f3c71d28-861e-4cf3-8907-48c713d3677e"
	data := make(map[string]string)
	data["money"] = "RMB"
	asset.Data = data
	asset.Divisible = true
	asset.Updatable = false
	asset.Refillable = false
	return asset
}

//the unspent asset only include 'freeze'
/*
return:
	flag:
		-1:request faild;
		0:no asset was frozen;
		1:get one frozen asset;
		2:the frozen asset had unfreeze;
		3:the frozen asset had transfer;
		4:get muti-frozen-asset;
*/
func GetFrozenUnspent(pubkey string, contractId string, taskId string, taskNum int) (inps []*model.Fulfillment, bal int, flag float64) {

	taskNumStr := strconv.Itoa(taskNum)
	param := "unspent=true&public_key=" + pubkey + "&contract_id=" + contractId + "&task_id=" + taskId + "&task_num=" + taskNumStr
	logs.Info(param)
	result, err := chain.GetFreezeUnspentTxs(param)
	if err != nil {
		logs.Error(err.Error())
		return nil, 0, -1
	}
	if result.Code != 200 {
		logs.Error(errors.New("request send failed"))
		return nil, 0, -1
	}

	inputs := []*model.Fulfillment{}
	var balance int
	logs.Info(result.Data)
	flag = result.Data.([]interface{})[0].(float64)
	unspendSlice := result.Data.([]interface{})[1].([]interface{})
	for index, unspend := range unspendSlice {
		//to map
		//logs.Info("unspend-map====",unspend)
		var unspentStruct model.UnSpentOutput
		mapObjBytes, _ := json.Marshal(unspend)
		json.Unmarshal(mapObjBytes, &unspentStruct)
		//logs.Info("unspend-stu====", common.StructSerialize(unspentStruct))
		//generate input
		inoutLink := model.ContractOutputLink{
			Cid:  unspentStruct.Cid,
			Txid: unspentStruct.Txid,
		}
		ownerbefore := []string{pubkey}
		input := model.Fulfillment{
			Fid:          index,
			OwnersBefore: ownerbefore,
			Fulfillment:  "cf:4:RtTtCxNf1Bq7MFeIToEosMAa3v_jKtZUtqiWAXyFz1ejPMv-t7vT6DANcrYvKFHAsZblmZ1Xk03HQdJbGiMyb5CmQqGPHwlgKusNu9N_IDtPn7y16veJ1RBrUP-up4YD",
			Input:        &inoutLink,
		}
		inputs = append(inputs, &input)
		balance += unspentStruct.Amount
		logs.Info("input====", common.StructSerialize(input))
	}
	//logs.Info(inputs)
	logs.Info(balance)
	logs.Info(flag)
	return inputs, balance, flag
}

func GetContractFromUnichain(contractId string) (model.ContractModel, error) {
	param := `{"contract_id":"` + contractId + `"}`
	result, err := chain.GetContractById(param)
	if err != nil {
		logs.Error(err.Error())
		return model.ContractModel{}, err
	}
	var contractStruct model.ContractModel
	res, ok := result.Data.([]interface{})
	if !ok {
		err = errors.New("type error")
		return contractStruct, err
	}
	if len(res) == 0 {
		err = errors.New("not find the contract")
		return contractStruct, err
	}
	for _, contract := range res {
		contractStruct = model.ContractModel{}
		contractBytes, _ := json.Marshal(contract)
		json.Unmarshal(contractBytes, &contractStruct)
		//logs.Info("index=", index, "----contract=", common.StructSerialize(contractStruct))
	}
	return contractStruct, nil
}

func NodeSign(contractOutput model.ContractOutput) model.ContractOutput {
	vote := &model.Vote{}
	vote.Id = common.GenerateUUID()
	vote.NodePubkey = config.Config.Keypair.PublicKey
	vote.VoteBody.Timestamp = common.GenTimestamp()
	vote.VoteBody.InvalidReason = "invalid reason"
	vote.VoteBody.IsValid = true
	vote.VoteBody.VoteFor = contractOutput.Id
	vote.VoteBody.VoteType = "TRANSACTION"
	// note: contractoutput(transaction) node signatrue : use the contractOutput.id
	vote.Signature = common.Sign(config.Config.Keypair.PrivateKey, contractOutput.Id)
	voters := contractOutput.Transaction.Relation.Voters
	votes := make([]*model.Vote, len(voters))
	for index, key := range voters {
		logs.Info("index::", index)
		if key == config.Config.Keypair.PublicKey {
			votes[index] = vote
		}
	}
	contractOutput.Transaction.Relation.Votes = votes
	return contractOutput
}

func IsOutputInUnichain(contractHashId string) (bool, error) {
	param := `{"contract_hash_id":"` + contractHashId + `"}`
	result, err := chain.GetTxByConHashId(param)
	if err != nil {
		logs.Error(err.Error())
		return false, err
	}
	logs.Info(result.Data)
	output, ok := result.Data.([]interface{})
	if !ok {
		err = errors.New("type error")
		logs.Error(err)
		return false, err
	}
	logs.Info(len(output))
	if len(output) > 0 {
		return true, nil
	}
	return false, err
}
