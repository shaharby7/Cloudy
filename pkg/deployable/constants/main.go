package constants

type CTX_ATTRIBUTES string

const (
	IS_DEPLOYABLE_CTX CTX_ATTRIBUTES = "is_deployable"
	DEPLOYABLE_REF                   = "deployable_ref"
	LOGGER_REF                       = "logger_ref"
	CONTROLLABLE_TYPE                = "controllable_type"
	CONTROLLABLE_NAME                = "controller_name"
	CTX_ID                           = "ctx_id"
)

type CONTROLLABLE_TYPES string

const (
	HTTP_SERVER CONTROLLABLE_TYPES = "HTTP_SERVER"
)
