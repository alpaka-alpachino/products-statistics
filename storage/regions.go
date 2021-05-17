package storage

type District struct {
	name  string
	index float64
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

var (
	Wheat      = Product{"пшениця", []District{Vinnytska, Kyivska, Kirovogradska, Ternopliska, Harkivska, Cherkaska}}
	Corn       = Product{"соняшник", []District{Vinnytska, IvanoFrankivska, Kirovogradska, Poltavska, Sumska, Harkivska, Hmelnytska, Cherkaska}}
	Milk       = Product{"молоко", []District{Vinnytska, Zhytomyrska, Zaporizka, Lvivska, Poltavska, Hmelnytska, Chernigivska}}
	Meat       = Product{"м'ясо", []District{Vinnytska, Volynska, Zaporizka, Dniprovska, Kyivska, Lvivska, Cherkaska}}
	Buckwheat  = Product{"гречка", []District{Vinnytska, Kyivska, Sumska, Hmelnytska, Cherkaska}}
	Vegetables = Product{"овочі", []District{Volynska, Dniprovska, Zhytomyrska, Zakarpatska, Kyivska, Mykolaivska, Poltavska, Ternopliska, Hersonska}}
	Millet     = Product{"пшоно", []District{Dniprovska, Zaporizka, Mykolaivska}}
	SugarBeets = Product{"цукрові буряки", []District{Zhytomyrska, IvanoFrankivska, Kyivska, Lvivska, Mykolaivska, Hmelnytska}}
	Potato     = Product{"картопля", []District{Zhytomyrska, Sumska, Hmelnytska}}
	Eggs       = Product{"яйця", []District{Kyivska, Harkivska, Hersonska, Hmelnytska}}
	Rice       = Product{"рис", []District{Odeska, Hersonska}}
	Fruits     = Product{"плоди та ягоди", []District{Odeska, Poltavska, Rivnenska, Chernivetsca}}
)
