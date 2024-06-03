package skyroom

import "encoding/json"

type getUserByNameRequest struct {
	Username string `json:"username"`
}

type getUserByIDRequest struct {
	UserID int `json:"user_id"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Status   int    `json:"status,omitempty"`
	IsPublic bool   `json:"is_public,omitempty"`
}

type UpdateUserRequest struct {
	UserID   int    `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Status   int    `json:"status,omitempty"`
	IsPublic bool   `json:"is_public,omitempty"`
}

type deleteUserRequest struct {
	UserID int `json:"user_id"`
}

type deleteUsersRequest struct {
	Users []int `json:"users"`
}

type getUserRoomsRequest struct {
	UserID int `json:"user_id"`
}

type addUserRoomsRequest struct {
	UserID int              `json:"user_id"`
	Rooms  []UserRoomAccess `json:"rooms"`
}

type UserRoomAccess struct {
	RoomID int `json:"room_id"`
	Access int `json:"access,omitempty"`
}

type removeUserRoomsRequest struct {
	UserID int   `json:"user_id"`
	Rooms  []int `json:"rooms"`
}

type CreateLoginURLRequest struct {
	RoomID     int    `json:"room_id"`
	UserID     int    `json:"user_id"`
	Nickname   string `json:"nickname"`
	Access     int    `json:"access,omitempty"`
	Concurrent int    `json:"concurrent,omitempty"`
	Language   string `json:"language,omitempty"`
	Ttl        int    `json:"ttl,omitempty"`
}

type GetUsersResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
}

type GetUserResponse struct {
	ID         int         `json:"id"`
	Username   string      `json:"username"`
	Nickname   string      `json:"nickname"`
	Email      string      `json:"email"`
	Firstname  string      `json:"fname"`
	Lastname   string      `json:"lname"`
	Gender     int         `json:"gender"`
	Status     int         `json:"status"`
	IsPublic   bool        `json:"is_public"`
	Concurrent int         `json:"concurrent"`
	TimeLimit  interface{} `json:"time_limit"`
	TimeUsage  int         `json:"time_usage"`
	TimeTotal  int         `json:"time_total"`
	ExpiryDate interface{} `json:"expiry_date"`
	CreateTime int         `json:"create_time"`
	UpdateTime int         `json:"update_time"`
}

type GetUserRoomsResponse struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Access int    `json:"access"`
}

func (sky *Skyroom) GetUsers() ([]GetUsersResponse, error) {
	result, err := sky.Post("getUsers", map[string]string{})
	if err != nil {
		return []GetUsersResponse{}, err
	}
	var users []GetUsersResponse

	if err := json.Unmarshal(result, &users); err != nil {
		return []GetUsersResponse{}, err
	}

	return users, nil
}

func (sky *Skyroom) GetUserByID(userID int) (GetUserResponse, error) {
	request := getUserByIDRequest{
		UserID: userID,
	}

	result, err := sky.Post("getUser", request)
	if err != nil {
		return GetUserResponse{}, err
	}
	var user GetUserResponse
	if err := json.Unmarshal(result, &user); err != nil {
		return GetUserResponse{}, err
	}
	return user, nil
}

func (sky *Skyroom) GetUserByUsername(username string) (GetUserResponse, error) {
	request := getUserByNameRequest{
		Username: username,
	}

	result, err := sky.Post("getUser", request)
	if err != nil {
		return GetUserResponse{}, err
	}
	var user GetUserResponse
	if err := json.Unmarshal(result, &user); err != nil {
		return GetUserResponse{}, err
	}
	return user, nil
}

func (sky *Skyroom) CreateUser(request CreateUserRequest) (int, error) {
	result, err := sky.Post("createUser", request)
	if err != nil {
		return 0, err
	}

	var user int

	if err := json.Unmarshal(result, &user); err != nil {
		return 0, err
	}
	return user, nil
}

func (sky *Skyroom) UpdateUser(request UpdateUserRequest) error {
	_, err := sky.Post("updateUser", request)

	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) DeleteUser(userId int) error {
	request := deleteUserRequest{
		UserID: userId,
	}

	_, err := sky.Post("deleteUser", request)
	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) DeleteUsers(users []int) error {
	request := deleteUsersRequest{
		Users: users,
	}

	_, err := sky.Post("deleteUsers", request)
	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) GetUserRooms(userId int) ([]GetUserRoomsResponse, error) {
	request := getUserRoomsRequest{
		UserID: userId,
	}
	result, err := sky.Post("getUserRooms", request)
	if err != nil {
		return []GetUserRoomsResponse{}, err
	}
	var userRooms []GetUserRoomsResponse
	if err := json.Unmarshal(result, &userRooms); err != nil {
		return []GetUserRoomsResponse{}, err
	}
	return userRooms, nil
}

func (sky *Skyroom) AddUserRooms(userID int, rooms []UserRoomAccess) error {
	request := addUserRoomsRequest{
		UserID: userID,
		Rooms:  rooms,
	}
	_, err := sky.Post("addUserRooms", request)
	if err != nil {
		return err
	}
	return nil
}

func (sky *Skyroom) removeUserRooms(userID int, rooms []int) error {
	request := removeUserRoomsRequest{
		UserID: userID,
		Rooms:  rooms,
	}
	_, err := sky.Post("removeUserRooms", request)
	if err != nil {
		return err
	}
	return nil
}

func (sky *Skyroom) CreateLoginURL(request CreateLoginURLRequest) (string, error) {
	result, err := sky.Post("createLoginUrl", request)
	if err != nil {
		return "", err
	}

	var url string
	if err := json.Unmarshal(result, &url); err != nil {
		return "", err
	}

	return url, nil
}
