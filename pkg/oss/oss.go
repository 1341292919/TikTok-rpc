package oss

import (
	"TikTok-rpc/config"
	"TikTok-rpc/pkg/errno"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func IsImage(data *multipart.FileHeader) error {
	// 打开文件
	file, err := data.Open()
	if err != nil {
		return errno.Errorf(errno.InterFileProcessErrorCode, "open file error")
	}
	defer func() {
		_ = file.Close()
	}()

	// 读取文件头
	buffer := make([]byte, 261)
	n, err := file.Read(buffer)
	if err != nil || n < 261 {
		return errno.Errorf(errno.InterFileProcessErrorCode, "read file error")
	}

	// 检查文件类型
	if filetype.IsImage(buffer) { //是图片返回true
		return nil
	}

	// 检查文件扩展名（可选）
	if strings.HasSuffix(strings.ToLower(data.Filename), ".webp") {
		return nil
	}

	return errno.Errorf(errno.ParamVerifyErrorCode, "file not image")
}

func IsVideo(data *multipart.FileHeader) error {
	// 打开文件
	file, err := data.Open()
	if err != nil {
		return errno.Errorf(errno.InterFileProcessErrorCode, "open file error")
	}
	defer func() {
		_ = file.Close()
	}()

	// 读取文件头
	buffer := make([]byte, 261)
	n, err := file.Read(buffer)
	if err != nil || n < 261 {
		return errno.Errorf(errno.InterFileProcessErrorCode, "read file error")
	}

	// 检查文件类型
	if filetype.IsVideo(buffer) {
		return nil // 是视频，返回 nil
	}

	// 如果文件不是视频，返回错误
	return errno.Errorf(errno.ParamVerifyErrorCode, "file not video")
}

func SaveFile(data *multipart.FileHeader, storePath, fileName string) (err error) {
	if _, err := os.Stat(storePath); os.IsNotExist(err) {
		// 路径不存在，创建路径
		err := os.MkdirAll(storePath, 0755) //0755 是一个八进制数，表示文件或目录的权限。它的二进制形式是 111 101 101，对应的权限
		if err != nil {
			return errno.Errorf(errno.InterFileProcessErrorCode, "mkdir error")
		}
	}

	//打开本地文件
	dist, err := os.OpenFile(filepath.Join(storePath, fileName), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return errno.Errorf(errno.InterFileProcessErrorCode, "open file error")
	}
	defer func(dist *os.File) {
		_ = dist.Close() //延迟关闭文件，防止资源泄漏
		//确保该语句在函数返回时执行
	}(dist) //// 立即调用匿名函数，并传入外部的 dist 作为参数

	//打开上传的文件
	src, err := data.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)
	// 复制文件内容
	_, err = io.Copy(dist, src)

	return nil
}

func Upload(localFile, filename, userid, origin string) (string, error) {
	key := fmt.Sprintf("%s/%s/%s", origin, userid, filename)

	putPolicy := storage.PutPolicy{
		Scope: config.Oss.Bucket,
	}

	mac := auth.New(config.Oss.AccessKey, config.Oss.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.Zone_z2
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}

	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		return "", errno.Errorf(errno.InterFileProcessErrorCode, "create file recorder failed")
	}

	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return "", errno.Errorf(errno.InterFileProcessErrorCode, "upload file  error%v", err.Error())
	}
	defer func() {
		err = os.Remove(localFile)
	}()
	if err != nil {
		return "", errno.Errorf(errno.InterFileProcessErrorCode, "remove file error")
	}
	return storage.MakePublicURL(config.Oss.Domain, ret.Key), nil
}

// ExtractFirstFrame 从视频文件中提取第一帧作为封面图片
func ExtractFirstFrame(videoPath, coverPath string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", videoPath, // 输入视频文件
		"-ss", "00:00:00", // 定位到视频开头
		"-vframes", "1", // 提取1帧
		"-q:v", "2", // 图像质量（2表示高质量）
		coverPath, // 输出封面路径
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errno.Errorf(errno.InterFileProcessErrorCode, "exec ffmpeg error")
	}

	return nil
}
