package repository

import (
	"net/http"

	"github.com/kzuabe/todolist-go-api/internal/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (repository *TaskRepository) Fetch(params model.TaskFetchParam) ([]model.Task, error) {
	tx := repository.DB.Session(&gorm.Session{})
	if userID := params.UserID; userID != "" {
		tx = tx.Where("user_id = ?", userID)
	}
	if status := params.Status; status != nil {
		tx = tx.Where("status = ?", status)
	}

	dbTasks := []Task{}
	result := tx.Find(&dbTasks)

	if err := result.Error; err != nil {
		return []model.Task{}, err
	}

	tasks := make([]model.Task, len(dbTasks))
	for i, t := range dbTasks {
		tasks[i] = t.toModel()
	}
	return tasks, nil
}

func (repository *TaskRepository) FetchByID(id string) (model.Task, error) {
	t := Task{}

	result := repository.DB.First(&t, "uuid = ?", id)

	if err := result.Error; err != nil {
		return model.Task{}, err
	}

	fetched := t.toModel()
	return fetched, nil
}

func (repository *TaskRepository) Create(task model.Task) (model.Task, error) {
	t := toDBTask(task)
	result := repository.DB.Create(&t)

	if result.Error != nil {
		e := &model.Error{StatusCode: http.StatusInternalServerError, Message: result.Error.Error()}
		return model.Task{}, e
	}

	created := t.toModel()
	return created, result.Error
}

func (repository *TaskRepository) Update(task model.Task) (model.Task, error) {
	t := Task{}
	result := repository.DB.First(&t, "uuid = ?", task.ID)

	if result.Error != nil {
		e := &model.Error{StatusCode: http.StatusInternalServerError, Message: result.Error.Error()}
		return model.Task{}, e
	}
	if t.UserID != task.UserID { // リクエストユーザーとタスクのユーザーが異なる場合
		e := &model.Error{StatusCode: http.StatusForbidden, Message: "許可されていないユーザー"}
		return model.Task{}, e
	}

	// データ更新
	t.updateFromModel(task)

	result = repository.DB.Save(&t)

	if result.Error != nil {
		e := &model.Error{StatusCode: http.StatusInternalServerError, Message: result.Error.Error()}
		return model.Task{}, e
	}

	updated := t.toModel()
	return updated, nil
}

func (repository *TaskRepository) Delete(id string, userID string) error {
	result := repository.DB.Delete(&Task{}, "uuid = ? AND user_id = ?", id, userID)
	if result.Error != nil {
		e := &model.Error{StatusCode: http.StatusInternalServerError, Message: result.Error.Error()}
		return e
	}
	return nil
}

type Task struct {
	gorm.Model
	UUID        string
	UserID      string
	Title       string
	Description string
	Status      int
}

func (task *Task) toModel() model.Task {
	t := model.Task{
		ID:          task.UUID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	return t
}

// 更新用
func (task *Task) updateFromModel(t model.Task) {
	task.Title = t.Title
	task.Description = t.Description
	task.Status = t.Status
}

func toDBTask(task model.Task) Task {
	t := Task{
		UUID:        task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
	return t
}
