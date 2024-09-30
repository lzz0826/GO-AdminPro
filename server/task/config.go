package task

var (

	//验证前四
	//Data: []byte{0x44, 0x5A, 0x50, 0x4B,}
	HEADER_INDICATER_0 byte = 'D'
	HEADER_INDICATER_1 byte = 'Z'
	HEADER_INDICATER_2 byte = 'P'
	HEADER_INDICATER_3 byte = 'K'

	/**
	 * ************************** Client Request Package ******************************
	 */

	RP_PROTOCOL            = 4
	RP_USER_ID_1           = 5
	RP_USER_ID_2           = 6
	RP_USER_ID_3           = 7
	RP_USER_ID_4           = 8
	RP_LANGUAGE_ID         = 9
	RP_CLIENT_PLATFORM     = 10
	RP_CLIENT_BUILD_NUMBER = 11
	RP_CUSTOM_ID_1         = 12
	RP_CUSTOM_ID_2         = 13
	RP_PRODUCT_ID_1        = 14
	RP_PRODUCT_ID_2        = 15
	RP_REQUEST_CODE_HIGH   = 16
	RP_REQUEST_CODE_LOW    = 17
	RP_SIZE_HIGH           = 18
	RP_SIZE_LOW            = 19
	RP_TOKEN_HIGH          = 20

	/**
	 * ************************** Server Response Package ****************************************
	 */
	//Server Response Package Header
	SRP_TYPE = 4

	//     SRP_SIZE_HIGH = 5;
	//     SRP_SIZE_LOW = 6;

	SRP_FEE_NEXT_YEAR           = 20
	SRP_FEE_NEXT_MONTH          = 21
	SRP_FEE_NEXT_DATE           = 22
	SRP_DATA_FEE_DATA_SIZE      = 31
	SRP_DATA_FEE_PAGE_DATA_SIZE = 33
	SRP_DATA_FEE_FILE_SIZE      = 35
	SRP_DATA_START              = 51
	SRP_PACKE_LEVEL             = 6
	SRP_REQUEST_HIGH            = 4
	SRP_REQUEST_LOW             = 5
	SRP_SIZE_HIGH               = 7
	SRP_SIZE_LOW                = 8

	/**
	 * ***************************** INPUT*******************************************************
	 *
	 */
	IMSI_CARD           = "imsi="
	SMS_CENTER_NUMBER   = "sms_sc="
	SMS_NUMBER          = "sms_num="
	SMS_CONTENT         = "sms_cnt="
	SMS_SEND_UP         = "sms_up="
	SMS_SEND_UP_SP      = "sp="
	SMS_SEND_UP_COUNT   = "count="
	SMS_SEND_UP_CONTENT = "content="

	/**
	 * ******************************* max fee size**************************************************
	 */
	MAX_FEE = 8

	/**
	 * ******************************* secondType return **************************************************
	 */
	SECOND_TYPE_NORMAL        = 1
	SECOND_TYPE_NORMAL_RETURN = "是"
	SECOND_TYPE_CHANGE        = 2
)
