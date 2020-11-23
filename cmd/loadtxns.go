package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var rsource = rand.NewSource(time.Now().UnixNano())

const (
	host     = "localhost"
	port     = 8001
	user     = "unicorn_user"
	password = "magical_password"
	dbname   = "rainbow_database"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	db.SetMaxOpenConns(50)
	fmt.Printf("Max Connections %d\n", db.Stats().MaxOpenConnections)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	type workItem string
	var workerCount int = 50
	fmt.Println("Successfully connected!")
	start := time.Now()

	//transaction_id |  transaction_time  | transaction_date | amount1 | amount2 | amount3 | amount4 | loc_id
	insertText := `INSERT INTO public.transactions VALUES ($1, now() AT TIME ZONE 'Australia/Sydney', (now() AT TIME ZONE 'Australia/Sydney')::date,$2,$3,$4,$5,$6) ON CONFLICT ON CONSTRAINT transactions_pkey DO NOTHING;`

	workChan := make(chan workItem, workerCount)
	var wg sync.WaitGroup

	for w := 0; w < workerCount; w++ {
		go func() {
			var wrkerSource = rand.New(rsource)
			for true {
				id, more := <-workChan
				if more {
					func() {
						_, err = db.Exec(insertText, id, rndAmount(wrkerSource), rndAmount(wrkerSource), rndAmount(wrkerSource), rndAmount(wrkerSource), wrkerSource.Intn(20))
						defer wg.Done()
						if err != nil {
							fmt.Printf("Couldn't insert :%s\nwith id :%s\n", insertText, id)
							panic(err)
						}
					}()
				} else {
					fmt.Printf("Terminating worker...")
					break
				}
			}
		}()
	}

	numRecords := 500000
	for i := 0; i < numRecords; i++ {
		wg.Add(1)
		workChan <- workItem(fmt.Sprintf("'id%d'", i))
	}
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Insert of %d records took %s", numRecords, elapsed)
	close(workChan)
	selectQuery := `SELECT
	    transaction_date AS day,
	    loc_id,
	    SUM(amount1) as amount1_sum,
	    SUM(amount2) as amount2_sum,
	    SUM(amount3) as amount3_sum,
	    SUM(amount4) as amount4_sum,
	    COUNT(loc_id) as item_count
	FROM public.transactions
	where day= TO_DATE('2020-11-20','YYYY-MM-DD')
	GROUP BY loc_id,transaction_date;`

	_, err = db.Exec(insertText)
}

func rndAmount(r *rand.Rand) string {
	return fmt.Sprintf("%d.%d", r.Intn(100), r.Intn(99))
}
