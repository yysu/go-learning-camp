# week9

#### 参考资料
https://aliliin.com/2021/12/17/TCP%20%E7%B2%98%E5%8C%85/

#### 问题一
总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

* 方式1: fix length

        发送方，每次发送固定长度的数据，并且不超过缓冲区，接受方每次按固定长度区接受数据
        应用场景： 运用于数据长度固定的粘包拆包场景, 比如指令发送

* 方式2: delimiter based

        发送方，在数据包添加特殊的分隔符，用来标记数据包边界
        应用场景：ftp、http

* 方式3: length field based

        发送方，在消息数据包头添加包长度信息
        应用场景：protobuf、goim

#### 问题二
实现一个从 socket connection 中解码出 goim 协议的解码器

* goim 协议结构

参考资料：https://github.com/Terry-Mao/goim/blob/e742c99ad7/api/protocol/protocol.go

4bytes PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
2bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
2bytes Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
4bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
4bytes Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码。