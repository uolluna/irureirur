	query := datastore.NewQuery("Task").EventualConsistency()
	it := client.Run(ctx, query)
	for {
		var t Task
		_, err := it.Next(&t)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(t)
	}  
