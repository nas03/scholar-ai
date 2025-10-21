package models

import (
	"database/sql"
	"time"
)

type TableCommon struct {
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
type User struct {
	UserID      string         `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username    string         `gorm:"uniqueIndex;not null" json:"username"`
	Email       string         `gorm:"uniqueIndex;not null" json:"email"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Password    string         `gorm:"not null" json:"password"`
	TableCommon

	// Relationships
	Courses []Course `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"courses,omitempty"`
}

type Course struct {
	ID          int            `gorm:"primaryKey" json:"id"`
	CourseID    string         `gorm:"not null" json:"course_id"`
	CourseName  string         `gorm:"not null" json:"course_name"`
	UserID      string         `gorm:"not null" json:"user_id"`
	Description sql.NullString `json:"description"`
	Lecturers   string         `json:"lecturers"`
	Credits     int            `gorm:"not null" json:"credits"`
	GPA         float32        `gorm:"not null" json:"gpa"`
	SemesterID  int            `gorm:"not null" json:"semester_id"`
	TableCommon

	// Relationships
	User     User     `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Semester Semester `gorm:"foreignKey:SemesterID;references:ID;constraint:OnDelete:RESTRICT" json:"semester,omitempty"`
	Tags     []Tag    `gorm:"many2many:course_tags;foreignKey:ID;joinForeignKey:CourseID;References:ID;joinReferences:TagID" json:"tags,omitempty"`
}

type Semester struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	TableCommon

	// Relationships
	Courses []Course `gorm:"foreignKey:SemesterID;references:ID;constraint:OnDelete:RESTRICT" json:"courses,omitempty"`
}

type Tag struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Color string `gorm:"default:#808080" json:"color"`
	TableCommon

	// Relationships
	Courses []Course `gorm:"many2many:course_tags;foreignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:CourseID" json:"courses,omitempty"`
}

type CourseTag struct {
	CourseID int `gorm:"primaryKey;column:course_id" json:"course_id"`
	TagID    int `gorm:"primaryKey;column:tag_id" json:"tag_id"`
}
