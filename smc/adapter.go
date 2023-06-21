package smc

import "github.com/sirupsen/logrus"

// IsAdapterEnabled returns whether the adapter is enabled.
func (c *Connection) IsAdapterEnabled() (bool, error) {
	logrus.Tracef("IsAdapterEnabled called")

	v, err := c.Read(AdapterKey)
	if err != nil {
		return false, err
	}

	ret := len(v.Bytes) == 1 && v.Bytes[0] == 0x0
	logrus.Tracef("IsAdapterEnabled returned %t", ret)

	return ret, nil
}

// EnableAdapter enables the adapter.
func (c *Connection) EnableAdapter() error {
	logrus.Tracef("EnableAdapter called")

	return c.Write(AdapterKey, []byte{0x0})
}

// DisableAdapter disables the adapter.
func (c *Connection) DisableAdapter() error {
	logrus.Tracef("DisableAdapter called")

	return c.Write(AdapterKey, []byte{0x1})
}
