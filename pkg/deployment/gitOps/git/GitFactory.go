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

package git

import (
	"crypto/tls"
	"fmt"
	"github.com/devtron-labs/devtron/api/bean/gitOps"
	"github.com/devtron-labs/devtron/pkg/deployment/gitOps/config"
	"github.com/devtron-labs/devtron/pkg/deployment/gitOps/git/adapter"
	"github.com/devtron-labs/devtron/util"
	"github.com/xanzy/go-gitlab"
	"go.uber.org/zap"
	"time"
)

type GitFactory struct {
	Client       GitOpsClient
	GitOpsHelper *GitOpsHelper
	logger       *zap.SugaredLogger
}

func (factory *GitFactory) Reload(gitOpsConfigReadService config.GitOpsConfigReadService) error {
	var err error
	start := time.Now()
	defer func() {
		util.TriggerGitOpsMetrics("Reload", "GitService", start, err)
	}()
	factory.logger.Infow("reloading gitops details")
	cfg, err := GetGitConfig(gitOpsConfigReadService)
	if err != nil {
		return err
	}
	factory.GitOpsHelper.SetAuth(cfg.GetAuth())
	client, err := NewGitOpsClient(cfg, factory.logger, factory.GitOpsHelper)
	if err != nil {
		return err
	}
	factory.Client = client
	factory.logger.Infow(" gitops details reload success")
	return nil
}

func (factory *GitFactory) GetGitLabGroupPath(gitOpsConfig *gitOps.GitOpsConfigDto) (string, error) {
	start := time.Now()
	var err error
	defer func() {
		util.TriggerGitOpsMetrics("GetGitLabGroupPath", "GitOpsHelper", start, err)
	}()

	var tlsConfig *tls.Config
	if gitOpsConfig.TLSConfig != nil {
		tlsConfig, err = util.GetTlsConfig(gitOpsConfig.TLSConfig.TLSKeyData, gitOpsConfig.TLSConfig.TLSCertData, gitOpsConfig.TLSConfig.CaData, GIT_TLS_DIR)
		if err != nil {
			factory.logger.Errorw("error in getting tls config", "err", err)
			return "", err
		}
	}

	gitLabClient, err := CreateGitlabClient(gitOpsConfig.Host, gitOpsConfig.Token, tlsConfig)
	if err != nil {
		factory.logger.Errorw("error in creating gitlab client", "err", err)
		return "", err
	}
	group, _, err := gitLabClient.Groups.GetGroup(gitOpsConfig.GitLabGroupId, &gitlab.GetGroupOptions{})
	if err != nil {
		factory.logger.Errorw("error in fetching gitlab group name", "err", err, "gitLab groupID", gitOpsConfig.GitLabGroupId)
		return "", err
	}
	if group == nil {
		factory.logger.Errorw("no matching groups found for gitlab", "gitLab groupID", gitOpsConfig.GitLabGroupId, "err", err)
		return "", fmt.Errorf("no matching groups found for gitlab group ID : %s", gitOpsConfig.GitLabGroupId)
	}
	return group.FullPath, nil
}

func (factory *GitFactory) NewClientForValidation(gitOpsConfig *gitOps.GitOpsConfigDto) (GitOpsClient, *GitOpsHelper, error) {
	start := time.Now()
	var err error
	defer func() {
		util.TriggerGitOpsMetrics("NewClientForValidation", "GitOpsHelper", start, err)
	}()
	cfg := adapter.ConvertGitOpsConfigToGitConfig(gitOpsConfig)
	//factory.GitOpsHelper.SetAuth(cfg.GetAuth())
	gitOpsHelper := NewGitOpsHelperImpl(cfg.GetAuth(), factory.logger, cfg.GetTLSConfig(), gitOpsConfig.EnableTLSVerification)

	client, err := NewGitOpsClient(cfg, factory.logger, gitOpsHelper)
	if err != nil {
		return client, gitOpsHelper, err
	}

	//factory.Client = client
	factory.logger.Infow("client changed successfully", "cfg", cfg)
	return client, gitOpsHelper, nil
}

func NewGitFactory(logger *zap.SugaredLogger, gitOpsConfigReadService config.GitOpsConfigReadService) (*GitFactory, error) {
	cfg, err := GetGitConfig(gitOpsConfigReadService)
	if err != nil {
		return nil, err
	}
	gitOpsHelper := NewGitOpsHelperImpl(cfg.GetAuth(), logger, cfg.GetTLSConfig(), cfg.EnableTLSVerification)
	client, err := NewGitOpsClient(cfg, logger, gitOpsHelper)
	if err != nil {
		logger.Errorw("error in creating gitOps client", "err", err, "gitProvider", cfg.GitProvider)
	}
	return &GitFactory{
		Client:       client,
		logger:       logger,
		GitOpsHelper: gitOpsHelper,
	}, nil
}
