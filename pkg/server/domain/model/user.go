package model

func init() {
	RegisterModel(&User{})
}

// DefaultAdminUserAlias default admin user alias
const DefaultAdminUserAlias = "Administrator"

// RoleAdmin admin role
const RoleAdmin = "admin"

// User is the model of user
type User struct {
	BaseModel
	Name     string `json:"name" gorm:"primaryKey"`
	Email    string `json:"email"`
	Alias    string `json:"alias,omitempty"`
	Password string `json:"password,omitempty"`
}

// TableName return custom table name
func (u *User) TableName() string {
	return tableNamePrefix + "user"
}

// ShortTableName return custom table name
func (u *User) ShortTableName() string {
	return "usr"
}

// PrimaryKey return custom primary key
func (u *User) PrimaryKey() string {
	return u.Name
}

// Index return custom index
func (u *User) Index() map[string]interface{} {
	index := make(map[string]interface{})
	if u.Name != "" {
		index["name"] = u.Name
	}
	if u.Email != "" {
		index["email"] = u.Email
	}

	return index
}
