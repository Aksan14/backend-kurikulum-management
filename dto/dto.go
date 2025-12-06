package dto

import "time"

// ============ AUTH DTOs ============

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64        `json:"expires_in"`
	TokenType    string       `json:"token_type"`
	User         UserResponse `json:"user"`
}

type RegisterRequest struct {
	Nama     string  `json:"nama" binding:"required,min=2"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8"`
	Role     string  `json:"role" binding:"required,oneof=kaprodi dosen"`
	NIP      *string `json:"nip"`
	Phone    *string `json:"phone"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// ============ USER DTOs ============

type UserResponse struct {
	ID        string     `json:"id"`
	Nama      string     `json:"nama"`
	Email     string     `json:"email"`
	NIP       *string    `json:"nip"`
	Role      string     `json:"role"`
	Status    string     `json:"status"`
	Phone     *string    `json:"phone"`
	AvatarURL *string    `json:"avatar_url"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time  `json:"created_at"`
}

type UpdateUserRequest struct {
	Nama   string  `json:"nama" binding:"omitempty,min=2"`
	Email  string  `json:"email" binding:"omitempty,email"`
	Role   string  `json:"role" binding:"omitempty,oneof=kaprodi dosen"`
	NIP    *string `json:"nip"`
	Phone  *string `json:"phone"`
	Status string  `json:"status" binding:"omitempty,oneof=active inactive"`
}

type UpdateProfileRequest struct {
	Nama      string  `json:"nama" binding:"omitempty,min=2"`
	Phone     *string `json:"phone"`
	AvatarURL *string `json:"avatar_url"`
}

type UserListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Search    string `form:"search"`
	Role      string `form:"role" binding:"omitempty,oneof=kaprodi dosen"`
	Status    string `form:"status" binding:"omitempty,oneof=active inactive"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=nama email created_at"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

// ============ CPL DTOs ============

type CPLRequest struct {
	Kode      string  `json:"kode" binding:"required,min=2,max=20"`
	Nama      string  `json:"nama" binding:"required,min=5,max=255"`
	Deskripsi *string `json:"deskripsi"`
	Status    string  `json:"status" binding:"omitempty,oneof=draft published archived"`
}

type CPLResponse struct {
	ID        string        `json:"id"`
	Kode      string        `json:"kode"`
	Nama      string        `json:"nama"`
	Deskripsi *string       `json:"deskripsi"`
	Status    string        `json:"status"`
	Version   int           `json:"version"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedBy string        `json:"created_by"`
	Creator   *UserResponse `json:"creator,omitempty"`
}

type CPLListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Search    string `form:"search"`
	Status    string `form:"status" binding:"omitempty,oneof=draft published archived"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=kode nama created_at"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type CPLStatusUpdateRequest struct {
	Status string `json:"status" binding:"required,oneof=draft published archived"`
}

type CPLStatisticsResponse struct {
	TotalCPL  int64 `json:"total_cpl"`
	Published int64 `json:"published"`
	Draft     int64 `json:"draft"`
	Archived  int64 `json:"archived"`
}

// ============ MATA KULIAH DTOs ============

// ============ MATA KULIAH DTOs ============

type CreateMataKuliahRequest struct {
	Kode            string   `json:"kode" binding:"required,min=3,max=20"`
	Nama            string   `json:"nama" binding:"required,max=255"`
	SKS             int      `json:"sks" binding:"required,min=1,max=6"`
	Semester        int      `json:"semester" binding:"required,min=1,max=8"`
	Jenis           string   `json:"jenis" binding:"omitempty,oneof=wajib pilihan"`
	Deskripsi       *string  `json:"deskripsi"`
	Prasyarat       []string `json:"prasyarat"`
	Status          string   `json:"status" binding:"omitempty,oneof=aktif nonaktif"`
	DosenPengampuID *string  `json:"dosen_pengampu_id" binding:"omitempty,uuid"`
	KoordinatorID   *string  `json:"koordinator_id" binding:"omitempty,uuid"`
}

type UpdateMataKuliahRequest struct {
	Kode            *string  `json:"kode" binding:"omitempty,min=3,max=20"`
	Nama            *string  `json:"nama" binding:"omitempty,max=255"`
	SKS             *int     `json:"sks" binding:"omitempty,min=1,max=6"`
	Semester        *int     `json:"semester" binding:"omitempty,min=1,max=8"`
	Jenis           *string  `json:"jenis" binding:"omitempty,oneof=wajib pilihan"`
	Deskripsi       *string  `json:"deskripsi"`
	Prasyarat       []string `json:"prasyarat"`
	Status          *string  `json:"status" binding:"omitempty,oneof=aktif nonaktif"`
	DosenPengampuID *string  `json:"dosen_pengampu_id" binding:"omitempty,uuid"`
	KoordinatorID   *string  `json:"koordinator_id" binding:"omitempty,uuid"`
}

type MataKuliahResponse struct {
	ID              string        `json:"id"`
	Kode            string        `json:"kode"`
	Nama            string        `json:"nama"`
	SKS             int           `json:"sks"`
	Semester        int           `json:"semester"`
	Jenis           string        `json:"jenis"`
	Deskripsi       *string       `json:"deskripsi"`
	Prasyarat       []string      `json:"prasyarat"`
	Status          string        `json:"status"`
	DosenPengampuID *string       `json:"dosen_pengampu_id"`
	KoordinatorID   *string       `json:"koordinator_id"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	CreatedBy       string        `json:"created_by"`
	DosenPengampu   *UserResponse `json:"dosen_pengampu,omitempty"`
	Koordinator     *UserResponse `json:"koordinator,omitempty"`
}

type MataKuliahListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Search    string `form:"search"`
	Semester  int    `form:"semester" binding:"omitempty,min=1,max=8"`
	Jenis     string `form:"jenis" binding:"omitempty,oneof=wajib pilihan"`
	Status    string `form:"status" binding:"omitempty,oneof=aktif nonaktif dihapus"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=kode nama sks semester created_at"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

// AssignDosenRequest untuk assign/update dosen pengampu atau koordinator
type AssignDosenRequest struct {
	DosenPengampuID *string `json:"dosen_pengampu_id" binding:"omitempty,uuid"`
	KoordinatorID   *string `json:"koordinator_id" binding:"omitempty,uuid"`
}

// UnassignDosenRequest untuk unassign dosen dari mata kuliah
type UnassignDosenRequest struct {
	Type string `json:"type" binding:"required,oneof=pengampu koordinator all"`
}

// ============ CPL ASSIGNMENT DTOs ============

type CPLAssignmentRequest struct {
	CPLID        string     `json:"cpl_id" binding:"required,uuid"`
	DosenID      string     `json:"dosen_id" binding:"required,uuid"`
	MataKuliah   *string    `json:"mata_kuliah"`
	MataKuliahID *string    `json:"mata_kuliah_id" binding:"omitempty,uuid"`
	Deadline     *time.Time `json:"deadline"`
	Catatan      *string    `json:"catatan"`
}

type CPLAssignmentResponse struct {
	ID              string              `json:"id"`
	CPLID           string              `json:"cpl_id"`
	DosenID         string              `json:"dosen_id"`
	MataKuliah      *string             `json:"mata_kuliah"`
	MataKuliahID    *string             `json:"mata_kuliah_id"`
	Deadline        *time.Time          `json:"deadline"`
	Status          string              `json:"status"`
	Catatan         *string             `json:"catatan"`
	RejectionReason *string             `json:"rejection_reason"`
	AssignedBy      string              `json:"assigned_by"`
	AssignedAt      time.Time           `json:"assigned_at"`
	ResponseAt      *time.Time          `json:"response_at"`
	CompletedAt     *time.Time          `json:"completed_at"`
	CPL             *CPLResponse        `json:"cpl,omitempty"`
	Dosen           *UserResponse       `json:"dosen,omitempty"`
	MataKuliahRef   *MataKuliahResponse `json:"mata_kuliah_ref,omitempty"`
	Assigner        *UserResponse       `json:"assigner,omitempty"`
}

type CPLAssignmentListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	CPLID     string `form:"cpl_id"`
	DosenID   string `form:"dosen_id"`
	Status    string `form:"status" binding:"omitempty,oneof=assigned accepted rejected completed"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=assigned_at deadline status"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type CPLAssignmentStatusRequest struct {
	Status          string  `json:"status" binding:"required,oneof=accepted rejected completed"`
	RejectionReason *string `json:"rejection_reason"`
	Catatan         *string `json:"catatan"`
}

// ============ RPS DTOs ============

type BobotNilaiRequest struct {
	Tugas     int `json:"tugas" binding:"omitempty,min=0,max=100"`
	UTS       int `json:"uts" binding:"omitempty,min=0,max=100"`
	UAS       int `json:"uas" binding:"omitempty,min=0,max=100"`
	Kehadiran int `json:"kehadiran" binding:"omitempty,min=0,max=100"`
	Praktikum int `json:"praktikum" binding:"omitempty,min=0,max=100"`
}

type RPSRequest struct {
	MataKuliahID  string            `json:"mata_kuliah_id" binding:"required,uuid"`
	TahunAkademik string            `json:"tahun_akademik" binding:"required"`
	Deskripsi     *string           `json:"deskripsi"`
	Tujuan        *string           `json:"tujuan"`
	Metode        []string          `json:"metode"`
	BobotNilai    BobotNilaiRequest `json:"bobot_nilai" binding:"required"`
}

type RPSResponse struct {
	ID                  string                           `json:"id"`
	MataKuliahID        string                           `json:"mata_kuliah_id"`
	MataKuliahNama      string                           `json:"mata_kuliah_nama"`
	KodeMK              string                           `json:"kode_mk"`
	SKS                 int                              `json:"sks"`
	Semester            int                              `json:"semester"`
	TahunAkademik       string                           `json:"tahun_akademik"`
	DosenID             string                           `json:"dosen_id"`
	DosenNama           string                           `json:"dosen_nama"`
	Deskripsi           *string                          `json:"deskripsi"`
	Tujuan              *string                          `json:"tujuan"`
	Metode              []string                         `json:"metode"`
	BobotNilai          BobotNilaiRequest                `json:"bobot_nilai"`
	Status              string                           `json:"status"`
	CreatedAt           time.Time                        `json:"created_at"`
	UpdatedAt           time.Time                        `json:"updated_at"`
	SubmittedAt         *time.Time                       `json:"submitted_at"`
	ReviewedAt          *time.Time                       `json:"reviewed_at"`
	PublishedAt         *time.Time                       `json:"published_at"`
	ReviewedBy          *string                          `json:"reviewed_by"`
	ReviewNotes         *string                          `json:"review_notes"`
	CPMK                []RPSCPMKResponse                `json:"cpmk,omitempty"`
	RencanaPembelajaran []RPSRencanaPembelajaranResponse `json:"rencana_pembelajaran,omitempty"`
	BahanBacaan         []RPSBahanBacaanResponse         `json:"bahan_bacaan,omitempty"`
	Evaluasi            []RPSEvaluasiResponse            `json:"evaluasi,omitempty"`
	Dosen               *UserResponse                    `json:"dosen,omitempty"`
	MataKuliah          *MataKuliahResponse              `json:"mata_kuliah,omitempty"`
}

type RPSListRequest struct {
	Page          int    `form:"page" binding:"omitempty,min=1"`
	Limit         int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Search        string `form:"search"`
	MataKuliahID  string `form:"mata_kuliah_id"`
	DosenID       string `form:"dosen_id"`
	Status        string `form:"status" binding:"omitempty,oneof=draft submitted approved rejected published"`
	TahunAkademik string `form:"tahun_akademik"`
	Semester      int    `form:"semester" binding:"omitempty,min=1,max=8"`
	SortBy        string `form:"sort_by" binding:"omitempty,oneof=mata_kuliah_nama created_at updated_at"`
	SortOrder     string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type RPSStatusUpdateRequest struct {
	Status      string  `json:"status" binding:"required,oneof=submitted approved rejected published"`
	ReviewNotes *string `json:"review_notes"`
}

// ============ RPS CPMK DTOs ============

type RPSCPMKRequest struct {
	Kode      string   `json:"kode" binding:"required"`
	Deskripsi string   `json:"deskripsi" binding:"required"`
	CPLIDs    []string `json:"cpl_ids"`
	Urutan    int      `json:"urutan" binding:"required,min=1"`
}

type RPSCPMKResponse struct {
	ID        string    `json:"id"`
	RPSID     string    `json:"rps_id"`
	Kode      string    `json:"kode"`
	Deskripsi string    `json:"deskripsi"`
	CPLIDs    []string  `json:"cpl_ids"`
	Urutan    int       `json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
}

// ============ RPS RENCANA PEMBELAJARAN DTOs ============

type RPSRencanaPembelajaranRequest struct {
	Pertemuan int      `json:"pertemuan" binding:"required,min=1,max=16"`
	Topik     string   `json:"topik" binding:"required"`
	SubTopik  []string `json:"sub_topik"`
	Metode    *string  `json:"metode"`
	Waktu     *int     `json:"waktu"`
	CPMKIDs   []string `json:"cpmk_ids"`
	Materi    *string  `json:"materi"`
}

type RPSRencanaPembelajaranResponse struct {
	ID        string    `json:"id"`
	RPSID     string    `json:"rps_id"`
	Pertemuan int       `json:"pertemuan"`
	Topik     string    `json:"topik"`
	SubTopik  []string  `json:"sub_topik"`
	Metode    *string   `json:"metode"`
	Waktu     *int      `json:"waktu"`
	CPMKIDs   []string  `json:"cpmk_ids"`
	Materi    *string   `json:"materi"`
	CreatedAt time.Time `json:"created_at"`
}

// ============ RPS BAHAN BACAAN DTOs ============

type RPSBahanBacaanRequest struct {
	Judul   string  `json:"judul" binding:"required"`
	Penulis *string `json:"penulis"`
	Tahun   *int    `json:"tahun"`
	Jenis   *string `json:"jenis" binding:"omitempty,oneof=buku jurnal artikel website modul"`
	URL     *string `json:"url"`
	ISBN    *string `json:"isbn"`
	Halaman *string `json:"halaman"`
	Urutan  *int    `json:"urutan"`
}

type RPSBahanBacaanResponse struct {
	ID        string    `json:"id"`
	RPSID     string    `json:"rps_id"`
	Judul     string    `json:"judul"`
	Penulis   *string   `json:"penulis"`
	Tahun     *int      `json:"tahun"`
	Jenis     *string   `json:"jenis"`
	URL       *string   `json:"url"`
	ISBN      *string   `json:"isbn"`
	Halaman   *string   `json:"halaman"`
	Urutan    *int      `json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
}

// ============ RPS EVALUASI DTOs ============

type RPSEvaluasiRequest struct {
	Jenis             string  `json:"jenis" binding:"required"`
	Bobot             int     `json:"bobot" binding:"required,min=0,max=100"`
	Deskripsi         *string `json:"deskripsi"`
	MingguPelaksanaan []int   `json:"minggu_pelaksanaan"`
	KriteriaPenilaian *string `json:"kriteria_penilaian"`
	RubrikPenilaian   *string `json:"rubrik_penilaian"`
}

type RPSEvaluasiResponse struct {
	ID                string    `json:"id"`
	RPSID             string    `json:"rps_id"`
	Jenis             string    `json:"jenis"`
	Bobot             int       `json:"bobot"`
	Deskripsi         *string   `json:"deskripsi"`
	MingguPelaksanaan []int     `json:"minggu_pelaksanaan"`
	KriteriaPenilaian *string   `json:"kriteria_penilaian"`
	RubrikPenilaian   *string   `json:"rubrik_penilaian"`
	CreatedAt         time.Time `json:"created_at"`
}

// ============ NOTIFICATION DTOs ============

type NotificationResponse struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Message     string     `json:"message"`
	Type        string     `json:"type"`
	IsRead      bool       `json:"is_read"`
	ActionURL   *string    `json:"action_url"`
	RelatedID   *string    `json:"related_id"`
	RelatedType *string    `json:"related_type"`
	Priority    string     `json:"priority"`
	CreatedAt   time.Time  `json:"created_at"`
	ReadAt      *time.Time `json:"read_at"`
}

type NotificationListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	UserID    string `form:"-"` // Set from context
	Type      string `form:"type" binding:"omitempty,oneof=assignment approval rejection document info deadline system"`
	IsRead    *bool  `form:"is_read"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

type CreateNotificationRequest struct {
	UserID      string  `json:"user_id" binding:"required,uuid"`
	Title       string  `json:"title" binding:"required"`
	Message     string  `json:"message" binding:"required"`
	Type        string  `json:"type" binding:"required,oneof=assignment approval rejection document info deadline system"`
	ActionURL   *string `json:"action_url"`
	RelatedID   *string `json:"related_id"`
	RelatedType *string `json:"related_type"`
	Priority    string  `json:"priority" binding:"omitempty,oneof=low normal high urgent"`
}

// ============ DOCUMENT TEMPLATE DTOs ============

type DocumentTemplateRequest struct {
	Nama      string   `json:"nama" binding:"required"`
	Deskripsi *string  `json:"deskripsi"`
	Sections  []string `json:"sections" binding:"required"`
	FileURL   *string  `json:"file_url"`
	Version   string   `json:"version"`
	IsActive  bool     `json:"is_active"`
}

type DocumentTemplateResponse struct {
	ID        string        `json:"id"`
	Nama      string        `json:"nama"`
	Deskripsi *string       `json:"deskripsi"`
	Sections  []string      `json:"sections"`
	FileURL   *string       `json:"file_url"`
	Version   string        `json:"version"`
	IsActive  bool          `json:"is_active"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	CreatedBy string        `json:"created_by"`
	Creator   *UserResponse `json:"creator,omitempty"`
}

type DocumentTemplateListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Search    string `form:"search"`
	IsActive  *bool  `form:"is_active"`
	SortBy    string `form:"sort_by" binding:"omitempty,oneof=nama created_at updated_at"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

// ============ GENERATED DOCUMENT DTOs ============

type GenerateDocumentRequest struct {
	TemplateID     string                 `json:"template_id" binding:"required,uuid"`
	Tahun          string                 `json:"tahun" binding:"required"`
	Sections       []string               `json:"sections" binding:"required"`
	FileType       string                 `json:"file_type" binding:"required,oneof=docx pdf xlsx"`
	GenerationData map[string]interface{} `json:"generation_data"`
}

type GeneratedDocumentResponse struct {
	ID             string                    `json:"id"`
	TemplateID     string                    `json:"template_id"`
	TemplateName   string                    `json:"template_name"`
	Tahun          string                    `json:"tahun"`
	Status         string                    `json:"status"`
	FileURL        *string                   `json:"file_url"`
	FileType       string                    `json:"file_type"`
	FileSize       *int64                    `json:"file_size"`
	Sections       []string                  `json:"sections"`
	GenerationData map[string]interface{}    `json:"generation_data"`
	Progress       int                       `json:"progress"`
	ErrorMessage   *string                   `json:"error_message"`
	CreatedAt      time.Time                 `json:"created_at"`
	CompletedAt    *time.Time                `json:"completed_at"`
	CreatedBy      string                    `json:"created_by"`
	Template       *DocumentTemplateResponse `json:"template,omitempty"`
}

type GeneratedDocumentListRequest struct {
	Page       int    `form:"page" binding:"omitempty,min=1"`
	Limit      int    `form:"limit" binding:"omitempty,min=1,max=100"`
	TemplateID string `form:"template_id"`
	Status     string `form:"status" binding:"omitempty,oneof=processing ready failed archived"`
	Tahun      string `form:"tahun"`
	FileType   string `form:"file_type" binding:"omitempty,oneof=docx pdf xlsx"`
	SortBy     string `form:"sort_by" binding:"omitempty,oneof=created_at completed_at"`
	SortOrder  string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

// ============ DASHBOARD DTOs ============

type DashboardKaprodiResponse struct {
	TotalCPL             int64 `json:"total_cpl"`
	PublishedCPL         int64 `json:"published_cpl"`
	DraftCPL             int64 `json:"draft_cpl"`
	TotalRPS             int64 `json:"total_rps"`
	ApprovedRPS          int64 `json:"approved_rps"`
	PendingReview        int64 `json:"pending_review"`
	RejectedRPS          int64 `json:"rejected_rps"`
	ActiveDosen          int64 `json:"active_dosen"`
	ActiveAssignments    int64 `json:"active_assignments"`
	CompletedAssignments int64 `json:"completed_assignments"`
	DocumentsGenerated   int64 `json:"documents_generated"`
}

type DashboardDosenResponse struct {
	TotalAssignments     int64 `json:"total_assignments"`
	AcceptedAssignments  int64 `json:"accepted_assignments"`
	PendingAssignments   int64 `json:"pending_assignments"`
	CompletedAssignments int64 `json:"completed_assignments"`
	TotalRPS             int64 `json:"total_rps"`
	ApprovedRPS          int64 `json:"approved_rps"`
	DraftRPS             int64 `json:"draft_rps"`
	SubmittedRPS         int64 `json:"submitted_rps"`
	RejectedRPS          int64 `json:"rejected_rps"`
}

// ============ FILE DTOs ============

type FileUploadResponse struct {
	ID           string    `json:"id"`
	OriginalName string    `json:"original_name"`
	StoredName   string    `json:"stored_name"`
	FilePath     string    `json:"file_path"`
	FileSize     int64     `json:"file_size"`
	MimeType     string    `json:"mime_type"`
	FileType     string    `json:"file_type"`
	URL          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
}

// ============ AUDIT LOG DTOs ============

type AuditLogResponse struct {
	ID        string                 `json:"id"`
	UserID    *string                `json:"user_id"`
	Action    string                 `json:"action"`
	TableName string                 `json:"table_name"`
	RecordID  *string                `json:"record_id"`
	OldValues map[string]interface{} `json:"old_values"`
	NewValues map[string]interface{} `json:"new_values"`
	IPAddress *string                `json:"ip_address"`
	UserAgent *string                `json:"user_agent"`
	CreatedAt time.Time              `json:"created_at"`
	User      *UserResponse          `json:"user,omitempty"`
}

type AuditLogListRequest struct {
	Page      int    `form:"page" binding:"omitempty,min=1"`
	Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
	UserID    string `form:"user_id"`
	Action    string `form:"action"`
	TableName string `form:"table_name"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	SortOrder string `form:"sort_order" binding:"omitempty,oneof=asc desc"`
}

// ============ SYSTEM SETTINGS DTOs ============

type SystemSettingRequest struct {
	Key         string      `json:"key" binding:"required"`
	Value       interface{} `json:"value" binding:"required"`
	Description *string     `json:"description"`
	Category    string      `json:"category"`
	IsPublic    bool        `json:"is_public"`
}

type SystemSettingResponse struct {
	ID          string      `json:"id"`
	Key         string      `json:"key"`
	Value       interface{} `json:"value"`
	Description *string     `json:"description"`
	Category    string      `json:"category"`
	IsPublic    bool        `json:"is_public"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// ============ ADDITIONAL DTOs ============

type CreateUserRequest struct {
	Nama     string  `json:"nama" binding:"required,min=2"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6"`
	NIP      *string `json:"nip"`
	Role     string  `json:"role" binding:"required,oneof=kaprodi dosen"`
	Phone    *string `json:"phone"`
}

type CreateCPLRequest struct {
	Kode      string  `json:"kode" binding:"required,min=2,max=20"`
	Nama      string  `json:"nama" binding:"required,min=5"`
	Deskripsi *string `json:"deskripsi"`
}

type UpdateCPLRequest struct {
	Kode      string  `json:"kode" binding:"omitempty,min=2,max=20"`
	Nama      string  `json:"nama" binding:"omitempty,min=5"`
	Deskripsi *string `json:"deskripsi"`
	Status    string  `json:"status" binding:"omitempty,oneof=draft published archived"`
}

type UpdateCPLStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=draft published archived"`
}

type CreateCPLAssignmentRequest struct {
	CPLID        string     `json:"cpl_id" binding:"required"`
	DosenID      string     `json:"dosen_id" binding:"required"`
	MataKuliah   *string    `json:"mata_kuliah"`
	MataKuliahID *string    `json:"mata_kuliah_id"`
	Deadline     *time.Time `json:"deadline"`
	Catatan      *string    `json:"catatan"`
}

type UpdateCPLAssignmentStatusRequest struct {
	Status          string  `json:"status" binding:"required,oneof=accepted rejected completed cancelled"`
	Catatan         *string `json:"catatan"`
	RejectionReason *string `json:"rejection_reason"`
}

type CreateRPSRequest struct {
	MataKuliahID        string   `json:"mata_kuliah_id" binding:"required"`
	TahunAjaran         string   `json:"tahun_ajaran" binding:"required"`
	SemesterType        string   `json:"semester_type" binding:"required,oneof=ganjil genap"`
	DeskripsiMK         *string  `json:"deskripsi_mk"`
	CapaianPembelajaran *string  `json:"capaian_pembelajaran"`
	MetodePembelajaran  []string `json:"metode_pembelajaran"`
	MediaPembelajaran   []string `json:"media_pembelajaran"`
}

type UpdateRPSRequest struct {
	TahunAjaran         string   `json:"tahun_ajaran"`
	SemesterType        string   `json:"semester_type" binding:"omitempty,oneof=ganjil genap"`
	DeskripsiMK         *string  `json:"deskripsi_mk"`
	CapaianPembelajaran *string  `json:"capaian_pembelajaran"`
	MetodePembelajaran  []string `json:"metode_pembelajaran"`
	MediaPembelajaran   []string `json:"media_pembelajaran"`
}

type ApproveRPSRequest struct {
	Catatan string `json:"catatan"`
}

type RejectRPSRequest struct {
	Alasan string `json:"alasan" binding:"required"`
}

type RequestRevisionRequest struct {
	Catatan string `json:"catatan" binding:"required"`
}

type CreateCPMKRequest struct {
	Kode      string   `json:"kode" binding:"required"`
	Deskripsi string   `json:"deskripsi" binding:"required"`
	Bobot     *float64 `json:"bobot"`
	Urutan    int      `json:"urutan"`
}

type UpdateCPMKRequest struct {
	Kode      string   `json:"kode"`
	Deskripsi string   `json:"deskripsi"`
	Bobot     *float64 `json:"bobot"`
	Urutan    int      `json:"urutan"`
}

type CreateRencanaPembelajaranRequest struct {
	Pertemuan          int      `json:"pertemuan" binding:"required,min=1"`
	KemampuanAkhir     string   `json:"kemampuan_akhir" binding:"required"`
	Indikator          *string  `json:"indikator"`
	Materi             string   `json:"materi" binding:"required"`
	MetodePembelajaran *string  `json:"metode_pembelajaran"`
	WaktuMenit         int      `json:"waktu_menit"`
	PengalamanBelajar  *string  `json:"pengalaman_belajar"`
	KriteriaPenilaian  *string  `json:"kriteria_penilaian"`
	BobotNilai         *float64 `json:"bobot_nilai"`
	Referensi          *string  `json:"referensi"`
}

type UpdateRencanaPembelajaranRequest struct {
	Pertemuan          int      `json:"pertemuan" binding:"omitempty,min=1"`
	KemampuanAkhir     string   `json:"kemampuan_akhir"`
	Indikator          *string  `json:"indikator"`
	Materi             string   `json:"materi"`
	MetodePembelajaran *string  `json:"metode_pembelajaran"`
	WaktuMenit         int      `json:"waktu_menit"`
	PengalamanBelajar  *string  `json:"pengalaman_belajar"`
	KriteriaPenilaian  *string  `json:"kriteria_penilaian"`
	BobotNilai         *float64 `json:"bobot_nilai"`
	Referensi          *string  `json:"referensi"`
}

type CreateBahanBacaanRequest struct {
	Jenis    string  `json:"jenis" binding:"required,oneof=utama pendukung"`
	Judul    string  `json:"judul" binding:"required"`
	Penulis  *string `json:"penulis"`
	Penerbit *string `json:"penerbit"`
	Tahun    *int    `json:"tahun"`
	ISBN     *string `json:"isbn"`
	URL      *string `json:"url"`
	Urutan   int     `json:"urutan"`
}

type UpdateBahanBacaanRequest struct {
	Jenis    string  `json:"jenis" binding:"omitempty,oneof=utama pendukung"`
	Judul    string  `json:"judul"`
	Penulis  *string `json:"penulis"`
	Penerbit *string `json:"penerbit"`
	Tahun    *int    `json:"tahun"`
	ISBN     *string `json:"isbn"`
	URL      *string `json:"url"`
	Urutan   int     `json:"urutan"`
}

type CreateEvaluasiRequest struct {
	Komponen          string  `json:"komponen" binding:"required"`
	TeknikPenilaian   *string `json:"teknik_penilaian"`
	Instrumen         *string `json:"instrumen"`
	Bobot             float64 `json:"bobot" binding:"required"`
	KriteriaPenilaian *string `json:"kriteria_penilaian"`
	Urutan            int     `json:"urutan"`
}

type UpdateEvaluasiRequest struct {
	Komponen          string  `json:"komponen"`
	TeknikPenilaian   *string `json:"teknik_penilaian"`
	Instrumen         *string `json:"instrumen"`
	Bobot             float64 `json:"bobot"`
	KriteriaPenilaian *string `json:"kriteria_penilaian"`
	Urutan            int     `json:"urutan"`
}

// ==================== CPL-MK MAPPING DTOs ====================

type CPLMKMappingListRequest struct {
	Page         int    `form:"page,default=1"`
	Limit        int    `form:"limit,default=10"`
	CPLID        string `form:"cpl_id"`
	MataKuliahID string `form:"mata_kuliah_id"`
	Level        string `form:"level"`
}

type UpsertCPLMKMappingRequest struct {
	CPLID        string `json:"cpl_id" binding:"required"`
	MataKuliahID string `json:"mata_kuliah_id" binding:"required"`
	Level        string `json:"level" binding:"required,oneof=tinggi sedang rendah"`
}

type CPLMKMappingResponse struct {
	ID           string             `json:"id"`
	CPLID        string             `json:"cpl_id"`
	MataKuliahID string             `json:"mata_kuliah_id"`
	Level        string             `json:"level"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	CPL          *CPLSimpleResponse `json:"cpl,omitempty"`
	MataKuliah   *MKSimpleResponse  `json:"mata_kuliah,omitempty"`
}

type CPLSimpleResponse struct {
	ID   string `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type MKSimpleResponse struct {
	ID       string `json:"id"`
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Semester int    `json:"semester"`
}
