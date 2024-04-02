/*
* Channels fazem comunicação entre as threads, são como filas de mensagens.
* Eles sabe quando a thread está pronta para receber ou enviar uma mensagem.
**/
package main

func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()
	<-forever
}
