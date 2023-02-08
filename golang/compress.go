package tools


// DeCompress 对.tar.gz 文件进行解压缩
// tarFile: 压缩文件的位置
// dest: 解压的存储位置
// return err
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := path.Join(dest, hdr.Name)
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	_ = os.RemoveAll(tarFile)
	return nil
}