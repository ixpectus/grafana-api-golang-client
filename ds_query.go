package gapi

import (
	"encoding/json"
	"time"
)

type DsDatasource struct {
	Type string `json:"type"`
	UID  string `json:"uid"`
}
type DsQuery struct {
	Datasource     DsDatasource `json:"datasource"`
	EditorMode     string       `json:"editorMode"`
	Expr           string       `json:"expr"`
	Format         string       `json:"format"`
	IntervalFactor int          `json:"intervalFactor"`
	LegendFormat   string       `json:"legendFormat"`
	Range          bool         `json:"range"`
	Exemplar       bool         `json:"exemplar"`
	RequestID      string       `json:"requestId"`
	UtcOffsetSec   int          `json:"utcOffsetSec"`
	Interval       string       `json:"interval"`
	DatasourceID   int          `json:"datasourceId"`
	IntervalMs     int          `json:"intervalMs"`
	MaxDataPoints  int          `json:"maxDataPoints"`
}

type DsRange struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
	Raw  struct {
		From *time.Time `json:"from,omitempty"`
		To   *time.Time `json:"to,omitempty"`
	} `json:"raw"`
}

type DsQueries struct {
	Queries []DsQuery `json:"queries"`
	Range   *DsRange  `json:"range"`
	From    *string   `json:"from,omitempty"`
	To      *string   `json:"to,omitempty"`
}

type Response struct {
	Status int     `json:"status"`
	Frames []Frame `json:"frames"`
}

type Frame struct {
	Schema Schema `json:"schema"`
	Data   Data   `json:"data"`
}

type Schema struct {
	Name   string  `json:"name"`
	RefID  string  `json:"refId"`
	Meta   Meta    `json:"meta"`
	Fields []Field `json:"fields"`
}

type Meta struct {
	Type                string `json:"type"`
	TypeVersion         []int  `json:"typeVersion"`
	Custom              Custom `json:"custom"`
	ExecutedQueryString string `json:"executedQueryString"`
}

type Custom struct {
	ResultType string `json:"resultType"`
}

type Field struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	TypeInfo TypeInfo    `json:"typeInfo"`
	Labels   Labels      `json:"labels"`
	Config   FieldConfig `json:"config"`
}

type TypeInfo struct {
	Frame string `json:"frame"`
}

type FieldConfig struct {
	Interval int    `json:"interval"`
	Frame    string `json:"frame"`
}

type Labels struct {
	Name                  string `json:"name"`
	CPU                   string `json:"cpu"`
	DeploymentEnvironment string `json:"deployment_environment"`
	Host                  string `json:"host"`
	Job                   string `json:"job"`
	PipelineID            string `json:"pipeline_id"`
	ServiceName           string `json:"service_name"`
}

type Data struct {
	Values [][]int64 `json:"values"`
}

type Results struct {
	Results map[string]Response `json:"results"`
}

func (c *Client) DatasourceQuery(q DsQueries) (*Results, error) {
	req, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	res := new(Results)
	if err := c.request("POST", "api/ds/query", nil, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
