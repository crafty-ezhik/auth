package auth

// HTTPContext - интерфейс, который определяет необходимые методы для работы middleware
type HTTPContext interface {
	// GetHeader - получение хедера Authorization
	GetHeader(name string) string

	// Unauthorized - отправка ответа клиенту, что он не авторизирован
	Unauthorized()

	// SetUser - передача значения userID, username, ... в контекст для дальнейшего взаимодействия
	SetUser(user any)

	// Next - передача управления следующему хендлеру
	Next()

	// Status - установка статуса ответа
	Status(code int)

	// Send - отправка сообщения в ответ
	Send(data []byte)
}
