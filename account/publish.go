package account

//func publish() {
//
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6389", // Connect to your custom RESP server
//		Password: "",
//		DB:       0,
//	})
//	defer rdb.Close()
//
//	_, err := rdb.Ping(ctx).Result()
//	if err != nil {
//		log.Fatalf("Could not connect to RESP server: %v", err)
//	}
//	fmt.Println("Connected to custom RESP server using go-redis!")
//
//	Subscribe(rdb)
//
//	//publishAccountUsers(rdb)
//	//publishAccountUser(rdb)
//	//publishPerson(rdb)
//
//	fmt.Println("Main application waiting for goroutines...")
//	time.Sleep(5 * time.Second) // Keep main alive to see subscriber output
//
//	err = rdb.Close()
//	if err != nil {
//		return
//	}
//
//	fmt.Println("Main application exiting.")
//}

//func Subscribe(rdb *redis.Client) {
//
//	// --- Subscriber Goroutine ---
//	go func() {
//		pubsub := rdb.Subscribe(ctx, "notifications", "account_channel")
//		defer pubsub.Close()
//		fmt.Println("[Subscriber] Listening for messages...")
//		ch := pubsub.Channel()
//		for msg := range ch {
//			fmt.Printf("[Subscriber] Channel: %s, Message: %s\n", msg.Channel, msg.Payload)
//		}
//		fmt.Println("[Subscriber] Goroutine finished.")
//	}()
//}

//func publishAccountUser(rdb *redis.Client) {
//
//	Users, err := model.CreateUserList()
//	if err != nil {
//		return
//	}
//
//	for _, user := range Users {
//
//		str, err2 := utils.ToStringJson(user)
//		if err2 != nil {
//			continue
//		}
//
//		ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Add a timeout for publish
//		_, err = rdb.Publish(ctx2, "account/user", str).Result()
//		cancel() // Release resources associated with the context
//
//		if err != nil {
//			log.Printf("[Publisher] Error publishing: %v", err)
//		}
//		time.Sleep(1 * time.Millisecond)
//	}
//}

//func publishAccountUsers(rdb *redis.Client) {
//
//	Users, err := model.CreateUserList()
//	if err != nil {
//		return
//	}
//
//	toJSON, err := utils.ToStringJson(Users)
//	if err != nil {
//		return
//	}
//
//	ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Add a timeout for publish
//	_, err = rdb.Publish(ctx2, "account/list", toJSON).Result()
//	cancel() // Release resources associated with the context
//
//	if err != nil {
//		log.Printf("[Publisher] Error publishing: %v", err)
//	}
//}
//
//func publishPerson(rdb *redis.Client) {
//
//	for i := 0; i < 2; i++ {
//		person := model.CreateFakePerson()
//		person.ID = i + 1
//		str, err := person.GetString()
//		if err != nil {
//			return
//		}
//
//		ctx2, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Add a timeout for publish
//		_, err = rdb.Publish(ctx2, "objects/person", str).Result()
//		cancel() // Release resources associated with the context
//
//		if err != nil {
//			log.Printf("[Publisher] Error publishing: %v", err)
//		}
//		time.Sleep(1 * time.Millisecond)
//	}
//}
