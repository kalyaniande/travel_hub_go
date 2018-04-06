package dcache

import (
	aredis "azri_hub/redis"
	"errors"
	_ "fmt"
	"log"
	"encoding/json"
)

type Agent struct {
	Wallet struct {
		Currency        interface{} `json:"currency"`
		AvailableCredit int         `json:"available_credit"`
	} `json:"wallet"`
	SuppliersType     string   `json:"suppliers_type"`
	Suppliers         []string `json:"suppliers"`
	Status            int      `json:"status"`
	StateCode         string   `json:"state_code"`
	PreferredCurrency string   `json:"preferred_currency"`
	PinCode           string   `json:"pin_code"`
	Markup            struct {
		MarkupType string  `json:"markup_type"`
		Markup     float64 `json:"markup"`
		Currency   string  `json:"currency"`
	} `json:"markup"`
	IsRebookingAgent bool   `json:"is_rebooking_agent"`
	InvoiceCurrency  string `json:"invoice_currency"`
	ID               int    `json:"id"`
	GstNumber        string `json:"gst_number"`
	Credit           struct {
		Currency        string  `json:"currency"`
		AvailableCredit float64 `json:"available_credit"`
	} `json:"credit"`
	CountryCode string `json:"country_code"`
	CityCode    string `json:"city_code"`
	APIKey      string `json:"api_key"`
	AgentConfig struct {
		Servicefee struct {
			Value    float64     `json:"value"`
			Type     int         `json:"type"`
			Currency interface{} `json:"currency"`
		} `json:"servicefee"`
	} `json:"agent_config"`
	Address string `json:"address"`
}

var ErrNoAgent = errors.New("Agent: could not be found")

func GetAgentInfo(api_key string) (Agent, error) {
	conn, err := aredis.DredisPool.Get()

	if err != nil {
		log.Fatal(err)
		return Agent{}, err
	}
	defer aredis.DredisPool.Put(conn)

	result, err := conn.Cmd("HGET", "agent", api_key).Str()

	if err != nil {
		return Agent{}, err
	} else if result == "" {
		return Agent{}, ErrNoAgent
	}
	agent := Agent{}
	json.Unmarshal([]byte(result), &agent)

	return agent, nil
}

func PutAgentInfo(api_key string, data Agent) error {
	conn, err := aredis.DredisPool.Get()
	if err != nil {
		return err
	}
	defer aredis.DredisPool.Put(conn)

	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp := conn.Cmd("HSET", "agent", api_key, json_data)
        if resp.Err != nil {
		return resp.Err
	}
	return nil
}

func DeleteAgentInfo(api_key string) error {
	conn, err := aredis.DredisPool.Get()
	if err != nil {
		return err
	}
	defer aredis.DredisPool.Put(conn)

	resp := conn.Cmd("HDEL", "agent", api_key)
	if resp.Err != nil {
		return resp.Err
	}
	return nil
}

func DeleteAllAgentsInfo() error {
	conn, err := aredis.DredisPool.Get()
	if err != nil {
		return err
	}
	defer aredis.DredisPool.Put(conn)

	resp := conn.Cmd("DEL", "agent")
	if resp.Err != nil {
		return resp.Err
	}
	return nil
}
