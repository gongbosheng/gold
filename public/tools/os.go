package tools

import (
	"archive/tar"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Mkdir(name string) {

	// 检查文件夹是否已经存在
	_, err := os.Stat(name)

	// 如果文件夹不存在，则创建
	if os.IsNotExist(err) {
		var mod fs.FileMode = 0755
		err = os.Mkdir(name, os.ModeDir|mod)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created:", name)
	} else if err == nil {
		return
	} else {
		// 其他错误
		fmt.Println("Error checking directory existence:", err)
	}
}

func GetTarGz(source, target string) error {
	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar.gz", filename))

	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func ConvertBytesToReadableFormat(size int64) string {
	const (
		KB = 1024 // 定义1KB为1024字节
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case size >= GB: // 文件大小在GB范围内
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB: // 文件大小在MB范围内
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB: // 文件大小在KB范围内
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d bytes", size) // 小于1KB，直接以字节为单位
	}
}
