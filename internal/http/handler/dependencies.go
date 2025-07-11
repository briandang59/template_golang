package handler

type Dependencies struct {
	Factory       *FactoryHandler
	Department    *DepartmentHandler
	EquipmentType *EquipementTypeHandler
	Equipment     *EquipementHandler
}
