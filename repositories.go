package harborcli

import (
	"fmt"
	"time"
)

type RepositoryAPI struct {
	client *HarborClient
}

const (
	RepositoryAPIPath = "api/repositories"
)

type RepoResp struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	ProjectID    int64     `json:"project_id"`
	Description  string    `json:"description"`
	PullCount    int64     `json:"pull_count"`
	StarCount    int64     `json:"star_count"`
	TagsCount    int64     `json:"tags_count"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}

type TagResp struct {
	Name    string
	Size    int64
	Digest  string
	Author  string
	Created time.Time
}

// List : Get repositories accompany with relevant project and repo name
func (r *RepositoryAPI) List(projectID int64) ([]*RepoResp, error) {
	err := r.client.authPing()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s?project_id=%d", RepositoryAPIPath, projectID)
	req, err := r.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var repos []*RepoResp
	_, err = r.client.do(req, &repos)

	return repos, err
}

// Delete : Delete a repository
func (r *RepositoryAPI) Delete(name string) error {
	err := r.client.authPing()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", RepositoryAPIPath, name)
	req, err := r.client.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = r.client.do(req, nil)
	return err
}

// DeleteTag : Delete a tag in a repository
func (r *RepositoryAPI) DeleteTag(repoName, tagName string) error {
	err := r.client.authPing()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s/tags/%s", RepositoryAPIPath, repoName, tagName)
	req, err := r.client.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = r.client.do(req, nil)
	return err
}

// GetTags : Get tags of the repository
func (r *RepositoryAPI) GetTags(name string) ([]*TagResp, error) {
	err := r.client.authPing()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/tags", RepositoryAPIPath, name)
	req, err := r.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var tags []*TagResp
	_, err = r.client.do(req, &tags)

	return tags, err
}
