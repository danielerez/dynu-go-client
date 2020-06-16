package dynu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type DynuClient struct {
	ApiKey string
}

type DnsRecordPayload struct {
	NodeName    string `json:"nodeName"`
	RecordType  string `json:"recordType"`
	TTL         int    `json:"ttl"`
	State       bool   `json:"state"`
	Group       string `json:"group"`
	Ipv4Address string `json:"ipv4Address"`
}

type DnsRecordA struct {
	StatusCode	int32     `json:"statusCode"`
	Id          int32     `json:"id,omitempty"`
	DomainId    int32     `json:"domainId,omitempty"`
	DomainName  string    `json:"domainName,omitempty"`
	NodeName    string    `json:"nodeName,omitempty"`
	Hostname    string    `json:"hostname,omitempty"`
	RecordType  string    `json:"recordType,omitempty"`
	State       bool      `json:"state,omitempty"`
	Content     string    `json:"content,omitempty"`
	Group       string    `json:"group,omitempty"`
	Ipv4Address string    `json:"ipv4Address,omitempty"`
}

func NewDynuClient(apiKey string) *DynuClient {
	return &DynuClient{
		ApiKey: apiKey,
	}
}

func (d *DynuClient) CreateRecordA(domainId, prefix, clusterName, ip string) (*DnsRecordA, error) {
	data := DnsRecordPayload{
		NodeName:    fmt.Sprintf("%s.%s", prefix, clusterName),
		RecordType:  "A",
		TTL:         300,
		State:       true,
		Group:       "",
		Ipv4Address: ip,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.dynu.com/v2/dns/%s/record", domainId)
	body := bytes.NewReader(payloadBytes)
	respBody, err := d.postRequest(url, body)
	if err != nil {
		return nil, err
	}

	var record DnsRecordA
	err = json.Unmarshal(respBody, &record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (d *DynuClient) postRequest(url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-Key", d.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return respBody, nil
}
