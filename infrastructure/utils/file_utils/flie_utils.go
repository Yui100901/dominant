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
func TraverseDirFiles(dir string, recursive bool) ([]string, []string, error) {
	var files []string
	var dirs []string

	if recursive {
		// 递归遍历目录
		err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				dirs = append(dirs, path)
			} else {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, nil, err
		}
	} else {
		// 非递归遍历目录
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, nil, err
		}
		for _, entry := range entries {
			fullPath := filepath.Join(dir, entry.Name())
			if entry.IsDir() {
				dirs = append(dirs, fullPath)
			} else {
				files = append(files, fullPath)
			}
		}
	}

	return dirs, files, nil
}
