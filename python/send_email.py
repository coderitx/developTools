# -*- coding: utf-8 -*-

import datetime
import os
import smtplib
from email.mime.text import MIMEText
from email.mime.application import MIMEApplication
from loguru import logger

def send_email(title,from_rmail,to_email,text,password,filepath=""):
    """
    :param: title: str 邮件的标题
    :param: from_email: str 发送方
    :param: to_email: list 接收方
    :param: password: str 发送邮箱的认证密码
    :param: text: str 发送的文本内容
    :param: filepath: str 如果发送个文件，文件的地址
    """
    m = MIMEMultipart()
    if filepath:
        file_apart = MIMEApplication(open(filepath, "rb").read())
        m.add_header("Content-Disposition", "attachment", filename=filepath)
        m.attach(file_apart)
    text_apart = MIMEText(text)
    m.attach(text_apart)
    m['Subject'] = f'{title}-{datetime.datetime.now().strftime("%Y%m%d%H%M%S")}'
    try:
        server = smtplib.SMTP('smtp.163.com')
        server.login(fromaddr, password)
        server.sendmail(fromaddr, toaddrs, m.as_string())
        logger.success("send email success")
        server.quit()
    except smtplib.SMTPException as e:
        logger.error('error:', e)  # 打印错误
   
