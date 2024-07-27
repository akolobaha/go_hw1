package main

// IsReflectMatrix
/*
1. Проверка на пустую матрицу:
- Временная сложность: O(1)

2. Циклы для проверки симметричности:
- Внешний цикл (по строкам) выполняется n раз.
- Внутренний цикл (по элементам строки) также выполняется n раз для каждой строки.
- Общее количество операций для двух вложенных циклов:
- Внешний цикл: i = 0 до n-1 (n итераций)
- Внутренний цикл: j = 0 до n-1 (n итераций для каждой итерации внешнего цикла)

Таким образом, общее количество операций в двойном цикле можно выразить как:
- T(n) = n * n = n²

Итак, общая временная сложность алгоритма:
- T(n) = O(1) + O(n²) = O(n²)

*/
func IsReflectMatrix(a [][]int) bool {
	n := len(a)
	if n == 0 {
		return true
	}

	for i := 0; i < n-1; i++ {
		if len(a[i]) != n {
			return false
		}
		for j := 0; j < n; j++ {
			if a[i][j] != a[j][i] {
				return false
			}
		}
	}
	return true
}

// MinEl
/*

Функция реализует рекурсивный подход, где каждый вызов функции обрабатывает последний элемент массива и вызывает себя для оставшейся части массива. Это приводит к следующему:
- Количество вызовов функции: Для среза длины n функция будет вызываться n раз, так как на каждом шаге она уменьшает размер среза на 1.
- Сравнения: На каждом уровне рекурсии функция выполняет одно сравнение (сравнение текущего элемента с минимальным элементом, найденным в предыдущих вызовах).
Таким образом, временная сложность T(n) может быть выражена как:

T(n) = T(n-1) + O(1)

Решая это уравнение, получаем:

T(n) = O(n)
*/
func MinEl(a []int) int {
	// только для первичной проверки
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	t := MinEl(a[:len(a)-1])
	if t <= a[len(a)-1] {
		return t
	}
	return a[len(a)-1]
}

/*
Представлен другой рекурсивный алгоритм для поиска наименьшего элемента слайса.
Вычислите его сложность, используя изученные формулы.
Сравните эффективность алгоритма с предыдущим вариантом.
Попробуйте увеличить скорость выполнения функции, используя инструменты Го.
Как изменится сложность алгоритма при увеличении скорости?
*/

/*
Рекурсивная функция `MinEl2` имеет следующую структуру:

1. Если длина массива равна 0, возвращается 0.
2. Если длина массива равна 1, возвращается единственный элемент.
3. В противном случае:
   - Разделяем массив на две половины.
   - Вызываем `MinEl2` для первой половины.
   - Вызываем `MinEl2` для второй половины.
   - Сравниваем результаты и возвращаем меньший.

Рекурсивное уравнение

Если обозначить T(n) как время выполнения алгоритма для массива длины n, то можно записать следующее уравнение:

T(n) = 2T(n/2) + O(1)

где O(1) — это время, необходимое для сравнения двух элементов и выполнения других операций.

Решение рекурсивного уравнения

Это уравнение можно решить с помощью метода подстановки или используя теорему о рекурсии (теорема мастер):

1. Известно, что T(n) = 2T(n/2) + O(1) соответствует случаю 2 в теореме о рекурсии, где a = 2, b = 2 и f(n) = O(1).
2. По теореме мастер, решение этого уравнения будет:

T(n) = O(n)

Итоговая сложность

Таким образом, алгоритм `MinEl2` имеет линейную временную сложность O(n), что означает, что время выполнения алгоритма будет пропорционально количеству элементов в массиве.


*/

func MinEl2(a []int) int {
	// только для первичной проверки
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return a[0]
	}
	t1 := MinEl2(a[:len(a)/2])
	t2 := MinEl2(a[len(a)/2:])
	if t1 <= t2 {
		return t1
	}
	return t2
}

// MinEl3
/*
Не создается доп слайсов и нет рекурсии

Временная сложность данного алгоритма составляет O(n), так как мы проходим по каждому элементу массива ровно один раз.
*/
func MinEl3(a []int) int {
	if len(a) == 0 {
		return 0
	}

	minimal := a[0]
	for _, value := range a {
		if value < minimal {
			minimal = value
		}
	}
	return minimal
}
