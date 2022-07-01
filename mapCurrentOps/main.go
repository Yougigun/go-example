package main

func main(){
	a := map[int]int{1:1}
	var c chan int
	for i := 0; i < 100; i++ {
		go func()int{
			var v int
			for{
				v=a[1]
				// a[1]=2
			}
			return v
		}()
	}
	<-c
}