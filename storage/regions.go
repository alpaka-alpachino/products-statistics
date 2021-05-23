package storage

type District struct {
	Name  string
	Index int
}

type Product struct {
	name       string
	goodPlaces []District
}

var (
	Vinnytska       = District{"Вінницька", 1977}
	Volynska        = District{"Волинська", 1200}
	Dniprovska      = District{"Дніпропетровська", 2220}
	Zhytomyrska     = District{"Житомирська", 993}
	Zakarpatska     = District{"Закарпатська", 2054}
	Zaporizka       = District{"Запоріжжя", 2256}
	IvanoFrankivska = District{"Івано-Франківська", 2000}
	Kyivska         = District{"Київська", 2858}
	Kirovogradska   = District{"Кіровоградська", 1760}
	Lvivska         = District{"Львівська", 2210}
	Mykolaivska     = District{"Миколаївська", 975}
	Odeska          = District{"Одеська", 2488}
	Poltavska       = District{"Полтавська", 4346}
	Rivnenska       = District{"Рівненська", 1202}
	Sumska          = District{"Сумська", 2943}
	Ternopliska     = District{"Тернопільська", 2013}
	Harkivska       = District{"Харківська", 2176}
	Hersonska       = District{"Херсонська", 1213}
	Hmelnytska      = District{"Хмельницька", 2153}
	Cherkaska       = District{"Черкаська", 2885}
	Chernivetsca    = District{"Чернівецька", 3020}
	Chernigivska    = District{"Чернігівська", 1300}
)

var ProductRegions = map[string][]District{
	"борошно":  {Vinnytska, Kyivska, Kirovogradska, Ternopliska, Harkivska, Cherkaska},
	"олія":     {Vinnytska, IvanoFrankivska, Kirovogradska, Poltavska, Sumska, Harkivska, Hmelnytska, Cherkaska},
	"Молочні":  {Vinnytska, Zhytomyrska, Zaporizka, Lvivska, Poltavska, Hmelnytska, Chernigivska},
	"М'ясні":   {Vinnytska, Volynska, Zaporizka, Dniprovska, Kyivska, Lvivska, Cherkaska},
	"гречка":   {Vinnytska, Kyivska, Sumska, Hmelnytska, Cherkaska},
	"Овочі":    {Volynska, Dniprovska, Zhytomyrska, Zakarpatska, Kyivska, Mykolaivska, Poltavska, Ternopliska, Hersonska},
	"пшоно":    {Dniprovska, Zaporizka, Mykolaivska},
	"цукор":    {Zhytomyrska, IvanoFrankivska, Kyivska, Lvivska, Mykolaivska, Hmelnytska},
	"картопля": {Zhytomyrska, Sumska, Hmelnytska},
	"Яйця":     {Kyivska, Harkivska, Hersonska, Hmelnytska},
	"рис":      {Odeska, Hersonska},
	"Фрукти":   {Odeska, Poltavska, Rivnenska, Chernivetsca},
}
