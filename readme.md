# Skyroom Online Meeting Service API SDK

Skyroom is a local online meeting service provider, For more detailed information, please see the [original package documentation](https://data.skyroom.online/help/webservice.html).

## Installation

To install the `skyroom` package, you can use `go get`:

```bash
go get github.com/mehranmf31/skyroom-sdk

## Examples

Then, you can import it in your project:

```go
import "github.com/mehranmf31/skyroom-sdk"
```

## Examples

Here's some basic examples to demonstrate how to use the skyroom-sdk package. First, you should create an instance of Skyroom:

```go
sky := skyroom.New("apikey")
```

### Get Services

```go
services, err := sky.GetServices()
```

### Get Rooms

```go
rooms, err := sky.GetRooms()
```

### Count Rooms
```go
roomsCount, err := sky.CountRooms()
```

### Get Room By ID Or Name
```go
room, err := sky.GetRoomByID(1)

room, err := sky.GetRoomByName("name")
```

### Get Room URL
```go
count, err := sky.GetRoomURL(1, "fa")
```

### Create Room
```go
roomID, err := sky.CreateRoom(skyroom.CreateRoomRequest{
    Name: "name",
    Title: "name",
    GuestLogin: true,
    OpLoginFirst: true,
    MaxUsers: 10,
})
```

### Update Room
```go
err := sky.UpdateRoom(skyroom.UpdateRoomRequest{
    RoomID:       0,
    Name:         "name",
    Title:        "title",
    GuestLogin:   false,
    OpLoginFirst: false,
    MaxUsers:     0,
})
```

### Delete Room
```go
err := sky.DeleteRoom(1)
```

### Get Room Users
```go
users, err := sky.GetRoomUsers(1)
```

### Add Room Users
```go
err := sky.AddRoomUsers(roomID, []skyroom.UserAccess{
    {
        UserID: 1,
        Access: 1,
    },
})
```

### Remove Room Users
```go
err := sky.RemoveRoomUsers(roomID, []int{1,2,3})
```

### Update Room Users
```go
err := sky.UpdateRoomUser(roomID, skyroom.UserAccess{
    UserID: 0,
    Access: 0,
})
```

### Get Users
```go
users, err := sky.GetUsers()
```

### Get User By ID Or Username
```go
userByID, err := sky.GetUserByID("name")

userByUsername, err := sky.GetUserByUsername("name")
```

### Create User
```go
userID, err := sky.CreateUser(skyroom.CreateUserRequest{
    Username: "",
    Password: "",
    Nickname: "",
    Status:   0,
    IsPublic: false,
})
```

### Update User
```go
err := sky.UpdateUser(skyroom.UpdateUserRequest{
    UserID:   0,
    Username: "",
    Password: "",
    Nickname: "",
    Status:   0,
    IsPublic: false,
})
```

### Delete User
```go
err := sky.DeleteUser(1)
```

### Delete Users
```go
err := sky.DeleteUsers([]int{1, 2, 3})
```

### Get User Rooms
```go
rooms, err := sky.GetUserRooms(1)
```

### Add User Rooms
```go
err := sky.AddUserRooms(1, []skyroom.UserRoomAccess{
    {
        RoomID: 0,
        Access: 0,
    },
})
```

### Remove User Rooms
```go
err := sky.RemoveUserRooms(1, []int{1, 2, 3})
```

### Create Login URL
```go
url, err := sky.CreateLoginURL(skyroom.CreateLoginURLRequest{
	RoomID:     0,
	UserID:     0,
	Nickname:   "",
	Access:     0,
	Concurrent: 0,
	Language:   "",
	Ttl:        0,
})
```
