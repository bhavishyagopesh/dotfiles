Vim�UnDo� ~����4���l�L�e��m&O�R`dÎ��D��                                     YR�    _�                             ����                                                                                                                                                                                                                                                                                                                                                             YR=     �                -from inmemory_zip_importer import ZipImporter5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YR@     �                �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YRN     �                *f = open('common_module_archive.zip','rb')5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YRO     �                data = f.read()5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YRP     �                	f.close()5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YRQ     �                zipbytes = io.BytesIO(data)5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YRR     �                +sys.meta_path.append(ZipImporter(zipbytes))5�_�      	                      ����                                                                                                                                                                                                                                                                                                                                                             YRS     �                print sys.meta_path5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                             YRT     �                 5�_�   	              
           ����                                                                                                                                                                                                                                                                                                                                                             YRT     �                 5�_�   
                         ����                                                                                                                                                                                                                                                                                                                                                             YR�     �                import zipimportx5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YR�     �                5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             YR�     �                import sys,io,zipfile5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             YR�     �                �             5�_�                       @    ����                                                                                                                                                                                                                                                                                                                                                             YR�     �               Azipimporter("mylib.zip").write_index(preload=["mymod*","mypkg*"])5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             YR�    �               &zipimporter("mylib.zip").write_index()5��