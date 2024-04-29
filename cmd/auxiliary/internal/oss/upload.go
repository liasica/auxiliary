// Copyright (C) auxiliary. 2024-present.
//
// Created at 2024-04-29, by liasica

package oss

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
)

// ProgressListener 定义进度条监听器
type ProgressListener struct {
	// bar   *uiprogress.Bar
	speed float64
	start time.Time
}

func NewProgressListener() *ProgressListener {
	uiprogress.Start()
	listener := &ProgressListener{}
	// listener.bar = uiprogress.AddBar(100).AppendCompleted().PrependFunc(func(b *uiprogress.Bar) string {
	// 	return strutil.PadLeft(fmt.Sprintf("%.2f Kb/s (%s)", listener.speed, strutil.PrettyTime(time.Since(b.TimeStarted))), 25, ' ')
	// })
	return listener
}

// ProgressChanged 定义进度变更事件处理函数
func (listener *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		// listener.bar.TimeStarted = time.Now()
		// _ = listener.bar.Set(0)
		listener.start = time.Now()
	case oss.TransferDataEvent:
		// kb := float64(event.ConsumedBytes) / 1024.0
		// past := time.Since(listener.bar.TimeStarted).Seconds()
		// listener.speed = kb / past
		// _ = listener.bar.Set(int(event.ConsumedBytes * 100 / event.TotalBytes))
	case oss.TransferCompletedEvent:
		kb := float64(event.ConsumedBytes) / 1024.0
		past := time.Since(listener.start).Seconds()
		fmt.Printf("\n上传完成, 文件大小: %d Bytes, 速度: %.2f Kb/s\n", event.TotalBytes, kb/past)
		// _ = listener.bar.Set(100)
	case oss.TransferFailedEvent:
		fmt.Printf("\n上传失败, 已上传: %d Bytes, 文件大小: %d.\n", event.ConsumedBytes, event.TotalBytes)
		os.Exit(1)
	default:
	}
}

func uploader() *cobra.Command {
	var (
		accessKeyId     string
		accessKeySecret string
		endpoint        string
	)

	cmd := &cobra.Command{
		Use:               "upload <file> <oss_path> [args]",
		Short:             "上传文件",
		Example:           "auxiliary oss upload file.txt oss://bucket/file.txt --accessKeyId=xxx --accessKeySecret=xxx --endpoint=https://oss-cn-beijing.aliyuncs.com",
		Args:              cobra.ExactArgs(2),
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		Run: func(_ *cobra.Command, args []string) {
			file := args[0]
			if _, err := os.Stat(file); err != nil {
				fmt.Printf("文件不存在: %s\n", file)
				os.Exit(1)
			}

			path := args[1]
			loc := regexp.MustCompile(`oss://.*?/`).FindStringIndex(path)
			if len(loc) != 2 {
				fmt.Println("oss_path格式错误, 正确格式例: oss://bucket/file.txt")
				os.Exit(1)
			}

			bucketName := path[6 : loc[1]-1]
			object := path[loc[1]:]
			fmt.Printf("Bucket: %s, Object: %s\n", bucketName, object)

			client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
			if err != nil {
				fmt.Printf("创建OSS客户端失败: %s\n", err)
				os.Exit(1)
			}

			var bucket *oss.Bucket
			bucket, err = client.Bucket(bucketName)
			if err != nil {
				fmt.Printf("获取Bucket失败: %s\n", err)
				os.Exit(1)
			}

			// 带进度条的上传。
			err = bucket.PutObjectFromFile(object, file, oss.Progress(NewProgressListener()))
			if err != nil {
				fmt.Printf("上传文件失败: %s\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVar(&accessKeyId, "accessKeyId", "", "Access Key ID")
	cmd.Flags().StringVar(&accessKeySecret, "accessKeySecret", "", "Access Key Secret")
	cmd.Flags().StringVar(&endpoint, "endpoint", "https://oss-cn-beijing.aliyuncs.com", "Aliyun Endpoint")

	_ = cmd.MarkFlagRequired("accessKeyId")
	_ = cmd.MarkFlagRequired("accessKeySecret")

	return cmd
}
