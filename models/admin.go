package models

type Admin struct {
	ID       int
	Username string `binding:"required"`
	Password string `binding:"required"`
	Email    string
}

var admins = []Admin{}

func (a Admin) ValidateCredentials() (bool, error) {
	for _, admin := range admins {
		if admin.Username == a.Username && admin.Password == a.Password {
			return true, nil
		}
	}
	return false, nil
}

func (a Admin) Save() error {
	admins = append(admins, a)
	return nil
}

func CreateAdmin(username, password, email string) error {
	admin := Admin{
		Username: username,
		Password: password,
		Email:    email,
	}
	return admin.Save()
}

// testing function
func GetAllAdmins() ([]Admin, error) {
	return admins, nil
}

func UpdateAdmin(id int, username, password, email string) error {
	for i, admin := range admins {
		if admin.ID == id {
			admins[i].Username = username
			admins[i].Password = password
			admins[i].Email = email
			return nil
		}
	}
	return nil
}

func DeleteAdmin(id int) error {
	for i, admin := range admins {
		if admin.ID == id {
			admins = append(admins[:i], admins[i+1:]...)
			return nil
		}
	}
	return nil
}
