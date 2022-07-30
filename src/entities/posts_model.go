package entities

// Post model
type Post struct {
	Model
	AuthorID    uint   `json:"authorId"`
	Author      *User  `json:"author" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Caption     string `json:"caption" gorm:"not null"`
	IsPublished bool   `json:"isPublished" gorm:"not null;default:true"`
}
