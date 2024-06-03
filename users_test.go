package skyroom

import "testing"

func TestSkyroom_GetUsers(t *testing.T) {
	sky := New(apiKey)

	result, err := sky.GetUsers()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_GetUserByID(t *testing.T) {
	sky := New(apiKey)

	result, err := sky.GetUserByID(userID)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_GetUserByUsername(t *testing.T) {
	sky := New(apiKey)

	result, err := sky.GetUserByUsername(userName)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_CreateUser(t *testing.T) {
	sky := New(apiKey)

	request := CreateUserRequest{
		Username: "test-create-users-to-delete",
		Password: "Test123123",
		Nickname: "ali ali",
		IsPublic: false,
		Status:   0,
	}
	result, err := sky.CreateUser(request)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_UpdateUser(t *testing.T) {
	sky := New(apiKey)
	request := UpdateUserRequest{
		UserID:   userID,
		Password: "Test123123",
	}
	err := sky.UpdateUser(request)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_DeleteUser(t *testing.T) {
	sky := New(apiKey)
	err := sky.DeleteUser(userToDelete)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_DeleteUsers(t *testing.T) {
	sky := New(apiKey)
	err := sky.DeleteUsers([]int{usersToDelete})

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_GetUserRooms(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.GetUserRooms(userID)

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_AddUserRooms(t *testing.T) {
	sky := New(apiKey)
	rooms := []UserRoomAccess{
		{
			RoomID: roomID,
		}}
	err := sky.AddUserRooms(userID, rooms)
	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_RemoveUserRooms(t *testing.T) {
	sky := New(apiKey)
	rooms := []int{roomID}
	err := sky.removeUserRooms(userID, rooms)
	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_CreateLoginURL(t *testing.T) {
	sky := New(apiKey)
	result, err := sky.CreateLoginURL(CreateLoginURLRequest{
		RoomID:     roomID,
		UserID:     userID,
		Nickname:   "Mehran",
		Access:     1,
		Concurrent: 1,
		Language:   "fa",
		Ttl:        3600,
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
