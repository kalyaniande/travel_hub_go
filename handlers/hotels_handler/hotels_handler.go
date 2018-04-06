package hotels_handler

import (
	"azri_hub/helper"
	"azri_hub/payloads/search_and_availability"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func SearchandAvailability(c *gin.Context) {
	req, _ := ioutil.ReadAll(c.Request.Body)
	req_struct := search_and_availability.RequestPayload{}
	json.Unmarshal([]byte(req), &req_struct)
	agent_details := c.MustGet("agent_details")
	request_id := c.MustGet("RequestId").(string)
	api_key := c.Request.Header.Get("api-key")
	search_id := helper.GenerateSearchID(request_id, api_key)
	fmt.Println(search_id)
	agent_suppliers := helper.GetAgentSuppliers(req_struct.Rooms, agent_details)
	fmt.Println(agent_suppliers)

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
