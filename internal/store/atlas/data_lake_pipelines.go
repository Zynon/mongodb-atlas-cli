// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated at 2023-04-12T16:00:33+01:00. Note: Manual updates are allowed, but may be overwritten.

package atlas

import (
	"time"

	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_data_lake_pipelines.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas PipelinesLister,PipelinesDescriber,PipelinesCreator,PipelinesUpdater,PipelinesDeleter,PipelineAvailableSnapshotsLister,PipelineAvailableSchedulesLister,PipelinesTriggerer,PipelinesPauser,PipelinesResumer

type PipelinesLister interface {
	Pipelines(string) ([]atlasv2.IngestionPipeline, error)
}

type PipelinesCreator interface {
	CreatePipeline(string, atlasv2.IngestionPipeline) (*atlasv2.IngestionPipeline, error)
}

type PipelinesDeleter interface {
	DeletePipeline(string, string) error
}

type PipelinesDescriber interface {
	Pipeline(string, string) (*atlasv2.IngestionPipeline, error)
}

type PipelinesUpdater interface {
	UpdatePipeline(string, string, atlasv2.IngestionPipeline) (*atlasv2.IngestionPipeline, error)
}

type PipelineAvailableSnapshotsLister interface {
	PipelineAvailableSnapshots(string, string, *time.Time, *atlas.ListOptions) (*atlasv2.PaginatedBackupSnapshot, error)
}

type PipelineAvailableSchedulesLister interface {
	PipelineAvailableSchedules(string, string) ([]atlasv2.PolicyItem, error)
}

type PipelinesTriggerer interface {
	PipelineTrigger(string, string, string) (*atlasv2.IngestionPipelineRun, error)
}

type PipelinesPauser interface {
	PipelinePause(string, string) (*atlasv2.IngestionPipeline, error)
}

type PipelinesResumer interface {
	PipelineResume(string, string) (*atlasv2.IngestionPipeline, error)
}

// Pipelines encapsulates the logic to manage different cloud providers.
func (s *Store) Pipelines(projectID string) ([]atlasv2.IngestionPipeline, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.ListPipelines(s.ctx, projectID).Execute()
	return result, err
}

// Pipeline encapsulates the logic to manage different cloud providers.
func (s *Store) Pipeline(projectID, id string) (*atlasv2.IngestionPipeline, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.GetPipeline(s.ctx, projectID, id).Execute()
	return result, err
}

// CreatePipeline encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePipeline(projectID string, opts atlasv2.IngestionPipeline) (*atlasv2.IngestionPipeline, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.CreatePipeline(s.ctx, projectID, &opts).Execute()
	return result, err
}

// UpdatePipeline encapsulates the logic to manage different cloud providers.
func (s *Store) UpdatePipeline(projectID, id string, opts atlasv2.IngestionPipeline) (*atlasv2.IngestionPipeline, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.UpdatePipeline(s.ctx, projectID, id, &opts).Execute()
	return result, err
}

// DeletePipeline encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePipeline(projectID, id string) error {
	_, _, err := s.clientv2.DataLakePipelinesApi.DeletePipeline(s.ctx, projectID, id).Execute()
	return err
}

// PipelineAvailableSchedules encapsulates the logic to manage different cloud providers.
func (s *Store) PipelineAvailableSchedules(projectID, pipelineName string) ([]atlasv2.PolicyItem, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.ListPipelineSchedules(s.ctx, projectID, pipelineName).Execute()
	return result, err
}

// PipelineAvailableSnapshots encapsulates the logic to manage different cloud providers.
func (s *Store) PipelineAvailableSnapshots(projectID, pipelineName string, completedAfter *time.Time, listOps *atlas.ListOptions) (*atlasv2.PaginatedBackupSnapshot, error) {
	request := s.clientv2.DataLakePipelinesApi.ListPipelineSnapshots(s.ctx, projectID, pipelineName)
	if completedAfter != nil {
		request = request.CompletedAfter(*completedAfter)
	}
	if listOps != nil {
		request = request.ItemsPerPage(listOps.ItemsPerPage)
		request = request.PageNum(listOps.PageNum)
	}
	result, _, err := request.Execute()
	return result, err
}

// PipelineTrigger encapsulates the logic to manage different cloud providers.
func (s *Store) PipelineTrigger(projectID, pipelineName, snapshotID string) (*atlasv2.IngestionPipelineRun, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.TriggerSnapshotIngestion(s.ctx, projectID, pipelineName,
		atlasv2.NewTriggerIngestionRequest(snapshotID)).Execute()
	return result, err
}

// PipelinePause encapsulates the logic to manage different cloud providers.
func (s *Store) PipelinePause(projectID, pipelineName string) (*atlasv2.IngestionPipeline, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.PausePipeline(s.ctx, projectID, pipelineName).Execute()
	return result, err
}

// PipelineResume encapsulates the logic to manage different cloud providers.
func (s *Store) PipelineResume(projectID, pipelineName string) (*atlasv2.IngestionPipeline, error) {
	result, _, err := s.clientv2.DataLakePipelinesApi.ResumePipeline(s.ctx, projectID, pipelineName).Execute()
	return result, err
}
