package main

import (
	imageprocessing "fanoutfanin/image_processing"
	"image"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) []Job {
	var jobs []Job

	for _, p := range paths {
		job := Job{
			InputPath: p,
			OutPath:   strings.Replace(p, "images/", "images/output/", 1),
		}

		job.Image = imageprocessing.ReadImage(p)
		jobs = append(jobs, job)
	}

	return jobs
}

func resize(jobs *[]Job) <-chan Job {
	out := make(chan Job, len(*jobs))

	for _, job := range *jobs {
		go func(job Job) {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}(job)
	}

	return out
}

func collectJobs(input <-chan Job, imageCnt int) []Job {
	var resizedJobs []Job

	for i := 0; i < imageCnt; i++ {
		job := <-input
		resizedJobs = append(resizedJobs, job)
	}

	return resizedJobs
}

func saveImages(jobs *[]Job) {
	for _, job := range *jobs {
		imageprocessing.WriteImage(job.OutPath, job.Image)
	}
}

func main() {
	imagePaths := []string{"images/image1.jpg", "images/image2.jpg", "images/image3.jpg"}

	jobs := loadImage(imagePaths)

	// Fan out this function to multiple goroutines
	out := resize(&jobs)

	// Collect / Fan in
	resizedJobs := collectJobs(out, len(imagePaths))

	saveImages(&resizedJobs)
}
