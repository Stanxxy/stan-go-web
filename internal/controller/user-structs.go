package controller

type (
	User          struct{}
	UserViewModel struct {
		Name string
		ID   string
	}
)

type (
	UserList          struct{}
	UserListViewModel struct {
		Users []UserViewModel
	}
)
