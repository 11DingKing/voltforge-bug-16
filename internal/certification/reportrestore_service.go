package certification

type ReportRestoreIndex struct{ buckets map[string][]string }

func NewReportRestoreIndex() *ReportRestoreIndex {
	return &ReportRestoreIndex{buckets: make(map[string][]string)}
}
func (i *ReportRestoreIndex) Restore(bucket, value string) {
	if i.buckets == nil {
		i.buckets = make(map[string][]string)
	}
	i.buckets[bucket] = append(i.buckets[bucket], value)
}
