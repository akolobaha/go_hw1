MESSAGE FORMATTER

Реализуйте интерфейс Formatter с методом Format, который возвращает отформатированную строку.
Определите структуру, удовлетворяющую интерфейсу Formatter: обычный текст (как есть), жирным шрифтом (** **), код (` `), курсив (_ _)

Опционально иметь возможность задать цепочку модификаторов
chainFormatter.AddFormatter(plainText)
chainFormatter.AddFormatter(bold)
chainFormatter.AddFormatter(code)