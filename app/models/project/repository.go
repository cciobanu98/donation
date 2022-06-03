package project

import (
	"github.com/jinzhu/gorm"
)

// Repository encapsulates the logic to access projects from the data source.
type Repository interface {
	// Get returns the project with the specified project ID.
	Get (id string) (Project, error)
	// GetAll return all projects
	GetAll () ([]Project, error)
	// Count returns the number of projects.
	Count () (uint, error)
	// Create saves a new project in the storage.
	Create (project *Project) error
	// Delete a record
	Delete (id string) error
}

type repository struct {
	db 		*gorm.DB
}

var repo repository;

// NewRepository creates a new project repository
func NewRepository(db *gorm.DB) Repository {
	repo = repository{db}
	return repo
}

func (r repository) Get(id string) (Project, error) {
	var project Project
	err := repo.db.Where("id = ?", id).First(&project).Error
	if (err != nil) {
		return Project{}, err
	}
	return project, nil
}

func (r repository) GetAll() ([]Project, error) {
	var projects []Project
	err := repo.db.Find(&projects).Error
	if (err != nil) {
		return nil, err
	}
	return projects, nil
}

func (r repository) Count() (uint, error) {
	return 0, nil
}

func (r repository) Create(project *Project) error {
	err := repo.db.Create(project).Error
	return err
}

func (r repository) Delete(id string) error {
	var project Project
	err := repo.db.Where("id = ?", id).First(&project).Error
	if err != nil {
		return err
	}
	repo.db.Delete(&project)
	return nil
}