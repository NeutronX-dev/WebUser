package WebUser

import (
	"encoding/json"
	"io/ioutil"
)

type UserList struct {
	Path string
	Data []interface{}
}

func (Users *UserList) CreateUser(username, password string) string {
	var str_token string = MakeToken(username, password)
	var hash_token string = Hash256(str_token)
	ToAdd := make(map[string]interface{})
	ToAdd["username"] = username
	ToAdd["password"] = Hash256(password)
	ToAdd["token"] = hash_token
	Users.Data = append(Users.Data, ToAdd)
	return hash_token
}

func (Users *UserList) CreateUserWithExtraColumns(username, password string, extra_columns map[string]interface{}) string {
	var str_token string = MakeToken(username, password)
	var hash_token string = Hash256(str_token)
	ToAdd := make(map[string]interface{})
	ToAdd["username"] = username
	ToAdd["password"] = Hash256(password)
	ToAdd["token"] = hash_token
	for key, value := range extra_columns {
		ToAdd[key] = value
	}
	Users.Data = append(Users.Data, ToAdd)
	return hash_token
}

func (Users *UserList) UserExists(username string) bool {
	var found bool
	for i := 0; i < len(Users.Data); i++ {
		user := Users.Data[i].(map[string]interface{})
		if user["username"] == username {
			found = true
		}
	}
	return found
}

func (Users *UserList) GetUserByUsername(username string) map[string]interface{} {
	var res map[string]interface{}
	for i := 0; i < len(Users.Data); i++ {
		user := Users.Data[i].(map[string]interface{})
		if user["username"] == username {
			res = user
		}
	}
	return res
}

func (Users *UserList) GetUserByToken(token string) map[string]interface{} {
	var res map[string]interface{}
	for i := 0; i < len(Users.Data); i++ {
		user := Users.Data[i].(map[string]interface{})
		if user["token"] == token {
			res = user
		}
	}
	return res
}

func (Users *UserList) PasswordMatch(username, decoded_password string) bool {
	user := Users.GetUserByUsername(username)
	if user["password"] != nil {
		return user["password"] == Hash256(decoded_password)
	}
	return false
}

func (Users *UserList) SaveData() error {
	data, err := json.MarshalIndent(Users.Data, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(Users.Path, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func Read(path string) (*UserList, error) {
	var res UserList = UserList{Path: path, Data: make([]interface{}, 0)}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return &res, err
	}

	// Initialize JSON if file is Empty.
	if string(data) == "" {
		data = []byte("[]")
	}

	Parsed := make([]interface{}, 0)
	err = json.Unmarshal(data, &Parsed)
	if err != nil {
		return &res, err
	}
	res.Data = Parsed
	return &res, nil
}
