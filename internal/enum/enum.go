package enum

type EquipmentStatus string

const (
	StatusAvailable EquipmentStatus = "available"
	StatusInUse     EquipmentStatus = "in_use"
	StatusBroken    EquipmentStatus = "broken"
	StatusUnknown   EquipmentStatus = "unknown"
)
