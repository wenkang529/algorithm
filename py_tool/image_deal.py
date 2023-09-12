import os
import random
import re
import shutil
import subprocess
from datetime import datetime

FILE_TYPES = ['mov', 'mp4', 'jpeg', 'jpg', 'png', 'heic']


def get_files(path):
    file_list = []
    for folder_name, sub_folders, filenames in os.walk(path):
        for filename in filenames:
            file_list.append(os.path.abspath(os.path.join(folder_name, filename)))
    return file_list


def get_image_creation_date(file_path):
    # exiftool -createdate ./new.jpg
    cmd = ["exiftool", "-createdate", file_path]
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    date_time_str = result.stdout.decode().strip()
    if "0000:00:00 00:00:00" in date_time_str:
        date_time_str = ""
    if "Create Date" in date_time_str:
        date_time_str = date_time_str[-19:]
    return date_time_str


def update_image_creation_date(file_path, date):
    # exiftool -AllDates='2023:06:11 12:23:46'  -overwrite_original ./new.jpg
    dateStr = date.strftime('%Y:%m:%d %H:%M:%S')
    cmd = ["exiftool", "-AllDates='" + dateStr + "'", "-overwrite_original", file_path]
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    date_str = result.stdout.decode().strip()
    return date_str


def update_time_by_name(filePath):
    ctime = get_image_creation_date(filePath)
    if ctime == "" or "20" not in ctime:
        # faceu_-413_20230611123046465.jpg
        match = re.search(r'^faceu.*(\d{17})\.jp.+', os.path.basename(filePath))
        if match:
            time_str = match.group(1)[:14]
            date_time = datetime.strptime(time_str, '%Y%m%d%H%M%S')
            update_image_creation_date(filePath, date_time)
            print(f"update image:{filePath} time:{date_time.strftime('%Y:%m:%d %H:%M:%S')}")
        else:
            print(f"not match, file:{filePath} ctime:{ctime}")


def rename_file_with_time(file_path):
    date_str = get_image_creation_date(file_path)
    print(f"file:{len(file_path)} date:{date_str}")
    if date_str != "":
        file_name = os.path.basename(file_path)
        date_time = datetime.strptime(date_str, '%Y:%m:%d %H:%M:%S')
        suffix = file_name.split(".")[-1]
        suffix = suffix.lower()
        if suffix == "jpg":
            suffix = "jpeg"
        new_file_name = date_time.strftime('%Y_%m_%d-%H_%M_%S') + "-" + str(random.randint(10, 99)) + "." + suffix
        dest_file = os.path.join(dest_dir, new_file_name)
        if os.path.exists(dest_file):
            print(f"file duplicate,{file_name}")
            return
        shutil.move(file_path, dest_file)
        print(f"file:[{file_path}] from:[{file_path}] dest:[{dest_file}]")


if __name__ == "__main__":
    file_dir = "/xxx"   # 需要处理的文件路径
    dest_dir = "/xxx"  # 重命名后，移动文件到指定目录
    files = get_files(file_dir)
    for file in files:
        # update_time_by_name(file)
        # rename_file_with_time(file)
        pass
