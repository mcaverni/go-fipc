package gofipc

type MessageType uint8

const MESSAGE_NEW_VALUE = MessageType(0)
const MESSAGE_DEFAULT = MessageType(1)
const MESSAGE_MIN = MessageType(2)
const MESSAGE_MAX = MessageType(3)
const MESSAGE_MU = MessageType(4)
const MESSAGE_DESCRIPTION = MessageType(5)
const MESSAGE_CODE = MessageType(6)
const MESSAGE_COMMAND = MessageType(7)
const MESSAGE_TO_SERVER = MessageType(8)
const MESSAGE_FILE_OPERATION = MessageType(9)
const MESSAGE_HISTORY = MessageType(10)
const MESSAGE_VIEW_CONDITION = MessageType(11)
const MESSAGE_DISABLED_CONDITION = MessageType(12)
const MESSAGE_PING = MessageType(200)
const MESSAGE_SYNC_REQUEST_NO_DESC = MessageType(248)
const MESSAGE_SYNC_REQUEST_LITE_NO_DESC = MessageType(249)
const MESSAGE_INITIALIZATION_COMPLETE = MessageType(250)
const MESSAGE_MAIN_CLIENT = MessageType(251)
const MESSAGE_SYNC_REQUEST = MessageType(252)
const MESSAGE_SYNC_COMPLETE = MessageType(253)
const MESSAGE_SYNC_REQUEST_LITE = MessageType(254)
const MESSAGE_UNKNOWN = MessageType(255)

type Message struct {
	Type MessageType
	Data []byte
}
