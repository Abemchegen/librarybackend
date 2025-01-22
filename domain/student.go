package domain

type Student struct {
	Name     string `json:"name" bson:"name"`
	SchoolID string `json:"id" bson:"id"`
}

type StudentUseCase interface {
	EnterLibrary(Student Student) (Student, ErrorResponse)
	LeaveLibrary(Student Student) (SuccessResponse, ErrorResponse)
	GetStudentActivity() ([]Activity, ErrorResponse)
	GetUniqueStudentCountPerDay() (map[string]int, ErrorResponse)
}

type StudentRepository interface {
	EnterLibrary(Student Student) (Student, error)
	LeaveLibrary(Student Student) error
	GetStudentActivity() ([]Activity, error)
	GetUniqueStudentCountPerDay() (map[string]int, error)
}
