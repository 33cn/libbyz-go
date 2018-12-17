# libbyz-go

封装libbyz库，暴露出几个go接口以供调用。`bft`下是原libbyz库，`client`下是封装好的client端的接口，`replica`下是封装好的replica端的接口。

## libbyz编译安装

在Ubuntu 16.04， gcc-5.4.0的环境下测试通过。

具体配置过程如下。

```shell
sudo apt-get install -y automake autoconf
sudo apt-get install -y gcc g++
sudo apt-get install -y libgmp-dev
sudo apt-get install -y libtool
sudo apt-get install -y flex bison
sudo apt-get install -y make

cd bft/sfslite-1.2
autoreconf -i
sh -x setup.gnu -f -i -s
mkdir install
SFSHOME=./bft/sfslite-1.2
./configure --prefix=$SFSHOME/install
make CFLAGS="-Werror=strict-aliasing" CXXFLAGS="-fpermissive -DHAVE_GMP_CXX_OPS"
make install

cd bft
ln -s sfslite-1.2/install sfs
ln -s /usr/lib gmp

cd bft/libbyz
sed -i '418s/^.*$/  th_assert(sizeof(t.tv_sec) <= sizeof(long), "tv_sec is too big");/' Node.cc
sed -i '420s/^.*$/  long int_bits = sizeof(long)*8;/' Node.cc
make CPPFLAGS="-I../gmp -I../sfs/include/sfslite -g -Wall -DRECOVERY -fpermissive -DHAVE_GMP_CXX_OPS"

(go demo: 运行replica，需要配置好config里的IP地址)
go build test.go
./test
```

