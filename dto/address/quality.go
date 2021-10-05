package address

const (
	// UniqQU Unique/Уникален: предлагается 1 эталонный адрес
	UniqQU UniqQ = iota
	// UniqQD Doubtful/Сомнителен: предлагается несколько эталонных адресов близких по написанию (возможен выбор)
	UniqQD
	// UniqQNU Not unique/Неуникален: есть несколько эталонных записей, в равной степени соответствующих исходному адресу
	UniqQNU

	// ActQA Accurate/Найдено по актуальной записи: название и административное подчинение, указанные в разбираемом адресе, соответствуют эталонному
	ActQA = iota
	// ActQR Rename/Переименование: устаревшее название одного из адресных элементов, указанных в разбираемом адресе
	ActQR
	// ActQRA Reassignment/Переподчинение: административное подчинение, указанное в разбираемом адресе, устарело
	ActQRA

	// UndfQN No/Нет
	UndfQN UndfQ = iota
	// UndfQIS Insignificant/Малозначимый: информация, не влияющая на результаты распознавания при ручной проверке
	UndfQIS
	// UndfQS Significant/Значимый: информация, которая при ручной проверке может повлиять на результат сравнения разбираемого адреса с эталоном
	UndfQS

	// LvlQRegion To the region(state)/До региона
	LvlQRegion LvlQ = iota + 1
	// LvlQDistrict To the district/До района
	LvlQDistrict
	// LvlQCity To the city/До города
	LvlQCity
	// LvlQCityArea To the district in the city/До района в городе
	LvlQCityArea
	// LvlQSettlement To the settlement/До населенного пункта
	LvlQSettlement
	// LvlQPlanStruct To the planning structure/До планировочной структуры
	LvlQPlanStruct
	// LvlQStreet To the street/До улицы
	LvlQStreet
	// LvlQHouse To the house/До дома
	LvlQHouse

	// HouseQNF Not found variants/Не найдено вариантов
	HouseQNF = 0
	// HouseQA Precise definition of the house by accurate entry/Точное определение дома по эталону
	HouseQA = 3
	// HouseQP Partial definition of a house by accurate entry/Частичное определение дома по эталону
	HouseQP = 4
	// HouseQNH The parsed address is missing a house number/В разбираемом адресе отсутствует номер дома
	HouseQNH = 9

	// GeoQRegion To the region(state)/До региона
	GeoQRegion GeoQ = iota + 1
	// GeoQDistrict To the district/До района
	GeoQDistrict
	// GeoQCity To the city/До города
	GeoQCity
	// GeoQCityArea To the district in the city/До района в городе
	GeoQCityArea
	// GeoQSettlement To the settlement/До населенного пункта
	GeoQSettlement
	// GeoQPlanStruct To the planning structure/До планировочной структуры
	GeoQPlanStruct
	// GeoQStreet To the street/До улицы
	GeoQStreet
	// GeoQHouse To the house/До дома
	GeoQHouse
)

// UniqQ Уровень уникальности найденного адреса
type UniqQ int

// Values Все возможные значения
func (q UniqQ) Values() []UniqQ {
	return []UniqQ{UniqQU, UniqQD, UniqQNU}
}

// ActQ Статус актуальности исходного адреса
type ActQ int

// Values Все возможные значения
func (q ActQ) Values() []ActQ {
	return []ActQ{ActQA, ActQR, ActQRA}
}

// UndfQ Разбор неадресной информации в исходном адресе
type UndfQ int

// Values Все возможные значения
func (q UndfQ) Values() []UndfQ {
	return []UndfQ{UndfQN, UndfQIS, UndfQS}
}

// LvlQ Уровень, до которого произведено сравнение исходного адреса с эталоном
type LvlQ int

// Values Все возможные значения
func (q LvlQ) Values() []LvlQ {
	return []LvlQ{
		LvlQRegion,
		LvlQDistrict,
		LvlQCity,
		LvlQCityArea,
		LvlQSettlement,
		LvlQPlanStruct,
		LvlQStreet,
		LvlQHouse,
	}
}

// HouseQ Степень совпадения номера дома в исходном адресе с эталоном
type HouseQ int

// Values Все возможные значения
func (q HouseQ) Values() []HouseQ {
	return []HouseQ{HouseQNF, HouseQA, HouseQP, HouseQNH}
}

// GeoQ Уровень, до которого разобраны координаты адреса
type GeoQ int

// Values Все возможные значения
func (q GeoQ) Values() []GeoQ {
	return []GeoQ{
		GeoQRegion,
		GeoQDistrict,
		GeoQCity,
		GeoQCityArea,
		GeoQSettlement,
		GeoQPlanStruct,
		GeoQStreet,
		GeoQHouse,
	}
}
