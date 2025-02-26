package sysinfo

type OSVariant uint8

const (
	OSVariantVxWorks5 OSVariant = 1
	OSVariantVxWorks7 OSVariant = 2
)

func (o OSVariant) String() string {
	switch o {
	case 1:
		return "VxWorks 5"
	case 2:
		return "VxWorks 7"
	}
	return "Unknown"
}

type TimeOnOffSetting int32

const (
	TimeOnOffSettingOff     TimeOnOffSetting = 0
	TimeOnOffSettingOn      TimeOnOffSetting = 1
	TimeOnOffSettingReset   TimeOnOffSetting = 2
	TimeOnOffSettingRequest TimeOnOffSetting = 3
)

type CPUUsageOnOffSetting int32

const (
	CPUUsageOnOffSettingSettingOff     CPUUsageOnOffSetting = 0
	CPUUsageOnOffSettingSettingOn      CPUUsageOnOffSetting = 1
	CPUUsageOnOffSettingSettingReset   CPUUsageOnOffSetting = 2
	CPUUsageOnOffSettingSettingRequest CPUUsageOnOffSetting = 3
)

type CPUUsageMeasurementMode int32

const (
	CPUUsageMeasurementModeContinuous CPUUsageMeasurementMode = 1
	CPUUsageMeasurementModeReset      CPUUsageMeasurementMode = 2
)
