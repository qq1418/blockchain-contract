package api

// contract relation
const (
	REQUEST_FIELD_AUTH_APPID           = "appId"
	REQUEST_FIELD_AUTH_TOKEN           = "token"
	REQUEST_FIELD_AUTH_TIMESTAMP       = "timestamp"
	REQUEST_FIELD_AUTH_SIGN            = "sign"
	REQUEST_FIELD_CONTRACT_ID          = "contractId"
	REQUEST_FIELD_CONTRACT_PRODUCT_ID  = "contractProductId"
	REQUEST_FIELD_CONTRACT_TEMPLATE_ID = "contractTemplateId"
	REQUEST_FIELD_CONTRACT_OWNER       = "contractOwner"
	REQUEST_FIELD_CONTRACT_STATE       = "contractState"
	REQUEST_FIELD_CONTRACT_NAME        = "contractName"
	REQUEST_FIELD_CONTRACT_STARTTIME   = "startTime"
	REQUEST_FIELD_CONTRACT_ENDTIME     = "endTime"
	REQUEST_FIELD_PAGE                 = "pageNum"
	REQUEST_FIELD_PAGE_SIZE            = "pageSize"
	REQUEST_FIELD_SORT                 = "sort"
)

var ALLOW_REQUEST_PARAMETERS_ALL = map[string]bool{
	/*************** FOR API AUTH ************/
	REQUEST_FIELD_AUTH_APPID:     true,
	REQUEST_FIELD_AUTH_TOKEN:     true,
	REQUEST_FIELD_AUTH_TIMESTAMP: true,
	REQUEST_FIELD_AUTH_SIGN:      true,
	/*************** FOR API QUERY ***********/
	//"format":   true, //sort=field1,field2,
	REQUEST_FIELD_SORT:      true,
	REQUEST_FIELD_PAGE:      true,
	REQUEST_FIELD_PAGE_SIZE: true,
	/*************** FOR API MODEL QUERY ***********/
	REQUEST_FIELD_CONTRACT_ID:          true,
	REQUEST_FIELD_CONTRACT_PRODUCT_ID:  true,
	REQUEST_FIELD_CONTRACT_TEMPLATE_ID: true,
	REQUEST_FIELD_CONTRACT_OWNER:       true,
	REQUEST_FIELD_CONTRACT_STATE:       true,
	REQUEST_FIELD_CONTRACT_NAME:        true,
	REQUEST_FIELD_CONTRACT_STARTTIME:   true,
	REQUEST_FIELD_CONTRACT_ENDTIME:     true,
}

// filter sort invalid fields
var ALLOW_REQUEST_PARAMETERS_MODEL = map[string]bool{
	/*************** FOR API MODEL QUERY ***********/
	"contractId":         true,
	"contractProductId":  true,
	"contractTemplateId": true,
	"contractOwner":      true,
	"contractState":      true,
	"contractName":       true,
	"startTime":          true,
	"endTime":            true,
}

// FOR API FILTER
var ALLOW_REQUEST_PARAMETERS_BASIC = map[string]bool{
	"appId":     true,
	"token":     true,
	"timestamp": true,
	"sign":      true,
}

var REQUEST_CONTRACT_STATE_MAP = map[string]bool{
	"Contract_Unknown":    true,
	"Contract_Create":     true,
	"Contract_Signature":  true,
	"Contract_In_Process": true,
	"Contract_Completed":  true,
	"Contract_Discarded":  true,
}
