package types

import "time"

type UserRepo interface {
	GetByEmail(email string) (*User, error)
	ExistsByEmail(email string) (bool, error)
	Create(u *User) (int, error)
}

type AnimalRepo interface {
	// GetAllByCategory retrieves only partial animal info (for cards)
	GetAllByCategory(category string) ([]Animal, error)
	// GetAllByUser retrieves only partial animal info (animals owned by user)
	GetAllByUser(userID int) ([]Animal, error)
	Get(id int) (*Animal, error)
	Create(a *Animal) (int, error)
	Delete(userID int, ID int) error
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Animal struct {
	ID          int       `json:"id"`
	UserID      int       `json:"-"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Sex         string    `json:"sex"`
	Breed       string    `json:"breed"`
	Size        string    `json:"size"`
	AgeInMonth  int       `json:"age_in_month"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"img_url"`
	Description string    `json:"description"`
	ContactInfo string    `json:"contact_info"`
	CreatedAt   time.Time `json:"created_at"`
}

type SignUpRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
