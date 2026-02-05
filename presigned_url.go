package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/database"
)

func generatePresignedURL(s3Client *s3.Client, bucket, key string, expireTime time.Duration) (string, error) {
	s3PresignClient := s3.NewPresignClient(s3Client)
	obj, err := s3PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(expireTime))
	if err != nil {
		return "", err
	}
	return obj.URL, nil
}

func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
	// take dbVideo and return dbVideo with videUrl set to presigned URL and error
	if video.VideoURL == nil {
		return video, nil
	}
	uSplit := strings.Split(*video.VideoURL, ",")
	fmt.Println("raw video URL:", *video.VideoURL, "parts:", uSplit)
	if len(uSplit) < 2 {
		return video, fmt.Errorf("Presigned url not generated")
	}
	bucket, key := uSplit[0], uSplit[1]

	fmt.Println("bucket:", bucket, "key:", key)

	signedURL, err := generatePresignedURL(cfg.s3Client, bucket, key, time.Hour)
	if err != nil {
		return video, fmt.Errorf("failed to generate presigned URL: %v", err)
	}
	video.VideoURL = &signedURL
	return video, nil
}
