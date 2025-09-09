package main

import "fmt"

//В работе с БД:
//драйверы БД часто адаптируются под стандартный database/sql интерфейс.
//В веб-разработке:
//адаптеры для различных логгеров (например, использование стороннего логгера через стандартный log.Logger).

//Плюсы:
//Позволяет использовать сторонний/старый код без его изменения.
//Упрощает интеграцию несовместимых компонентов.
//Поддерживает принцип Single Responsibility (разделение обязанностей).
//Следует принципу Open/Closed (открыт для расширения, закрыт для изменения)

// Минусы:
// Увеличивает количество классов и усложняет структуру проекта.
// Может снизить производительность при большом количестве обёрток.
// При чрезмерном использовании код становится труднее сопровождать
type iPhone struct {
}

func (p iPhone) lighting() string {
	return fmt.Sprintf("Charging>>>")
}

type USB interface {
	canCharge() string
}

type lightingAdapter struct {
	adapter *iPhone
}

func (a lightingAdapter) canCharge() string {
	return a.adapter.lighting()
}

func main() {

	cable := &iPhone{}

	var adapter USB = &lightingAdapter{adapter: cable}

	fmt.Printf("Your iPhone: %v", adapter.canCharge())
}
