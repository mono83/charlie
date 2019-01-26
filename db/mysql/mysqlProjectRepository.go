package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mono83/charlie/db"
	"github.com/mono83/charlie/model"
	"log"
)

type mysqlProjectRepository struct {
	Conn *sql.DB
}

func NewMysqlProjectRepository(Conn *sql.DB) db.ProjectRepository {
	return &mysqlProjectRepository{Conn: Conn}
}

func (r *mysqlProjectRepository) fetch(query string, args ...interface{}) ([]*model.Project, error) {

	rows, err := r.Conn.Query(query, args...)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	projects := make([]*model.Project, 0)
	for rows.Next() {
		p := new(model.Project)
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		projects = append(projects, p)
		if err != nil {
			log.Fatal(err)
			return projects, err
		}
	}

	return projects, nil
}

func (r *mysqlProjectRepository) GetByID(id int64) (*model.Project, error) {
	projects, err := r.fetch("SELECT id, name, description FROM `project` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return r.getSingleResult(projects)
}

func (r *mysqlProjectRepository) GetByName(name string) (*model.Project, error) {
	projects, err := r.fetch("SELECT id, name, description FROM `project` WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	return r.getSingleResult(projects)
}

func (r *mysqlProjectRepository) getSingleResult(projects []*model.Project) (*model.Project, error) {
	if len(projects) > 1 {
		return nil, errors.New(fmt.Sprintf("Expected not more than 1 item, but was %d", len(projects)))
	}
	if len(projects) == 0 {
		return nil, nil
	}

	return projects[0], nil
}
