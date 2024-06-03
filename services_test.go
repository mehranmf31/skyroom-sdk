package skyroom

import "testing"

func TestSkyroom_GetServices(t *testing.T) {
	sky := New(apiKey)
	services, err := sky.GetServices()

	if err != nil {
		t.Error(err)
	}
	t.Log(services)
}
