// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the operators client
type API interface {
	/*
	   ListOfClusterOperators Lists operators to be monitored for a cluster.*/
	ListOfClusterOperators(ctx context.Context, params *ListOfClusterOperatorsParams) (*ListOfClusterOperatorsOK, error)
	/*
	   ListOperatorProperties Lists properties for an operator.*/
	ListOperatorProperties(ctx context.Context, params *ListOperatorPropertiesParams) (*ListOperatorPropertiesOK, error)
	/*
	   ListSupportedOperators Retrieves the list of supported operators.*/
	ListSupportedOperators(ctx context.Context, params *ListSupportedOperatorsParams) (*ListSupportedOperatorsOK, error)
	/*
	   ReportMonitoredOperatorStatus Controller API to report of monitored operators.*/
	ReportMonitoredOperatorStatus(ctx context.Context, params *ReportMonitoredOperatorStatusParams) (*ReportMonitoredOperatorStatusOK, error)
}

// New creates a new operators API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for operators API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ListOfClusterOperators Lists operators to be monitored for a cluster.
*/
func (a *Client) ListOfClusterOperators(ctx context.Context, params *ListOfClusterOperatorsParams) (*ListOfClusterOperatorsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListOfClusterOperators",
		Method:             "GET",
		PathPattern:        "/v1/clusters/{cluster_id}/monitored_operators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOfClusterOperatorsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOfClusterOperatorsOK), nil

}

/*
ListOperatorProperties Lists properties for an operator.
*/
func (a *Client) ListOperatorProperties(ctx context.Context, params *ListOperatorPropertiesParams) (*ListOperatorPropertiesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListOperatorProperties",
		Method:             "GET",
		PathPattern:        "/v1/supported-operators/{operator_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOperatorPropertiesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOperatorPropertiesOK), nil

}

/*
ListSupportedOperators Retrieves the list of supported operators.
*/
func (a *Client) ListSupportedOperators(ctx context.Context, params *ListSupportedOperatorsParams) (*ListSupportedOperatorsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListSupportedOperators",
		Method:             "GET",
		PathPattern:        "/v1/supported-operators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSupportedOperatorsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSupportedOperatorsOK), nil

}

/*
ReportMonitoredOperatorStatus Controller API to report of monitored operators.
*/
func (a *Client) ReportMonitoredOperatorStatus(ctx context.Context, params *ReportMonitoredOperatorStatusParams) (*ReportMonitoredOperatorStatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReportMonitoredOperatorStatus",
		Method:             "PUT",
		PathPattern:        "/v1/clusters/{cluster_id}/monitored_operators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReportMonitoredOperatorStatusReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReportMonitoredOperatorStatusOK), nil

}
