package skyroom

import (
	"testing"
)

func TestSkyroom_GetRooms(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.GetRooms()

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_CountRooms(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.CountRooms()

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_GetRoomByID(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.GetRoomByID(RoomID)

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_GetRoomByName(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.GetRoomByName(RoomName)

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_GetRoomURL(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.GetRoomURL(RoomID, "fa")

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_CreateRoom(t *testing.T) {
	sky := New(APIKey)
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
	sky := New(APIKey)
	err := sky.UpdateRoom(UpdateRoomRequest{
		RoomID: RoomID,
		Title:  "test room",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_DeleteRoom(t *testing.T) {
	sky := New(APIKey)
	err := sky.DeleteRoom(RoomToDelete)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_GetRoomUsers(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.GetRoomUsers(RoomID)

	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_AddRoomUsers(t *testing.T) {
	sky := New(APIKey)

	users := []UserAccess{
		{UserID: UserID, Access: 1},
	}

	err := sky.AddRoomUsers(RoomID, users)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_RemoveRoomUsers(t *testing.T) {
	sky := New(APIKey)

	users := []int{UserID}

	err := sky.RemoveRoomUsers(RoomID, users)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_UpdateRoomUser(t *testing.T) {
	sky := New(APIKey)

	user := UserAccess{UserID: UserID, Access: 3}

	err := sky.UpdateRoomUser(RoomID, user)

	if err != nil {
		t.Error(err)
	}
}
