
import subprocess
from loguru import logger

def run_shell(cmd):
  """
  run_shell: 运行shell命令并实时输出，返回成功与否
  :param cmd: 允许的shell命令
  :return: 成功或者失败
  """
  p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
  while p.poll() is None:
    while True:
      if p.wait() != 0:
        logger.error(f"run shell [{cmd}] failed")
        return False
      else:
        line = p.stdout.readline()
        if not line:
          break
        print(line.rstrip().decode("utf-8"))
  logger.success(f"run shell [{cmd}] success")
  return True
