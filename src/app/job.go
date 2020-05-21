package app

import (
	"github.com/oojob/service-job/src/model"
)

// CreateJob :- creates a job
func (ctx *Context) CreateJob(job *model.Job) (string, error) {
	return ctx.Database.CreateJob(job)
}

// ReadJob :- read a job
func (ctx *Context) ReadJob(job *model.Job) (string, error) {
	return ctx.Database.ReadJob(job)
}

// UpdateJob :- update a job
func (ctx *Context) UpdateJob(job *model.Job) (string, error) {
	return ctx.Database.UpdateJob(job)
}

// DeleteJob :- delete a job
func (ctx *Context) DeleteJob(job *model.Job) (string, error) {
	return ctx.Database.DeleteJob(job)
}

// ReadAllJobsByCompany :- read all jobs by company
func (ctx *Context) ReadAllJobsByCompany(job *model.Job) (string, error) {
	return ctx.Database.ReadAllJobsByCompany(job)
}

// ReadAllJobs :- read all jobs
func (ctx *Context) ReadAllJobs(job *model.Job) (string, error) {
	return ctx.Database.ReadAllJobs(job)
}
