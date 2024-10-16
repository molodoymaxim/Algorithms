package sorting

//	Алгоритм:
//	1. Проходи по списку от начала до конца.
//	2. Сравнивай каждый элемент с его следующим соседом.
//	3. Если элементы находятся в неправильном порядке (первый больше второго), меняй их местами.
//	4. Повторяй этот процесс, пока не останется элементов для сравнения.
//	5. За каждый проход самый большой элемент оказывается на своем правильном месте.
//	6. Алгоритм останавливается, когда за один проход не произошло ни одного обмена.
//
//	Сложность O(n^2)

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//	Алгоритм:
//	1. Выбор опорного элемента (pivot):
//		Выбери элемент из массива, который будет использоваться для разделения.
//		Это может быть первый, последний, случайный элемент или медиана.
//	2. Разделение (partitioning):
//		Массив перестраивается так, чтобы все элементы, меньшие опорного,
//		оказались слева, а все элементы, большие опорного, — справа.
//		Сам опорный элемент окажется на своей правильной позиции.
//	3. Рекурсия:
//		Теперь для двух частей массива (левее и правее от опорного элемента) применяется
//		та же самая процедура — выбирается новый опорный элемент, выполняется разделение,
//		и так продолжается до тех пор, пока в каждой подгруппе не останется один или ноль элементов.
//	4. Слияние:
//		После завершения всех рекурсивных вызовов массив будет полностью отсортирован.

// 	Сложность: O(n log(n))

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := []int{}
	right := []int{}

	for _, val := range arr[:len(arr)-1] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

//	Aлгоритм:
//	1. Начинаем с первого элемента (он считается отсортированным).
//	2. Берем следующий элемент и перемещаем его влево в отсортированную часть массива, пока не найдем правильное место для вставки.
//	3. Повторяем этот процесс для каждого элемента массива, пока весь массив не будет отсортирован.

// 	Сложность: O(n^2)

func InsertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		elem := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > elem {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = elem
	}
	return arr
}

//	Aлгоритм:
//	1. Разделить массив пополам на два подмассива.
//	2. Рекурсивно отсортировать каждый подмассив.
//	3. После того как два подмассива отсортированы, объединить их в один отсортированный массив.

// 	Сложность: O(n log(n))

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

//	Алгоритм:
//	1. Построение максимальной кучи:
//		Массив преобразуется в максимальную кучу, где каждый элемент больше или равен своим потомкам.
//	2. Извлечение максимального элемента:
//		Максимальный элемент (корень кучи) перемещается в конец массива.
//	3. Восстановление кучи:
//		Оставшаяся часть массива снова преобразуется в кучу, и процесс повторяется для оставшихся элементов.

// 	Сложность: O(n log(n))

func HeapSort(arr []int) {
	n := len(arr)

	// Построение максимальной кучи
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Извлечение элементов из кучи один за другим
	for i := n - 1; i > 0; i-- {
		// Перемещаем текущий корень (максимальный элемент) в конец
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

// Восстанавление свойства кучи для поддерева, корень которого находится в node
func heapify(arr []int, n int, node int) {
	largest := node     // Инициализируем наибольший элемент как корень
	left := 2*node + 1  // Левый потомок
	right := 2*node + 2 // Правый потомок

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != node {
		arr[node], arr[largest] = arr[largest], arr[node]

		heapify(arr, n, largest)
	}
}
