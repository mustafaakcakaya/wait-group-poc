package main

import (
	"fmt"
	"sync"
)

// Customer havuzda depolanacak müşteri veri yapısıdır
type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	var wg sync.WaitGroup

	// sync.Pool nesnesi oluşturulur
	pool := &sync.Pool{
		New: func() interface{} {
			return &Customer{}
		},
	}

	// Havuzda müşteri verilerini depolamak için bir WaitGroup kullanılır
	wg.Add(1)
	go func() {
		// Havuza bir Customer nesnesi eklenir
		customer := pool.Get().(*Customer)
		customer.ID = 1
		customer.FirstName = "John"
		customer.LastName = "Doe"
		customer.Address = "123 Main St"

		// Havuzdaki müşteri bilgileri kullanılır
		fmt.Println("Havuzdaki Müşteri:", customer)

		// Müşteri bilgileri havuza geri döndürülür
		pool.Put(customer)

		wg.Done()
	}()

	// Havuzdaki müşteri verilerini kullanmak için bir WaitGroup daha kullanılır
	wg.Add(1)
	go func() {
		// Havuzdan bir müşteri bilgisi alınır
		customerFromPool := pool.Get().(*Customer)
		fmt.Println("Havuzdan Alınan Müşteri:", customerFromPool)

		// Havuzdaki müşteri bilgileri kullanılır
		// ...

		// Müşteri bilgileri havuza geri döndürülür
		pool.Put(customerFromPool)

		wg.Done()
	}()

	// Tüm go routin'lerin tamamlanmasını bekler
	wg.Wait()
}
