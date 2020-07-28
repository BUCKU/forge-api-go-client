package dm

import (
	"encoding/json"
	"errors"
	"fmt"
	"forge-api-go-client/oauth"
	"net/http"
	"strconv"
	"time"
)

type IssuesContainerData struct {
	Data []struct {
		ID    string `json:"id, omitempty"`
		Type  string `json:"type, omitempty"`
		Links struct {
			Self string `json:"self, omitempty"`
		} `json:"links, omitempty"`
		Attributes struct {
			CreatedAt       time.Time   `json:"created_at, omitempty"`
			SyncedAt        time.Time   `json:"synced_at, omitempty"`
			UpdatedAt       time.Time   `json:"updated_at, omitempty"`
			CloseVersion    interface{} `json:"close_version, omitempty"`
			ClosedAt        time.Time   `json:"closed_at, omitempty"`
			ClosedBy        string      `json:"closed_by, omitempty"`
			CreatedBy       string      `json:"created_by, omitempty"`
			StartingVersion int         `json:"starting_version, omitempty"`
			Title           string      `json:"title, omitempty"`
			Description     string      `json:"description, omitempty"`
			TargetUrn       string      `json:"target_urn, omitempty"`
			TargetUrnPage   interface{} `json:"target_urn_page, omitempty"`
		} `json:"attributes, omitempty"`
	} `json:"data, omitempty"`
}

type IssuesAPI struct {
	oauth.TwoLeggedAuth
	IssuesAPIPath string
}

func NewIssuesAPIWithCredentials(clientID, clientSecret string) *IssuesAPI {
	return &IssuesAPI{
		TwoLeggedAuth: oauth.NewTwoLeggedClient(clientID, clientSecret),
		// /issues/v1/containers/:containerId/  https://forge.autodesk.com/en/docs/bim360/v1/reference/http/field-issues-GET/
		IssuesAPIPath: fmt.Sprintf("/issues/v1/containers"),
	}
}

func (api *IssuesAPI) GetIssues(issueContainerId string) (result *IssuesContainerData, err error) {
	bearer, err := api.AuthenticateIfNecessary("data:read")
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s%s/%s", api.Host, api.IssuesAPIPath, issueContainerId)
	result, err = getIssues(path, bearer.AccessToken)
	return result, err
}

func getIssues(path string, token string) (result *IssuesContainerData, err error) {
	task := http.Client{}

	req, err := http.NewRequest("GET", path+"/quality-issues", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	response, err := task.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New(strconv.Itoa(response.StatusCode))
		return nil, err
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(result)
	return result, nil
}
