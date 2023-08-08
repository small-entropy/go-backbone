package error

const (
	ErrStoreInsert  string = "CanNotInsert"
	ErrStoreDecode  string = "CanNotDecode"
	ErrStoreDelete  string = "CanNotDelete"
	ErrStoreCount   string = "CanNotGetCount"
	ErrStoreRead    string = "CanNotRead"
	ErrStoreUpdate  string = "CanNotUpdate"
	ErrStoreUnknown string = "Unknown"
)

const (
	ErrControllerCreate        string = "CanNotCreate"
	MsgControllerCreate        string = "Can not create entity in database"
	ErrControllerDelete        string = "CanNotDelete"
	MsgControllerDelete        string = "Can not mark to delete entity"
	ErrControllerErase         string = "CanNotErase"
	MsgControllerErase         string = "Can not erase entity from database"
	ErrControllerRead          string = "CanNotRead"
	MsgControllerRead          string = "Can not read entity(es)"
	ErrControllerTotal         string = "CanNotGetTotal"
	MsgControllerTotal         string = "Can not get entities total"
	ErrControllerMarshalData   string = "CanNotMarshalFilter"
	MsgControllerMarshalData   string = "Can not marshal update struct"
	ErrControllerUnMarshalData string = "CanNotUnmarshalFilter"
	MsgControllerUnMarshalData string = "Can not marshal update struct"
	ErrControllerUpdate        string = "CanNotUpdate"
	MsgControllerUpdate        string = "Can not update entity"
)

const (
	ErrHandlerDto       string = "CanNotGetDTO"
	MsgHandlerDto       string = "Can not get Data Transfer Object from request"
	ErrHandlerFill      string = "CanNotFill"
	MsgHandlerFill      string = "Can not fill Data Object by Data Transfer Object"
	ErrHandlerProvider  string = "CanNotGetProvider"
	MsgHandlerProvider  string = "Can not get store provider"
	ErrHandlerCreate    string = "CanNotCreate"
	MsgHandlerCreate    string = "Can not create entity"
	ErrHandlerDelete    string = "CanNotDelete"
	MsgHandlerDelete    string = "Can not delete entity(es)"
	ErrHandlerConvertID string = "CanNotConvertId"
	MsgHandlerConvertID string = "Can not convert identifier from URL params"
	ErrHandlerParams    string = "CanNotGetParam"
	MsgHandlerParams    string = "Can not get param from URL"
	ErrHandlerErase     string = "CanNotErase"
	MsgHandlerErase     string = "Can not erase entity"
	ErrHandlerUpdate    string = "CanNotUpdate"
	MsgHandlerUpdate    string = "Can not update entity"
)
