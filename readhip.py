import zipfile
import sys
import os

here = os.path.dirname(os.path.abspath(__file__))
print('Here:', here)
# import py7zr
# import pyunpack
import shutil
import tempfile

def main(file=None):
    path_to_libs = os.path.dirname(os.__file__)
    path_to_scripts = path_to_libs.replace('lib','Scripts')
    os.environ['PATH'] += os.pathsep+path_to_scripts
    
    path_to_file = here+os.sep+file
    
    spl = file.split('.')
    ext = spl.pop(-1)
    file_wo_ext = '.'.join(spl)
    print('file:', file_wo_ext, '||ext:', ext)

    sys.exit()

    if spl[1] not in ['zip', 'ZIP','Zip']:
        tempdir = tempfile.gettempdir()
        path_to_temp = os.sep.join(tempdir,spl)

        shutil.copy2()
        os.rename()
    # os.environ()
    arh = pyunpack.Archive( path_to_file )
    for a in dir(arh):
        print(a)
    arh.extractall('.')
    print('Archive:', arh)

if __name__ == '__main__':
    file = None
    if len(sys.argv) >1:
        file = sys.argv[1]
    main(file)