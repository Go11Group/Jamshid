package postgres

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"lesson29Main/model"
)

type WriterRepository struct {
	Db *gorm.DB
}

func NewWriterRepository(db *gorm.DB) *WriterRepository {
	return &WriterRepository{Db: db}
}

func (w *WriterRepository) InsetInto(writer model.Writer) error {
	err := w.Db.Create(&writer)
	if err.Error != nil {
		return err.Error
	}
	return nil

}
func (w *WriterRepository) Update(writer model.Writer) error {
	err := w.Db.Model(&model.Writer{}).Where("id=?").Update("first_name", "Jon")
	if err.Error != nil {
		return err.Error
	}
	return nil

}
func (w *WriterRepository) DeleteId(id int) error {
	err := w.Db.Delete(&model.Writer{}, "id=?", id)
	if err.Error != nil {
		return (err.Error)
	}
	return nil
}

func (w *WriterRepository) FindByName(firstName string) ([]model.Writer, error) {
	var writers []model.Writer
	if err := w.Db.Where("first_name = ?", firstName).Find(&writers).Error; err != nil {
		return nil, err
	}
	return writers, nil
}

func (w *WriterRepository) FindByLastName(lastName string) ([]model.Writer, error) {
	var writers []model.Writer
	err := w.Db.Where("last_name = ?", lastName).Find(&writers)
	if err.Error != nil {
		return nil, err.Error
	}
	return writers, nil
}
func (w *WriterRepository) FindByLastEmail(email string) ([]model.Writer, error) {
	var writers []model.Writer
	err := w.Db.Where("email=?", email).Find(&writers)
	if err.Error != nil {
		return nil, err.Error
	}

	return writers, nil
}
func (w *WriterRepository) FindByAgeFromAgeTo(age1, age2 int) ([]model.Writer, error) {
	writers := []model.Writer{}
	err := w.Db.Where("age >= ? and age <= ?", age1, age2).Find(&writers)
	if err.Error != nil {
		return nil, err.Error
	}
	return writers, nil

}

func (w *WriterRepository) FindByGender(gender string) ([]model.Writer, error) {
	writers := []model.Writer{}
	err := w.Db.Raw("SELECT where  gender = ?", gender).Scan(&writers)
	if err.Error != nil {
		return nil, err.Error
	}
	return writers, nil
}
func (w *WriterRepository) FindByIsEmployee(isEmployee bool) ([]model.Writer, error) {
	writers := []model.Writer{}

	err := w.Db.Where("is_employee = ?", isEmployee).Find(&writers).Error
	if err != nil {
		return nil, err
	}
	return writers, nil
}

func (w *WriterRepository) FindALl(isEmployee bool) ([]model.Writer, error) {
	writers := []model.Writer{}
	err := w.Db.Where("last_name = ?", isEmployee).Find(&writers).Error
	if err != nil {
		return nil, err
	}
	return writers, nil
}
