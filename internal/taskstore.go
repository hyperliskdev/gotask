

type Task struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Tags []string `json:"tags"`
	Due time.Time `json:"due"`

}

func New() *TaskStore

//Create a new task
func (ts *TaskStore) CreateTask(text string, tags[]string, due time.Time) int

//Get a task by its id
func (ts *TaskStore) GetTask(id int) (Task, error)

// 
func (ts *TaskStore) DeleteTask(id int) error



