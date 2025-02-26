package m1client

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConnection(t *testing.T) {
	t.Parallel()

	client := NewClient(net.IP{0, 0, 0, 0}, 0)

	conn := client.GetConnection(3000)
	assert.NotNil(t, conn)
	assert.Equal(t, client.connections[3000], conn)

	conn2 := client.GetConnection(3001)
	assert.NotNil(t, conn2)
	assert.Equal(t, client.connections[3001], conn2)

	conn3 := client.GetConnection(3000)
	assert.NotNil(t, conn3)
	assert.Equal(t, conn, conn3)

	client.Close()
	assert.Empty(t, client.connections)
}
