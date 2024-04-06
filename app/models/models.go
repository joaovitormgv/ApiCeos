package models

import "encoding/json"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// função de desserializar dados
func UnmarshalData(data []byte) (*User, error) {
	user := &User{}
	err := json.Unmarshal(data, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// função de serializar dados
func MarshalData(user *User) ([]byte, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}
