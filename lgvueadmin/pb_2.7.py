import os
import io
import re
import shutil

source_path = os.path.abspath(r'./src/pb/proto')
target_path = os.path.abspath(r'./src/pb/temop')

if not os.path.exists(target_path):
    os.makedirs(target_path)

if os.path.exists(source_path):
    # root 所指的是当前正在遍历的这个文件夹的本身的地址
    # dirs 是一个 list，内容是该文件夹中所有的目录的名字(不包括子目录)
    # files 同样是 list, 内容是该文件夹中所有的文件(不包括子目录)
    for root, dirs, files in os.walk(source_path):
        for file in files:
            folder = os.path.basename(root)
            if folder == "proto":
              src_file = os.path.join(root, file)
              shutil.copy(src_file, target_path)
            elif folder != "protobuf":
                src_file = os.path.join(root, file)
                out_file = os.path.join(target_path, file)
                file_data = ""
                print(src_file)
                with io.open(src_file, "r", encoding='utf-8') as f:
                  for line in f:
                    if 'import' in line:
                      cite = re.findall(r"\"(.+?)\"",line)[0]
                      cpaths = cite.split("/")
                      if len(cpaths) == 2:
                        line = re.sub(
                          cite, cpaths[1], line)
                        print(line)
                    file_data += line
                with io.open(out_file, "w", encoding='utf-8') as f:
                  f.write(file_data)
cmdstr = 'protoc --proto_path=src/pb/temop --js_out=import_style=commonjs,binary:./src/pb/js ./src/pb/temop/*.proto'
print(cmdstr)
os.system(cmdstr)
print('copy files finished!')
