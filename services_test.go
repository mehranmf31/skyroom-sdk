package skyroom

import "testing"

func TestSkyroom_GetServices(t *testing.T) {
	sky := New(APIKey)
	services, err := sky.GetServices()

	if err != nil {
		t.Error(err)
	}
	t.Log(services)
}
