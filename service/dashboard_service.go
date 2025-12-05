package service

import (
	"backend-kurikulum-apps/dto"
	"backend-kurikulum-apps/repository"
)

type DashboardService interface {
	GetKaprodiDashboard() (*dto.DashboardKaprodiResponse, error)
	GetDosenDashboard(dosenID string) (*dto.DashboardDosenResponse, error)
}

type dashboardService struct {
	cplRepo        repository.CPLRepository
	rpsRepo        repository.RPSRepository
	assignmentRepo repository.CPLAssignmentRepository
	userRepo       repository.UserRepository
	docRepo        repository.GeneratedDocumentRepository
}

func NewDashboardService(
	cplRepo repository.CPLRepository,
	rpsRepo repository.RPSRepository,
	assignmentRepo repository.CPLAssignmentRepository,
	userRepo repository.UserRepository,
	docRepo repository.GeneratedDocumentRepository,
) DashboardService {
	return &dashboardService{
		cplRepo:        cplRepo,
		rpsRepo:        rpsRepo,
		assignmentRepo: assignmentRepo,
		userRepo:       userRepo,
		docRepo:        docRepo,
	}
}

func (s *dashboardService) GetKaprodiDashboard() (*dto.DashboardKaprodiResponse, error) {
	totalCPL, _ := s.cplRepo.CountByStatus("")
	publishedCPL, _ := s.cplRepo.CountByStatus("published")
	draftCPL, _ := s.cplRepo.CountByStatus("draft")

	totalRPS, _ := s.rpsRepo.CountByStatus("")
	approvedRPS, _ := s.rpsRepo.CountByStatus("approved")
	pendingReview, _ := s.rpsRepo.CountByStatus("submitted")
	rejectedRPS, _ := s.rpsRepo.CountByStatus("rejected")

	activeDosen, _ := s.userRepo.CountByRole("dosen")

	activeAssignments, _ := s.assignmentRepo.CountByStatus("accepted")
	assigned, _ := s.assignmentRepo.CountByStatus("assigned")
	activeAssignments += assigned

	completedAssignments, _ := s.assignmentRepo.CountByStatus("done")

	documentsGenerated, _ := s.docRepo.CountByStatus("ready")

	return &dto.DashboardKaprodiResponse{
		TotalCPL:             totalCPL,
		PublishedCPL:         publishedCPL,
		DraftCPL:             draftCPL,
		TotalRPS:             totalRPS,
		ApprovedRPS:          approvedRPS,
		PendingReview:        pendingReview,
		RejectedRPS:          rejectedRPS,
		ActiveDosen:          activeDosen,
		ActiveAssignments:    activeAssignments,
		CompletedAssignments: completedAssignments,
		DocumentsGenerated:   documentsGenerated,
	}, nil
}

func (s *dashboardService) GetDosenDashboard(dosenID string) (*dto.DashboardDosenResponse, error) {
	totalAssignments, _ := s.assignmentRepo.CountByDosenAndStatus(dosenID, "")
	acceptedAssignments, _ := s.assignmentRepo.CountByDosenAndStatus(dosenID, "accepted")
	pendingAssignments, _ := s.assignmentRepo.CountByDosenAndStatus(dosenID, "assigned")
	completedAssignments, _ := s.assignmentRepo.CountByDosenAndStatus(dosenID, "done")

	totalRPS, _ := s.rpsRepo.CountByDosenAndStatus(dosenID, "")
	approvedRPS, _ := s.rpsRepo.CountByDosenAndStatus(dosenID, "approved")
	draftRPS, _ := s.rpsRepo.CountByDosenAndStatus(dosenID, "draft")
	submittedRPS, _ := s.rpsRepo.CountByDosenAndStatus(dosenID, "submitted")
	rejectedRPS, _ := s.rpsRepo.CountByDosenAndStatus(dosenID, "rejected")

	return &dto.DashboardDosenResponse{
		TotalAssignments:     totalAssignments,
		AcceptedAssignments:  acceptedAssignments,
		PendingAssignments:   pendingAssignments,
		CompletedAssignments: completedAssignments,
		TotalRPS:             totalRPS,
		ApprovedRPS:          approvedRPS,
		DraftRPS:             draftRPS,
		SubmittedRPS:         submittedRPS,
		RejectedRPS:          rejectedRPS,
	}, nil
}
