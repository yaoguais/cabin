package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"gopkg.in/redis.v5"
)

var (
	rootCmd = &cobra.Command{
		Use: "benchmark [command]",
	}
	genCmd = &cobra.Command{
		Use:     "gen",
		Example: "benchmark gen -f geo.txt -n 1000000",
		RunE:    commandGen,
	}
	redisCmd = &cobra.Command{
		Use: "redis",
		Example: `benchmark redis init -f geo.txt
benchmark redis -f geo_search.txt`,
		RunE: commandRedis,
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())

	genCmd.Flags().StringP("file", "f", "geo.txt", "output file for saving location data")
	genCmd.Flags().IntP("number", "n", 1000000, "number of record for generate")
	rootCmd.AddCommand(genCmd)

	redisCmd.Flags().StringP("file", "f", "geo.txt", "file saved test data")
	redisCmd.Flags().IntP("port", "p", 6379, "server port")
	redisCmd.Flags().StringP("host", "H", "127.0.0.1", "server hostname")
	redisCmd.Flags().IntP("concurrency", "c", 50, "concurrency client number")
	rootCmd.AddCommand(redisCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("OK")
	}
}

func commandGen(cmd *cobra.Command, args []string) error {
	file := cmd.Flags().Lookup("file").Value.String()
	number := cast.ToInt(cmd.Flags().Lookup("number").Value.String())

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < number; i++ {
		// -180 to 180
		lng := float64(rand.Intn(360)-180) + float64(rand.Intn(1000000))/1000000
		// -85.05112878 to 85.05112878
		lat := float64(rand.Intn(168)-84) + float64(rand.Intn(1000000))/1000000
		f.WriteString(strconv.FormatFloat(lng, 'f', 6, 64))
		f.WriteString(" ")
		f.WriteString(strconv.FormatFloat(lat, 'f', 6, 64))
		f.WriteString("\n")
	}

	return nil
}

func commandRedis(cmd *cobra.Command, args []string) error {
	if len(args) > 1 || (len(args) == 1 && args[0] != "init") {
		return errors.New("invalid parameters")
	}

	if len(args) == 1 {
		return commandRedisInit(cmd, args)
	}

	return commandRedisBenchmark(cmd, args)
}

func newRedisClient(opt *redis.Options) *redis.Client {
	c := redis.NewClient(opt)
	var wg sync.WaitGroup
	for i := 0; i < 2*opt.PoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				c.Ping()
			}
		}()
	}
	wg.Wait()

	fmt.Print("pool stat\n")
	fmt.Printf("totalConn: %d\n", c.PoolStats().TotalConns)
	fmt.Printf("freeConn: %d\n", c.PoolStats().FreeConns)
	fmt.Print("\n")

	return c
}

func readRedisGeoLocationFromFile(file string) ([]*redis.GeoLocation, error) {
	var geoLocations []*redis.GeoLocation

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	idx := 0
	r := bufio.NewReader(f)
	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if isPrefix {
			return nil, errors.New("line is too long")
		}

		columns := strings.Split(string(line), " ")
		if len(columns) == 2 {
			lng, err := strconv.ParseFloat(columns[0], 64)
			if err != nil {
				return nil, err
			}
			lat, err := strconv.ParseFloat(columns[1], 64)
			if err != nil {
				return nil, err
			}
			idx++
			geoLocation := &redis.GeoLocation{
				Name:      cast.ToString(idx),
				Longitude: lng,
				Latitude:  lat,
			}
			geoLocations = append(geoLocations, geoLocation)
		}
	}
	return geoLocations, nil
}

func commandRedisInit(cmd *cobra.Command, args []string) error {
	concurrency := cast.ToInt(cmd.Flags().Lookup("concurrency").Value.String())
	host := cmd.Flags().Lookup("host").Value.String()
	port := cast.ToInt(cmd.Flags().Lookup("port").Value.String())
	file := cmd.Flags().Lookup("file").Value.String()

	fmt.Print("GeoAdd\n")

	geoLocations, err := readRedisGeoLocationFromFile(file)
	if err != nil {
		return err
	}

	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		PoolSize: concurrency,
	}

	c := newRedisClient(opt)
	defer c.Close()

	key := "geo:benchmark"
	c.Del(key)

	info, err := c.Info().Result()
	if err != nil {
		return err
	}

	start := time.Now()
	for _, v := range geoLocations {
		if err := c.GeoAdd(key, v).Err(); err != nil {
			return err
		}
	}
	finish := time.Now()

	fmt.Print("redis server(before)\n")
	fmt.Printf("info: %s\n", info)
	fmt.Print("\n")

	info, err = c.Info().Result()
	if err != nil {
		return err
	}
	fmt.Print("redis server(after)\n")
	fmt.Printf("info: %s\n", info)
	fmt.Print("\n")

	fmt.Print("benchmark GeoAdd\n")
	fmt.Printf("concurrency: %d\n", concurrency)
	fmt.Printf("request: %d\n", len(geoLocations))
	fmt.Printf("start: %s\n", start.Format("2006-01-02 15:04:05.999999999"))
	fmt.Printf("finish: %s\n", start.Format("2006-01-02 15:04:05.999999999"))
	fmt.Printf("spend: %d ms\n", finish.Sub(start)/time.Millisecond)
	fmt.Printf("rps: %d r/s\n", int(float64(len(geoLocations))/(float64(finish.Sub(start))/float64(time.Second))))
	fmt.Print("\n")

	return nil
}

func commandRedisBenchmark(cmd *cobra.Command, args []string) error {
	concurrency := cast.ToInt(cmd.Flags().Lookup("concurrency").Value.String())
	host := cmd.Flags().Lookup("host").Value.String()
	port := cast.ToInt(cmd.Flags().Lookup("port").Value.String())
	file := cmd.Flags().Lookup("file").Value.String()

	fmt.Print("GeoRadius\n")

	geoLocations, err := readRedisGeoLocationFromFile(file)
	if err != nil {
		return err
	}

	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		PoolSize: concurrency,
	}

	c := newRedisClient(opt)
	defer c.Close()

	info, err := c.Info().Result()
	if err != nil {
		return err
	}

	hit := 0
	key := "geo:benchmark"
	geoQuery := &redis.GeoRadiusQuery{
		Radius:   5.0,
		Unit:     "km",
		WithDist: true,
		Sort:     "ASC",
	}

	start := time.Now()
	for _, v := range geoLocations {
		gs, err := c.GeoRadius(key, v.Longitude, v.Latitude, geoQuery).Result()
		if err != nil {
			return err
		}
		hit += len(gs)
	}
	finish := time.Now()

	fmt.Print("redis server(before)\n")
	fmt.Printf("info: %s\n", info)
	fmt.Print("\n")

	info, err = c.Info().Result()
	if err != nil {
		return err
	}
	fmt.Print("redis server(after)\n")
	fmt.Printf("info: %s\n", info)
	fmt.Print("\n")

	fmt.Print("benchmark\n")
	fmt.Printf("hit: %d\n", hit)
	fmt.Printf("concurrency: %d\n", concurrency)
	fmt.Printf("request: %d\n", len(geoLocations))
	fmt.Printf("start: %s\n", start.Format("2006-01-02 15:04:05.999999999"))
	fmt.Printf("finish: %s\n", start.Format("2006-01-02 15:04:05.999999999"))
	fmt.Printf("spend: %d ms\n", finish.Sub(start)/time.Millisecond)
	fmt.Printf("rps: %d r/s\n", int(float64(len(geoLocations))/(float64(finish.Sub(start))/float64(time.Second))))
	fmt.Print("\n")

	return nil
}
