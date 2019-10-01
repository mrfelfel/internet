# SCP Transport Protocol
SCP(SabzCityProtocol) transport protocol is SabzCity experimental open-source transport protocol! It is very simple protocol influenced by [QUIC](https://en.wikipedia.org/wiki/QUIC) but base on IPv6 and have some difference in internal service call, security, pipelining, multiplexing, ... . It will introduce huge advantage and improve app network performance! It is based on code-generation but also can use by some function in this package to test in easy way.
In [PersiaOS](https://github.com/SabzCity/PersiaOS) IPv6 standard every service have one (or more) unique IPv6, So SCP don't need any layer between SCP & IPv6 unlike QUIC, ...!

Some idea get from these protocols:
- https://github.com/alecthomas/go_serialization_benchmarks
- https://tools.ietf.org/html/rfc4506
- https://github.com/google/flatbuffers
- https://capnproto.org/
- https://www.semanticscholar.org/paper/Blockchain-models-for-universal-connectivity-Navarro-Castro/788b7a634b369d98e72ed37c5fdf71f7fd62ef0b
https://pdfs.semanticscholar.org/788b/7a634b369d98e72ed37c5fdf71f7fd62ef0b.pdf?_ga=2.260489549.1562006812.1569054619-1995410782.1569054619
