package api

import (
	"context"

	protobuf "github.com/oojob/protobuf"
	job "github.com/oojob/protorepo-job-go"
	model "github.com/oojob/service-job/src/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getIdentity(identity *protobuf.Identifier) model.IdentifierModel {
	identityModel := model.IdentifierModel{
		Identifier:                identity.GetIdentifier(),
		Name:                      identity.GetName(),
		AlternateName:             identity.GetAlternateName(),
		Type:                      identity.GetType(),
		AdditionalType:            identity.GetAdditionalType(),
		Description:               identity.GetDescription(),
		DisambiguatingDescription: identity.GetDisambiguatingDescription(),
		Headline:                  identity.GetHeadline(),
		Slogan:                    identity.GetSlogan(),
	}

	return identityModel
}

func getRange(vRange *protobuf.Range) model.Range {
	rangeModel := model.Range{
		Min: vRange.GetMin(),
		Max: vRange.GetMax(),
	}

	return rangeModel
}

// func getTime(time *protobuf.Time) model.Time {
// 	timeModel := model.Time{
// 		ValidFrom:    time.GetValidFrom(),
// 		ValidThrough: time.GetValidThrough(),
// 	}

// 	return timeModel
// }

func getAddress(address *protobuf.Address) model.AddressModel {
	addressModel := model.AddressModel{
		Country:    address.GetCountry(),
		Locality:   address.GetLocality(),
		Region:     address.GetRegion(),
		PostalCode: address.GetPostalCode(),
		Street:     address.GetStreet(),
	}

	return addressModel
}

// CreateJob creates a Job
func (c *API) CreateJob(ctx context.Context, in *job.Job) (*protobuf.Id, error) {
	return nil, nil
}

// ReadJob :- read job
func (c *API) ReadJob(ctx context.Context, in *protobuf.Id) (*job.Job, error) {
	return nil, nil
}

// UpdateJob :- update a job
func (c *API) UpdateJob(ctx context.Context, in *job.Job) (*protobuf.Id, error) {
	return nil, nil
}

// DeleteJob :- delete a job
func (c *API) DeleteJob(ctx context.Context, in *protobuf.Id) (*protobuf.Id, error) {
	return nil, nil
}

// ReadAllJobsByCompany :- read all jobs by company
func (c *API) ReadAllJobsByCompany(ctx context.Context, in *protobuf.Pagination) (*job.JobAllResponse, error) {
	return nil, nil
}

// ReadAllJobs :- read all jobs
func (c *API) ReadAllJobs(ctx context.Context, in *protobuf.Pagination) (*job.JobAllResponse, error) {
	return nil, nil
}

// Check check the context
func (c *API) Check(ctx context.Context, in *protobuf.HealthCheckRequest) (*protobuf.HealthCheckResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if in.Service == "" {
		// check the server overall health status.
		return &protobuf.HealthCheckResponse{
			Status: protobuf.HealthCheckResponse_NOT_SERVING,
		}, nil
	}

	statusMap := make(map[string]protobuf.HealthCheckResponse_ServingStatus)
	if status, ok := statusMap[in.Service]; ok {
		return &protobuf.HealthCheckResponse{
			Status: status,
		}, nil
	}

	return nil, status.Errorf(codes.Internal, "unknown service")
}

// Watch watch the serving status
func (c *API) Watch(_ *protobuf.HealthCheckRequest, stream job.JobService_WatchServer) error {
	stream.Send(&protobuf.HealthCheckResponse{
		Status: protobuf.HealthCheckResponse_SERVING,
	})

	return nil
}
