package codelist

const (
	ESV_SetI       = 0x60
	ESV_SetC       = 0x61
	ESV_Get        = 0x62
	ESV_INF_REQ    = 0x63
	ESV_SetGet     = 0x6E
	ESV_Set_Res    = 0x71
	ESV_Get_Res    = 0x72
	ESV_INF        = 0x73
	ESV_INFC       = 0x74
	ESV_INFC_Res   = 0x7A
	ESV_SetGet_Res = 0x7E
	ESV_SetI_SNA   = 0x50
	ESV_SetC_SNA   = 0x51
	ESV_Get_SNA    = 0x52
	ESV_INF_SNA    = 0x53
	ESV_SetGet_SNA = 0x5E
)

const (
	Group_SensorRelatedDevice              = 0x00
	Group_AirConditionerRelatedDevice      = 0x01
	Group_HousingFacilitiesRelatedDevice   = 0x02
	Group_CookingHouseholdRelatedDevice    = 0x03
	Group_HealthRelatedDevice              = 0x04
	Group_ManagementOperationRelatedDevice = 0x05
	Group_AudiovisualRelatedDevice         = 0x06
	Group_ProfileObject                    = 0x0E
)

const (
	Class_HotWaterGenerator = 0x72 // Group: 0x02
	Class_NodeProfile       = 0xF0 // Group: 0x0E
)
