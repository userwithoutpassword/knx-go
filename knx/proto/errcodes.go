// Copyright 2017 Ole Krüger.

package proto

import "fmt"
import "reflect"

// What follows are errors codes defined in the KNX standard.
const (
	// NoError indicates a successful operation (E_NO_ERROR).
	NoError = 0x00

	// ErrHostProtocolType indicates an unsupported host protocol (E_HOST_PROTOCOL_TYPE).
	ErrHostProtocolType = 0x01

	// ErrVersionNotSupported indicates an unsupported KNXnet/IP protocol version.
	ErrVersionNotSupported = 0x02

	// ErrSequenceNumber indicates that an out-of-order sequence number has been received.
	ErrSequenceNumber = 0x04

	// ErrConnectionID indicates that there is no active data connection with given ID.
	ErrConnectionID = 0x21

	// ErrConnectionType indicates an unsupported connection type.
	ErrConnectionType = 0x22

	// ErrConnectionOption indicates an unsupported connection option.
	ErrConnectionOption = 0x23

	// ErrNoMoreConnections is returned by a Tunnelling Server when it cannot accept more
	// connections.
	ErrNoMoreConnections = 0x24

	// ErrNoMoreUniqueConnections is returned by a Tunnelling Server when it has no free Individual
	// Address available that could be used by the connection.
	ErrNoMoreUniqueConnections = 0x25

	// ErrDataConnection indicates an error with a data connection.
	ErrDataConnection = 0x26

	// ErrKNXConnection indicates an error with a KNX connection.
	ErrKNXConnection = 0x27

	// ErrTunnellingLayer indicates an unsupported tunnelling layer.
	ErrTunnellingLayer = 0x29
)

func toInt(x interface{}) (int, bool) {
	value := reflect.ValueOf(x)

	switch value.Kind() {
	case reflect.Int:
		return int(value.Int()), true

	case reflect.Int8:
		return int(value.Int()), true

	case reflect.Int16:
		return int(value.Int()), true

	case reflect.Int32:
		return int(value.Int()), true

	case reflect.Int64:
		return int(value.Int()), true

	case reflect.Uint:
		return int(value.Uint()), true

	case reflect.Uint8:
		return int(value.Uint()), true

	case reflect.Uint16:
		return int(value.Uint()), true

	case reflect.Uint32:
		return int(value.Uint()), true

	case reflect.Uint64:
		return int(value.Uint()), true
	}

	return 0, false
}

// ErrString returns a string representation of the error code.
func ErrString(err interface{}) string {
	val, ok := toInt(err)
	if !ok {
		return fmt.Sprintf("Unknown error code type %T", err)
	}

	switch val {
	case NoError:
		return "No error"

	case ErrHostProtocolType:
		return "Host protocol is not supported"

	case ErrVersionNotSupported:
		return "KNXnet/IP version is not supported"

	case ErrSequenceNumber:
		return "Sequence number is out-of-order"

	case ErrConnectionID:
		return "No active data connection"

	case ErrConnectionType:
		return "Unsupported connection type"

	case ErrConnectionOption:
		return "Unsupported connection option"

	case ErrNoMoreConnections:
		return "No more connections available"

	case ErrNoMoreUniqueConnections:
		return "No more unique connections available"

	case ErrDataConnection:
		return "Data connection error"

	case ErrKNXConnection:
		return "KNX connection error"

	case ErrTunnellingLayer:
		return "Unsupported tunnelling layer"

	default:
		return fmt.Sprintf("Unknown error code %#x", err)
	}
}
