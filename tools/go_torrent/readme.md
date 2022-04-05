# 来源
[bilibili](https://www.bilibili.com/video/BV13P4y1378F?spm_id_from=333.1007.top_right_bar_window_history.content.click)
[github](https://github.com/archeryue/go-torrent)
# Bit Torrent
1. p2p传输协议，找就近资源，人人为我，我为人人(peers 同等的)，去中心化
2. 如何找到peers  -> 找到tracker
3. 如何与peers协作完成下载

# Torrent File格式
- announce:string(tracker url)
- announce-list:[]string(备用tracker的列表)
- info:dict(文件具体信息)
  - name:string
  - length:int(文件总长度)
  - pieces:[][20]byte(每个文件片的SHA-1值,无论文件多长都转换为20byte的哈希值，用于校验)
  - piece length:int(每个文件片的长度)
  - files:[]dict(多文件时会有这个，下面时没有dict里的结构)
    - path:string(文件的子路径名)
    - length:int（长度）

# Bencode协议（序列化协议）
四种格式
## string
len:data
即：字符串长度:字符串
如：5:abcde
## int
'i'num'e'
即：'i'表示是数字，num表示实际的整数（可以使负数）,'e'表示结尾
如：i123456789e
## list
'l'
    ...
'e'
即：'l'表示list，...表示任意，可以是四种中的一种
## dict
'd'
    ...
'e'
即：'d'表示字典，...表示任意

# 结构
## bencode库
实现bencode的序列化与反序列化
## torrent
基于bencode库的torrent的解析库，得到tracker和种子信息
## tracker
获取peers信息
## download
文件片的下载与校验
## assembeer
将文件片拼成file