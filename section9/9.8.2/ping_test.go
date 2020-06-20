package main

import "testing"

func BenchmarkPing(b *testing.B) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	done := make(chan struct{})
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch1
			ch2 <- "ping"
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < b.N; i++ {
			<-ch1
			ch2 <- "pong"
		}
		done <- struct{}{}
	}()

	<-done
	<-done

	close(ch1)
	close(ch2)
	close(done)
}
