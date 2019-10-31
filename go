#!/usr/bin/env python3
from prompt_toolkit import prompt, HTML
from prompt_toolkit import print_formatted_text as print
import os

if __name__ == '__main__':
    path = os.path.dirname(os.path.realpath(__file__))
    dirs = [d for d in os.listdir(path) if os.path.isdir(os.path.join(path, d)) and not d.startswith('.')]
    count = 0
    question = {}
    for dir in dirs:
        print(HTML("<b fg='orange'>{}</b>".format(dir.upper())))
        filespath = os.path.join(path, dir)
        files = [f for f in os.listdir(filespath) if os.path.isfile(os.path.join(filespath, f))]
        for file in files:
            print(HTML("<b fg='red'>{}</b> - {}".format(file.rsplit(' | ',1)[0],file.rsplit(' | ',1)[1])))
            # print(HTML("<b fg='red'>{}</b> - {}".format(count,file)))
            question[file.rsplit(' | ',1)[0]] = os.path.join(filespath, file)
            count += 1
    
    #print(question)
    answer = prompt(HTML("<b fg='green'>Choose:</b> "))
    executable = question[answer]
    print('running... {}'.format(executable))
    os.system('"{}"'.format(executable))