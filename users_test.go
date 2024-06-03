package skyroom

import "testing"

func TestSkyroom_GetUsers(t *testing.T) {
	sky := New(APIKey)

	result, err := sky.GetUsers()
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_GetUserByID(t *testing.T) {
	sky := New(APIKey)

	result, err := sky.GetUserByID(UserID)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_GetUserByUsername(t *testing.T) {
	sky := New(APIKey)

	result, err := sky.GetUserByUsername(UserName)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}

func TestSkyroom_CreateUser(t *testing.T) {
	sky := New(APIKey)

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
	sky := New(APIKey)
	request := UpdateUserRequest{
		UserID:   UserID,
		Password: "Test123123",
	}
	err := sky.UpdateUser(request)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_DeleteUser(t *testing.T) {
	sky := New(APIKey)
	err := sky.DeleteUser(UserToDelete)

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_DeleteUsers(t *testing.T) {
	sky := New(APIKey)
	err := sky.DeleteUsers([]int{UsersToDelete})

	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_GetUserRooms(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.GetUserRooms(UserID)

	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSkyroom_AddUserRooms(t *testing.T) {
	sky := New(APIKey)
	rooms := []UserRoomAccess{
		{
			RoomID: RoomID,
		}}
	err := sky.AddUserRooms(UserID, rooms)
	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_RemoveUserRooms(t *testing.T) {
	sky := New(APIKey)
	rooms := []int{RoomID}
	err := sky.removeUserRooms(UserID, rooms)
	if err != nil {
		t.Error(err)
	}
}

func TestSkyroom_CreateLoginURL(t *testing.T) {
	sky := New(APIKey)
	result, err := sky.CreateLoginURL(CreateLoginURLRequest{
		RoomID:     RoomID,
		UserID:     UserID,
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
