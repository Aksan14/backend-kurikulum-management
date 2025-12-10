package model

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// JSON type for MySQL
type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	}
	*j = bytes
	return nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return nil
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// Base Model
type BaseModel struct {
	ID        string         `gorm:"type:char(36);primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}

// User Model
type User struct {
	BaseModel
	Nama         string     `gorm:"type:varchar(255);not null" json:"nama"`
	Email        string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string     `gorm:"type:varchar(255);column:password_hash;not null" json:"-"`
	NIP          *string    `gorm:"type:varchar(50);column:nip;uniqueIndex" json:"nip"`
	Role         string     `gorm:"type:varchar(20);not null" json:"role"`
	Status       string     `gorm:"type:varchar(20);default:'active'" json:"status"`
	Phone        *string    `gorm:"type:varchar(20);column:phone" json:"phone"`
	AvatarURL    *string    `gorm:"type:varchar(500);column:avatar_url" json:"avatar_url"`
	LastLogin    *time.Time `gorm:"column:last_login" json:"last_login"`
}

func (User) TableName() string {
	return "users"
}

// RefreshToken Model
type RefreshToken struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	UserID    string    `gorm:"type:char(36);not null;index" json:"user_id"`
	Token     string    `gorm:"type:varchar(500);column:token;not null" json:"-"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (r *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// CPL Model
type CPL struct {
	ID        string     `gorm:"type:char(36);primary_key" json:"id"`
	Kode      string     `gorm:"type:varchar(20);uniqueIndex;not null" json:"kode"`
	Nama      string     `gorm:"type:varchar(255);column:nama;not null" json:"nama"`
	Deskripsi *string    `gorm:"type:text" json:"deskripsi"`
	Status    string     `gorm:"type:varchar(20);default:'draft'" json:"status"`
	Version   int        `gorm:"default:1" json:"version"`
	CreatedBy string     `gorm:"type:char(36);not null" json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Creator   User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (CPL) TableName() string {
	return "cpl"
}

func (c *CPL) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

// MataKuliah Model
type MataKuliah struct {
	ID              string     `gorm:"type:char(36);primary_key" json:"id"`
	Kode            string     `gorm:"type:varchar(20);uniqueIndex;not null" json:"kode"`
	Nama            string     `gorm:"type:varchar(255);not null" json:"nama"`
	SKS             int        `gorm:"not null" json:"sks"`
	Semester        int        `gorm:"not null" json:"semester"`
	Jenis           string     `gorm:"type:enum('wajib','pilihan');default:'wajib'" json:"jenis"`
	Deskripsi       *string    `gorm:"type:text" json:"deskripsi"`
	Prasyarat       JSON       `gorm:"type:json" json:"prasyarat"`
	DosenPengampuID *string    `gorm:"type:char(36);column:dosen_pengampu_id" json:"dosen_pengampu_id"`
	KoordinatorID   *string    `gorm:"type:char(36);column:koordinator_id" json:"koordinator_id"`
	IsActive        bool       `gorm:"column:is_active;default:true" json:"-"`
	Status          string     `gorm:"type:enum('aktif','nonaktif','dihapus');default:'aktif'" json:"status"`
	CreatedBy       string     `gorm:"type:char(36);not null" json:"created_by"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Creator         User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	DosenPengampu   *User      `gorm:"foreignKey:DosenPengampuID" json:"dosen_pengampu,omitempty"`
	Koordinator     *User      `gorm:"foreignKey:KoordinatorID" json:"koordinator,omitempty"`
}

func (MataKuliah) TableName() string {
	return "mata_kuliah"
}

func (m *MataKuliah) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	if m.Jenis == "" {
		m.Jenis = "wajib"
	}
	if m.Status == "" {
		m.Status = "aktif"
	}
	return nil
}

// CPLAssignment Model
type CPLAssignment struct {
	ID              string     `gorm:"type:char(36);primary_key" json:"id"`
	CPLIDs          JSON       `gorm:"type:json;not null" json:"cpl_ids"`
	DosenID         string     `gorm:"type:char(36);not null;index" json:"dosen_id"`
	MataKuliah      *string    `gorm:"column:mata_kuliah;type:varchar(255)" json:"mata_kuliah"`
	MataKuliahID    *string    `gorm:"column:mata_kuliah_id;type:char(36)" json:"mata_kuliah_id"`
	Deadline        *time.Time `gorm:"column:deadline" json:"deadline"`
	Status          string     `gorm:"type:varchar(20);default:'assigned'" json:"status"`
	Catatan         *string    `gorm:"column:catatan;type:text" json:"catatan"`
	RejectionReason *string    `gorm:"column:rejection_reason;type:text" json:"rejection_reason"`
	AssignedBy      string     `gorm:"column:assigned_by;type:char(36);not null" json:"assigned_by"`
	AssignedAt      time.Time  `gorm:"column:assigned_at;default:CURRENT_TIMESTAMP" json:"assigned_at"`
	ResponseAt      *time.Time `gorm:"column:response_at" json:"response_at"`
	CompletedAt     *time.Time `gorm:"column:completed_at" json:"completed_at"`
	CreatedAt       time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;index" json:"deleted_at"`
	Dosen           User       `gorm:"foreignKey:DosenID" json:"dosen,omitempty"`
	MataKuliahRef   MataKuliah `gorm:"foreignKey:MataKuliahID" json:"mata_kuliah_ref,omitempty"`
	Assigner        User       `gorm:"foreignKey:AssignedBy" json:"assigner,omitempty"`
}

func (CPLAssignment) TableName() string {
	return "cpl_assignments"
}

func (c *CPLAssignment) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

// BobotNilai struct for RPS
type BobotNilai struct {
	Tugas     int `json:"tugas"`
	UTS       int `json:"uts"`
	UAS       int `json:"uas"`
	Kehadiran int `json:"kehadiran"`
	Praktikum int `json:"praktikum"`
}

// RPS Model
type RPS struct {
	ID                  string         `gorm:"type:char(36);primary_key" json:"id"`
	MataKuliahID        string         `gorm:"type:char(36);not null;index" json:"mata_kuliah_id"`
	TahunAjaran         string         `gorm:"column:tahun_ajaran;type:varchar(20);not null" json:"tahun_ajaran"`
	SemesterType        string         `gorm:"column:semester_type;type:enum('ganjil','genap')" json:"semester_type"`
	SemesterTipe        string         `gorm:"column:semester_tipe;type:enum('ganjil','genap');default:'ganjil'" json:"semester_tipe"`
	TanggalPenyusunan   *time.Time     `gorm:"type:date" json:"tanggal_penyusunan"`
	DosenID             string         `gorm:"type:char(36);not null;index" json:"dosen_id"`
	DosenNama           string         `gorm:"type:varchar(255);not null" json:"dosen_nama"`
	PenyusunID          *string        `gorm:"type:char(36)" json:"penyusun_id"`
	PenyusunNama        *string        `gorm:"type:varchar(255)" json:"penyusun_nama"`
	PenyusunNIDN        *string        `gorm:"column:penyusun_nidn;type:varchar(50)" json:"penyusun_nidn"`
	KoordinatorRMKID    *string        `gorm:"type:char(36)" json:"koordinator_rmk_id"`
	KoordinatorRMKNama  *string        `gorm:"type:varchar(255)" json:"koordinator_rmk_nama"`
	KoordinatorRMKNIDN  *string        `gorm:"column:koordinator_rmk_nidn;type:varchar(50)" json:"koordinator_rmk_nidn"`
	KaprodiID           *string        `gorm:"type:char(36)" json:"kaprodi_id"`
	KaprodiNama         *string        `gorm:"type:varchar(255)" json:"kaprodi_nama"`
	KaprodiNIDN         *string        `gorm:"column:kaprodi_nidn;type:varchar(50)" json:"kaprodi_nidn"`
	Fakultas            *string        `gorm:"type:varchar(255)" json:"fakultas"`
	ProgramStudi        *string        `gorm:"type:varchar(255)" json:"program_studi"`
	DeskripsiMK         *string        `gorm:"column:deskripsi_mk;type:text" json:"deskripsi_mk"`
	CapaianPembelajaran *string        `gorm:"column:capaian_pembelajaran;type:text" json:"capaian_pembelajaran"`
	MetodePembelajaran  JSON           `gorm:"column:metode_pembelajaran;type:json" json:"metode_pembelajaran"`
	MediaPembelajaran   JSON           `gorm:"column:media_pembelajaran;type:json" json:"media_pembelajaran"`
	Status              string         `gorm:"type:varchar(20);default:'draft'" json:"status"`
	Version             int            `gorm:"default:1" json:"version"`
	ReviewerID          *string        `gorm:"column:reviewer_id;type:char(36)" json:"reviewer_id"`
	ReviewCatatan       *string        `gorm:"column:review_catatan;type:text" json:"review_catatan"`
	ReviewedAt          *time.Time     `gorm:"column:reviewed_at" json:"reviewed_at"`
	ApprovedAt          *time.Time     `gorm:"column:approved_at" json:"approved_at"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	// Relations
	MataKuliah           MataKuliah                   `gorm:"foreignKey:MataKuliahID" json:"mata_kuliah,omitempty"`
	Dosen                User                         `gorm:"foreignKey:DosenID" json:"dosen,omitempty"`
	Reviewer             *User                        `gorm:"foreignKey:ReviewerID" json:"reviewer,omitempty"`
	CPMK                 []RPSCPMK                    `gorm:"foreignKey:RPSID" json:"cpmk,omitempty"`
	RencanaPembelajaran  []RPSRencanaPembelajaran     `gorm:"foreignKey:RPSID" json:"rencana_pembelajaran,omitempty"`
	BahanBacaan          []RPSBahanBacaan             `gorm:"foreignKey:RPSID" json:"bahan_bacaan,omitempty"`
	RencanaTugas         []RPSRencanaTugas            `gorm:"foreignKey:RPSID" json:"rencana_tugas,omitempty"`
	AnalisisKetercapaian []RPSAnalisisKetercapaianCPL `gorm:"foreignKey:RPSID" json:"analisis_ketercapaian,omitempty"`
}

func (RPS) TableName() string {
	return "rps"
}

func (r *RPS) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// RPSCPMK Model
type RPSCPMK struct {
	ID          string           `gorm:"type:char(36);primary_key" json:"id"`
	RPSID       string           `gorm:"type:char(36);not null;index" json:"rps_id"`
	Kode        string           `gorm:"type:varchar(20);not null" json:"kode"`
	Deskripsi   string           `gorm:"type:text;not null" json:"deskripsi"`
	Bobot       *float64         `gorm:"type:decimal(5,2)" json:"bobot"`
	Urutan      int              `gorm:"not null;default:1" json:"urutan"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	SubCPMKs    []SubCPMK        `gorm:"foreignKey:CPMKID" json:"sub_cpmks,omitempty"`
	CPLMappings []CPMKCPLMapping `gorm:"foreignKey:CPMKID" json:"cpl_mappings,omitempty"`
}

func (RPSCPMK) TableName() string {
	return "rps_cpmk"
}

func (r *RPSCPMK) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// SubCPMK Model - Kemampuan Akhir Tiap Tahapan
type SubCPMK struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	CPMKID    string    `gorm:"type:char(36);not null;index" json:"cpmk_id"`
	Kode      string    `gorm:"type:varchar(50);not null" json:"kode"`
	Deskripsi string    `gorm:"type:text;not null" json:"deskripsi"`
	Urutan    int       `gorm:"not null;default:1" json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Relations
	CPMK RPSCPMK `gorm:"foreignKey:CPMKID" json:"cpmk,omitempty"`
}

func (SubCPMK) TableName() string {
	return "sub_cpmk"
}

func (s *SubCPMK) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// CPMKCPLMapping Model - Relasi CPMK ke CPL
type CPMKCPLMapping struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	CPMKID    string    `gorm:"type:char(36);not null;index" json:"cpmk_id"`
	CPLID     string    `gorm:"type:char(36);not null;index" json:"cpl_id"`
	CreatedAt time.Time `json:"created_at"`
	// Relations
	CPMK RPSCPMK `gorm:"foreignKey:CPMKID" json:"cpmk,omitempty"`
	CPL  CPL     `gorm:"foreignKey:CPLID" json:"cpl,omitempty"`
}

func (CPMKCPLMapping) TableName() string {
	return "cpmk_cpl_mapping"
}

func (c *CPMKCPLMapping) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

type RPSRencanaPembelajaran struct {
	ID                 string    `gorm:"type:char(36);primary_key" json:"id"`
	RPSID              string    `gorm:"type:char(36);not null;index" json:"rps_id"`
	MingguKe           int       `gorm:"column:minggu_ke;not null" json:"minggu_ke"`                      // MG KE-
	SubCPMKID          *string   `gorm:"column:sub_cpmk_id;type:char(36)" json:"sub_cpmk_id"`             // KEMAMPUAN AKHIR TIAP TAHAPAN (SUB-CPMK) - FK ke sub_cpmk
	Topik              string    `gorm:"type:varchar(500);not null" json:"topik"`                         // TOPIK
	SubTopik           JSON      `gorm:"column:sub_topik;type:json" json:"sub_topik"`                     // SUB-TOPIK MATERI
	MetodePembelajaran *string   `gorm:"column:metode_pembelajaran;type:text" json:"metode_pembelajaran"` // METODE PEMBELAJARAN (SKEMA BLENDED LEARNING)
	WaktuMenit         *int      `gorm:"column:waktu_menit" json:"waktu_menit"`                           // WAKTU (MENIT)
	TeknikKriteria     *string   `gorm:"column:teknik_kriteria;type:text" json:"teknik_kriteria"`         // PENILAIAN - TEKNIK & KRITERIA
	BobotPersen        *float64  `gorm:"column:bobot_persen;type:decimal(5,2)" json:"bobot_persen"`       // PENILAIAN - BOBOT (%)
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	// Relations - SubCPMK dengan foreign key (indikator = deskripsi dari sub_cpmk)
	SubCPMK *SubCPMK `gorm:"foreignKey:SubCPMKID" json:"sub_cpmk,omitempty"`
}

func (RPSRencanaPembelajaran) TableName() string {
	return "rps_rencana_pembelajaran"
}

func (r *RPSRencanaPembelajaran) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// RPSRencanaTugas Model - Rencana Tugas
type RPSRencanaTugas struct {
	ID                    string    `gorm:"type:char(36);primary_key" json:"id"`
	RPSID                 string    `gorm:"type:char(36);not null;index" json:"rps_id"`
	NomorTugas            int       `gorm:"not null" json:"nomor_tugas"`
	Judul                 string    `gorm:"type:varchar(500);not null" json:"judul"`
	IndikatorKeberhasilan *string   `gorm:"type:text" json:"indikator_keberhasilan"`
	BatasWaktuMinggu      *int      `json:"batas_waktu_minggu"`
	PetunjukPengerjaan    *string   `gorm:"type:text" json:"petunjuk_pengerjaan"`
	JenisTugas            string    `gorm:"type:enum('individu','kelompok');default:'individu'" json:"jenis_tugas"`
	LuaranTugas           *string   `gorm:"type:text" json:"luaran_tugas"`
	KriteriaPenilaian     *string   `gorm:"type:text" json:"kriteria_penilaian"`
	TeknikPenilaian       *string   `gorm:"type:varchar(255)" json:"teknik_penilaian"`
	Bobot                 int       `gorm:"not null;default:0" json:"bobot"`
	SubCPMKID             *string   `gorm:"column:sub_cpmk_id;type:char(36)" json:"sub_cpmk_id"` // FK ke sub_cpmk (Sub-CPMK)
	DaftarRujukan         *string   `gorm:"type:text" json:"daftar_rujukan"`                     // Daftar Rujukan
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	// Relations - SubCPMK dengan foreign key (Indikator = deskripsi dari sub_cpmk)
	SubCPMK *SubCPMK `gorm:"foreignKey:SubCPMKID" json:"sub_cpmk,omitempty"`
}

func (RPSRencanaTugas) TableName() string {
	return "rps_rencana_tugas"
}

func (r *RPSRencanaTugas) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// RPSAnalisisKetercapaianCPL Model - Analisis Ketercapaian CPL
type RPSAnalisisKetercapaianCPL struct {
	ID              string    `gorm:"type:char(36);primary_key" json:"id"`
	RPSID           string    `gorm:"type:char(36);not null;index" json:"rps_id"`
	MingguMulai     int       `gorm:"not null" json:"minggu_mulai"`
	MingguSelesai   *int      `json:"minggu_selesai"`
	CPLID           string    `gorm:"type:char(36);not null;index" json:"cpl_id"`
	CPMKIDs         JSON      `gorm:"type:json" json:"cpmk_ids"`
	SubCPMKIDs      JSON      `gorm:"type:json" json:"sub_cpmk_ids"`
	TopikMateri     *string   `gorm:"type:varchar(500)" json:"topik_materi"`
	JenisAssessment *string   `gorm:"type:varchar(255)" json:"jenis_assessment"`
	BobotKontribusi int       `gorm:"not null;default:0" json:"bobot_kontribusi"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	// Relations
	CPL CPL `gorm:"foreignKey:CPLID" json:"cpl,omitempty"`
}

func (RPSAnalisisKetercapaianCPL) TableName() string {
	return "rps_analisis_ketercapaian_cpl"
}

func (r *RPSAnalisisKetercapaianCPL) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// RPSBahanBacaan Model
type RPSBahanBacaan struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	RPSID     string    `gorm:"type:char(36);not null;index" json:"rps_id"`
	Judul     string    `gorm:"type:varchar(500);not null" json:"judul"`
	Penulis   *string   `gorm:"type:varchar(255)" json:"penulis"`
	Penerbit  *string   `gorm:"type:varchar(255)" json:"penerbit"`
	Tahun     *int      `json:"tahun"`
	Jenis     *string   `gorm:"type:varchar(50)" json:"jenis"`
	URL       *string   `gorm:"type:varchar(1000)" json:"url"`
	ISBN      *string   `gorm:"type:varchar(20)" json:"isbn"`
	Halaman   *string   `gorm:"type:varchar(50)" json:"halaman"`
	Urutan    *int      `json:"urutan"`
	CreatedAt time.Time `json:"created_at"`
}

func (RPSBahanBacaan) TableName() string {
	return "rps_bahan_bacaan"
}

func (r *RPSBahanBacaan) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// Notification Model
type Notification struct {
	ID            string     `gorm:"type:char(36);primary_key" json:"id"`
	UserID        string     `gorm:"type:char(36);not null;index" json:"user_id"`
	Title         string     `gorm:"type:varchar(255);not null" json:"title"`
	Message       string     `gorm:"type:text;not null" json:"message"`
	Type          string     `gorm:"type:varchar(50);not null" json:"type"`
	ReferenceType *string    `gorm:"column:reference_type;type:varchar(50)" json:"reference_type"`
	ReferenceID   *string    `gorm:"column:reference_id;type:char(36)" json:"reference_id"`
	IsRead        bool       `gorm:"column:is_read;default:false" json:"is_read"`
	ReadAt        *time.Time `gorm:"column:read_at" json:"read_at"`
	ActionURL     *string    `gorm:"column:action_url;type:varchar(500)" json:"action_url"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"created_at"`
	User          User       `gorm:"foreignKey:UserID" json:"-"`
}

func (Notification) TableName() string {
	return "notifications"
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == "" {
		n.ID = uuid.New().String()
	}
	return nil
}

// DocumentTemplate Model
type DocumentTemplate struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	Nama      string    `gorm:"type:varchar(255);not null" json:"nama"`
	Deskripsi *string   `gorm:"type:text" json:"deskripsi"`
	Sections  JSON      `gorm:"type:json;not null" json:"sections"`
	FileURL   *string   `gorm:"type:varchar(1000)" json:"file_url"`
	Version   string    `gorm:"type:varchar(20);default:'1.0'" json:"version"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `gorm:"type:char(36);not null" json:"created_by"`
	Creator   User      `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (DocumentTemplate) TableName() string {
	return "document_templates"
}

func (d *DocumentTemplate) BeforeCreate(tx *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.New().String()
	}
	return nil
}

// GeneratedDocument Model
type GeneratedDocument struct {
	ID             string           `gorm:"type:char(36);primary_key" json:"id"`
	TemplateID     string           `gorm:"type:char(36);not null;index" json:"template_id"`
	TemplateName   string           `gorm:"type:varchar(255);not null" json:"template_name"`
	Tahun          string           `gorm:"type:varchar(20);not null" json:"tahun"`
	Status         string           `gorm:"type:varchar(20);default:'processing'" json:"status"`
	FileURL        *string          `gorm:"type:varchar(1000)" json:"file_url"`
	FileType       string           `gorm:"type:varchar(10);not null" json:"file_type"`
	FileSize       *int64           `json:"file_size"`
	Sections       JSON             `gorm:"type:json;not null" json:"sections"`
	GenerationData JSON             `gorm:"type:json" json:"generation_data"`
	Progress       int              `gorm:"default:0" json:"progress"`
	ErrorMessage   *string          `gorm:"type:text" json:"error_message"`
	CreatedAt      time.Time        `json:"created_at"`
	CompletedAt    *time.Time       `json:"completed_at"`
	CreatedBy      string           `gorm:"type:char(36);not null" json:"created_by"`
	Template       DocumentTemplate `gorm:"foreignKey:TemplateID" json:"template,omitempty"`
	Creator        User             `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (GeneratedDocument) TableName() string {
	return "generated_documents"
}

func (g *GeneratedDocument) BeforeCreate(tx *gorm.DB) error {
	if g.ID == "" {
		g.ID = uuid.New().String()
	}
	return nil
}

// RPSCPLMapping Model
type RPSCPLMapping struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	RPSID     string    `gorm:"type:char(36);not null;index" json:"rps_id"`
	CPMKID    string    `gorm:"type:char(36);not null;index" json:"cpmk_id"`
	CPLID     string    `gorm:"type:char(36);not null;index" json:"cpl_id"`
	Level     string    `gorm:"type:varchar(20);not null" json:"level"`
	Bobot     int       `gorm:"default:1" json:"bobot"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `gorm:"type:char(36);not null" json:"created_by"`
	RPS       RPS       `gorm:"foreignKey:RPSID" json:"rps,omitempty"`
	CPMK      RPSCPMK   `gorm:"foreignKey:CPMKID" json:"cpmk,omitempty"`
	CPL       CPL       `gorm:"foreignKey:CPLID" json:"cpl,omitempty"`
}

func (RPSCPLMapping) TableName() string {
	return "rps_cpl_mapping"
}

func (r *RPSCPLMapping) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}

// AuditLog Model
type AuditLog struct {
	ID         string    `gorm:"type:char(36);primary_key" json:"id"`
	UserID     *string   `gorm:"type:char(36);index" json:"user_id"`
	Action     string    `gorm:"type:varchar(100);not null" json:"action"`
	AuditTable string    `gorm:"column:table_name;type:varchar(100);not null" json:"table_name"`
	RecordID   *string   `gorm:"type:char(36)" json:"record_id"`
	OldValues  JSON      `gorm:"type:json" json:"old_values"`
	NewValues  JSON      `gorm:"type:json" json:"new_values"`
	IPAddress  *string   `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent  *string   `gorm:"type:text" json:"user_agent"`
	CreatedAt  time.Time `json:"created_at"`
	User       *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}

func (a *AuditLog) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}

// SystemSetting Model
type SystemSetting struct {
	ID          string    `gorm:"type:char(36);primary_key" json:"id"`
	Key         string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"key"`
	Value       JSON      `gorm:"type:json" json:"value"`
	Description *string   `gorm:"type:text" json:"description"`
	Category    string    `gorm:"type:varchar(50);default:'general'" json:"category"`
	IsPublic    bool      `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   *string   `gorm:"type:char(36)" json:"updated_by"`
}

func (SystemSetting) TableName() string {
	return "system_settings"
}

func (s *SystemSetting) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// File Model
type File struct {
	ID           string         `gorm:"type:char(36);primary_key" json:"id"`
	OriginalName string         `gorm:"type:varchar(255);not null" json:"original_name"`
	StoredName   string         `gorm:"type:varchar(255);not null" json:"stored_name"`
	FilePath     string         `gorm:"type:varchar(1000);not null" json:"file_path"`
	FileSize     int64          `gorm:"not null" json:"file_size"`
	MimeType     string         `gorm:"type:varchar(100);not null" json:"mime_type"`
	FileType     string         `gorm:"type:varchar(50);not null" json:"file_type"`
	RelatedID    *string        `gorm:"type:char(36)" json:"related_id"`
	RelatedType  *string        `gorm:"type:varchar(50)" json:"related_type"`
	UploadedBy   string         `gorm:"type:char(36);not null" json:"uploaded_by"`
	IsTemporary  bool           `gorm:"default:false" json:"is_temporary"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Uploader     User           `gorm:"foreignKey:UploadedBy" json:"uploader,omitempty"`
}

func (File) TableName() string {
	return "files"
}

func (f *File) BeforeCreate(tx *gorm.DB) error {
	if f.ID == "" {
		f.ID = uuid.New().String()
	}
	return nil
}

type CPLMKMapping struct {
	ID           string     `gorm:"type:char(36);primary_key" json:"id"`
	CPLID        string     `gorm:"column:cpl_id;type:char(36);not null" json:"cpl_id"`
	MataKuliahID string     `gorm:"column:mata_kuliah_id;type:char(36);not null" json:"mata_kuliah_id"`
	Level        string     `gorm:"type:enum('tinggi','sedang','rendah');not null" json:"level"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	CPL          CPL        `gorm:"foreignKey:CPLID" json:"cpl,omitempty"`
	MataKuliah   MataKuliah `gorm:"foreignKey:MataKuliahID" json:"mata_kuliah,omitempty"`
}

func (CPLMKMapping) TableName() string {
	return "cpl_mk_mappings"
}

func (m *CPLMKMapping) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	return nil
}
