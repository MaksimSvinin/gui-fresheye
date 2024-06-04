# gui-fresheye

Программа для поиска близкорасположенных слов на русском языке.

За основу была взята https://gitlab.com/opennota/fresheye

## Как использовать программу

Левое поле предназначено для вставки и редактирования текста.Текст для анализа можно либо вставить в левое поле,либо импортировать через меню file -> Open File

Сейчас поддерживается только файлы формата txt (поддержка doc docs в планах)

Затем для анализа текста нужно нажать кнопку "Анализ" в области справа.

При этом будут выведены чекбоксы с найденными словами.

Затем можно выбрать интересующие слова и при нажатии кнопки "Показать выделенные слова" на правое текстовое поле будет выведен текст с цветовым отображением найденных слов.

Усли активировать чекбокс "Показывать только близкие слова" то отображены будут только слова расположенные в пределах количества символов указанных в верхнем поле "Растояние между парами слов в символах"

Чекбокс "Выделить всё" позволяет выбрать все найденные слова

Чекбокс "win1251 кодировка" нужен для чтения файлов созданных в системе windows (в кодировке win1251).Без этого чекбокса будет использоваться unicode (utf-8)
