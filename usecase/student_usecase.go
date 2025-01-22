package usecase

import (
	"librarybackend/domain"
)

type StudentUseCase struct {
	StudentRepository domain.StudentRepository
}

func NewStudentUseCase(StudentRepository domain.StudentRepository) domain.StudentUseCase {
	return &StudentUseCase{
		StudentRepository: StudentRepository,
	}
}

func (uc *StudentUseCase) EnterLibrary(Student domain.Student) (domain.Student, domain.ErrorResponse) {

	res, err := uc.StudentRepository.EnterLibrary(Student)

	if err != nil {
		return domain.Student{}, domain.ErrorResponse{Message: "Failed to enter library", Status: 500}
	}

	return res, domain.ErrorResponse{}
}

func (uc *StudentUseCase) LeaveLibrary(Student domain.Student) (domain.SuccessResponse, domain.ErrorResponse) {

	err := uc.StudentRepository.LeaveLibrary(Student)
	if err != nil {
		return domain.SuccessResponse{}, domain.ErrorResponse{Message: "error leaving library", Status: 404}
	}

	return domain.SuccessResponse{
		Message: "library left successfully",
		Status:  200,
	}, domain.ErrorResponse{}
}

func (uc *StudentUseCase) GetStudentActivity() ([]domain.Activity, domain.ErrorResponse) {
	Students, err := uc.StudentRepository.GetStudentActivity()

	if err != nil {
		return []domain.Activity{}, domain.ErrorResponse{Message: "activity not found", Status: 404}
	}

	return Students, domain.ErrorResponse{}
}
