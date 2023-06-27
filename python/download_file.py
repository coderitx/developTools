
import os
import time
import tempfile
import requests
from zipfile import ZipFile
from loguru import logger


def clear_old_file(filename: str) -> bool:
    """
    清除旧的文件或文件夹
    :param filename:
    :return:
    """
    if not os.path.exists(filename):
        return True
    try:
        if os.path.isdir(filename):
            shutil.rmtree(filename)
        else:
            os.remove(filename)
        logger.warning(f"清除旧的文件 {filename} 成功")
        return True
    except Exception as e:
        logger.error(f"clear {filename} failed error: {e}")
        return False

def _write_response_to_file(response: requests.Response, filename: str, path: str, start: float):
    """
    :param response: 
    :param filename: 
    :param path: 
    :param start: 
    :return: 
    """
    size = 0  # 初始化已下载大小
    chunk_size = 1024  # 每次下载的数据大小
    content_size = int(response.headers['content-length'])  # 下载文件总大小
    logger.info(f"start download {filename}")
    try:
        if response.status_code == 200:  # 判断是否响应成功
            print('Start download,[File size]:{size:.2f} MB'.format(
                size=content_size / chunk_size / 1024))  # 开始下载，显示下载文件大小

            filepath = os.path.join(path,filename)  # 设置图片name，注：必须加上扩展名
            with open(filepath, 'wb') as file:  # 显示进度条
                for data in response.iter_content(chunk_size=chunk_size):
                    file.write(data)
                    size += len(data)
                    print('\r' + f'[{filepath} 下载进度]: %s%.2f%%' % (
                        '>' * int(size * 50 / content_size), float(size / content_size * 100)), end=' ')
        end = time.time()  # 下载结束时间
        print('Download completed!,times: %.2f秒' % (end - start))  # 输出下载用时时间
        return True
    except:
        logger.error('Download Error!')
        return False


def download_file(url, path) -> bool:
    """
    下载zip文件并解压
    :param url: 
    :param path: 
    :return: 
    """
    start = time.time()  # 下载开始时间
    filename = _get_file_name(url)
    # 删除旧的包
    if os.path.exists(os.path.join(path)):
        clear_old_file(path)
        logger.info(f"删除旧的文件 {path}")
    if not os.path.exists(path):  # 看是否有该文件夹，没有则创建文件夹
        os.makedirs(path)
    response = requests.get(url, stream=True, headers=_header(),timeout=(20,30))
    return _write_response_to_file(response, filename, path, start)

def download_zipfile(url:str,path:str)->bool:
    clear_old_file(path)
    try:
        response = requests.get(url, params=params)
        _tmp_file = tempfile.TemporaryFile()
        _tmp_file.write(response.content)
        zf = ZipFile(_tmp_file, mode='r')
        for names in zf.namelist():
            zf.extract(names, path)
        zf.close()
    except Exception as e:
        logger.error(f"download failed {repr(e)}")
  
