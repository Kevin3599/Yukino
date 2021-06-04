class video(object):
 def __init__(self,path):
        self.path = path
       
        def play(self):
           from os import startfile
           startfile(self.path)  

class Movie_mp4(video):
    type='MP4'

movie = Movie_mp4(r'')
movie.paly()