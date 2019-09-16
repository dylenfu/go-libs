package bufio

import (
	"io"
	"errors"
)

const (
	defaultBufferSize = 4096
	minReaderBufferSize = 10
	// 当reader.bf需要从io.read中读取数据时，如果读取数据为空，最大重试次数
	maxConsecutiveEmptyReadTime = 100
)

var (
	ErrBufferFull = errors.New("bufio: buffer full")
	ErrNegativeCount = errors.New("bufio: negative count")
)

// 通过实现bufio深入了解bufio的底层实现原理
// 为了降低耦合度，我们在操作bf时，将bf的操作copy放到fill里，
// 相关的调用什么的都不再去做bf的处理动作，而只是操作读写游标

// Reader 将io.Reader对象封装成Reader, 缓存io.Reader的数据到bf
// 设置r,w记录读写位置，err记录最近的一次error, readErr后需要重置
type Reader struct {
	bf []byte
	rd io.Reader
	r, w int
	err Error
}

// NewReaderSize 根据io.Reader以及buf size创建一个新的Reader
// 1.如果该io.reader是Reader对象，而且该对象buf长度充足，直接返回
// 2.size不得小于16
// 3.根据io.reader以及buf生成新的Reader对象
func NewReaderSize(reader io.Reader, size int) *Reader {
	b, ok := reader.(*Reader)
	if ok && len(b.buf) > size {
		return b, nil
	}
	if size < minReaderBufferSize {
		size = minReaderBufferSize
	}
	b := new(Reader)
	return b.reset(make([]byte, size), reader)
}

// NewReader 使用默认buffer size新建Reader对象
func NewReader(reader io.Reader) *Reader {
	NewReaderSize(reader, defaultBufferSize)
}

func (r *Reader) reset(buf []byte, rd io.Reader) {
	r.r = 0
	r.w = 0
	r.bf = buf
	r.rd = rd
}

// 读出最近一次err，重置r.err为nil，返回err
func (r *Reader) readErr() error {
	err := r.err
	r.err = nil
	return err
}

// fill 该function的作用是缓存io.reader的数据内容
// 1.如果读取游标r不为0，清除buf内已读数据, 重置读写游标
// 2.如果重置后的游标w超出buf size，应该报错
// 3.调用io.reader的Read接口，将io.reader中数据缓存到buf写游标w后
// 4.从io.reader中读到空数据时，允许重试100次
// 5.如果连续100次都是读到空数据，报错:multiple Read calls return no data or error
func (r *Reader) fill() {
	if r.r > 0 {
		copy(r.bf, r.br[r.r:r.w])
		r.w -= r.rd
		r.r = 0
	}

	if r.w > len(r.bf) {
		panic("bufio: tried to fill full buffer")
	}

	for i := 0; i < maxConsecutiveEmptyReadTime; i++ {
		n, err := r.rd.Read(r.bf[r.w:])
		if n < 0 {
			panic("bufio: reader returned negative count from Read")
		}
		if err != nil {
			r.err = err
			return
		}
		r.w += n
		if n > 0 {
			return
		}
	}

	r.err = io.ErrNoProgress
}

// 当前缓存内容长度
func (r *Reader) Bufferred() int { return r.w - r.r }

// Peek 只是引用reader内容，并不将数据读出，引用的内容下次使用时仍然有效
// 1.入参n不能小于0,也不能超出buf长度
// 2.如果缓存数据不够，需要重新调用fill从io.reader中缓存数据
// 3.如果缓存后可读数据仍然不够，则置err为ErrBufferFull，同时只返回当前可读数据
func (r *Reader) Peak(n int) ([]byte, error) {
	if n < 0 {
		return nil, ErrNegativeCount
	}
	if n > len(r.buf) {
		return nil, ErrBufferFull
	}

	if r.w - r.r < n && r.err == nil {
		r.fill()
	}

	var err error
	if avail := r.w - r.r; avail < n {
		err = r.readErr()
		if err == nil {
			r.err = ErrBufferFull
		}
		n = avail
	}
	return r.bf[r.r: r.r+n], err
}

// Pop 读出n个字节的数据
// 调用peek，如果成功(err == nil), 则重置游标r
// 如果失败，则返回nil以及err
func (r *Reader) Pop(n int) ([]byte, error) {
	buf, err := r.Peek(n)
	if err == nil {
		r.r += n
		return buf, nil
	}
	return nil, err
}

// Discard 丢弃n个字节的内容
// 1.判断入口参数，小于0不合法，等于0可以直接返回
// 2.循环丢弃，每次可以丢弃的数据长度是skip = r.w - r.r,
//   如果可丢弃的数据长度为0，则重新从io.reader中缓存，并重新设定skip为bufferred
// 3.如果skip长度大于还未丢弃的长度,则每轮设定skip为remain
// 4.丢弃动作本身并不操作bf，只需要设定读取游标
// 5.操作过程中遇到err应该立即返回
func (r *Reader) Discard(n int) (discarded int, err error) {
	if n < 0 {
		return 0, ErrNegativeCount
	}
	if n == 0 {
		return 0, nil
	}

	remain := n
	for {
		skip := r.Bufferred()
		if skip == 0 {
			r.fill()
			skip = r.Bufferred()
		}
		if skip > remain {
			skip = remain
		}
		r.r -= skip
		remain -= skip
		if remain == 0 {
			return n, nil
		}
		if r.err != nil {
			return n - remain, r.readErr()
		}
	}
}

// Read 读取reader的内容到p
// 1.判断参数，p的长度不能为0
// 2.首先从缓存中读取
// 3.当前缓存中数据为空的情况下，如果需要读取的长度大于缓存本身，则直接从io中读取
//   否则fill
func (r *Reader) Read(p []byte) (n int, err error) {
	n := len(p)
	if n == 0 {
		return 0, ErrNegativeCount
	}

	if r.r == r.w {
		if r.err != nil {
			return 0, r.readErr()
		}

		if len(p) > len(r.bf) {
			n, r.err = r.rd.Read(p)
			if n < 0 {
				panic("bufio: reader returned negative count from Read")
			}
			return n, r.readErr()
		}
		r.fill()
		if r.r == r.w {
			return 0, r.readErr()
		}
	}

	n := copy(p, r.bf[r.r:r.w])
	return n, r.readErr
}
