package routes

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkGetAllData(b *testing.B) {
	url := "http://localhost:3000/materials?limit=1000" // ubah limit untuk pengecekan performa
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := go_wrk.Init().
			HTTPTimeout(5000).
			Request("GET", url).
			Build().
			Run()
		if err != nil {
			panic(err)
		}
	}
	b.StopTimer()
	fmt.Printf("Benchmark for %d requests per second with average time per request %dms, total time taken %s\n",
		b.N, b.T, time.Duration(b.T.Nanoseconds()))
}

func BenchmarkGetOneData(b *testing.B) {
	url := "http://localhost:3000/material/6431580d92b73bcd2e66eb71?limit=10" // ubah limit untuk pengecekan performa
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := go_wrk.Init().
			HTTPTimeout(5000).
			Request("GET", url).
			Build().
			Run()
		if err != nil {
			panic(err)
		}
	}
	b.StopTimer()
	fmt.Printf("Benchmark for  %d requests per second with average time per request %dms, total time taken %s\n",
		b.N, b.T, time.Duration(b.T.Nanoseconds()))
}

func BenchmarkUpdateMaterial(b *testing.B) {
	url := "http://localhost:3000/material/64317e27826ebe25725d710f?limit=10" // ubah :id dengan id material yang ingin diupdate
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		material := model.Material{
			CourseName:   faker.Lorem().Word(),
			MaterialName: faker.Lorem().Word(),
			Description:  faker.Lorem().Sentence(10),
			UpdatedAt:    time.Now(),
		}
		payload, _ := json.Marshal(material)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
	}
	b.StopTimer()
	fmt.Printf("Benchmark for UpdateMaterial %d requests per second with average time per request %dms, total time taken %s\n",
		b.N, b.T, time.Duration(b.T.Nanoseconds()))
}

func BenchmarkUpdateCourse(b *testing.B) {
	url := "http://localhost:3000/course/6436586cc62bf3e17177ddb5?limit=10" // ubah :id dengan id material yang ingin diupdate
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		course := model.Course{
			CourseName:  faker.Lorem().Word(),
			Description: faker.Lorem().Sentence(10),
			UpdatedAt:   time.Now(),
		}
		payload, _ := json.Marshal(material)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
	}
	b.StopTimer()
	fmt.Printf("Benchmark for UpdateCourse %d requests per second with average time per request %dms, total time taken %s\n",
		b.N, b.T, time.Duration(b.T.Nanoseconds()))
}

func BenchmarkDelete(b *testing.B) {
	url := "http://localhost:3000/material/6436586cc62bf3e17177ddb5?limit=10" // ubah :id dengan id material yang ingin diupd
	// run delete benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id := materialIDs[i]
		_, err := go_wrk.Init().
			HTTPTimeout(5000).
			Request("DELETE", fmt.Sprintf("http://localhost:3000/materials/%s", id)).
			Build().
			Run()
		if err != nil {
			panic(err)
		}
	}
	b.StopTimer()

	fmt.Printf("Benchmark for %d requests per second with average time per request %dms, total time taken %s\n",
		b.N, b.T.Nanoseconds()/int64(b.N)/1000000, time.Duration(b.T.Nanoseconds()))
}
