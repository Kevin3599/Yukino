class Video(object):
    def __init__(self,path):
        self.path = path

    def play(self):
        from os import startfile
        startfile(self.path)

class Movie_MP4(Video):
    type = 'MP4'

movie = Movie_MP4(r'D:\彩蛋\【4K60帧】经典老歌：瑞克·埃斯利《Never-Gonna-Give-You-Up》1987-AI修复补帧版-_P1.-【4K60 帧】AI修复补帧版.mp4')
movie.play()
