// Мы получаем сообщения из источника
// У каждого сообщения есть тип (префикс из пяти символов)
// Нужно запоминать самые свежие сообщения каждого типа

package main

type Type struct {
	str string
}

func main() {
	consumeMessages()
}

func consumeMessages() {
	for {
		msg := receiveMessage()                // получаем сообщение из источника
		msgType := getMessageType(msg)         // определяем тип сообщения
		storeMessageType(string(msg), msgType) // сохраняем сообщение
	}
}

func storeMessageType(message string, msgType Type) {
	// ...
}

func getMessageType(msg []rune) Type {
	return Type{
		str: string(msg),
	}
}

func receiveMessage() []rune {
	return []rune("rand string")
}
