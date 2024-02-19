# reimagined-octo-tribble
# 经过几次失败,3.py可以传输一些简单的数据.
# 2021/7/25更新，除了个别文件外都可以传输一些简单的数据.
# 2021/7/29更新,function.go 是一个伟大的进步！！！
# 2021/7/31更新，在7月最后一天，我成功了！！！successul.go是最终成果，之前的文件其实只是出了一点小错，今天终于知道错在哪里了！！！
# header:=make([]byte,22)
# len,_:=conn.Read(header)
# //此时len显示为3，但是打印header为{0x05,0x01,0x00,0,0......},0会占位，导致host和port获得时会将后面的0代进去，导致程序编译没问题，但是一运行就出错。通过io.ReadFull()函数进行读取header可以避免一系列的错误！！！
# https://gist.github.com/felix021/7f9d05fa1fd9f8f62cbce9edbdb19253和https://github.com/WilliamColton/reimagined-octo-tribble/blob/main/successful.go可以用来参考。
# 时隔数年，再度更新（2024 2/20 早2：42（没错是凌晨） 在昨日发现一个很好用的框架Goframe后，彻夜研究，终于实现了第一个带加解密的半成功品，实属不易，接下来不会在此更新了，而是在第一个成功品出来之后重新开一个仓库正式更新） ProxyServer和ProxyLocal文件夹下面的就是我的半成功品，chatgpt4.0在过程中出了很大的力，从中受益颇多，当然也和我这些年的进步密不可分
