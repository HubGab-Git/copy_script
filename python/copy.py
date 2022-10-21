import argparse
from distutils.dir_util import copy_tree

parser = argparse.ArgumentParser(
    description='Copy files form one folder to another',
    prog='copy_script'
)
def main():
    parser.add_argument(
        '-s', 
        '--source-dir', 
        default=".",
        help="Source directory from where files should be copied"
    )
    parser.add_argument(
        '-d', 
        '--destination-dir', 
        default=".", 
        help="Source directory from where files should be copied"
    )
    args = parser.parse_args()
    if args.source_dir == "." and args.destination_dir == "." :
        parser.print_help()
    else:
        copy_tree(args.source_dir,args.destination_dir)
        print(f"Copied files from {args.source_dir} to {args.destination_dir}")

if __name__ == "__main__":
    main()