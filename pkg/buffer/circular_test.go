package buffer

import (
	"testing"

	"github.com/zikwall/app_metrica/pkg/x"
)

func TestIsEmpty(t *testing.T) {
	// Создаем экземпляр CircularBuffer
	buf := NewCircularBuffer(5)

	// Проверяем, что новый буфер является пустым
	if !buf.IsEmpty() {
		t.Errorf("Expected buffer to be empty, but it is not empty")
	}
}

func TestEnqueue(t *testing.T) {
	// Создаем экземпляр CircularBuffer
	buf := NewCircularBuffer(5)

	// Вводим данные в буфер
	err := buf.Enqueue([]byte("data"))
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	// Проверяем, что буфер теперь не является пустым
	if buf.IsEmpty() {
		t.Errorf("Expected buffer to be non-empty after enqueue, but it is empty")
	}
}

func TestDequeue(t *testing.T) {
	// Создаем экземпляр CircularBuffer
	buf := NewCircularBuffer(5)

	// Пытаемся удалить элемент из пустого буфера
	_, err := buf.Dequeue()
	if err != ErrBufferIsEmpty {
		t.Errorf("Expected error: %v, but got: %v", ErrBufferIsEmpty, err)
	}

	// Вводим данные в буфер
	data := []byte("data")
	err = buf.Enqueue(data)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	// Удаляем элемент из буфера
	dequeuedData, err := buf.Dequeue()
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	// Проверяем, что удаленный элемент соответствует введенным данным
	if string(dequeuedData) != string(data) {
		t.Errorf("Expected dequeued data: %s, but got: %s", string(data), string(dequeuedData))
	}

	// Проверяем, что буфер теперь снова пустой
	if !buf.IsEmpty() {
		t.Errorf("Expected buffer to be empty after dequeue, but it is not empty")
	}
}

func TestCircularBuffer(t *testing.T) {
	// Создаем новый буфер
	b := NewCircularBuffer(2)

	// Проверяем, что буфер пустой
	if !b.IsEmpty() {
		t.Errorf("expected buffer to be empty")
	}

	// Добавляем элементы в буфер
	err := b.Enqueue([]byte("data"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = b.Enqueue([]byte("data_2"))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Пробуем добавить элемент в полный буфер
	err = b.Enqueue([]byte("data_3"))
	if err != ErrBufferIsFull {
		t.Errorf("expected ErrBufferIsFull, got: %v", err)
	}

	// Проверяем, что буфер не пустой
	if b.IsEmpty() {
		t.Errorf("expected buffer to not be empty")
	}

	// Получаем элементы из буфера
	item, err := b.Dequeue()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if string(item) != "data" {
		t.Errorf("expected item to be 1, got: %v", item)
	}

	// Получаем пачку элементов из буфера
	items, err := b.DequeueBatch(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got := x.Map[[]byte, string](items, func(i []byte, _ int) string { return string(i) })
	expectedItems := []string{"data_2"}
	if !isEqualSlice(got, expectedItems) {
		t.Errorf("expected items to be %v, got: %v", expectedItems, got)
	}

	// Получаем элемент из пустого буфера
	_, err = b.Dequeue()
	if err != ErrBufferIsEmpty {
		t.Errorf("expected ErrBufferIsEmpty, got: %v", err)
	}
}

// Функция для сравнения двух слайсов
func isEqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
