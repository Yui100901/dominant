package file_utils

import (
	"os"
	"path/filepath"
)

//
// @Author yfy2001
// @Date 2024/9/14 09 57
//

// TraverseDirFiles 遍历给定目录并返回文件路径列表
// recursive 参数表明是否递归遍历子目录
func TraverseDirFiles(dir string, recursive bool) ([]string, error) {
	var files []string

	if recursive {
		// 递归遍历目录
		err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	} else {
		// 非递归遍历目录
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				files = append(files, filepath.Join(dir, entry.Name()))
			}
		}
	}

	return files, nil
}
