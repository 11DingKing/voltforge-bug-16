package certification

type ReportRestoreIndex struct{ buckets map[string][]string }

func NewReportRestoreIndex() *ReportRestoreIndex { return &ReportRestoreIndex{} }
func (i *ReportRestoreIndex) Restore(bucket, value string) {
	i.buckets[bucket] = append(i.buckets[bucket], value)
}
