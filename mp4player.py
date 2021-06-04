import numpy as np
import cv2

cap = cv2.VideoCapture('')

while (cap.isOpened()):
    ret, frame = cap.read()
    cv2.imshow('you were folled!', frame)
    if cv2.waitKey(40) & 0xFF == ord('q'):
        break

cap.release()
cv2.destroyAllWindows()
#本代码来源自opnecv官网