package account

//func (manager *ClientManager) InitRespService() {
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
//	// --- Subscriber Goroutine ---
//	go func() {
//
//		pubsub := rdb.Subscribe(ctx, "account/user", "account/list", "account/search")
//		defer pubsub.Close()
//
//		Subscriber(pubsub)
//	}()
//
//	time.Sleep(100 * time.Minute) // Give subscriber time to connect
//
//	// --- Publisher Operations ---
//	//fmt.Println("\n--- Publisher Operations ---")
//	//channelsToPublish := []string{"my_custom_channel", "notifications", "account_channel"}
//	//messagesToPublish := []string{"Hello from publisher 1!", "System alert: Update available!", "This is a second message!"}
//	//
//	//for i, channel := range channelsToPublish {
//	//	message := messagesToPublish[i]
//	//	fmt.Printf("[Publisher] Publishing '%s' to channel '%s'\n", message, channel)
//	//	_, err := rdb.Publish(ctx, channel, message).Result()
//	//	if err != nil {
//	//		log.Printf("[Publisher] Error publishing: %v", err)
//	//	}
//	//	time.Sleep(2.json * time.Second)
//	//}
//
//}
