package main

import (
	"fmt"
	"sync"

	geometry "github.com/raihaninkam/weeklyTaskGo/internals/Geometry"
	processnumber "github.com/raihaninkam/weeklyTaskGo/internals/ProcessNumber"
	usermanager "github.com/raihaninkam/weeklyTaskGo/internals/UserManager"
	webfetcer "github.com/raihaninkam/weeklyTaskGo/internals/WebFetcer"
)

func main() {

	// Test 1: input nil
	fmt.Println("Test 1: Input nil")
	result, err := processnumber.ProcessNumber(nil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	// Test 2: input empty list - wrap in separate function to catch panic
	fmt.Println("Test 2: Input empty list")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Panic ditangani:", r)
			}
		}()

		emptyList := []int{}
		result, err := processnumber.ProcessNumber(emptyList)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	}()
	fmt.Println()

	// Test 3: input valid
	fmt.Println("Test 3: Input valid")
	validInput := []int{1, 2, 3, 4, 5}
	result, err = processnumber.ProcessNumber(validInput)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	// Test 4: another valid input
	fmt.Println("Test 4: Another valid input")
	successInput := []int{10, 20, 30}
	result, err = processnumber.ProcessNumber(successInput)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()
	///////////////////////////////////////////////////////////////////////////////

	urls := []string{
		"https://raihan.com",
		"https://insan.com",
		"https://kamil.com",
		"https://inkam.com",
		"https://jaenab.com",
	}

	resultChan := make(chan string, len(urls))

	done := make(chan bool)

	var wg sync.WaitGroup

	go webfetcer.Receiver(resultChan, done)

	for _, url := range urls {
		wg.Add(1)
		go webfetcer.Webfetcer(url, resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)

	<-done
	fmt.Println()

	//////////////////////////////////////////////////////////////

	um := usermanager.NewUserManager()

	fmt.Println("Menambahkan user baru:")
	um.AddUser(1, "Raihan Insan")
	um.AddUser(2, "Kamil Jaenab")
	um.AddUser(3, "Inkam Project")
	fmt.Println()

	fmt.Println("Menambahkan user ID yang sudah ada")
	err = um.AddUser(1, "User Duplikat")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	fmt.Println("Mencari data user berdasarkan ID:")
	user, err := um.GetUser(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("User ditemukan: ID %d - %s\n", user.ID, user.Name)
	}

	user, err = um.GetUser(99)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("User ditemukan: ID %d - %s\n", user.ID, user.Name)
	}
	um.DisplayUsers()
	fmt.Println()

	/////////////////////////////////////////

	circle1 := geometry.Cirlce{Radius: 5.0}
	circle2 := geometry.Cirlce{Radius: 3.0}
	rectangle1 := geometry.Rectangle{Width: 4.0, Height: 6.0}
	rectangle2 := geometry.Rectangle{Width: 2.0, Height: 8.0}

	// slice of interface geometry
	shapes := []geometry.Geometry{circle1, circle2, rectangle1, rectangle2}

	fmt.Println("Area masing - masing bentuk")
	for i, shape := range shapes {
		switch s := shape.(type) {
		case geometry.Cirlce:
			fmt.Printf("Circle %d (Radius: %.1f): Area = %.2f\n", i+1, s.Radius, s.Area())
		case geometry.Rectangle:
			fmt.Printf("Rectangle %d (Width: %.1f, Height: %.1f): Area = %.2f\n", i+1, s.Width, s.Height, s.Area())
		}
	}
	fmt.Println()

	// hitung total area dengan calculator fungsi
	totalArea := geometry.CalculatorArea(shapes)
	fmt.Printf("Total area dari semua bentuk: %.2f\n", totalArea)

	// Contoh tambahan dengan slice yang berbeda
	fmt.Println("\n=== Contoh Lain ===")
	
	// Hanya circles
	circles := []geometry.Geometry{
		geometry.Cirlce{Radius: 7.0},
		geometry.Cirlce{Radius: 2.5},
	}
	
	fmt.Println("Area circles:")
	for i, circle := range circles {
		c := circle.(geometry.Cirlce)
		fmt.Printf("Circle %d: Area = %.2f\n", i+1, c.Area())
	}
	fmt.Printf("Total area circles: %.2f\n", geometry.CalculatorArea(circles))
	
	// Hanya rectangles
	rectangles := []geometry.Geometry{
		geometry.Rectangle{Width: 10.0, Height: 5.0},
		geometry.Rectangle{Width: 3.0, Height: 7.0},
	}
	
	fmt.Println("\nArea rectangles:")
	for i, rect := range rectangles {
		r := rect.(geometry.Rectangle)
		fmt.Printf("Rectangle %d: Area = %.2f\n", i+1, r.Area())
	}
	fmt.Printf("Total area rectangles: %.2f\n", geometry.CalculatorArea(rectangles))
}


