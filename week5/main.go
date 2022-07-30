package main

import (
	"fmt"
	"time"
)

/*
作业：参考 Hystrix 实现一个滑动窗口计数器。
答：把map简化为了切片
*/

type event struct {
	timestamp int64
}

type SlidingWindow struct {
	buckets    []*event
	cap        int
	timeWindow int64
}

func main() {
	// create a sliding window with capacity 5 and time window 10 seconds
	// which means that there can be 5 events maximum, and
	// it will accept new events after 10 seconds
	s := CreateSlidingWindow(5, 10)

	go func() {
		fmt.Println("current time is", time.Now().Unix())
		c := time.Tick(10 * time.Second)
		after := <-c
		fmt.Println("10 seconds time window reached", after.Unix())
	}()

	// create some dummy events that will send after certain seconds
	// only the last two events (12000 and 14000) will be accepted after the time window
	timeSeries := []int64{500, 1500, 3000, 5000, 5800, 7000, 8200, 10900, 12000, 14000}
	for i := 1; i < len(timeSeries); i++ {
		time.Sleep((time.Duration(timeSeries[i]-timeSeries[i-1]) * time.Millisecond))
		fmt.Print(s.Pass())
	}
}

func CreateSlidingWindow(cap int, timeWindow int64) *SlidingWindow {
	return &SlidingWindow{
		timeWindow: timeWindow,
		buckets:    []*event{},
		cap:        cap,
	}
}

func (s *SlidingWindow) Pass() string {
	now := time.Now().Unix()
	// if the buckets are not full, the new event can pass
	if len(s.buckets) < s.cap {
		s.buckets = append(s.buckets, &event{timestamp: now})
		return fmt.Sprintln("event", now, "accepted")
	}
	earliestBucket := s.buckets[0]
	// if the time difference is bigger than the window, the new event can pass
	if now-earliestBucket.timestamp > s.timeWindow {
		s.buckets = append(s.buckets[1:], &event{timestamp: now})
		return fmt.Sprintln("event", now, "accepted")
	}
	return fmt.Sprintln("event", now, "declined")
}
