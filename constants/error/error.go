package error

const (
	ERR_STORE_INSERT  string = "CanNotInsert"
	ERR_STORE_DECODE  string = "CanNotDecode"
	ERR_STORE_DELETE  string = "CanNotDelete"
	ERR_STORE_COUNT   string = "CanNotGetCount"
	ERR_STORE_READ    string = "CanNotRead"
	ERR_STORE_UPDATE  string = "CanNotUpdate"
	ERR_STORE_UNKNOWN string = "Unknown"
)

const (
	ERR_CONTROLLER_CREATE         string = "CanNotCreate"
	MSG_CONTROLLER_CREATE         string = "Can not create entity in database"
	ERR_CONTROLLER_DELETE         string = "CanNotDelete"
	MSG_CONTROLLER_DELETE         string = "Can not mark to delete entity"
	ERR_CONTROLLER_ERASE          string = "CanNotErase"
	MSG_CONTROLLER_ERASE          string = "Can not erase entity from database"
	ERR_CONTROLLER_READ           string = "CanNotRead"
	MSG_CONTROLLER_READ           string = "Can not read entity(es)"
	ERR_CONTROLLER_TOTAL          string = "CanNotGetTotal"
	MSG_CONTROLLER_TOTAL          string = "Can not get entities total"
	ERR_CONTROLLER_MARSHAL_DATA   string = "CanNotMarshalFilter"
	MSG_CONTROLLER_MARSHAL_DATA   string = "Can not marshal update struct"
	ERR_CONTROLLER_UNMARSHAL_DATA string = "CanNotUnmarshalFilter"
	MSG_CONTROLLER_UNMARSHAL_DATA string = "Can not marshal update struct"
	ERR_CONTROLLER_UPDATE         string = "CanNotUpdate"
	MSG_CONTROLLER_UPDATE         string = "Can not update entity"
)

const (
	ERR_HANDLER_DTO        string = "CanNotGetDTO"
	MSG_HANDLER_DTO        string = "Can not get Data Transfer Object from request"
	ERR_HANDLER_FILL       string = "CanNotFill"
	MSG_HANDLER_FILL       string = "Can not fill Data Object by Data Transfer Object"
	ERR_HANDLER_PROVIDER   string = "CanNotGetProvider"
	MSG_HANDLER_PROVIDER   string = "Can not get store provider"
	ERR_HANDLER_CREATE     string = "CanNotCreate"
	MSG_HANDLER_CREATE     string = "Can not create entity"
	ERR_HANDLER_DELETE     string = "CanNotDelete"
	MSG_HANDLER_DELETE     string = "Can not delete entity(es)"
	ERR_HANDLER_CONVERT_ID string = "CanNotConvertId"
	MSG_HANDLER_CONVERT_ID string = "Can not convert identifier from URL params"
	ERR_HANDLER_PARAMS     string = "CanNotGetParam"
	MSG_HANDLER_PARAMS     string = "Can not get param from URL"
	ERR_HANDLER_ERASE      string = "CanNotErase"
	MSG_HANDLER_ERASE      string = "Can not erase entity"
	ERR_HANDLER_UPDATE     string = "CanNotUpdate"
	MSG_HANDLER_UPDATE     string = "Can not update entity"
)
