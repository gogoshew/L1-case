package tasks

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.


Паттерн Adapter относится к структурным паттернам уровня класса.
Часто в новом проекте разработчики хотят повторно использовать уже существующий код.
Например, имеющиеся классы могут обладать нужной функциональностью и иметь при этом несовместимые интерфейсы.
В таких случаях следует использовать паттерн Adapter.
*/

// Пытаемся подключить чайник к MacBook

// TeaPoter Создаем интерфейс и присваиваем ему метод TeapotConnect
type TeaPoter interface {
	TeapotConnect()
}

// Electricity Структура электросети от которой обычно работает чайник
type Electricity struct {
}

// TeapotConnectToElectricity принимает интерфейс TeaPoter
func (e *Electricity) TeapotConnectToElectricity(tea TeaPoter) {
	fmt.Println("User connecting teapot to electricity")
	tea.TeapotConnect()
}

type Socket struct {
}

func (s *Socket) TeapotConnect() {
	fmt.Println("Teapot successfully working with electricity socket!")
}

type MacBook struct {
}

func (mb *MacBook) InsertToTypeCPort() {
	fmt.Println("Type-C port plugged to MacBook and now teapot is boiling!")
}

type MacBookAdapter struct {
	teapotAdapter *MacBook
}

func (ma *MacBookAdapter) TeapotConnect() {
	fmt.Println("Adapter converts electricity signal to Type-C")
	ma.teapotAdapter.InsertToTypeCPort()
}

/*
У нас есть структура Electricity с методом TeapotConnectToElectricity, которая принимает интерфейс чайника
Затем мы создали 2 структуры Socket и MacBook, Socket ожидает электричество, чтобы можно было подключить чайник
и методы Socket удовлетворяют интерфейсу чайника, MacBook же имеет только Type-C порт, соответственно, чтобы
подключить к нему чайник, нам нужен MacBookAdapter, эти Макбуки, только адаптеры им и подавай...

Создаем MacBookAdapter, в котором поле teapotAdapter имеет тип MacBook и присваиваем адаптеру метод TeapotConnect,
чтобы он тоже удовлетворял интерфейсу.

Таким образом, мы теперь можем вскипятить чайник от любого электричества методом TeapotConnectToElectricity,
поскольку Socket и так был для этого предназначен, а благодаря MacBookAdapter, мы теперь можем приготовить чаек
и от MacBook :)
*/

func Task21() {
	electronic := &Electricity{}
	socket := &Socket{}

	electronic.TeapotConnectToElectricity(socket)

	mac := &MacBook{}
	macAdapter := &MacBookAdapter{
		teapotAdapter: mac,
	}

	electronic.TeapotConnectToElectricity(macAdapter)
}
