package harborcli

import (
	"fmt"
	"time"
)

const (
	ProjectAPIPath = "api/projects"
)

type ProjectAPI struct {
	client *HarborClient
}

// ProjectRequest holds informations that need for creating project API
type ProjectRequest struct {
	Name     string            `json:"project_name"`
	Public   *int              `json:"public"` // deprecated, reserved for project creation in replication
	Metadata map[string]string `json:"metadata"`
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64             `orm:"pk;auto;column(project_id)" json:"project_id"`
	OwnerID      int               `orm:"column(owner_id)" json:"owner_id"`
	Name         string            `orm:"column(name)" json:"name"`
	CreationTime time.Time         `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time         `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool              `orm:"column(deleted)" json:"deleted"`
	OwnerName    string            `orm:"-" json:"owner_name"`
	Togglable    bool              `orm:"-" json:"togglable"`
	Role         int               `orm:"-" json:"current_user_role_id"`
	RepoCount    int64             `orm:"-" json:"repo_count"`
	ChartCount   uint64            `orm:"-" json:"chart_count"`
	Metadata     map[string]string `orm:"-" json:"metadata"`
}

func (p *ProjectAPI) Create(project *ProjectRequest) error {
	err := p.client.authPing()
	if err != nil {
		return err
	}
	req, err := p.client.newRequest("POST", ProjectAPIPath, project)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	_, err = p.client.do(req, nil)
	return err
}

// Check if the project name user provided already exists
func (p *ProjectAPI) Check(name string) error {
	err := p.client.authPing()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s?project_name=%s", ProjectAPIPath, name)
	req, err := p.client.newRequest("HEAD", path, nil)
	if err != nil {
		return err
	}

	_, err = p.client.do(req, nil)
	return err
}

// Return specific project detail infomation
func (p *ProjectAPI) Get(id int64) (*Project, error) {
	err := p.client.authPing()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/%d", ProjectAPIPath, id)
	req, err := p.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	project := &Project{}
	_, err = p.client.do(req, project)

	return project, err
}

// Update project by id
func (p *ProjectAPI) Update(id int64, project *ProjectRequest) error {
	err := p.client.authPing()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%d", ProjectAPIPath, id)
	req, err := p.client.newRequest("PUT", path, project)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	_, err = p.client.do(req, nil)
	return err
}

// Delete project by id
func (p *ProjectAPI) Delete(id int64) error {
	err := p.client.authPing()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%d", ProjectAPIPath, id)
	req, err := p.client.newRequest("DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = p.client.do(req, nil)
	return err
}

// List projects
func (p *ProjectAPI) List(name string) ([]*Project, error) {
	err := p.client.authPing()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s?name=%s", ProjectAPIPath, name)
	req, err := p.client.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var projects []*Project
	_, err = p.client.do(req, &projects)

	return projects, err
}
