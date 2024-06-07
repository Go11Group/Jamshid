package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type GroupRepository struct {
	Db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		Db: db,
	}
}

func (gr *GroupRepository) CreateGroup(group model.Group) error {
	_, err := gr.Db.Exec("insert into  groups(group_name,student_number) values ($1,$2)", group.GroupName, group.Student_count)
	if err != nil {
		return err
	}
	return nil
}

func (gr *GroupRepository) UpdateGroup(id string, group model.Group) error {
	_, err := gr.Db.Exec("update groups set group_name=$1,student_number=$2 where id=$3", group.GroupName, group.Student_count, id)
	if err != nil {
		return err
	}
	return nil
}

func (gr *GroupRepository) DeleteGroup(id string) error {
	_, err := gr.Db.Exec("delete  from  groups where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (gr *GroupRepository) ReadAllGroup() ([]model.Group, error) {
	rows, err := gr.Db.Query("select id,group_name,student_number from groups")
	groups := []model.Group{}
	if err != nil {
		return nil, err
	}
	group := model.Group{}
	for rows.Next() {
		err := rows.Scan(&group.Id, &group.GroupName, &group.Student_count)

		if err != nil {
			fmt.Println("has", err)
			return nil, err
		}

		groups = append(groups, group)
	}
	fmt.Println("salom", groups)

	return groups, err
}
