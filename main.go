package main

import (
	"fmt"
)

// Структура для честного 3D-объекта (Блока)
type Block3D struct {
	Name    string
	X, Y, Z float64 // Позиция в 3D-пространстве
	W, H, D float64 // Размеры блока (Ширина, Высота, Глубина) для растягивания
}

// Структура для 2D-элемента (PNG-кнопка)
type Element2D struct {
	TexturePath string // Путь к PNG
	ActionName  string // Привязанное действие из скрипта ccgljh
}

// Состояние окон и вкладок редактора
type EditorState struct {
	ActiveTab    string       // "3D", "2D" или "SCRIPT"
	Blocks3D     []Block3D    // Все заспавненные 3D блоки
	Elements2D   []Element2D  // Все перетащенные PNG-картинки
	LeftPanel    []string     // Экранчик слева: список всех скриптов
	BottomPanel  []string     // Экранчик снизу: Explorer (файлы)
}

// Запуск визуального движка экранов
func StartGoEditor() {
	// Инициализируем интерфейс по твоей раскладке
	state := EditorState{
		ActiveTab:   "3D", // По умолчанию открывается честное 3D
		LeftPanel:   []string{"player.gd", "enemy.gd", "game_butter.cppc"},
		BottomPanel: []string{"game.logo", "project.cppc", "icon.png"},
	}

	fmt.Println("=================================================================")
	fmt.Println("  [Go-Script] МОДУЛЬ ВИЗУАЛЬНЫХ ЭКРАНОВ И 3D/2D РЕДАКТОРА ЗАПУЩЕН")
	fmt.Println("=================================================================")

	// Выводим левую и нижнюю панели, которые ты просил
	renderLayout(state)

	// 1. ДЕМОНСТРАЦИЯ ВКЛАДКИ 3D (Честное 3D: спавн и растягивание блоков)
	fmt.Printf("\n[Вкладка] Переключаемся на: %s\n", state.ActiveTab)
	state.SpawnBlock("Cube_1", 0, 0, 0)
	state.StretchBlock("Cube_1", 5.5, 2.0, 5.5) // Растягиваем блок по осям X, Y, Z

	// 2. ДЕМОНСТРАЦИЯ ВКЛАДКИ 2D (Перетаскивание PNG и создание кнопок действия)
	state.ActiveTab = "2D"
	fmt.Printf("\n[Вкладка] Переключаемся на: %s\n", state.ActiveTab)
	state.DragAndDropPNG("res://play_button.png", "START_GAME_ACTION")

	// 3. ДЕМОНСТРАЦИЯ ВКЛАДКИ СКРИПТОВ
	state.ActiveTab = "SCRIPT"
	fmt.Printf("\n[Вкладка] Переключаемся на: %s\n", state.ActiveTab)
	fmt.Println("[Вкладка Скрипт] Открыт главный скрипт gd для написания логики игры.")
	fmt.Println("=================================================================\n")
}

// Функция отрисовки интерфейса по твоей схеме (Лево, Снизу, Центр)
func renderLayout(s EditorState) {
	fmt.Println("\n--- РАСКЛАДКА ЭКРАНОВ РЕДАКТОРА ---")
	
	// Левая сторона
	fmt.Println("[ЛЕВАЯ ПАНЕЛЬ (Список скриптов)]:")
	for _, script := range s.LeftPanel {
		fmt.Printf("  • [Файл] %s\n", script)
	}

	// Нижняя сторона
	fmt.Println("\n[НИЖНЯЯ ПАНЕЛЬ (Экранчик Explorer)]:")
	for _, file := range s.BottomPanel {
		fmt.Printf("  📁 %s\n", file)
	}
	fmt.Println("----------------------------------")
}

// Функция спавна полноценного 3D блока
func (s *EditorState) SpawnBlock(name string, x, y, z float64) {
	newBlock := Block3D{Name: name, X: x, Y: y, Z: z, W: 1, H: 1, D: 1} // Изначальный размер 1x1x1
	s.Blocks3D = append(s.Blocks3D, newBlock)
	fmt.Printf("[3D-Редактор] Заспавнен новый 3D Блок: '%s' в координатах (X:%.1f, Y:%.1f, Z:%.1f)\n", name, x, y, z)
}

// Функция РАСТАЩИВАНИЯ / РАСТЯГИВАНИЯ 3D-блока
func (s *EditorState) StretchBlock(name string, newW, newH, newD float64) {
	for i, b := range s.Blocks3D {
		if b.Name == name {
			s.Blocks3D[i].W = newW
			s.Blocks3D[i].H = newH
			s.Blocks3D[i].D = newD
			fmt.Printf("[3D-Редактор] Блок '%s' УСПЕШНО РАСТЯНУТ! Новые размеры 3D -> Ширина: %.1f, Высота: %.1f, Глубина: %.1f\n", name, newW, newH, newD)
			return
		}
	}
}

// Функция перетаскивания PNG во вкладке 2D
func (s *EditorState) DragAndDropPNG(pngPath string, action string) {
	newElem := Element2D{TexturePath: pngPath, ActionName: action}
	s.Elements2D = append(s.Elements2D, newElem)
	fmt.Printf("[2D-Редактор] Перетащили картинку '%s' -> Она превращена в кнопку! Привязано действие скрипта: %s\n", pngPath, action)
}

func main() {
	StartGoEditor()
}
