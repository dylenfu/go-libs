package leveldb

func Route(sub string) {
	switch sub {
	case "simple-put-get":
		SimplePutAndGet()
		break

	case "simple-batch":
		SimpleBatch()
		break

	case "simple-batch-load":
		SimpleBatchLoad()
		break

	case "simple-get-property":
		SimpleGetProperty()
		break

	case "simple-get-snapshot":
		SimpleGetSnapshot()
		break

	case "simple-db-iterator":
		SimpleNewDBIterator()
		break

	case "simple-iterator-seek":
		SimpleDBIteratorSeek()
		break

	case "simple-iterator-prefix":
		SimpleIteratorWithPrefix()
		break

	case "simple-filter":
		SimpleFilter()
		break
	}
}
