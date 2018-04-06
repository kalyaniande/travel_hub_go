package helper

import (
	"azri_hub/dcache"
	"azri_hub/payloads/search_and_availability"
	"azri_hub/webservice/rest"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	_ "fmt"
	"strconv"
	"time"
)

var AgentEndPoint = "http://localhost:4003/api/v3/agents"
var HubtoServerAPIKey = "i5je4q2ikvbfit2tivjfmrksgiydcnrnga4c2mrveaytkorsge5denq"

func GetAgentInfo(api_key string) dcache.Agent {
	agent_info, err := dcache.GetAgentInfo(api_key)
	if err != nil {
		return MakeAgentInfoReq(api_key)
	}
	return agent_info
}

func MakeAgentInfoReq(api_key string) dcache.Agent {
	headers := map[string]string{
		"Content-Type": "application/json",
		"api-key":      HubtoServerAPIKey,
		"user-id":      api_key,
	}
	req_data := rest.RequestData{
		AgentEndPoint + "/" + api_key,
		headers,
		make(map[string]string),
		"",
		"",
		make(map[string]string),
	}
	resp := rest.MakeRequest("GET", req_data)
	agent := dcache.Agent{}
	json.Unmarshal([]byte(resp), &agent)
	return agent
}

func GenerateSearchID(req_id string, api_key string) string {
	timestamp := GetCurrentTimestamp()
	search_id := api_key + req_id + strconv.FormatInt(timestamp, 10)
	return GetMD5String(search_id)
}

func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetMD5String(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetAgentSuppliers(rooms []search_and_availability.Room, agent interface{}) []string {
	//suppliers := make([]string)
	agent_info := agent.(dcache.Agent)
	if len(rooms) == 1 || agent_info.SuppliersType == "" || agent_info.SuppliersType == "A" {
		return agent_info.Suppliers
	}
	return agent_info.Suppliers
}
