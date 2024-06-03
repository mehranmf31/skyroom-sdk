package skyroom

import (
	"encoding/json"
	"fmt"
)

type getRoomByIDRequest struct {
	RoomID int `json:"room_id"`
}

type getRoomByNameRequest struct {
	Name string `json:"name"`
}

type getRoomURLRequest struct {
	RoomID   int    `json:"room_id"`
	Language string `json:"language,omitempty"`
}

type CreateRoomRequest struct {
	Name         string `json:"name"`
	Title        string `json:"title"`
	GuestLogin   bool   `json:"guest_login,omitempty"`
	OpLoginFirst bool   `json:"op_login_first,omitempty"`
	MaxUsers     int    `json:"max_users,omitempty"`
}

type UpdateRoomRequest struct {
	RoomID       int    `json:"room_id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	GuestLogin   bool   `json:"guest_login,omitempty"`
	OpLoginFirst bool   `json:"op_login_first,omitempty"`
	MaxUsers     int    `json:"max_users,omitempty"`
}

type deleteRoomRequest struct {
	RoomID int `json:"room_id"`
}

type getRoomUsersRequest struct {
	RoomID int `json:"room_id"`
}

type addRoomUsersRequest struct {
	RoomID int          `json:"room_id"`
	Users  []UserAccess `json:"users"`
}

type UserAccess struct {
	UserID int `json:"user_id"`
	Access int `json:"access,omitempty"`
}

type removeRoomUsersRequest struct {
	RoomID int   `json:"room_id"`
	Users  []int `json:"users"`
}

type updateRoomUsersRequest struct {
	RoomID int `json:"room_id"`
	UserID int `json:"user_id"`
	Access int `json:"access"`
}

type GetRoomsResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
	ServiceID int    `json:"service_id"`
}

type GetRoomResponse struct {
	Id              int         `json:"id"`
	ServiceId       int         `json:"service_id"`
	Name            string      `json:"name"`
	Title           string      `json:"title"`
	Description     interface{} `json:"description"`
	Status          int         `json:"status"`
	CreateTime      int         `json:"create_time"`
	UpdateTime      int         `json:"update_time"`
	TimeLimit       interface{} `json:"time_limit"`
	TimeUsage       int         `json:"time_usage"`
	TimeTotal       int         `json:"time_total"`
	GuestLogin      bool        `json:"guest_login"`
	GuestLimit      interface{} `json:"guest_limit"`
	OpLoginFirst    bool        `json:"op_login_first"`
	MaxUsers        interface{} `json:"max_users"`
	SessionDuration interface{} `json:"session_duration"`
}

type CountRoomsResponse int

type CreateRoomResponse int

type GetRoomUsersResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Access   int    `json:"access"`
}

func (sky *Skyroom) GetRooms() ([]GetRoomsResponse, error) {
	result, err := sky.Post("getRooms", make(map[string]string))
	if err != nil {
		fmt.Println(err)
	}

	var rooms []GetRoomsResponse

	if err := json.Unmarshal(result, &rooms); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rooms, nil
}

func (sky *Skyroom) CountRooms() (CountRoomsResponse, error) {
	result, err := sky.Post("countRooms", make(map[string]string))
	if err != nil {
		fmt.Println(err)
	}

	var roomsCount CountRoomsResponse

	if err := json.Unmarshal(result, &roomsCount); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return roomsCount, nil
}

func (sky *Skyroom) GetRoomByID(id int) (GetRoomResponse, error) {
	request := getRoomByIDRequest{RoomID: id}

	result, err := sky.Post("getRoom", request)
	if err != nil {
		fmt.Println(err)
	}

	var room GetRoomResponse

	if err := json.Unmarshal(result, &room); err != nil {
		fmt.Println(err)
		return GetRoomResponse{}, err
	}

	return room, nil
}

func (sky *Skyroom) GetRoomByName(name string) (GetRoomResponse, error) {
	request := getRoomByNameRequest{Name: name}

	result, err := sky.Post("getRoom", request)
	if err != nil {
		return GetRoomResponse{}, err
	}

	var room GetRoomResponse

	if err := json.Unmarshal(result, &room); err != nil {
		fmt.Println(err)
		return GetRoomResponse{}, err
	}

	return room, nil
}

func (sky *Skyroom) GetRoomURL(roomID int, language ...string) (string, error) {
	request := getRoomURLRequest{RoomID: roomID, Language: language[0]}

	result, err := sky.Post("getRoomUrl", request)
	if err != nil {
		return "", err
	}

	var room string

	if err := json.Unmarshal(result, &room); err != nil {
		fmt.Println(err)
		return "", err
	}

	return room, nil
}

func (sky *Skyroom) CreateRoom(room CreateRoomRequest) (CreateRoomResponse, error) {
	result, err := sky.Post("createRoom", room)
	if err != nil {
		return 0, err
	}

	var roomID CreateRoomResponse

	if err := json.Unmarshal(result, &roomID); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return roomID, nil
}

func (sky *Skyroom) UpdateRoom(room UpdateRoomRequest) error {
	_, err := sky.Post("updateRoom", room)
	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) DeleteRoom(roomID int) error {
	room := deleteRoomRequest{RoomID: roomID}
	_, err := sky.Post("deleteRoom", room)
	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) GetRoomUsers(roomID int) ([]GetRoomUsersResponse, error) {
	room := getRoomUsersRequest{RoomID: roomID}
	result, err := sky.Post("getRoomUsers", room)
	if err != nil {
		return []GetRoomUsersResponse{}, err
	}

	var roomUsers []GetRoomUsersResponse

	if err := json.Unmarshal(result, &roomUsers); err != nil {
		fmt.Println(err)
		return []GetRoomUsersResponse{}, err
	}

	return roomUsers, nil
}

func (sky *Skyroom) AddRoomUsers(roomID int, users []UserAccess) error {
	request := addRoomUsersRequest{RoomID: roomID, Users: users}

	_, err := sky.Post("addRoomUsers", request)

	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) RemoveRoomUsers(roomID int, users []int) error {
	request := removeRoomUsersRequest{RoomID: roomID, Users: users}
	_, err := sky.Post("removeRoomUsers", request)

	if err != nil {
		return err
	}

	return nil
}

func (sky *Skyroom) UpdateRoomUser(roomID int, user UserAccess) error {
	request := updateRoomUsersRequest{RoomID: roomID, UserID: user.UserID, Access: user.Access}

	_, err := sky.Post("updateRoomUser", request)

	if err != nil {
		return err
	}

	return nil
}
