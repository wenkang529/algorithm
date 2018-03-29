# Tensorflow
## Tensorflow read data efficiently  
[link](http://ycszen.github.io/2016/08/17/TensorFlow%E9%AB%98%E6%95%88%E8%AF%BB%E5%8F%96%E6%95%B0%E6%8D%AE/)

There are three methods:  
1. feeding by python  
2. read from file
3. define all the data

The second proposal is recommended : TFRecords (use queue provided by TF)  
### Creat TFRecords file

* create TFrecord file : `use tf.train.Example` to define data format which we want to input 
* use `tf.python_io.TFRecordWtiter` to write
```
import os
import tensorflow as tf 
from PIL import Image
cwd = os.getcwd()
'''python
此处我加载的数据目录如下：
0 -- img1.jpg
     img2.jpg
     img3.jpg
     ...
1 -- img1.jpg
     img2.jpg
     ...
2 -- ...
...
'''
writer = tf.python_io.TFRecordWriter("train.tfrecords")
for index, name in enumerate(classes):
    class_path = cwd + name + "/"
    for img_name in os.listdir(class_path):
        img_path = class_path + img_name
            img = Image.open(img_path)
            img = img.resize((224, 224))
        img_raw = img.tobytes()              #将图片转化为原生bytes
        example = tf.train.Example(features=tf.train.Features(feature={
            "label": tf.train.Feature(int64_list=tf.train.Int64List(value=[index])),
            'img_raw': tf.train.Feature(bytes_list=tf.train.BytesList(value=[img_raw]))
        }))
        writer.write(example.SerializeToString())  #序列化为字符串
writer.close()
```

### Read TFRecords file(use queue)
```
def read_and_decode(filename):
	#根据文件名生成一个队列
    filename_queue = tf.train.string_input_producer([filename])
    reader = tf.TFRecordReader()
    _, serialized_example = reader.read(filename_queue)   #返回文件名和文件
    features = tf.parse_single_example(serialized_example,
                                       features={
                                           'label': tf.FixedLenFeature([], tf.int64),
                                           'img_raw' : tf.FixedLenFeature([], tf.string),
                                       })
    img = tf.decode_raw(features['img_raw'], tf.uint8)
    img = tf.reshape(img, [224, 224, 3])
    img = tf.cast(img, tf.float32) * (1. / 255) - 0.5
    label = tf.cast(features['label'], tf.int32)
    return img, label
```
### use in train
```
img, label = read_and_decode("train.tfrecords")
# 使用shuffle_batch可以随机打乱输入
img_batch, label_batch = tf.train.shuffle_batch([img, label],
                                                batch_size=30, capacity=2000,
                                                min_after_dequeue=1000)
init = tf.initialize_all_variables()
with tf.Session() as sess:
    sess.run(init)
    threads = tf.train.start_queue_runners(sess=sess)
    for i in range(3):
        val, l= sess.run([img_batch, label_batch])
        #我们也可以根据需要对val， l进行处理
        #l = to_categorical(l, 12) 
        print(val.shape, l)
```

第一，tensorflow里的graph能够记住状态（state），这使得TFRecordReader能够记住tfrecord的位置，并且始终能返回下一个。而这就要求我们在使用之前，必须初始化整个graph，这里我们使用了函数tf.initialize_all_variables()来进行初始化。

第二，tensorflow中的队列和普通的队列差不多，不过它里面的operation和tensor都是符号型的（symbolic），在调用sess.run()时才执行。

第三， TFRecordReader会一直弹出队列中文件的名字，直到队列为空。

### Conclution
生成tfrecord文件
定义record reader解析tfrecord文件
构造一个批生成器（batcher）
构建其他的操作
初始化所有的操作
启动QueueRunner



## 张量的打印输出

```
print(sess.run(c,feed_dict={a:data}))
print(c.eval(feed_dict={a:data})) 
```

