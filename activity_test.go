package SendPriorityMessageToAssets

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())
	staffObj:= "{\"Address\":null,\"AlarisStatus\":null,\"AlertStatus\":\"Blue\",\"AssociatedDevices\":[{\"Assigned\":\"Assigned\",\"BatteryLevel\":88,\"BatteryStatus\":\"Green\",\"ConfigId\":54,\"CustomerId\":2047,\"CustomerName\":\"\",\"DateCreated\":\"2023-09-21T10:04:15.28\",\"DateUpdated\":\"2023-11-28T16:56:01.857\",\"DefaultConfig\":false,\"Description\":\"\",\"DeviceHealth\":\"{\\r\\n\\t\\\"BLE\\\": false, \\r\\n\\t\\\"Button1\\\": false, \\r\\n\\t\\\"Button2\\\": false,\\r\\n\\t\\\"Button3\\\": false, \\r\\n\\t\\\"Display\\\": false, \\r\\n\\t\\\"Motion\\\": false, \\r\\n\\t\\\"Temperature\\\": false, \\r\\n\\t\\\"SafetySwitch\\\": false\\r\\n\\t}\",\"ElapsedTimeInMillseconds\":0,\"EnableHygiene\":false,\"EnableSDCT\":false,\"EnableTenancy\":false,\"ErrorMessage\":\"\",\"ExpirationDate\":\"0001-01-01T00:00:00\",\"FirmwareVersion\":\"6.169.589\",\"GroupId\":300000103,\"HasError\":false,\"Id\":92208,\"InfraGroupId\":0,\"InfraName\":\"\",\"IntegrationConfigId\":0,\"ItemId\":400001909,\"Latitude\":0,\"LocationUpdated\":\"2023-11-28T21:56:01.983\",\"Longitude\":0,\"MaintenanceUpdated\":\"2023-11-28T21:43:05.373\",\"MapId\":4796,\"ModelDeviceType\":\"\",\"ModelName\":\"B4\",\"Name\":\"C4:CB:6B:23:24:AC\",\"PendingConfigId\":0,\"PersonAssociated\":\"COE Security Guard Brandon 1\",\"PhoneNumber\":\"\",\"PurchaseDate\":\"0001-01-01T00:00:00\",\"PurchaseOrderNumber\":\"\",\"ShippedDate\":\"0001-01-01T00:00:00\",\"Status\":\"In Use\",\"StatusMessage\":\"{\\\"DataBusReports\\\":[],\\\"DataCorrelations\\\":[],\\\"DeviceReports\\\":[{\\\"Attributes\\\":[],\\\"AUTOID_EPC\\\":\\\"MAC_C4:CB:6B:23:24:AC\\\",\\\"DataTimestamp\\\":\\\"1970-01-01T00:00:00\\\",\\\"DeviceType\\\":\\\"RFID_ACTIVE_802_11\\\",\\\"DeviceUniqueID\\\":\\\"C4:CB:6B:23:24:AC\\\",\\\"DeviceUniqueID_DisplayName\\\":\\\"MAC\\\",\\\"Events\\\":[],\\\"Item\\\":{\\\"DateCreated\\\":\\\"1970-01-01T00:00:00\\\",\\\"DateUpdated\\\":\\\"1970-01-01T00:00:00\\\"},\\\"LastEvent\\\":{\\\"StartDateTime\\\":\\\"2023-11-28T16:56:01.9601264-05:00\\\",\\\"SystemName\\\":\\\"ARC\\\"},\\\"LastEventIndex\\\":-1,\\\"MessageId\\\":\\\"1701208606913\\\",\\\"MessageType\\\":\\\"RFID_ACTIVE_802_11\\\",\\\"ReceivedTimestamp\\\":\\\"2023-11-28T21:56:01.9591405Z\\\",\\\"RFID\\\":{\\\"Readers\\\":[]},\\\"RTLSAddress\\\":{},\\\"RTLSContact\\\":{},\\\"RTLSGeo\\\":{},\\\"RTLSGPS\\\":{},\\\"RTLSModel2D\\\":{\\\"IsValid\\\":true,\\\"PosMapID\\\":4796,\\\"PosModelID\\\":4796,\\\"PosX\\\":2131.0,\\\"PosY\\\":1036.0,\\\"PosZoneIDs\\\":\\\"[{\\\\\\\"ZoneID\\\\\\\":44841,\\\\\\\"ZoneName\\\\\\\":\\\\\\\"Services\\\\\\\",\\\\\\\"ZoneType\\\\\\\":\\\\\\\"Open\\\\\\\"}]\\\",\\\"Timestamp\\\":\\\"2023-11-28T21:56:01.9601264Z\\\"},\\\"Sensor\\\":{},\\\"SequenceNumber\\\":-1,\\\"SiteID\\\":\\\"Xpert AoA Receiver\\\",\\\"Status\\\":{\\\"BatteryLevel1\\\":88.0,\\\"DeviceReportReason\\\":3,\\\"IsValid\\\":true,\\\"Timestamp\\\":\\\"2023-11-28T21:56:01.9601264Z\\\"},\\\"DateCreated\\\":\\\"2023-11-28T21:56:01.9591405Z\\\",\\\"DateUpdated\\\":\\\"2023-11-28T21:56:01.9591405Z\\\"}],\\\"ItemInfo\\\":{\\\"DateCreated\\\":\\\"1970-01-01T00:00:00\\\",\\\"DateUpdated\\\":\\\"1970-01-01T00:00:00\\\"},\\\"ProximityReports\\\":[],\\\"SchemaName\\\":\\\"XpertSchema.XpertMessage.XpertMessage\\\",\\\"SchemaVersion\\\":\\\"1\\\",\\\"ZoneID\\\":0}\",\"SuccessMessage\":\"\",\"Temperature\":-273,\"TenantId\":\"\",\"UnderWarranty\":false,\"UniqueId\":\"C4:CB:6B:23:24:AC\",\"XCoordinate\":\"2131.0\",\"YCoordinate\":\"1036.0\"}],\"BatteryLevel\":0,\"BedStatus\":\"\",\"CurrentBuildingID\":4795,\"CurrentBuildingName\":\"Sparks HQ\",\"CurrentSiteName\":\"\",\"CurrentFloorName\":\"New Sparks Office Floorplan Zoomed\",\"CurrentMapId\":4796,\"CurrentModelId\":4796,\"CurrentSiteID\":4794,\"CurrentTimestamp\":\"2023-11-28T21:56:01.96\",\"CurrentX\":2131,\"CurrentY\":1036,\"CurrentZones\":\"[{\\\"ZoneID\\\":44841,\\\"ZoneName\\\":\\\"Services\\\",\\\"ZoneType\\\":\\\"Open\\\"}]\",\"DepartmentID\":100000018,\"DepartmentName\":\"Security\",\"DeviceID\":92208,\"DeviceLogID\":24557164,\"DeviceName\":\"C4:CB:6B:23:24:AC\",\"Email\":\"\",\"EnableAlerts\":true,\"EnableHygiene\":false,\"EnableSDCT\":false,\"EventCountAcknowledged\":0,\"EventCountClosed\":1169,\"EventCountNew\":0,\"EventCountOpen\":0,\"FromLDAP\":false,\"GroupID\":300000103,\"GroupName\":\"Security\",\"HealthStatus\":\"\",\"Icon\":\"assets\\\\icons\\\\star.png\",\"ImageData\":\"\",\"ImageType\":\"\",\"IsTestMode\":false,\"OldTamper\":false,\"OldMotion\":false,\"Latitude\":0,\"Longitude\":0,\"LocationUpdated\":\"2023-11-28T21:56:01.993\",\"ModelName\":\"\",\"OldBuildingID\":4795,\"OldMapId\":4796,\"OldModelId\":4796,\"OldSiteID\":4794,\"OldLocationUpdated\":\"2023-11-28T21:55:48.71\",\"OldX\":2359,\"OldY\":778,\"OldZones\":\"[{\\\"ZoneID\\\":44829,\\\"ZoneName\\\":\\\"Conference Room\\\",\\\"ZoneType\\\":\\\"Open\\\"}]\",\"PendingDepartmentDateUpdated\":\"0001-01-01T00:00:00\",\"PendingDepartmentId\":0,\"PhoneNumber\":\"\",\"Portrait\":\"assets\\\\icons\\\\PATIENT_F_4.png\",\"StaffID\":\"9876543\",\"StaffSettings\":null,\"Temperature\":\"\",\"UseCases\":[],\"MultiAssign\":false,\"AssocItemID\":0,\"PendingDepartmentName\":null,\"AssocItemName\":null,\"CustomerId\":2047,\"DateCreated\":\"2023-09-29T13:40:12.837\",\"DateUpdated\":\"2023-09-29T13:40:12.837\",\"Description\":\"\",\"EnableTenancy\":false,\"Name\":\"COE Security Guard Brandon 1\",\"TenantId\":\"\",\"ElapsedTimeInMillseconds\":0,\"ErrorMessage\":\"\",\"SuccessMessage\":\"\",\"HasError\":false,\"Id\":400001909}"
	input := &Input{IP: "3.212.201.170:802", CustomerId: "2047", Username: "afadmin", Password: "admin", StaffIdList: staffObj, Message: "Test Prio Msg OBJ 2"}
	// StaffIdList "9064" OR "9064,37685" OR staffObj
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{} 
	err = tc.GetOutputObject(output)
	assert.Nil(t, err)

	assert.Equal(t, "true", output.Status)
}
