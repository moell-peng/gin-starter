package requests

import "github.com/thedevsaddam/govalidator"

type UserRequest struct {
	Action string
}

func (r *UserRequest) Rules() govalidator.MapData {
	rules := govalidator.MapData{
		"name":  []string{"required", "between:1,20"},
		"phone": []string{"size:11"},
		"email": []string{"email"},
		"sex":   []string{"in:0,1"},
	}

	if r.Action == "create" {
		rules["username"] = []string{"required", "between:5,20"}
		rules["password"] = []string{"required", "between:8,20"}
	}

	return rules
}

func (r *UserRequest) Messages() govalidator.MapData {
	return govalidator.MapData{}
}

type LoginRequest struct {
}

func (l *LoginRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"username": []string{"required", "between:5,20"},
		"password": []string{"required", "between:8,20"},
	}
}

func (l *LoginRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"username": []string{"required:用户名不能为空", "between:用户名必须为5-20位"},
		"password": []string{"required:密码不能为空", "between:用户名必须为8-20位"},
	}
}

type RegisterRequest struct {
}

func (r *RegisterRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"name":     []string{"required", "between:1,20"},
		"username": []string{"required", "between:5,20"},
		"password": []string{"required", "between:8,20"},
	}
}

func (r *RegisterRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"name":     []string{"required:昵称不能为空", "between:昵称必须为1-20位"},
		"username": []string{"required:用户名不能为空", "between:用户名必须为5-20位"},
		"password": []string{"required:密码不能为空", "between:用户名必须为8-20位"},
	}
}

type FrozenAccountRequest struct {
}

func (r *FrozenAccountRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"password": []string{"required", "between:8,20"},
	}
}

func (r *FrozenAccountRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"password": []string{"required:密码不能为空", "between:用户名必须为8-20位"},
	}
}

type ChangePasswordRequest struct {
}

func (r *ChangePasswordRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"old_password": []string{"required", "between:8,20"},
		"password":     []string{"required", "between:8,20"},
	}
}

func (r *ChangePasswordRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"old_password": []string{"required:密码不能为空", "between:用户名必须为8-20位"},
		"password":     []string{"required:密码不能为空", "between:用户名必须为8-20位"},
	}
}
