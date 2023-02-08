package tools


// isExists 判断文件夹是否存在，不存在则创建，存在则删除重新创建
// dir
// return 
func isExists(dir string) bool {
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			logrus.Errorf("%s not found create failed %s", dir, err.Error())
			return false
		}
		return true
	} else {
		_ = os.RemoveAll(dir)
		_ = os.MkdirAll(dir, 0755)
		return true
	}
}


// Download 从指定url下载到指定位置，返回下载存储路径
// urlStr: 指定url
// dest: 存储的文件位置
// return: 解压结果，存储路径
func Download(urlStr, dest string) (bool, string) {
	rsp, err := http.Get(urlStr)
	if err != nil {
		logrus.Errorf("request %s error %s", urlStr, err.Error())
		return false, ""
	}
	urls := strings.Split(urlStr, "/")
	file := urls[len(urls)-1]
	exists := isExists(dest)
	if !exists {
		logrus.Errorf("create dir failed")
		return exists, ""
	}
	f, _ := os.Create(path.Join(dest, file))
	_, err = io.Copy(f, rsp.Body)
	if err != nil {
		logrus.Errorf("write request body error %s", err.Error())
		return false, ""
	}
	return true, path.Join(dest, file)
}
