/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package repository

import (
	"fmt"
	"github.com/devtron-labs/devtron/internal/sql/repository/pipelineConfig/bean/workflow/cdWorkflow"
	"github.com/devtron-labs/devtron/pkg/sql"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"time"
)

type InstalledAppVersionHistoryRepository interface {
	CreateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error)
	UpdateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error)
	GetInstalledAppVersionHistory(id int) (*InstalledAppVersionHistory, error)
	GetInstalledAppVersionHistoryByVersionId(installAppVersionId int) ([]*InstalledAppVersionHistory, error)
	GetLatestInstalledAppVersionHistory(installAppVersionId int) (*InstalledAppVersionHistory, error)
	GetLatestInstalledAppVersionHistoryByGitHash(gitHash string) (*InstalledAppVersionHistory, error)
	GetAppIdAndEnvIdWithInstalledAppVersionId(id int) (int, int, error)
	GetLatestInstalledAppVersionHistoryByInstalledAppId(installedAppId int) (*InstalledAppVersionHistory, error)
	FindPreviousInstalledAppVersionHistoryByStatus(installedAppId int, installedAppVersionHistoryId int, status []string) ([]*InstalledAppVersionHistory, error)
	UpdateInstalledAppVersionHistoryWithTxn(models []*InstalledAppVersionHistory, tx *pg.Tx) error
	GetAppStoreApplicationVersionIdByInstalledAppVersionHistoryId(installedAppVersionHistoryId int) (int, error)
	UpdateGitHash(installedAppVersionHistoryId int, gitHash string, tx *pg.Tx) error
	UpdateDeploymentHistoryMessage(installedAppVersionHistoryId int, message string) error
	GetConnection() *pg.DB
}

type InstalledAppVersionHistoryRepositoryImpl struct {
	dbConnection *pg.DB
	Logger       *zap.SugaredLogger
}

func NewInstalledAppVersionHistoryRepositoryImpl(Logger *zap.SugaredLogger, dbConnection *pg.DB) *InstalledAppVersionHistoryRepositoryImpl {
	return &InstalledAppVersionHistoryRepositoryImpl{dbConnection: dbConnection, Logger: Logger}
}

type InstalledAppVersionHistory struct {
	TableName               struct{}  `sql:"installed_app_version_history" pg:",discard_unknown_columns"`
	Id                      int       `sql:"id,pk"`
	InstalledAppVersionId   int       `sql:"installed_app_version_id,notnull"`
	ValuesYamlRaw           string    `sql:"values_yaml_raw"`
	Status                  string    `sql:"status"`
	Message                 string    `sql:"message"`
	GitHash                 string    `sql:"git_hash"`
	StartedOn               time.Time `sql:"started_on,type:timestamptz"`
	FinishedOn              time.Time `sql:"finished_on,type:timestamptz"`
	HelmReleaseStatusConfig string    `sql:"helm_release_status_config"`
	sql.AuditLog
}

func (model *InstalledAppVersionHistory) SetStartedOn() {
	model.StartedOn = time.Now()
}

func (model *InstalledAppVersionHistory) SetFinishedOn() {
	model.FinishedOn = time.Now()
}

// SetStatus sets the status of the in InstalledAppVersionHistory
// To update failed status, use MarkDeploymentFailed instead as it sets the message as well
func (model *InstalledAppVersionHistory) SetStatus(status string) {
	model.Status = status
}

// MarkDeploymentFailed marks the deployment as failed and sets the message in InstalledAppVersionHistory
func (model *InstalledAppVersionHistory) MarkDeploymentFailed(err error) {
	model.SetStatus(cdWorkflow.WorkflowFailed)
	model.Message = fmt.Sprintf("Deployment failed: %v", err)
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) CreateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error) {
	err := tx.Insert(model)
	if err != nil {
		impl.Logger.Error(err)
		return model, err
	}
	return model, nil
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) UpdateInstalledAppVersionHistory(model *InstalledAppVersionHistory, tx *pg.Tx) (*InstalledAppVersionHistory, error) {
	if tx == nil {
		err := impl.dbConnection.Update(model)
		if err != nil {
			impl.Logger.Errorw("error in updating installed app version history", "err", err, "InstalledAppVersionHistory", model)
			return nil, err
		}
		return model, nil
	} else {
		err := tx.Update(model)
		if err != nil {
			impl.Logger.Error(err)
			return model, err
		}
		return model, nil
	}
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetInstalledAppVersionHistory(id int) (*InstalledAppVersionHistory, error) {
	model := &InstalledAppVersionHistory{}
	err := impl.dbConnection.Model(model).
		Column("installed_app_version_history.*").
		Where("installed_app_version_history.id = ?", id).Select()
	return model, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetAppStoreApplicationVersionIdByInstalledAppVersionHistoryId(installedAppVersionHistoryId int) (int, error) {
	appStoreApplicationVersionId := 0
	query := "SELECT iav.app_store_application_version_id " +
		" FROM installed_app_version_history iavh " +
		" INNER JOIN installed_app_versions iav " +
		" ON iav.id=iavh.installed_app_version_id " +
		"WHERE iavh.id=%d;"

	query = fmt.Sprintf(query, installedAppVersionHistoryId)
	_, err := impl.dbConnection.Query(&appStoreApplicationVersionId, query)
	return appStoreApplicationVersionId, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetInstalledAppVersionHistoryByVersionId(installAppVersionId int) ([]*InstalledAppVersionHistory, error) {
	var model []*InstalledAppVersionHistory
	err := impl.dbConnection.Model(&model).
		Column("installed_app_version_history.*").
		Where("installed_app_version_history.installed_app_version_id = ?", installAppVersionId).
		Order("installed_app_version_history.id desc").
		Select()
	return model, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetLatestInstalledAppVersionHistory(installAppVersionId int) (*InstalledAppVersionHistory, error) {
	model := &InstalledAppVersionHistory{}
	err := impl.dbConnection.Model(model).
		Column("installed_app_version_history.*").
		Where("installed_app_version_history.installed_app_version_id = ?", installAppVersionId).
		Order("installed_app_version_history.id desc").Limit(1).
		Select()
	return model, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetLatestInstalledAppVersionHistoryByGitHash(gitHash string) (*InstalledAppVersionHistory, error) {
	model := &InstalledAppVersionHistory{}
	err := impl.dbConnection.Model(model).
		Column("installed_app_version_history.*").
		Where("installed_app_version_history.git_hash = ?", gitHash).
		Order("installed_app_version_history.id desc").Limit(1).
		Select()
	return model, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetAppIdAndEnvIdWithInstalledAppVersionId(id int) (int, int, error) {
	type appEnvId struct {
		AppId int `json:"app_id"`
		EnvId int `json:"env_id"`
	}
	//TODO: use explain analyse
	model := appEnvId{}
	query := "select ia.app_id,ia.environment_id as env_id" +
		" from installed_apps ia  " +
		" INNER JOIN installed_app_versions iav ON ia.id = iav.installed_app_id " +
		" where iav.id = ?;"
	_, err := impl.dbConnection.Query(&model, query, id)
	return model.AppId, model.EnvId, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetLatestInstalledAppVersionHistoryByInstalledAppId(installedAppId int) (*InstalledAppVersionHistory, error) {
	model := &InstalledAppVersionHistory{}
	query := `select iavh.* from installed_app_version_history iavh inner join installed_app_versions iav on iavh.installed_app_version_id=iav.id
				inner join installed_apps ia on iav.installed_app_id=ia.id where ia.id=? and iav.active=? order by iavh.id desc limit ?;`
	_, err := impl.dbConnection.Query(model, query, installedAppId, true, 1)
	if err != nil {
		impl.Logger.Errorw("error in GetLatestInstalledAppVersionHistoryByInstalledAppId", "err", err)
		return nil, err
	}
	return model, nil
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) FindPreviousInstalledAppVersionHistoryByStatus(installedAppId int, installedAppVersionHistoryId int, status []string) ([]*InstalledAppVersionHistory, error) {
	var iavr []*InstalledAppVersionHistory
	err := impl.dbConnection.
		Model(&iavr).
		Column("installed_app_version_history.*").
		Where("installed_app_version_history.installed_app_version_id in ( select id from installed_app_versions where installed_app_id in (?))", installedAppId).
		Where("installed_app_version_history.id < ?", installedAppVersionHistoryId).
		Where("installed_app_version_history.status not in (?) ", pg.In(status)).
		Order("installed_app_version_history.id DESC").
		Select()
	return iavr, err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) UpdateGitHash(installedAppVersionHistoryId int, gitHash string, tx *pg.Tx) error {
	query := "update installed_app_version_history set git_hash=? where id=?"
	_, err := tx.Exec(query, gitHash, installedAppVersionHistoryId)
	if err != nil {
		return err
	}
	return nil
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) UpdateDeploymentHistoryMessage(installedAppVersionHistoryId int, message string) error {
	_, err := impl.dbConnection.Model((*InstalledAppVersionHistory)(nil)).
		Set("message = ?", message).
		Where("id = ?", installedAppVersionHistoryId).
		Update()
	return err
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) GetConnection() *pg.DB {
	return impl.dbConnection
}

func (impl *InstalledAppVersionHistoryRepositoryImpl) UpdateInstalledAppVersionHistoryWithTxn(models []*InstalledAppVersionHistory, tx *pg.Tx) error {
	_, err := tx.Model(&models).Update()
	return err
}
