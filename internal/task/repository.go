package task

type TaskRepository struct {
	db any
}

func NewTaskRepository(db any) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}
