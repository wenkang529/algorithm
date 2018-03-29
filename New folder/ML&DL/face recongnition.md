# face-recongnition


## 关键步骤
1.人脸检测&关键点标记  
2.人脸对齐  
3.embedding
4.recongnition

[总体介绍的文章](https://zhuanlan.zhihu.com/p/25335957)  
[Introduction about some face recongnition company](https://www.zhihu.com/question/37060782)

## dlib 人脸检测 关键点标记
[dlib](http://blog.csdn.net/sunmc1204953974/article/details/49976045)


## opencv 仿射变换
[python下的实现](http://blog.csdn.net/ying86615791/article/details/71217273)

关键计算：旋转中心（人眼） 旋转角度（根据两眼进行计算）
```
RotateMatrix = cv2.getRotationMatrix2D(eye_center, angle, scale=1) # 计算仿射矩阵
RotImg = cv2.warpAffine(img, RotateMatrix, (img.shape[0], img.shape[1])) # 进行放射变换，即旋转
```

## facenet
[山人七](https://zhuanlan.zhihu.com/p/24837264)  
[facenet-github](https://github.com/davidsandberg/facenet)

## mtcnn
github 与facenet在一起
[csdn](http://blog.csdn.net/mr_evanchen/article/details/77650883)  
[文章](https://kpzhang93.github.io/MTCNN_face_detection_alignment/)

## 直接一个步骤lenet5 识别

[csdn](http://blog.csdn.net/jrrrj/article/details/76550576)

## PCA & SVM 
[link](http://www.voidcn.com/article/p-dpjuwhqa-da.html)
## one sample: creat one neural network let it know who i am 
[code link](http://tumumu.cn/2017/05/02/deep-learning-face/)

