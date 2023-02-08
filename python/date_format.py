import time
import datetime


def time_stamp2format_time(time_stamp, hms: bool = False):
    """
    :param time_stamp:  时间戳
    :param hms:         日期格式
    :return:            标准日期
    """
    format_time = "%Y-%m-%d %H:%M:%S" if hms else "%Y-%m-%d"
    time_stamp = int(time_stamp) if len(str(time_stamp)) == 10 else int(time_stamp) / 1000
    format_time = time.strftime(format_time, time.localtime(time_stamp))

    return format_time


def format_time2time_stamp(format_time, milli: bool = True):
    """
    :param format_time: 标准日期
    :param milli:       秒/毫秒标志
    :return:            时间戳
    """
    format_flag = "%Y-%m-%d" if len(format_time) == 10 else "%Y-%m-%d %H:%M:%S"
    time_array = time.strptime(format_time, format_flag)
    time_stamp = int(time.mktime(time_array))

    return time_stamp * 1000 if milli else time_stamp


def utc2format_time(utc_time, utc_format="%Y-%m-%dT%H:%M:%S.%fZ", hms: bool = False):
    """
    :param utc_time:    格林威治标准时间（GMT）、协调世界时（UTC）
    :param utc_format:  时间格式
    :param hms:         日期格式
    :return:
    """

    format_flag = "%Y-%m-%d %H:%M:%S" if hms else "%Y-%m-%d"

    utc_time = datetime.datetime.strptime(utc_time, utc_format)
    format_time = utc_time + datetime.timedelta(hours=8)

    return format_time.strftime(format_flag)
