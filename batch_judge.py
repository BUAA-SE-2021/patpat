from os import walk
from zipfile import ZipFile
import subprocess
proj_num=input()
base_addr='./'+proj_num
_, _, zipfile_names = next(walk(base_addr))
print(zipfile_names)
print()
for zipfile_name in zipfile_names:
    with ZipFile(base_addr+'/'+zipfile_name, 'r') as zip_obj:
         zip_obj.extractall(base_addr+'/'+zipfile_name[:-4])
for zipfile_name in zipfile_names:
    subprocess.call(['./bin/patpat-windows-amd64', 'ta', '-judge', zipfile_name[:-4]])
    print()