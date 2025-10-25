package response

const (
	CodeSuccess = 200

	// User related codes
	CodeRegisterInternalError = 2001
	CodeUserAlreadyExists     = 2002
	CodeUserNotFound          = 2003
	CodeFailedGetUser         = 2004
	CodeFailedUpdateUser      = 2005
	CodeInvalidInput          = 2006
	CodeInvalidOTP            = 2007
	CodeOTPExpired            = 2008
	CodeEmailNotVerified      = 2009
	CodePhoneNotVerified      = 2010
	CodeInvalidEmail          = 2011
	CodeInvalidUsername       = 2012
	CodeEmptyPassword         = 2013

	// Mail related codes
	CodeMailConfigMissing    = 3001
	CodeMailUsernameMissing  = 3002
	CodeMailPasswordMissing  = 3003
	CodeMailClientCreation   = 3004
	CodeMailConnectionFailed = 3005
	CodeMailSendFailed       = 3006

	//
)
