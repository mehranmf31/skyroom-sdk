package skyroom

import (
	"testing"
)

func TestSkyroom_GetRooms(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.GetRooms()

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_CountRooms(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.CountRooms()

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_GetRoomByID(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.GetRoomByID(roomID)

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_GetRoomByName(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.GetRoomByName(roomName)

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_GetRoomURL(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.GetRoomURL(roomID, "fa")

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_CreateRoom(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.CreateRoom(CreateRoomRequest{
		Name:  "test-create-room",
		Title: "test room",
	})

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_UpdateRoom(t *testing.T) {
	sky := New(apiKey)
	err := sky.UpdateRoom(UpdateRoomRequest{
		RoomID: roomID,
		Title:  "test room",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_DeleteRoom(t *testing.T) {
	sky := New(apiKey)
	err := sky.DeleteRoom(roomToDelete)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_GetRoomUsers(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.GetRoomUsers(roomID)

	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_AddRoomUsers(t *testing.T) {
	sky := New(apiKey)

	users := []UserAccess{
		{UserID: userID, Access: 1},
	}

	err := sky.AddRoomUsers(roomID, users)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_RemoveRoomUsers(t *testing.T) {
	sky := New(apiKey)

	users := []int{userID}

	err := sky.RemoveRoomUsers(roomID, users)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_UpdateRoomUser(t *testing.T) {
	sky := New(apiKey)

	user := UserAccess{UserID: userID, Access: 3}

	err := sky.UpdateRoomUser(roomID, user)

	if err != nil {
		t.Error(err)
	}
}
