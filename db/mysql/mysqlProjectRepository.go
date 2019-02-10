package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mono83/charlie/model"
	"log"
)

type mysqlProjectRepository struct {
	Conn *sql.DB
}

func NewMysqlProjectRepository(Conn *sql.DB) *mysqlProjectRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlProjectRepository{Conn: Conn}
}

func (repo *mysqlProjectRepository) fetch(query string, args ...interface{}) ([]*model.Project, error) {

	rows, err := repo.Conn.Query(query, args...)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	projects := make([]*model.Project, 0)
	for rows.Next() {
		p := new(model.Project)
		var description sql.NullString
		err := rows.Scan(&p.ID, &p.Name, &description)
		p.Description = description.String
		projects = append(projects, p)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return projects, nil
}

func (repo *mysqlProjectRepository) GetByID(id int64) (*model.Project, error) {
	projects, err := repo.fetch("SELECT id, name, description FROM `project` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return repo.getSingleResult(projects)
}

func (repo *mysqlProjectRepository) GetByName(name string) (*model.Project, error) {
	projects, err := repo.fetch("SELECT id, name, description FROM `project` WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	return repo.getSingleResult(projects)
}

func (repo *mysqlProjectRepository) getSingleResult(projects []*model.Project) (*model.Project, error) {
	if len(projects) > 1 {
		return nil, errors.New(fmt.Sprintf("Expected not more than 1 item, but was %d", len(projects)))
	}
	if len(projects) == 0 {
		return nil, nil
	}

	return projects[0], nil
}
